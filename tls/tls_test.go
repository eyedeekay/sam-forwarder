package i2ptls

import "testing"

func TestTLSGenerate(t *testing.T) {
	config, err := TLSConfig("cert.pem", "key.pem", []string{"idk.i2p"})
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(config.ServerName)
}
