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

func TestClientOptionHost(t *testing.T) {
	client, err := NewSAMSSUClientForwarderFromOptions(SetClientHost("127.0.0.1"))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestClientOptionPort(t *testing.T) {
	client, err := NewSAMSSUClientForwarderFromOptions(SetClientSAMPort("7656"))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestClientOptionInLength(t *testing.T) {
	client, err := NewSAMSSUClientForwarderFromOptions(SetClientInLength(3))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestClientOptionOutLength(t *testing.T) {
	client, err := NewSAMSSUClientForwarderFromOptions(SetClientInLength(3))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestClientOptionInVariance(t *testing.T) {
	client, err := NewSAMSSUClientForwarderFromOptions(SetClientInVariance(1))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestClientOptionOutVariance(t *testing.T) {
	client, err := NewSAMSSUClientForwarderFromOptions(SetClientOutVariance(1))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestClientOptionInQuantity(t *testing.T) {
	client, err := NewSAMSSUClientForwarderFromOptions(SetClientInQuantity(6))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestClientOptionOutQuantity(t *testing.T) {
	client, err := NewSAMSSUClientForwarderFromOptions(SetClientOutQuantity(6))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestClientOptionInBackups(t *testing.T) {
	client, err := NewSAMSSUClientForwarderFromOptions(SetClientInBackups(5))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestClientOptionOutBackups(t *testing.T) {
	client, err := NewSAMSSUClientForwarderFromOptions(SetClientOutBackups(5))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestClientOptionReduceIdleQuantity(t *testing.T) {
	client, err := NewSAMSSUClientForwarderFromOptions(SetClientReduceIdleQuantity(4))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestClientOptionReduceIdleTimeMs(t *testing.T) {
	client, err := NewSAMSSUClientForwarderFromOptions(SetClientReduceIdleTimeMs(300000))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestClientOptionReduceIdleTime(t *testing.T) {
	client, err := NewSAMSSUClientForwarderFromOptions(SetClientReduceIdleTime(6))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestClientOptionCloseIdleTimeMs(t *testing.T) {
	client, err := NewSAMSSUClientForwarderFromOptions(SetClientCloseIdleTimeMs(300000))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestClientOptionCloseIdleTime(t *testing.T) {
	client, err := NewSAMSSUClientForwarderFromOptions(SetClientCloseIdleTime(6))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestClientOptionEncryptLease(t *testing.T) {
	client, err := NewSAMSSUClientForwarderFromOptions(SetClientEncrypt(true))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestClientOptionSaveFile(t *testing.T) {
	client, err := NewSAMSSUClientForwarderFromOptions(SetClientSaveFile(true))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}
