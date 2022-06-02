package samforwardertest

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	samtunnel "github.com/eyedeekay/sam-forwarder/interface"

	samoptions "github.com/eyedeekay/sam-forwarder/options"

	samforwarder "github.com/eyedeekay/sam-forwarder/tcp"

	samforwarderudp "github.com/eyedeekay/sam-forwarder/udp"
)

var (
	port               = "8100"
	cport              = "8101"
	UDPServerPort      = "8102"
	UDPClientPort      = "8103"
	SSUServerPort      = "8104"
	udpserveraddr      *net.UDPAddr
	udplocaladdr       *net.UDPAddr
	ssulocaladdr       *net.UDPAddr
	udpserverconn      *net.UDPConn
	directory          = "./www"
	err                error
	forwarder          samtunnel.SAMTunnel
	forwarderclient    samtunnel.SAMTunnel
	ssuforwarder       samtunnel.SAMTunnel
	ssuforwarderclient samtunnel.SAMTunnel
)

func serve() {
	flag.Parse()

	forwarder, err = samforwarder.NewSAMForwarderFromOptions(
		samoptions.SetHost("127.0.0.1"),
		samoptions.SetPort(port),
		samoptions.SetSAMHost("127.0.0.1"),
		samoptions.SetSAMPort("7656"),
		samoptions.SetName("testserver"),
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
		samoptions.SetHost("127.0.0.1"),
		samoptions.SetPort(cport),
		samoptions.SetSAMHost("127.0.0.1"),
		samoptions.SetSAMPort("7656"),
		samoptions.SetName("testclient"),
		samoptions.SetDestination(forwarder.Base32()),
	)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Printf("Connecting %s HTTP port: %s %s\n", cport, "to",
		forwarder.Base32())
	go forwarderclient.Serve()
}

func echo() {
	/* Lets prepare a address at any address at port UDPServerPort */
	udpserveraddr, err = net.ResolveUDPAddr("udp", "127.0.0.1:"+UDPServerPort)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("starting UDP echo server listening on :", UDPServerPort)

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

func echoclient() {
	udpclientaddr, e := net.Dial("udp", "127.0.0.1:"+UDPServerPort)
	if e != nil {
		log.Fatal(e)
	}
	udpclientaddr.Write([]byte("test"))
}

func serveudp() {
	ssuforwarder, err = samforwarderudp.NewSAMDGForwarderFromOptions(
		samoptions.SetHost("127.0.0.1"),
		samoptions.SetPort(UDPServerPort),
		samoptions.SetSAMHost("127.0.0.1"),
		samoptions.SetSAMPort("7656"),
		samoptions.SetName("testudpserver"),
	)
	if err != nil {
		log.Fatal(err.Error())
	}
	f, e := ssuforwarder.Load()
	if e != nil {
		log.Fatal(e.Error())
	}
	go f.Serve()
	log.Printf("Serving on UDP port: %s  and on %s\n", UDPServerPort, ssuforwarder.Base32())
}

func clientudp() {
	ssuforwarderclient, err = samforwarderudp.NewSAMDGClientForwarderFromOptions(
		samoptions.SetHost("127.0.0.1"),
		samoptions.SetPort(UDPClientPort),
		samoptions.SetSAMHost("127.0.0.1"),
		samoptions.SetSAMPort("7656"),
		samoptions.SetName("testudpclient"),
		samoptions.SetDestination(ssuforwarder.Base32()),
	)
	if err != nil {
		log.Fatal(err.Error())
	}

	f, e := ssuforwarderclient.Load()
	if e != nil {
		log.Fatal(e.Error())
	}
	go f.Serve()
	log.Printf("Connecting UDP port: %s to %s\n", SSUServerPort, ssuforwarder.Base32())
}

func setupudp() {
	ssulocaladdr, err = net.ResolveUDPAddr("udp", "127.0.0.1:"+SSUServerPort)
	if err != nil {
		log.Fatal(err)
	}

	udplocaladdr, err = net.ResolveUDPAddr("udp", "127.0.0.1:"+UDPClientPort)
	if err != nil {
		log.Fatal(err)
	}
}
