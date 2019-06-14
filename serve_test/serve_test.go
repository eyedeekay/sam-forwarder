package samforwardertest

import (
	"log"
	"net"
	//"net/http"
	"testing"
	"time"
)

func countdown(i int) {
	for i > 0 {
		time.Sleep(1 * time.Second)
		i--
		if i%10 == 0 {
			log.Println("Waiting", i, "more seconds.")
		}
	}
}

/*
func TestTCP(t *testing.T) {
	go serve()
	countdown(61)
	go client()
	countdown(61)
	resp, err := http.Get("http://127.0.0.1:" + cport + "/test.html")
	log.Println("requesting http://127.0.0.1:" + cport + "/test.html")
	if err != nil {
		t.Fatal(err)
	}
	log.Println(resp)
}
*/
func TestUDP(t *testing.T) {
	go echo()
	countdown(11)
	echoclient()
	countdown(11)
	go serveudp()
	countdown(61)
	go clientudp()
	countdown(61)
	setupudp()

	conn, err := net.DialUDP("udp", udplocaladdr, ssulocaladdr)
	if err != nil {
		t.Fatal(err)
	}
	message := []byte("Hello UDP")
	_, err = conn.Write(message)
	if err != nil {
		t.Fatal("UDP error", err)
	}
	log.Println(string(message))
}
