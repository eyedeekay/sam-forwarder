package i2ptunconf

import (
	"log"
	"testing"
)

func TestConf(t *testing.T) {
	log.Println("testing configuration loader")
	if config, err := NewI2PTunConf("../etc/sam-forwarder/tunnels.ini"); err != nil {
		log.Fatal(err)
	} else {
		log.Println(config.Print())
	}
}
