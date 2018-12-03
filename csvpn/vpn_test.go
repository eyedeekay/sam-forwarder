package i2pvpn

import (
	"log"
	//	"net/http"
	"github.com/eyedeekay/sam-forwarder/config"
	"testing"
	"time"
)

func TestVPN(t *testing.T) {
	log.Println("Setting up VPN")
	if config, err := i2ptunconf.NewI2PTunConf("../etc/i2pvpn/i2pvpn.ini"); err == nil {
		if vpn, err := NewSAMClientServerVPN(config); err != nil {
			t.Fatal(err)
		} else {
			time.Sleep(time.Duration(30 * time.Second))
			log.Println(&vpn)
		}
	} else {
		t.Fatal(err)
	}
}
