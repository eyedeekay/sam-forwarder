package sammanager

import (
	"log"
	"testing"
)

func TestOptionHost(t *testing.T) {
	client, err := NewSAMManagerFromOptions(SetManagerHost("127.0.0.1"))
	if err != nil {
		t.Fatalf("NewSAMManager() Error: %q\n", err)
	}
	log.Println(client.config.Labels)
}

func TestOptionPort(t *testing.T) {
	client, err := NewSAMManagerFromOptions(SetManagerPort("7957"))
	if err != nil {
		t.Fatalf("NewSAMManager() Error: %q\n", err)
	}
	log.Println(client.config.Labels)
}

func TestOptionSAMHost(t *testing.T) {
	client, err := NewSAMManagerFromOptions(SetManagerSAMHost("127.0.0.1"))
	if err != nil {
		t.Fatalf("NewSAMManager() Error: %q\n", err)
	}
	log.Println(client.config.Labels)
}

func TestOptionSAMPort(t *testing.T) {
	client, err := NewSAMManagerFromOptions(SetManagerSAMPort("7957"))
	if err != nil {
		t.Fatalf("NewSAMManager() Error: %q\n", err)
	}
	log.Println(client.config.Labels)
}
