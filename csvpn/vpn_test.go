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
	if sconfig, err := i2ptunconf.NewI2PTunConf("../etc/i2pvpn/i2pvpn.ini"); err == nil {
		if vpn, err := NewSAMClientServerVPN(sconfig); err != nil {
			t.Fatal(err)
		} else {
			time.Sleep(time.Duration(30 * time.Second))
			log.Println(&vpn)
			if config, err := i2ptunconf.NewI2PTunConf("../etc/i2pvpn/i2pvpnclient.ini"); err == nil {
				if vpnc, err := NewSAMClientVPN(config, vpn.Base32()); err != nil {
					t.Fatal(err)
				} else {
					time.Sleep(time.Duration(30 * time.Second))
					log.Println(&vpnc)
				}
			} else {
				t.Fatal(err)
			}
		}
	} else {
		t.Fatal(err)
	}

}
