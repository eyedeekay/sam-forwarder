package samforwardervpn

import (
	"github.com/eyedeekay/sam-forwarder/config"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
	"time"
)

func echo() {
	http.Handle("/", http.FileServer(http.Dir("../server_test/www")))
	log.Fatal(http.ListenAndServe("0.0.0.0:8022", nil))
}

func TestVPN(t *testing.T) {
	log.Println("Setting up VPN")
	go echo()

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
					log.Println(&vpnc)
					time.Sleep(time.Duration(30 * time.Second))
					resp, err := http.Get("http://10.76.0.2:8022/test.html")
					if err != nil {
						t.Fatal(err)
					}
					defer resp.Body.Close()
					if resp.StatusCode == http.StatusOK {
						bodyBytes, err := ioutil.ReadAll(resp.Body)
						if err != nil {
							t.Fatal(err)
						}
						log.Println(string(bodyBytes))
					}
				}
			} else {
				t.Fatal(err)
			}
		}
	} else {
		t.Fatal(err)
	}

}
