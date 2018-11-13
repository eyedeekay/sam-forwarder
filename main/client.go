package main

import (
	"log"
	"os"
	"os/signal"
)

import "github.com/eyedeekay/sam-forwarder/config"

func clientMode() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	if *udpMode {
		log.Println("Proxying udp", *targetHost+":"+*targetPort, "to", *targetDestination)
		forwarder, err := i2ptunconf.NewSAMSSUClientForwarderFromConf(config)
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
		log.Println("Proxying tcp", *targetHost+":"+*targetPort, "to", *targetDestination)
		forwarder, err := i2ptunconf.NewSAMClientForwarderFromConf(config)
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
