package hashhash

import (
	"io/ioutil"
	"strings"
)

import (
	"github.com/boreq/friendlyhash"
)

type Hasher struct {
	FileName string
	split    []string
	length   int
	hasher   *friendlyhash.FriendlyHash
}

func (h *Hasher) dictionary() []string {
	if len(h.split) > 0 {
		return h.split
	}
	bytes, _ := ioutil.ReadFile(h.FileName)
	h.split = strings.Split(string(bytes), "\n")
	return h.split
}

func (h *Hasher) Friendly(input string) (string, error) {
	slice, err := h.hasher.Humanize([]byte(input))
	if err != nil {
		return "", err
	}
	return strings.Join(slice, " "), nil
}

func (h *Hasher) Unfriendly(input string) (string, error) {
	slice := strings.Split(input, " ")
	hash, err := h.hasher.Dehumanize(slice)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (h *Hasher) Unfriendlyslice(input []string) (string, error) {
	hash, err := h.hasher.Dehumanize(input)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func NewHasher(length int) (*Hasher, error) {
	var h Hasher
	var err error
	h.FileName = ""
	h.length = length
	h.split = Default()
	h.hasher, err = friendlyhash.New(h.dictionary(), length)
	if err != nil {
		return nil, err
	}
	return &h, nil
}

func Init(FileName string, length int) (*Hasher, error) {
	var h Hasher
	var err error
	h.FileName = FileName
	h.length = length
	h.hasher, err = friendlyhash.New(h.dictionary(), length)
	if err != nil {
		return nil, err
	}
	return &h, nil
}
