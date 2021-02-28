package samforwarderudp

import (
	"log"
	"testing"

	"github.com/eyedeekay/sam-forwarder/options"
)

func TestOptionUDPHost(t *testing.T) {
	client, err := NewSAMDGForwarderFromOptions(samoptions.SetHost("127.0.0.1"))
	if err != nil {
		t.Fatalf("NewSAMDGForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestOptionUDPPort(t *testing.T) {
	client, err := NewSAMDGForwarderFromOptions(samoptions.SetPort("7656"))
	if err != nil {
		t.Fatalf("NewSAMDGForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestOptionUDPInLength(t *testing.T) {
	client, err := NewSAMDGForwarderFromOptions(samoptions.SetInLength(3))
	if err != nil {
		t.Fatalf("NewSAMDGForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestOptionUDPOutLength(t *testing.T) {
	client, err := NewSAMDGForwarderFromOptions(samoptions.SetInLength(3))
	if err != nil {
		t.Fatalf("NewSAMDGForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestOptionUDPInVariance(t *testing.T) {
	client, err := NewSAMDGForwarderFromOptions(samoptions.SetInVariance(1))
	if err != nil {
		t.Fatalf("NewSAMDGForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestOptionUDPOutVariance(t *testing.T) {
	client, err := NewSAMDGForwarderFromOptions(samoptions.SetOutVariance(1))
	if err != nil {
		t.Fatalf("NewSAMDGForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestOptionUDPInQuantity(t *testing.T) {
	client, err := NewSAMDGForwarderFromOptions(samoptions.SetInQuantity(6))
	if err != nil {
		t.Fatalf("NewSAMDGForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestOptionUDPOutQuantity(t *testing.T) {
	client, err := NewSAMDGForwarderFromOptions(samoptions.SetOutQuantity(6))
	if err != nil {
		t.Fatalf("NewSAMDGForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestOptionUDPInBackups(t *testing.T) {
	client, err := NewSAMDGForwarderFromOptions(samoptions.SetInBackups(5))
	if err != nil {
		t.Fatalf("NewSAMDGForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestOptionUDPOutBackups(t *testing.T) {
	client, err := NewSAMDGForwarderFromOptions(samoptions.SetOutBackups(5))
	if err != nil {
		t.Fatalf("NewSAMDGForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestOptionUDPReduceIdleQuantity(t *testing.T) {
	client, err := NewSAMDGForwarderFromOptions(samoptions.SetReduceIdleQuantity(4))
	if err != nil {
		t.Fatalf("NewSAMDGForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestOptionUDPEncryptLease(t *testing.T) {
	client, err := NewSAMDGForwarderFromOptions(samoptions.SetEncrypt(true))
	if err != nil {
		t.Fatalf("NewSAMDGForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestOptionUDPSaveFile(t *testing.T) {
	client, err := NewSAMDGForwarderFromOptions(samoptions.SetSaveFile(true))
	if err != nil {
		t.Fatalf("NewSAMDGForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestClientOptionHost(t *testing.T) {
	client, err := NewSAMDGClientForwarderFromOptions(samoptions.SetHost("127.0.0.1"))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestClientOptionPort(t *testing.T) {
	client, err := NewSAMDGClientForwarderFromOptions(samoptions.SetSAMPort("7656"))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestClientOptionInLength(t *testing.T) {
	client, err := NewSAMDGClientForwarderFromOptions(samoptions.SetInLength(3))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestClientOptionOutLength(t *testing.T) {
	client, err := NewSAMDGClientForwarderFromOptions(samoptions.SetInLength(3))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestClientOptionInVariance(t *testing.T) {
	client, err := NewSAMDGClientForwarderFromOptions(samoptions.SetInVariance(1))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestClientOptionOutVariance(t *testing.T) {
	client, err := NewSAMDGClientForwarderFromOptions(samoptions.SetOutVariance(1))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestClientOptionInQuantity(t *testing.T) {
	client, err := NewSAMDGClientForwarderFromOptions(samoptions.SetInQuantity(6))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestClientOptionOutQuantity(t *testing.T) {
	client, err := NewSAMDGClientForwarderFromOptions(samoptions.SetOutQuantity(6))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestClientOptionInBackups(t *testing.T) {
	client, err := NewSAMDGClientForwarderFromOptions(samoptions.SetInBackups(5))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestClientOptionOutBackups(t *testing.T) {
	client, err := NewSAMDGClientForwarderFromOptions(samoptions.SetOutBackups(5))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestClientOptionReduceIdleQuantity(t *testing.T) {
	client, err := NewSAMDGClientForwarderFromOptions(samoptions.SetReduceIdleQuantity(4))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestClientOptionReduceIdleTimeMs(t *testing.T) {
	client, err := NewSAMDGClientForwarderFromOptions(samoptions.SetReduceIdleTimeMs(300000))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestClientOptionReduceIdleTime(t *testing.T) {
	client, err := NewSAMDGClientForwarderFromOptions(samoptions.SetReduceIdleTime(6))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestClientOptionCloseIdleTimeMs(t *testing.T) {
	client, err := NewSAMDGClientForwarderFromOptions(samoptions.SetCloseIdleTimeMs(300000))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestClientOptionCloseIdleTime(t *testing.T) {
	client, err := NewSAMDGClientForwarderFromOptions(samoptions.SetCloseIdleTime(6))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestClientOptionEncryptLease(t *testing.T) {
	client, err := NewSAMDGClientForwarderFromOptions(samoptions.SetEncrypt(true))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestClientOptionSaveFile(t *testing.T) {
	client, err := NewSAMDGClientForwarderFromOptions(samoptions.SetSaveFile(true))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}
