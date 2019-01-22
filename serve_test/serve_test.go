package samforwardertest

import (
	"log"
	"net"
	"net/http"
	"testing"
	"time"
)

var longcount = 60

func countdown(i int) {
	for i > 0 {
		time.Sleep(1 * time.Second)
		i--
		log.Println("waiting", i, "more seconds.")
	}
}

func TestTCP(t *testing.T) {
	go serve()
	countdown(longcount)
	go client()
	countdown(longcount)
	resp, err := http.Get("http://127.0.0.1:" + cport + "/test.html")
	log.Println("requesting http://127.0.0.1:" + cport + "/test.html")
	if err != nil {
		t.Fatal(err)
	}
	log.Println(resp)
}

func TestUDP(t *testing.T) {
	go echo()
	countdown(3)
	go serveudp()
	countdown(longcount)
	go clientudp()
	countdown(longcount)
	setupudp()

	conn, err := net.DialUDP("udp", udplocaladdr, ssulocaladdr)
	if err != nil {
		t.Fatal(err)
	}
	message := []byte("Hello SSU")
	_, err = conn.Write(message)
	if err != nil {
		t.Fatal("SSU error", err)
	}
	log.Println(string(message))
}

/*
func TestUDPeasy(t *testing.T) {
	go echo()
	time.Sleep(time.Duration(1 * time.Second))
    conn, err := net.DialUDP("udp", udplocaladdr, udpserveraddr)
	//defer conn.Close()
	if err != nil {
		t.Fatal(err)
	}
	_, err = conn.Write([]byte("Hello SSU"))
	if err != nil {
		t.Fatal("SSU error", err)
	}
}
*/
