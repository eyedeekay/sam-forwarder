package sfi2pkeys

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/eyedeekay/sam-forwarder/i2pkeys/keys"
	//"github.com/eyedeekay/sam-forwarder/i2pkeys/password"
	"github.com/eyedeekay/sam3"
	"github.com/eyedeekay/sam3/i2pkeys"
)

func Encrypt(i2pkeypath, aeskeypath string) error {
	return i2pkeyscrypt.EncryptKey(i2pkeypath, aeskeypath)
}

func Decrypt(i2pkeypath, aeskeypath string) error {
	return i2pkeyscrypt.DecryptKey(i2pkeypath, aeskeypath)
}

func Save(FilePath, TunName, passfile string, SamKeys i2pkeys.I2PKeys) error {
	if _, err := os.Stat(filepath.Join(FilePath, TunName+".i2pkeys")); os.IsNotExist(err) {
		file, err := os.Create(filepath.Join(FilePath, TunName+".i2pkeys"))
		if err != nil {
			return err
		}
		err = i2pkeys.StoreKeysIncompat(SamKeys, file)
		if err != nil {
			return err
		}
		//err = Encrypt(filepath.Join(FilePath, TunName+".i2pkeys"), passfile)
		//if err != nil {
		//return err
		//}
		return nil
	}
	file, err := os.Open(filepath.Join(FilePath, TunName+".i2pkeys"))
	if err != nil {
		return err
	}
	//err = Decrypt(filepath.Join(FilePath, TunName+".i2pkeys"), passfile)
	//if err != nil {
	//return err
	//}
	SamKeys, err = i2pkeys.LoadKeysIncompat(file)
	if err != nil {
		return err
	}
	//SamKeys = &tempkeys
	//err = Encrypt(filepath.Join(FilePath, TunName+".i2pkeys"), passfile)
	//if err != nil {
	//return err
	//}
	return nil
}

func Load(FilePath, TunName, passfile string, samConn *sam3.SAM, save bool) (i2pkeys.I2PKeys, error) {
	if !save {
		return samConn.NewKeys()
	}
	if _, err := os.Stat(filepath.Join(FilePath, TunName+".i2pkeys")); os.IsNotExist(err) {
		log.Println("Generating keys from SAM bridge")
		SamKeys, err := samConn.NewKeys()
		if err != nil {
			return i2pkeys.I2PKeys{}, err
		}
		return SamKeys, nil
	}
	log.Println("Generating keys from disk")
	file, err := os.Open(filepath.Join(FilePath, TunName+".i2pkeys"))
	if err != nil {
		return i2pkeys.I2PKeys{}, err
	}
	//err = Decrypt(filepath.Join(FilePath, TunName+".i2pkeys"), passfile)
	//if err != nil {
	//return i2pkeys.I2PKeys{}, err
	//}
	return i2pkeys.LoadKeysIncompat(file)
}

func Prop(in string) (string, string) {
	k := "unset"
	v := "unset"
	vals := strings.SplitN(in, "=", 2)
	if len(vals) >= 1 {
		k = vals[0]
	}
	if len(vals) >= 2 {
		v = vals[1]
	}
	return k, v
}
