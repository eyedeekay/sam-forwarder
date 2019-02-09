package i2pkeyspass

import (
//"io/ioutil"
//"os"
//"log"
//"path/filepath"

//"github.com/eyedeekay/sam3"
//"github.com/gtank/cryptopasta"
)

func bytes(k [32]byte) []byte {
	var r []byte
	for _, v := range k {
		r = append(r, v)
	}
	return r
}

func key(k []byte) *[32]byte {
	var r [32]byte
	for i, v := range k {
		r[i] = v
	}
	return &r
}

func EncryptPassword(i2pkeypath, password string) error {
	return nil
}

func DecryptPassword(i2pkeypath, password string) error {
	return nil
}
