package sammanager

import (
	"log"
	"testing"
)

func TestOption0(t *testing.T) {
	client, err := NewSAMManagerFromOptions(
		SetManagerHost("127.0.0.1"),
		SetManagerSAMHost("127.0.0.1"),
		SetManagerPort("8080"),
		SetManagerSAMPort("7656"),
		SetManagerWebHost("127.0.0.1"),
		SetManagerWebPort("7958"),
		SetManagerFilePath("../etc/sam-forwarder/tunnels.ini"),
	)
	if err != nil {
		t.Fatalf("NewSAMManager() Error: %q\n", err)
	}
	log.Println(client.List())
}

func TestOption1(t *testing.T) {
	client, err := NewSAMManagerFromOptions(
		SetManagerHost("127.0.0.1"),
		SetManagerSAMHost("127.0.0.1"),
		SetManagerPort("8081"),
		SetManagerSAMPort("7656"),
		SetManagerWebHost("127.0.0.1"),
		SetManagerWebPort("7959"),
		SetManagerFilePath("../etc/sam-forwarder/tunnels.ini"),
	)
	if err != nil {
		t.Fatalf("NewSAMManager() Error: %q\n", err)
	}
	log.Println(client.List(""))
}

func TestOption2(t *testing.T) {
	client, err := NewSAMManagerFromOptions(
		SetManagerHost("127.0.0.1"),
		SetManagerSAMHost("127.0.0.1"),
		SetManagerPort("8082"),
		SetManagerSAMPort("7656"),
		SetManagerWebHost("127.0.0.1"),
		SetManagerWebPort("7960"),
		SetManagerFilePath("../etc/sam-forwarder/tunnels.ini"),
	)
	if err != nil {
		t.Fatalf("NewSAMManager() Error: %q\n", err)
	}
	log.Println(client.List("asdgrepgbutwhrsgfbxv"))
}

func TestOption3(t *testing.T) {
	client, err := NewSAMManagerFromOptions(
		SetManagerHost("127.0.0.1"),
		SetManagerSAMHost("127.0.0.1"),
		SetManagerPort("8083"),
		SetManagerSAMPort("7656"),
		SetManagerWebHost("127.0.0.1"),
		SetManagerWebPort("7961"),
		SetManagerFilePath("none"),
		SetManagerFilePath("../etc/sam-forwarder/tunnels.ini"),
	)
	if err != nil {
		t.Fatalf("NewSAMManager() Error: %q\n", err)
	}
	log.Println(client.List("server"))
}

func TestOption4(t *testing.T) {
	client, err := NewSAMManagerFromOptions(
		SetManagerHost("127.0.0.1"),
		SetManagerSAMHost("127.0.0.1"),
		SetManagerPort("8083"),
		SetManagerSAMPort("7656"),
		SetManagerWebHost("127.0.0.1"),
		SetManagerWebPort("7961"),
		SetManagerFilePath("none"),
		SetManagerFilePath("../etc/sam-forwarder/tunnels.ini"),
	)
	if err != nil {
		t.Fatalf("NewSAMManager() Error: %q\n", err)
	}
	log.Println(client.List(""))
}

func TestOption5(t *testing.T) {
	client, err := NewSAMManagerFromOptions(
		SetManagerHost("127.0.0.1"),
		SetManagerSAMHost("127.0.0.1"),
		SetManagerPort("8083"),
		SetManagerSAMPort("7656"),
		SetManagerWebHost("127.0.0.1"),
		SetManagerWebPort("7961"),
		SetManagerFilePath("none"),
		SetManagerFilePath("../etc/samcatd/tunnels.ini"),
	)
	if err != nil {
		t.Fatalf("NewSAMManager() Error: %q\n", err)
	}
	log.Println(client.List(""))
}
