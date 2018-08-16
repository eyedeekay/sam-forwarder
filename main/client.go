package main

import "log"

import "github.com/eyedeekay/sam-forwarder/config"

func clientMode() {
	if *udpMode {
		log.Fatal("UDP client mode not implemented yet.")
	} else {
		log.Println("Proxying tcp", *targetHost+":"+*targetPort, "to", *targetDestination)
		forwarder, err := i2ptunconf.NewSAMClientForwarderFromConf(config)
		if err == nil {
			forwarder.Serve(*targetDestination)
		} else {
			log.Println(err.Error())
		}
	}
}
