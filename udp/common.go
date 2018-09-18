package samforwarderudp

import (
	"github.com/gtank/cryptopasta"
	"io/ioutil"
	"os"
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

func Encrypt(i2pkeypath, aeskeypath string) error {
	if aeskeypath != "" {
		if r, e := ioutil.ReadFile(i2pkeypath); e != nil {
			return e
		} else {
			if _, err := os.Stat(aeskeypath + ".key"); os.IsNotExist(err) {
				key := cryptopasta.NewEncryptionKey()
				ioutil.WriteFile(aeskeypath, bytes(*key), 644)
			} else if err != nil {
				return err
			}
			if ra, re := ioutil.ReadFile(aeskeypath + ".key"); re != nil {
				return e
			} else {
				crypted, err := cryptopasta.Encrypt(r, key(ra))
				if err != nil {
					return err
				}
				ioutil.WriteFile(i2pkeypath, crypted, 644)
			}
		}
	}
	return nil
}

func Decrypt(i2pkeypath, aeskeypath string) error {
	if aeskeypath != "" {
		if r, e := ioutil.ReadFile(i2pkeypath); e != nil {
			return e
		} else {
			if _, err := os.Stat(aeskeypath + ".key"); os.IsNotExist(err) {
				return nil
			}
			if ra, re := ioutil.ReadFile(aeskeypath + ".key"); re != nil {
				return e
			} else {
				crypted, err := cryptopasta.Decrypt(r, key(ra))
				if err != nil {
					return err
				}
				ioutil.WriteFile(i2pkeypath, crypted, 644)
			}
		}
	}
	return nil
}