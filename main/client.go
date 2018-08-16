package main

import "log"

import "github.com/eyedeekay/sam-forwarder/config"

func ClientMode() {
	if *udpMode {

	} else {
		log.Println("Proxying tcp", *TargetHost+":"+*TargetPort, "to", *TargetDestination)
		forwarder, err := i2ptunconf.NewSAMClientForwarderFromConf(config)
		if err == nil {
			forwarder.Serve(*TargetDestination)
		} else {
			log.Println(err.Error())
		}
	}
}
