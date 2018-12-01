package samforwardertest

import (
	"flag"
	"log"
	"net/http"
)

import "github.com/eyedeekay/sam-forwarder"
import "github.com/eyedeekay/sam-forwarder/udp"

var (
	port               = "8100"
	cport              = "8101"
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
		forwarder.Base32()+".b32.i2p")
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
	go forwarderclient.Serve()
}

func serveudp() {
	flag.Parse()

	ssuforwarder, err = samforwarderudp.NewSAMSSUForwarderFromOptions(
		samforwarderudp.SetHost("127.0.0.1"),
		samforwarderudp.SetPort(port),
		samforwarderudp.SetSAMHost("127.0.0.1"),
		samforwarderudp.SetSAMPort("7656"),
		samforwarderudp.SetName("testserver"),
	)
	if err != nil {
		log.Fatal(err.Error())
	}
	go forwarder.Serve()

	http.Handle("/", http.FileServer(http.Dir(directory)))

	log.Printf("Serving %s on HTTP port: %s %s %s\n", directory, port, "and on",
		forwarder.Base32()+".b32.i2p")
	log.Fatal(http.ListenAndServe("127.0.0.1:"+port, nil))
}

func clientudp() {
	flag.Parse()

	ssuforwarderclient, err = samforwarderudp.NewSAMSSUClientForwarderFromOptions(
		samforwarderudp.SetClientHost("127.0.0.1"),
		samforwarderudp.SetClientPort(cport),
		samforwarderudp.SetClientSAMHost("127.0.0.1"),
		samforwarderudp.SetClientSAMPort("7656"),
		samforwarderudp.SetClientName("testclient"),
		samforwarderudp.SetClientDestination(forwarder.Base32()),
	)
	if err != nil {
		log.Fatal(err.Error())
	}
	go forwarderclient.Serve()
}
