package samforwardertest

import (
	"log"
	"net"
	"net/http"
	"testing"
	"time"
)

func TestTCP(t *testing.T) {
	go serve()
	time.Sleep(time.Duration(30 * time.Second))
	go client()
	time.Sleep(time.Duration(30 * time.Second))
	resp, err := http.Get("http://127.0.0.1:" + cport + "/test.html")
	if err != nil {
		t.Fatal(err)
	}
	log.Println(resp)
}

func TestUDP(t *testing.T) {
	go echo()
	time.Sleep(time.Duration(30 * time.Second))
	defer udpserverconn.Close()
	go serveudp()
	time.Sleep(time.Duration(30 * time.Second))
	go clientudp()
	time.Sleep(time.Duration(30 * time.Second))
//	conn, err := net.DialUDP("udp", udpserveraddr, udplocaladdr)
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
