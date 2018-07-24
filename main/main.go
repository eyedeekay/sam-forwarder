package main

import (
	"flag"
	"log"
)

import "github.com/eyedeekay/ephemeral-eepSite-SAM"

func main() {
	samlistener.Target = *flag.String("addr", "127.0.0.1:8081", "Target host:port")
	flag.Parse()
	log.Println("Redirecting", samlistener.Target, "to i2p")
	samlistener.Serve()
}
