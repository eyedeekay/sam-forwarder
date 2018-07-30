package samforwarderudp

import (
	"log"
	"testing"
)

func TestOptionUDPHost(t *testing.T) {
	client, err := NewSAMSSUForwarderFromOptions(SetHost("127.0.0.1"))
	if err != nil {
		t.Fatalf("NewSAMSSUForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestOptionUDPPort(t *testing.T) {
	client, err := NewSAMSSUForwarderFromOptions(SetPort("7656"))
	if err != nil {
		t.Fatalf("NewSAMSSUForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestOptionUDPInLength(t *testing.T) {
	client, err := NewSAMSSUForwarderFromOptions(SetInLength(3))
	if err != nil {
		t.Fatalf("NewSAMSSUForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestOptionUDPOutLength(t *testing.T) {
	client, err := NewSAMSSUForwarderFromOptions(SetInLength(3))
	if err != nil {
		t.Fatalf("NewSAMSSUForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestOptionUDPInVariance(t *testing.T) {
	client, err := NewSAMSSUForwarderFromOptions(SetInVariance(1))
	if err != nil {
		t.Fatalf("NewSAMSSUForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestOptionUDPOutVariance(t *testing.T) {
	client, err := NewSAMSSUForwarderFromOptions(SetOutVariance(1))
	if err != nil {
		t.Fatalf("NewSAMSSUForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestOptionUDPInQuantity(t *testing.T) {
	client, err := NewSAMSSUForwarderFromOptions(SetInQuantity(6))
	if err != nil {
		t.Fatalf("NewSAMSSUForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestOptionUDPOutQuantity(t *testing.T) {
	client, err := NewSAMSSUForwarderFromOptions(SetOutQuantity(6))
	if err != nil {
		t.Fatalf("NewSAMSSUForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestOptionUDPInBackups(t *testing.T) {
	client, err := NewSAMSSUForwarderFromOptions(SetInBackups(5))
	if err != nil {
		t.Fatalf("NewSAMSSUForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestOptionUDPOutBackups(t *testing.T) {
	client, err := NewSAMSSUForwarderFromOptions(SetOutBackups(5))
	if err != nil {
		t.Fatalf("NewSAMSSUForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestOptionUDPReduceIdleQuantity(t *testing.T) {
	client, err := NewSAMSSUForwarderFromOptions(SetReduceIdleQuantity(4))
	if err != nil {
		t.Fatalf("NewSAMSSUForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestOptionUDPEncryptLease(t *testing.T) {
	client, err := NewSAMSSUForwarderFromOptions(SetEncrypt(true))
	if err != nil {
		t.Fatalf("NewSAMSSUForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestOptionUDPSaveFile(t *testing.T) {
	client, err := NewSAMSSUForwarderFromOptions(SetSaveFile(true))
	if err != nil {
		t.Fatalf("NewSAMSSUForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}
