package i2pkeys

import (
	"github.com/eyedeekay/sam3"
	"github.com/gtank/cryptopasta"
	"io/ioutil"
	"os"
	"path/filepath"
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

func Decrypt(i2pkeypath, aeskeypath string) error {
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

func Save(FilePath, TunName, passfile string, SamKeys *sam3.I2PKeys) error {
	if _, err := os.Stat(filepath.Join(FilePath, TunName+".i2pkeys")); os.IsNotExist(err) {
		file, err := os.Create(filepath.Join(FilePath, TunName+".i2pkeys"))
		if err != nil {
			return err
		}
		err = sam3.StoreKeysIncompat(*SamKeys, file)
		if err != nil {
			return err
		}
		err = Encrypt(filepath.Join(FilePath, TunName+".i2pkeys"), passfile)
		if err != nil {
			return err
		}
	}
	file, err := os.Open(filepath.Join(FilePath, TunName+".i2pkeys"))
	if err != nil {
		return err
	}
	err = Decrypt(filepath.Join(FilePath, TunName+".i2pkeys"), passfile)
	if err != nil {
		return err
	}
	tempkeys, err := sam3.LoadKeysIncompat(file)
	if err != nil {
		return err
	}
	SamKeys = &tempkeys
	err = Encrypt(filepath.Join(FilePath, TunName+".i2pkeys"), passfile)
	if err != nil {
		return err
	}
	return nil
}
