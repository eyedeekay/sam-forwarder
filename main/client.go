package main

import "log"

import "github.com/eyedeekay/sam-forwarder/config"

func clientMode() {
	if *udpMode {
		log.Println("Proxying udp", *targetHost+":"+*targetPort, "to", *targetDestination)
		forwarder, err := i2ptunconf.NewSAMSSUClientForwarderFromConf(config)
		if err == nil {
			forwarder.Serve()
		} else {
			log.Println(err.Error())
		}
	} else {
		log.Println("Proxying tcp", *targetHost+":"+*targetPort, "to", *targetDestination)
		forwarder, err := i2ptunconf.NewSAMClientForwarderFromConf(config)
		if err == nil {
			forwarder.Serve()
		} else {
			log.Println(err.Error())
		}
	}
}
