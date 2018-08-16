package main

import "log"

import "github.com/eyedeekay/sam-forwarder/config"

func ClientMode() {
	if *udpMode {

	} else {
		log.Println("Redirecting tcp", *TargetHost+":"+*TargetPort, "to i2p")
		forwarder, err := i2ptunconf.NewSAMClientForwarderFromConf(config)
		if err == nil {
			forwarder.Serve()
		} else {
			log.Println(err.Error())
		}
	}
}
