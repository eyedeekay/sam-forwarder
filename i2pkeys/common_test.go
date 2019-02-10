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
	sk, err := Load("./", "test", "", sc, true)
	if err != nil {
		t.Fatal(err)
	}
	log.Println("Loaded tunnel keys")
	err = Save("./", "test", "", sk)
	if err != nil {
		t.Fatal(err)
	}
}

func TestKeysGenLoadAgain(t *testing.T) {
	sc, err := sam3.NewSAM("127.0.0.1:7656")
	if err != nil {
		t.Fatal(err)
	}
	log.Println("Saved tunnel keys")
	sk, err := Load("./", "test", "", sc, true)
	if err != nil {
		t.Fatal(err)
	}
	log.Println("Loaded tunnel keys 2")
	err = Save("./", "test2", "", sk)
	if err != nil {
		t.Fatal(err)
	}
	log.Println("Saved tunnel keys 2")
}
