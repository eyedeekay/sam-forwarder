package main

import (
	"flag"
	"log"
)

import "github.com/eyedeekay/sam-forwarder"

func main() {
	TargetHost := *flag.String("host", "127.0.0.1", "Target host")
    TargetPort := *flag.String("port", "8081", "Target port")
    SamHost := *flag.String("samhost", "127.0.0.1", "SAM host")
    SamPort := *flag.String("samport", "7656", "SAM port")
    TunName := *flag.String("name", "forwarder", "Tunnel name")
	flag.Parse()
	log.Println("Redirecting", TargetHost+":"+TargetPort, "to i2p")
    forwarder := samforwarder.NewSAMForwarderFromOptions(
        samforwarder.SetHost(TargetHost),
        samforwarder.SetPort(TargetPort),
        samforwarder.SetSAMHost(SamHost),
        samforwarder.SetSAMPort(SamPort),
        samforwarder.SetName(TunName),
    )
	forwarder.Serve()
}
