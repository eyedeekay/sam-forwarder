package samforwarderudp

import (
	"log"
	"testing"
)

func TestOptionUDPHost(t *testing.T) {
	client, err := NewSAMDGForwarderFromOptions(SetHost("127.0.0.1"))
	if err != nil {
		t.Fatalf("NewSAMDGForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestOptionUDPPort(t *testing.T) {
	client, err := NewSAMDGForwarderFromOptions(SetPort("7656"))
	if err != nil {
		t.Fatalf("NewSAMDGForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestOptionUDPInLength(t *testing.T) {
	client, err := NewSAMDGForwarderFromOptions(SetInLength(3))
	if err != nil {
		t.Fatalf("NewSAMDGForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestOptionUDPOutLength(t *testing.T) {
	client, err := NewSAMDGForwarderFromOptions(SetInLength(3))
	if err != nil {
		t.Fatalf("NewSAMDGForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestOptionUDPInVariance(t *testing.T) {
	client, err := NewSAMDGForwarderFromOptions(SetInVariance(1))
	if err != nil {
		t.Fatalf("NewSAMDGForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestOptionUDPOutVariance(t *testing.T) {
	client, err := NewSAMDGForwarderFromOptions(SetOutVariance(1))
	if err != nil {
		t.Fatalf("NewSAMDGForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestOptionUDPInQuantity(t *testing.T) {
	client, err := NewSAMDGForwarderFromOptions(SetInQuantity(6))
	if err != nil {
		t.Fatalf("NewSAMDGForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestOptionUDPOutQuantity(t *testing.T) {
	client, err := NewSAMDGForwarderFromOptions(SetOutQuantity(6))
	if err != nil {
		t.Fatalf("NewSAMDGForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestOptionUDPInBackups(t *testing.T) {
	client, err := NewSAMDGForwarderFromOptions(SetInBackups(5))
	if err != nil {
		t.Fatalf("NewSAMDGForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestOptionUDPOutBackups(t *testing.T) {
	client, err := NewSAMDGForwarderFromOptions(SetOutBackups(5))
	if err != nil {
		t.Fatalf("NewSAMDGForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestOptionUDPReduceIdleQuantity(t *testing.T) {
	client, err := NewSAMDGForwarderFromOptions(SetReduceIdleQuantity(4))
	if err != nil {
		t.Fatalf("NewSAMDGForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestOptionUDPEncryptLease(t *testing.T) {
	client, err := NewSAMDGForwarderFromOptions(SetEncrypt(true))
	if err != nil {
		t.Fatalf("NewSAMDGForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestOptionUDPSaveFile(t *testing.T) {
	client, err := NewSAMDGForwarderFromOptions(SetSaveFile(true))
	if err != nil {
		t.Fatalf("NewSAMDGForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestClientOptionHost(t *testing.T) {
	client, err := NewSAMDGClientForwarderFromOptions(SetClientHost("127.0.0.1"))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestClientOptionPort(t *testing.T) {
	client, err := NewSAMDGClientForwarderFromOptions(SetClientSAMPort("7656"))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestClientOptionInLength(t *testing.T) {
	client, err := NewSAMDGClientForwarderFromOptions(SetClientInLength(3))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestClientOptionOutLength(t *testing.T) {
	client, err := NewSAMDGClientForwarderFromOptions(SetClientInLength(3))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestClientOptionInVariance(t *testing.T) {
	client, err := NewSAMDGClientForwarderFromOptions(SetClientInVariance(1))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestClientOptionOutVariance(t *testing.T) {
	client, err := NewSAMDGClientForwarderFromOptions(SetClientOutVariance(1))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestClientOptionInQuantity(t *testing.T) {
	client, err := NewSAMDGClientForwarderFromOptions(SetClientInQuantity(6))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestClientOptionOutQuantity(t *testing.T) {
	client, err := NewSAMDGClientForwarderFromOptions(SetClientOutQuantity(6))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestClientOptionInBackups(t *testing.T) {
	client, err := NewSAMDGClientForwarderFromOptions(SetClientInBackups(5))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestClientOptionOutBackups(t *testing.T) {
	client, err := NewSAMDGClientForwarderFromOptions(SetClientOutBackups(5))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestClientOptionReduceIdleQuantity(t *testing.T) {
	client, err := NewSAMDGClientForwarderFromOptions(SetClientReduceIdleQuantity(4))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestClientOptionReduceIdleTimeMs(t *testing.T) {
	client, err := NewSAMDGClientForwarderFromOptions(SetClientReduceIdleTimeMs(300000))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestClientOptionReduceIdleTime(t *testing.T) {
	client, err := NewSAMDGClientForwarderFromOptions(SetClientReduceIdleTime(6))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestClientOptionCloseIdleTimeMs(t *testing.T) {
	client, err := NewSAMDGClientForwarderFromOptions(SetClientCloseIdleTimeMs(300000))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestClientOptionCloseIdleTime(t *testing.T) {
	client, err := NewSAMDGClientForwarderFromOptions(SetClientCloseIdleTime(6))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestClientOptionEncryptLease(t *testing.T) {
	client, err := NewSAMDGClientForwarderFromOptions(SetClientEncrypt(true))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestClientOptionSaveFile(t *testing.T) {
	client, err := NewSAMDGClientForwarderFromOptions(SetClientSaveFile(true))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}
