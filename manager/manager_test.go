package sammanager

import (
	"log"
	"testing"
)

func TestOptionFile(t *testing.T) {
	client, err := NewSAMManagerFromOptions(
		SetManagerFilePath("../etc/sam-forwarder/tunnels.ini"),
	)
	if err != nil {
		t.Fatalf("NewSAMManager() Error: %q\n", err)
	}
	log.Println(client.config.Type)
}

func TestOptionHost(t *testing.T) {
	client, err := NewSAMManagerFromOptions(
		SetManagerHost("127.0.0.1"),
		SetManagerFilePath("../etc/sam-forwarder/tunnels.ini"),
	)
	if err != nil {
		t.Fatalf("NewSAMManager() Error: %q\n", err)
	}
	log.Println(client.config.Type)
}

func TestOptionPort(t *testing.T) {
	client, err := NewSAMManagerFromOptions(
		SetManagerPort("7957"),
		SetManagerFilePath("../etc/sam-forwarder/tunnels.ini"),
	)
	if err != nil {
		t.Fatalf("NewSAMManager() Error: %q\n", err)
	}
	log.Println(client.config.Type)
}

func TestOptionSAMHost(t *testing.T) {
	client, err := NewSAMManagerFromOptions(
		SetManagerSAMHost("127.0.0.1"),
		SetManagerFilePath("../etc/sam-forwarder/tunnels.ini"),
	)
	if err != nil {
		t.Fatalf("NewSAMManager() Error: %q\n", err)
	}
	log.Println(client.config.Type)
}

func TestOptionSAMPort(t *testing.T) {
	client, err := NewSAMManagerFromOptions(
		SetManagerSAMPort("7656"),
		SetManagerFilePath("../etc/sam-forwarder/tunnels.ini"),
	)
	if err != nil {
		t.Fatalf("NewSAMManager() Error: %q\n", err)
	}
	log.Println(client.config.Type)
}
