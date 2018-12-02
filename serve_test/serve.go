package samforwardertest

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"
)

import ".."
import "../udp"

var (
	port               = "8100"
	cport              = "8101"
	uport              = "8102"
	ucport             = "8103"
	ssuport            = "8104"
	udpserveraddr      *net.UDPAddr
	udplocaladdr       *net.UDPAddr
	ssulocaladdr       *net.UDPAddr
	udpserverconn      *net.UDPConn
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

	/* Now listen at selected port */
	udpserverconn, err = net.ListenUDP("udp", udpserveraddr)
	if err != nil {
		log.Fatal(err)
	}

	buf := make([]byte, 1024)

	for {
		n, addr, err := udpserverconn.ReadFromUDP(buf)
		log.Println("received:", string(buf[0:n]), "from: ", addr)
		if err != nil {
			fmt.Println("error: ", err)
		}

		udpserverconn.WriteTo(buf[0:n], addr)
		time.Sleep(time.Duration(1 * time.Second))
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
	go ssuforwarder.Serve()
	log.Printf("Serving on UDP port: %s  and on %s\n", uport, ssuforwarder.Base32())
}

func clientudp() {
	flag.Parse()

	ssuforwarderclient, err = samforwarderudp.NewSAMSSUClientForwarderFromOptions(
		samforwarderudp.SetClientHost("127.0.0.1"),
		samforwarderudp.SetClientPort(ssuport),
		samforwarderudp.SetClientSAMHost("127.0.0.1"),
		samforwarderudp.SetClientSAMPort("7656"),
		samforwarderudp.SetClientName("testudpclient"),
		samforwarderudp.SetClientDestination(ssuforwarder.Base32()),
	)
	if err != nil {
		log.Fatal(err.Error())
	}
	go ssuforwarderclient.Serve()
	log.Printf("Connecting UDP port: %s to %s\n", ssuport, ssuforwarder.Base32())
}

func setupudp() {
	ssulocaladdr, err = net.ResolveUDPAddr("udp", "127.0.0.1:"+ssuport)
	if err != nil {
		log.Fatal(err)
	}

	udplocaladdr, err = net.ResolveUDPAddr("udp", "127.0.0.1:"+ucport)
	if err != nil {
		log.Fatal(err)
	}
}
