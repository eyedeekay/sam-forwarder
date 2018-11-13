package main

import (
	"log"
	"os"
	"os/signal"
)

import "github.com/eyedeekay/sam-forwarder/config"

func serveMode() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	if *udpMode {
		log.Println("Redirecting udp", *targetHost+":"+*targetPort, "to i2p")
		forwarder, err := i2ptunconf.NewSAMSSUForwarderFromConf(config)
		if err == nil {
			forwarder.Serve()
		} else {
			log.Println(err.Error())
		}
		go func() {
			for sig := range c {
				if sig == os.Interrupt {
					forwarder.Cleanup()
				}
			}
		}()
	} else {
		log.Println("Redirecting tcp", *targetHost+":"+*targetPort, "to i2p")
		forwarder, err := i2ptunconf.NewSAMForwarderFromConf(config)
		if err == nil {
			forwarder.Serve()
		} else {
			log.Println(err.Error())
		}
		go func() {
			for sig := range c {
				if sig == os.Interrupt {
					forwarder.Cleanup()
				}
			}
		}()
	}
}
