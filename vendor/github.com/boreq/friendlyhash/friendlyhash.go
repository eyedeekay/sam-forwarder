// Package friendlyhash implements human-readable and reversible representation
// of known-length byte slices.
package friendlyhash

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"math"
)

// New creates a struct responsible for encoding and decoding the data using
// the provided dictionary and hash length. The dictionary should be a list of
// at least two unique strings, usually words. To decode once encoded data an
// identical dictionary has to be used. Hash length is required to avoid
// including the length of the data in the resulting representation or using
// padding characters.
func New(dictionary []string, hashLength int) (*FriendlyHash, error) {
	if err := validateDictionary(dictionary); err != nil {
		return nil, err
	}
	rv := &FriendlyHash{
		dictionary: dictionary,
		hashLength: hashLength,
	}
	return rv, nil
}

// FriendlyHash creates a human-friendly representation of byte slices.
type FriendlyHash struct {
	dictionary []string
	hashLength int
}

// Humanize encodes the provided byte slice and returns its representation as a
// list of words from the dictionary.
func (h *FriendlyHash) Humanize(hash []byte) ([]string, error) {
	if len(hash) != h.hashLength {
		return nil, errors.New("invalid hash length")
	}

	bitsPerWord := howManyBitsPerWord(len(h.dictionary))

	indexes, err := h.splitIntoIndexes(hash, bitsPerWord)
	if err != nil {
		return nil, err
	}

	var words []string
	for _, index := range indexes {
		words = append(words, h.dictionary[index])
	}
	return words, nil
}

// Dehumanize converts the provided list of words previously created using the
// humanize function back to its byte slice equivalent.
func (h *FriendlyHash) Dehumanize(words []string) ([]byte, error) {
	bitsPerWord := howManyBitsPerWord(len(h.dictionary))

	if len(words) != howManyWords(bitsPerWord, h.hashLength) {
		return nil, errors.New("invalid words length")
	}

	to := make([]byte, h.hashLength)
	for i, word := range words {
		// Get the byte representation of this word
		wordIndex, err := findIndex(word, h.dictionary)
		if err != nil {
			return nil, err
		}

		buf := &bytes.Buffer{}
		if err := binary.Write(buf, binary.BigEndian, uint32(wordIndex)); err != nil {
			return nil, err
		}

		// Copy the relevant bits to the output slice
		from := buf.Bytes()
		fromI := len(from)*8 - bitsPerWord

		toI := i * bitsPerWord

		n := bitsPerWord
		if toI+n > len(to)*8 { // In case we run out of space in the to slice
			n = len(to)*8 - toI // Use only as many bits as can fit in that slice
		}

		if err := copyBits(from, fromI, to, toI, n); err != nil {
			return nil, err
		}
	}
	return to, nil
}

// NumberOfWords returns the number of words returned by the humanize function.
func (h *FriendlyHash) NumberOfWords() int {
	bitsPerWord := howManyBitsPerWord(len(h.dictionary))
	return howManyWords(bitsPerWord, h.hashLength)
}

// NumberOfBytes returns the hash length returned by the dehumanize function.
func (h *FriendlyHash) NumberOfBytes() int {
	return h.hashLength
}

func findIndex(element string, slice []string) (int, error) {
	for i, sliceElement := range slice {
		if sliceElement == element {
			return i, nil
		}
	}
	return -1, fmt.Errorf("word is not in the dictionary: %s", element)
}

func (h *FriendlyHash) splitIntoIndexes(data []byte, bitsPerWord int) ([]int, error) {
	var indexes []int
	for bitI := 0; bitI < len(data)*8; bitI += bitsPerWord {
		index, err := getBits(data, bitI, bitsPerWord)
		if err != nil {
			return nil, err
		}
		indexes = append(indexes, int(index))
	}
	return indexes, nil
}

// getBits attempts to extract "n" bits starting from the bit position "i" in
// the data slice and return them as a number.
func getBits(data []byte, i int, n int) (uint32, error) {
	var tmpData = make([]byte, 4)

	length := len(data) * 8
	nBits := n
	if i+nBits > length { // If there are no n bits available
		nBits = length - i // Copy all remaining bits
	}

	if err := copyBits(data, i, tmpData, len(tmpData)*8-n, nBits); err != nil {
		return 0, err
	}

	var rv uint32
	if err := binary.Read(bytes.NewBuffer(tmpData), binary.BigEndian, &rv); err != nil {
		return 0, err
	}
	return rv, nil
}

// copyBits copies n bits between the specified byte slices, starting with the
// specified bit indexes.
func copyBits(from []byte, fromStartBit int, to []byte, toStartBit int, n int) error {
	if fromStartBit+n > len(from)*8 {
		return errors.New("not enough bits in the from slice")
	}

	if toStartBit+n > len(to)*8 {
		return errors.New("not enough bits in the to slice")
	}

	for i := 0; i < n; i++ {
		// read bit
		fromBitI := fromStartBit + i
		currentFromByte := fromBitI / 8
		currentFromBit := fromBitI % 8
		bitState := checkBit(from[currentFromByte], currentFromBit)

		// write bit
		toBitI := toStartBit + i
		currentToByte := toBitI / 8
		currentToBit := toBitI % 8
		if bitState {
			to[currentToByte] = setBit(to[currentToByte], currentToBit)
		} else {
			to[currentToByte] = clearBit(to[currentToByte], currentToBit)
		}
	}
	return nil
}

// checkBit returns true if the specified bit is set to 1 and false otherwise.
func checkBit(b byte, i int) bool {
	var mask byte = 1 << uint(7-i)
	return b&mask != 0
}

// setBit returns the provided byte with the specified bit set to 1.
func setBit(b byte, i int) byte {
	return b | 1<<uint(7-i)
}

// clearBit returns the provided byte with the specified bit set to 0.
func clearBit(b byte, i int) byte {
	return b & ^(1 << uint(7-i))
}

// howManyWords returns the number of words needed to encode a hash of a
// specified length if each word represents a specified amount of bits.
func howManyWords(bitsPerWord int, hashLength int) int {
	return int(math.Ceil(float64(hashLength) * 8.0 / float64(bitsPerWord)))
}

// howManyBitsPerWord returns the number of bits that can be represented using
// a single word from a dictionary of the specified size.
func howManyBitsPerWord(numberOfWords int) int {
	return int(math.Floor(math.Log2(float64(numberOfWords))))
}

func validateDictionary(dictionary []string) error {
	// Check length
	if len(dictionary) < 2 {
		return errors.New("dictionary must have at least 2 entries")
	}

	// Check duplicates
	m := make(map[string]bool)
	for _, word := range dictionary {
		m[word] = true
	}
	if len(m) != len(dictionary) {
		return errors.New("dictionary entries must be unique")
	}

	return nil
}
