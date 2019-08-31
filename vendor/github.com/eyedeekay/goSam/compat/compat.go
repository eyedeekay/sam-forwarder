package samkeys

import (
	"fmt"
	"github.com/eyedeekay/sam3/i2pkeys"
	"strings"
)

func DestToKeys(dest string) (i2pkeys.I2PKeys, error) {
	addr, err := i2pkeys.NewI2PAddrFromString(dest)
	if err != nil {
		return i2pkeys.I2PKeys{}, err
	}
	return i2pkeys.NewKeys(addr, dest), nil
}

func KeysToDest(keys i2pkeys.I2PKeys) (string, error) {
	pksk := strings.SplitN(keys.String(), "\n", 2)
	if len(pksk) != 2 {
		return "", fmt.Errorf("Error converting from keys to destination")
	}
	return pksk[1], nil
}
