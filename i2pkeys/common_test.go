package i2pkeys

import (
	//"os"
	"log"
	"testing"
	//"path/filepath"

	"github.com/eyedeekay/sam3"
)

func TestKeysGenLoad(t *testing.T) {
	sc, err := sam3.NewSAM("127.0.0.1:7656")
	if err != nil {
		t.Fatal(err)
	}
	log.Println("Initialized SAM connection")
	sk, err := Load("./", "test", "", sc)
	if err != nil {
		t.Fatal(err)
	}
	err = Save("./", "test", "", &sk)
	if err != nil {
		t.Fatal(err)
	}
}
