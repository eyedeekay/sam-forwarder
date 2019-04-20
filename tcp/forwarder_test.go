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
	client, err := NewSAMForwarderFromOptions(SetPort("8080"))
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

func TestClientOptionHost(t *testing.T) {
	client, err := NewSAMClientForwarderFromOptions(SetClientHost("127.0.0.1"))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestClientOptionPort(t *testing.T) {
	client, err := NewSAMClientForwarderFromOptions(SetClientSAMPort("7656"))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestClientOptionInLength(t *testing.T) {
	client, err := NewSAMClientForwarderFromOptions(SetClientInLength(3))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestClientOptionOutLength(t *testing.T) {
	client, err := NewSAMClientForwarderFromOptions(SetClientInLength(3))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestClientOptionInVariance(t *testing.T) {
	client, err := NewSAMClientForwarderFromOptions(SetClientInVariance(1))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestClientOptionOutVariance(t *testing.T) {
	client, err := NewSAMClientForwarderFromOptions(SetClientOutVariance(1))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestClientOptionInQuantity(t *testing.T) {
	client, err := NewSAMClientForwarderFromOptions(SetClientInQuantity(6))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestClientOptionOutQuantity(t *testing.T) {
	client, err := NewSAMClientForwarderFromOptions(SetClientOutQuantity(6))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestClientOptionInBackups(t *testing.T) {
	client, err := NewSAMClientForwarderFromOptions(SetClientInBackups(5))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestClientOptionOutBackups(t *testing.T) {
	client, err := NewSAMClientForwarderFromOptions(SetClientOutBackups(5))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestClientOptionReduceIdleQuantity(t *testing.T) {
	client, err := NewSAMClientForwarderFromOptions(SetClientReduceIdleQuantity(4))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestClientOptionReduceIdleTimeMs(t *testing.T) {
	client, err := NewSAMClientForwarderFromOptions(SetClientReduceIdleTimeMs(300000))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestClientOptionReduceIdleTime(t *testing.T) {
	client, err := NewSAMClientForwarderFromOptions(SetClientReduceIdleTime(6))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestClientOptionCloseIdleTimeMs(t *testing.T) {
	client, err := NewSAMClientForwarderFromOptions(SetClientCloseIdleTimeMs(300000))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestClientOptionCloseIdleTime(t *testing.T) {
	client, err := NewSAMClientForwarderFromOptions(SetClientCloseIdleTime(6))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestClientOptionEncryptLease(t *testing.T) {
	client, err := NewSAMClientForwarderFromOptions(SetClientEncrypt(true))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

func TestClientOptionSaveFile(t *testing.T) {
	client, err := NewSAMClientForwarderFromOptions(SetClientSaveFile(true))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}

/*func TestOptionTargetForPort443(t *testing.T) {
	client, err := NewSAMForwarderFromOptions(SetTargetForPort443("443"))
	if err != nil {
		t.Fatalf("NewSAMForwarder() Error: %q\n", err)
	}
	log.Println(client.Base32())
}*/
