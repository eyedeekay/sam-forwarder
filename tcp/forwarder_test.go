package samforwarder

import (
	"log"
	"testing"

	"github.com/eyedeekay/sam-forwarder/options"
)

func TestOptionHost(t *testing.T) {
	client, err := NewSAMForwarderFromOptions(samoptions.SetHost("127.0.0.1"))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestOptionPort(t *testing.T) {
	client, err := NewSAMForwarderFromOptions(samoptions.SetPort("8080"))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestOptionInLength(t *testing.T) {
	client, err := NewSAMForwarderFromOptions(samoptions.SetInLength(3))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestOptionOutLength(t *testing.T) {
	client, err := NewSAMForwarderFromOptions(samoptions.SetInLength(3))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestOptionInVariance(t *testing.T) {
	client, err := NewSAMForwarderFromOptions(samoptions.SetInVariance(1))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestOptionOutVariance(t *testing.T) {
	client, err := NewSAMForwarderFromOptions(samoptions.SetOutVariance(1))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestOptionInQuantity(t *testing.T) {
	client, err := NewSAMForwarderFromOptions(samoptions.SetInQuantity(6))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestOptionOutQuantity(t *testing.T) {
	client, err := NewSAMForwarderFromOptions(samoptions.SetOutQuantity(6))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestOptionInBackups(t *testing.T) {
	client, err := NewSAMForwarderFromOptions(samoptions.SetInBackups(5))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestOptionOutBackups(t *testing.T) {
	client, err := NewSAMForwarderFromOptions(samoptions.SetOutBackups(5))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestOptionReduceIdleQuantity(t *testing.T) {
	client, err := NewSAMForwarderFromOptions(samoptions.SetReduceIdleQuantity(4))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestOptionReduceIdleTimeMs(t *testing.T) {
	client, err := NewSAMForwarderFromOptions(samoptions.SetReduceIdleTimeMs(300000))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestOptionReduceIdleTime(t *testing.T) {
	client, err := NewSAMForwarderFromOptions(samoptions.SetReduceIdleTime(6))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestOptionCloseIdleTimeMs(t *testing.T) {
	client, err := NewSAMForwarderFromOptions(samoptions.SetCloseIdleTimeMs(300000))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestOptionCloseIdleTime(t *testing.T) {
	client, err := NewSAMForwarderFromOptions(samoptions.SetCloseIdleTime(6))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestOptionEncryptLease(t *testing.T) {
	client, err := NewSAMForwarderFromOptions(samoptions.SetEncrypt(true))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestOptionSaveFile(t *testing.T) {
	client, err := NewSAMForwarderFromOptions(samoptions.SetSaveFile(true))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestClientOptionHost(t *testing.T) {
	client, err := NewSAMClientForwarderFromOptions(samoptions.SetHost("127.0.0.1"))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestClientOptionPort(t *testing.T) {
	client, err := NewSAMClientForwarderFromOptions(samoptions.SetSAMPort("7656"))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestClientOptionInLength(t *testing.T) {
	client, err := NewSAMClientForwarderFromOptions(samoptions.SetInLength(3))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestClientOptionOutLength(t *testing.T) {
	client, err := NewSAMClientForwarderFromOptions(samoptions.SetInLength(3))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestClientOptionInVariance(t *testing.T) {
	client, err := NewSAMClientForwarderFromOptions(samoptions.SetInVariance(1))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestClientOptionOutVariance(t *testing.T) {
	client, err := NewSAMClientForwarderFromOptions(samoptions.SetOutVariance(1))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestClientOptionInQuantity(t *testing.T) {
	client, err := NewSAMClientForwarderFromOptions(samoptions.SetInQuantity(6))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestClientOptionOutQuantity(t *testing.T) {
	client, err := NewSAMClientForwarderFromOptions(samoptions.SetOutQuantity(6))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestClientOptionInBackups(t *testing.T) {
	client, err := NewSAMClientForwarderFromOptions(samoptions.SetInBackups(5))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestClientOptionOutBackups(t *testing.T) {
	client, err := NewSAMClientForwarderFromOptions(samoptions.SetOutBackups(5))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestClientOptionReduceIdleQuantity(t *testing.T) {
	client, err := NewSAMClientForwarderFromOptions(samoptions.SetReduceIdleQuantity(4))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestClientOptionReduceIdleTimeMs(t *testing.T) {
	client, err := NewSAMClientForwarderFromOptions(samoptions.SetReduceIdleTimeMs(300000))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestClientOptionReduceIdleTime(t *testing.T) {
	client, err := NewSAMClientForwarderFromOptions(samoptions.SetReduceIdleTime(6))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestClientOptionCloseIdleTimeMs(t *testing.T) {
	client, err := NewSAMClientForwarderFromOptions(samoptions.SetCloseIdleTimeMs(300000))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestClientOptionCloseIdleTime(t *testing.T) {
	client, err := NewSAMClientForwarderFromOptions(samoptions.SetCloseIdleTime(6))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestClientOptionEncryptLease(t *testing.T) {
	client, err := NewSAMClientForwarderFromOptions(samoptions.SetEncrypt(true))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestClientOptionSaveFile(t *testing.T) {
	client, err := NewSAMClientForwarderFromOptions(samoptions.SetSaveFile(true))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

/*func TestOptionTargetForPort443(t *testing.T) {
	client, err := NewSAMForwarderFromOptions(samoptions.SetTargetForPort443("443"))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}*/
