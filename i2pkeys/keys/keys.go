package i2pkeyscrypt

import (
	"io/ioutil"
	"os"

	"github.com/gtank/cryptopasta"
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

func EncryptKey(i2pkeypath, aeskeypath string) error {
	if aeskeypath != "" {
		if r, e := ioutil.ReadFile(i2pkeypath); e != nil {
			return e
		} else {
			if _, err := os.Stat(aeskeypath); os.IsNotExist(err) {
				key := cryptopasta.NewEncryptionKey()
				ioutil.WriteFile(aeskeypath, bytes(*key), 644)
			} else if err != nil {
				return err
			}
			if ra, re := ioutil.ReadFile(aeskeypath); re != nil {
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

func DecryptKey(i2pkeypath, aeskeypath string) error {
	if aeskeypath != "" {
		if r, e := ioutil.ReadFile(i2pkeypath); e != nil {
			return e
		} else {
			if _, err := os.Stat(aeskeypath); os.IsNotExist(err) {
				return err
			}
			if ra, re := ioutil.ReadFile(aeskeypath); re != nil {
				return e
			} else {
				crypted, err := cryptopasta.Decrypt(r, key(ra))
				if err != nil {
					return err
				}
				ioutil.WriteFile(i2pkeypath, crypted, 644)
			}
			//crypted
		}
	}
	return nil
}
