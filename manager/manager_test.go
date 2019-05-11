package sammanager

import (
	"fmt"
	"log"
	"strings"
	"testing"
)

func stringify(s []string) string {
	var p string
	for _, x := range s {
		if x != "ntcpserver" && x != "httpserver" && x != "ssuserver" && x != "ntcpclient" && x != "ssuclient" {
			p += x + ","
		}
	}
	r := strings.Trim(strings.Trim(strings.Replace(p, ",,", ",", -1), " "), "\n")
	return r
}

func (s *SAMManager) List(search ...string) *[]string {
	var r []string
	if search == nil {
		for index, element := range s.handlerMux.Tunnels() {
			r = append(r, fmt.Sprintf("  %v. %s", index, element.Print()))
		}
		return &r
	} else if len(search) > 0 {
		switch search[0] {
		case "":
			for index, element := range s.handlerMux.Tunnels() {
				r = append(r, fmt.Sprintf("  %v. %s", index, element.Print()))
			}
			return &r
		case "ntcpserver":
			for index, element := range s.handlerMux.Tunnels() {
				r = append(r, fmt.Sprintf("  %v. %s", index, element.Search(stringify(search))))
			}
			return &r
		case "httpserver":
			for index, element := range s.handlerMux.Tunnels() {
				if element.GetType() == "http" {
					r = append(r, fmt.Sprintf("  %v. %s", index, element.Search(stringify(search))))
				}
			}
			return &r
		case "ntcpclient":
			for index, element := range s.handlerMux.Tunnels() {
				r = append(r, fmt.Sprintf("  %v. %s", index, element.Search(stringify(search))))
			}
			return &r
		case "ssuserver":
			for index, element := range s.handlerMux.Tunnels() {
				r = append(r, fmt.Sprintf("  %v. %s", index, element.Search(stringify(search))))
			}
			return &r
		case "ssuclient":
			for index, element := range s.handlerMux.Tunnels() {
				r = append(r, fmt.Sprintf("  %v. %s", index, element.Search(stringify(search))))
			}
			return &r
		default:
			for index, element := range s.handlerMux.Tunnels() {
				if element.Search(stringify(search)) != "" {
					r = append(r, fmt.Sprintf("  %v. %s", index, element.Search(stringify(search))))
				}
			}
			for index, element := range s.handlerMux.Tunnels() {
				if element.Search(stringify(search)) != "" {
					r = append(r, fmt.Sprintf("  %v. %s", index, element.Search(stringify(search))))
				}
			}
			for index, element := range s.handlerMux.Tunnels() {
				if element.Search(stringify(search)) != "" {
					r = append(r, fmt.Sprintf("  %v. %s", index, element.Search(stringify(search))))
				}
			}
			for index, element := range s.handlerMux.Tunnels() {
				if element.Search(stringify(search)) != "" {
					r = append(r, fmt.Sprintf("  %v. %s", index, element.Search(stringify(search))))
				}
			}
			return &r
		}
	}
	return &r
}

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
