package main

import (
	"flag"
	"log"
	"net/http"
)

import "github.com/eyedeekay/sam-forwarder"

func main() {
	port := flag.String("p", "8100", "port to serve on")
	directory := flag.String("d", ".", "the directory of static file to host")
	flag.Parse()

	forwarder, err := samforwarder.NewSAMForwarderFromOptions(
		samforwarder.SetHost("127.0.0.1"),
		samforwarder.SetPort(*port),
		samforwarder.SetSAMHost("127.0.0.1"),
		samforwarder.SetSAMPort("7656"),
		samforwarder.SetName("staticfiles"),
	)
	if err != nil {
		log.Fatal(err.Error())
	}
	go forwarder.Serve()

	http.Handle("/", http.FileServer(http.Dir(*directory)))

	log.Printf("Serving %s on HTTP port: %s\n", *directory, *port, "and on",
		forwarder.Base32()+".b32.i2p")
	log.Fatal(http.ListenAndServe("127.0.0.1:"+*port, nil))
}
