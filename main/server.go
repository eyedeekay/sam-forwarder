package main

import "log"

import (
	"github.com/eyedeekay/sam-forwarder"
	"github.com/eyedeekay/sam-forwarder/config"
	"github.com/eyedeekay/sam-forwarder/udp"
)

func ServeMode() {
	if *udpMode {
		var forwarder *samforwarderudp.SAMSSUForwarder
		log.Println("Redirecting udp", *TargetHost+":"+*TargetPort, "to i2p")
		forwarder, err = i2ptunconf.NewSAMSSUForwarderFromConf(config)
		if err == nil {
			forwarder.Serve()
		} else {
			log.Println(err.Error())
		}
	} else {
		var forwarder *samforwarder.SAMForwarder
		log.Println("Redirecting tcp", *TargetHost+":"+*TargetPort, "to i2p")
		forwarder, err := i2ptunconf.NewSAMForwarderFromConf(config)
		if err == nil {
			forwarder.Serve()
		} else {
			log.Println(err.Error())
		}
	}
}
