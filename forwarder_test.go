package samforwarder

import (
	"log"
	"testing"
)

func TestOptionHost(t *testing.T) {
	client, err := NewSAMForwarderFromOptions(SetHost("127.0.0.1"))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestOptionPort(t *testing.T) {
	client, err := NewSAMForwarderFromOptions(SetPort("7656"))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestOptionInLength(t *testing.T) {
	client, err := NewSAMForwarderFromOptions(SetInLength(3))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestOptionOutLength(t *testing.T) {
	client, err := NewSAMForwarderFromOptions(SetInLength(3))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestOptionInVariance(t *testing.T) {
	client, err := NewSAMForwarderFromOptions(SetInVariance(1))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestOptionOutVariance(t *testing.T) {
	client, err := NewSAMForwarderFromOptions(SetOutVariance(1))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestOptionInQuantity(t *testing.T) {
	client, err := NewSAMForwarderFromOptions(SetInQuantity(6))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestOptionOutQuantity(t *testing.T) {
	client, err := NewSAMForwarderFromOptions(SetOutQuantity(6))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestOptionInBackups(t *testing.T) {
	client, err := NewSAMForwarderFromOptions(SetInBackups(5))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestOptionOutBackups(t *testing.T) {
	client, err := NewSAMForwarderFromOptions(SetOutBackups(5))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestOptionReduceIdleQuantity(t *testing.T) {
	client, err := NewSAMForwarderFromOptions(SetReduceIdleQuantity(4))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestOptionReduceIdleTimeMs(t *testing.T) {
	client, err := NewSAMForwarderFromOptions(SetReduceIdleTimeMs(300000))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestOptionReduceIdleTime(t *testing.T) {
	client, err := NewSAMForwarderFromOptions(SetReduceIdleTime(6))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestOptionCloseIdleTimeMs(t *testing.T) {
	client, err := NewSAMForwarderFromOptions(SetCloseIdleTimeMs(300000))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestOptionCloseIdleTime(t *testing.T) {
	client, err := NewSAMForwarderFromOptions(SetCloseIdleTime(6))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestOptionEncryptLease(t *testing.T) {
	client, err := NewSAMForwarderFromOptions(SetEncrypt(true))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestOptionSaveFile(t *testing.T) {
	client, err := NewSAMForwarderFromOptions(SetSaveFile(true))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}
