package sammanager

import (
	"log"
	"testing"
)

func TestOptionSAMHost(t *testing.T) {
	client, err := NewSAMManagerFromOptions(
		SetManagerHost("127.0.0.1"),
		SetManagerSAMHost("127.0.0.1"),
		SetManagerPort("7957"),
		SetManagerSAMPort("7656"),
		SetManagerFilePath("../etc/sam-forwarder/tunnels.ini"),
	)
	if err != nil {
		t.Fatalf("NewSAMManager() Error: %q\n", err)
	}
	log.Println(client.config.Type)
}
