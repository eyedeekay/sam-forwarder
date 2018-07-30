package samforwarderudp

import (
	"log"
	"testing"
)

func TestOptionHost(t *testing.T) {
	client, err := NewSAMSSUForwarderFromOptions(SetHost("127.0.0.1"))
	if err != nil {
		t.Fatalf("NewDefaultClient() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestOptionPort(t *testing.T) {
	client, err := NewSAMSSUForwarderFromOptions(SetPort("7656"))
	if err != nil {
		t.Fatalf("NewDefaultClient() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestOptionInLength(t *testing.T) {
	client, err := NewSAMSSUForwarderFromOptions(SetInLength(3))
	if err != nil {
		t.Fatalf("NewDefaultClient() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestOptionOutLength(t *testing.T) {
	client, err := NewSAMSSUForwarderFromOptions(SetInLength(3))
	if err != nil {
		t.Fatalf("NewDefaultClient() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestOptionInVariance(t *testing.T) {
	client, err := NewSAMSSUForwarderFromOptions(SetInVariance(1))
	if err != nil {
		t.Fatalf("NewDefaultClient() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestOptionOutVariance(t *testing.T) {
	client, err := NewSAMSSUForwarderFromOptions(SetOutVariance(1))
	if err != nil {
		t.Fatalf("NewDefaultClient() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestOptionInQuantity(t *testing.T) {
	client, err := NewSAMSSUForwarderFromOptions(SetInQuantity(6))
	if err != nil {
		t.Fatalf("NewDefaultClient() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestOptionOutQuantity(t *testing.T) {
	client, err := NewSAMSSUForwarderFromOptions(SetOutQuantity(6))
	if err != nil {
		t.Fatalf("NewDefaultClient() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestOptionInBackups(t *testing.T) {
	client, err := NewSAMSSUForwarderFromOptions(SetInBackups(5))
	if err != nil {
		t.Fatalf("NewDefaultClient() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestOptionOutBackups(t *testing.T) {
	client, err := NewSAMSSUForwarderFromOptions(SetOutBackups(5))
	if err != nil {
		t.Fatalf("NewDefaultClient() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestOptionReduceIdleQuantity(t *testing.T) {
	client, err := NewSAMSSUForwarderFromOptions(SetReduceIdleQuantity(4))
	if err != nil {
		t.Fatalf("NewDefaultClient() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestOptionEncryptLease(t *testing.T) {
	client, err := NewSAMSSUForwarderFromOptions(SetEncrypt(true))
	if err != nil {
		t.Fatalf("NewDefaultClient() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestOptionSaveFile(t *testing.T) {
	client, err := NewSAMSSUForwarderFromOptions(SetSaveFile(true))
	if err != nil {
		t.Fatalf("NewDefaultClient() Error: %q\n", err)
	}
	log.Println(client.Base32())
}
