package samforwardertest

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
)

import "github.com/eyedeekay/sam-forwarder"
import "github.com/eyedeekay/sam-forwarder/udp"

var (
	port               = "8100"
	cport              = "8101"
	uport              = "8102"
	ucport             = "8103"
	udpserveraddr      *net.UDPAddr
	udplocaladdr       *net.UDPAddr
	directory          = "./www"
	err                error
	forwarder          *samforwarder.SAMForwarder
	forwarderclient    *samforwarder.SAMClientForwarder
	ssuforwarder       *samforwarderudp.SAMSSUForwarder
	ssuforwarderclient *samforwarderudp.SAMSSUClientForwarder
)

func serve() {
	flag.Parse()

	forwarder, err = samforwarder.NewSAMForwarderFromOptions(
		samforwarder.SetHost("127.0.0.1"),
		samforwarder.SetPort(port),
		samforwarder.SetSAMHost("127.0.0.1"),
		samforwarder.SetSAMPort("7656"),
		samforwarder.SetName("testserver"),
	)
	if err != nil {
		log.Fatal(err.Error())
	}
	go forwarder.Serve()

	http.Handle("/", http.FileServer(http.Dir(directory)))

	log.Printf("Serving %s on HTTP port: %s %s %s\n", directory, port, "and on",
		forwarder.Base32())
	log.Fatal(http.ListenAndServe("127.0.0.1:"+port, nil))
}

func client() {
	flag.Parse()

	forwarderclient, err = samforwarder.NewSAMClientForwarderFromOptions(
		samforwarder.SetClientHost("127.0.0.1"),
		samforwarder.SetClientPort(cport),
		samforwarder.SetClientSAMHost("127.0.0.1"),
		samforwarder.SetClientSAMPort("7656"),
		samforwarder.SetClientName("testclient"),
		samforwarder.SetClientDestination(forwarder.Base32()),
	)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Printf("Connecting %s HTTP port: %s %s\n", cport, "to",
		forwarder.Base32())
	go forwarderclient.Serve()
}

func echo() {
	/* Lets prepare a address at any address at port 10001*/
	udpserveraddr, err = net.ResolveUDPAddr("udp", "127.0.0.1:"+uport)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("listening on :", uport)
	udplocaladdr, err = net.ResolveUDPAddr("udp", "127.0.0.1:"+ucport)
	if err != nil {
		log.Fatal(err)
	}
	/* Now listen at selected port */
	ServerConn, err := net.ListenUDP("udp", udpserveraddr)
	if err != nil {
		log.Fatal(err)
	}
	defer ServerConn.Close()

	buf := make([]byte, 1024)

	for {
		n, addr, err := ServerConn.ReadFromUDP(buf)
		fmt.Printf("received: %s from: %s\n", string(buf[0:n]), addr)

		if err != nil {
			fmt.Println("error: ", err)
		}

		ServerConn.WriteTo(buf[0:n], addr)
	}
}

func serveudp() {
	flag.Parse()

	ssuforwarder, err = samforwarderudp.NewSAMSSUForwarderFromOptions(
		samforwarderudp.SetHost("127.0.0.1"),
		samforwarderudp.SetPort(uport),
		samforwarderudp.SetSAMHost("127.0.0.1"),
		samforwarderudp.SetSAMPort("7656"),
		samforwarderudp.SetName("testudpserver"),
	)
	if err != nil {
		log.Fatal(err.Error())
	}
	go forwarder.Serve()

	log.Printf("Serving %s on UDP port: %s %s\n", uport, "and on",
		forwarder.Base32())
	log.Fatal(http.ListenAndServe("127.0.0.1:"+uport, nil))
}

func clientudp() {
	flag.Parse()

	ssuforwarderclient, err = samforwarderudp.NewSAMSSUClientForwarderFromOptions(
		samforwarderudp.SetClientHost("127.0.0.1"),
		samforwarderudp.SetClientPort(ucport),
		samforwarderudp.SetClientSAMHost("127.0.0.1"),
		samforwarderudp.SetClientSAMPort("7656"),
		samforwarderudp.SetClientName("testudpclient"),
		samforwarderudp.SetClientDestination(ssuforwarder.Base32()),
	)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Printf("Connecting %s UDP port: %s %s\n", ucport, "to",
		forwarder.Base32())
	go forwarderclient.Serve()
}
