package i2pkeys

import (
	"log"
	"os"
	"path/filepath"

	"github.com/eyedeekay/sam-forwarder/i2pkeys/keys"
	//"github.com/eyedeekay/sam-forwarder/i2pkeys/password"
	"github.com/eyedeekay/sam3"
)

func Encrypt(i2pkeypath, aeskeypath string) error {
	return i2pkeyscrypt.EncryptKey(i2pkeypath, aeskeypath)
}

func Decrypt(i2pkeypath, aeskeypath string) error {
	return i2pkeyscrypt.DecryptKey(i2pkeypath, aeskeypath)
}

func Save(FilePath, TunName, passfile string, SamKeys sam3.I2PKeys) error {
	if _, err := os.Stat(filepath.Join(FilePath, TunName+".i2pkeys")); os.IsNotExist(err) {
		file, err := os.Create(filepath.Join(FilePath, TunName+".i2pkeys"))
		if err != nil {
			return err
		}
		err = sam3.StoreKeysIncompat(SamKeys, file)
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
	SamKeys, err = sam3.LoadKeysIncompat(file)
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

func Load(FilePath, TunName, passfile string, samConn *sam3.SAM, save bool) (sam3.I2PKeys, error) {
    if ! save {
        return samConn.NewKeys()
    }
	if _, err := os.Stat(filepath.Join(FilePath, TunName+".i2pkeys")); os.IsNotExist(err) {
		log.Println("Generating keys from SAM bridge")
		SamKeys, err := samConn.NewKeys()
		if err != nil {
			return sam3.I2PKeys{}, err
		}
		return SamKeys, nil
	}
	log.Println("Generating keys from disk")
	file, err := os.Open(filepath.Join(FilePath, TunName+".i2pkeys"))
	if err != nil {
		return sam3.I2PKeys{}, err
	}
	//err = Decrypt(filepath.Join(FilePath, TunName+".i2pkeys"), passfile)
	//if err != nil {
	//return sam3.I2PKeys{}, err
	//}
	return sam3.LoadKeysIncompat(file)
}
