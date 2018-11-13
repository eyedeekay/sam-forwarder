package main

import "log"

import "github.com/eyedeekay/sam-forwarder/config"

func serveMode() {
	if *udpMode {
		log.Println("Redirecting udp", *targetHost+":"+*targetPort, "to i2p")
		forwarder, err := i2ptunconf.NewSAMSSUForwarderFromConf(config)
		if err == nil {
			forwarder.Serve()
		} else {
			log.Println(err.Error())
		}
		forwarder.Cleanup()
	} else {
		log.Println("Redirecting tcp", *targetHost+":"+*targetPort, "to i2p")
		forwarder, err := i2ptunconf.NewSAMForwarderFromConf(config)
		if err == nil {
			forwarder.Serve()
		} else {
			log.Println(err.Error())
		}
		forwarder.Cleanup()
	}
}
