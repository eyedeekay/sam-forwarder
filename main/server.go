package main

import "log"

import "github.com/eyedeekay/sam-forwarder/config"

func ServeMode() {
	if *udpMode {
		log.Println("Redirecting udp", *TargetHost+":"+*TargetPort, "to i2p")
		forwarder, err := i2ptunconf.NewSAMSSUForwarderFromConf(config)
		if err == nil {
			forwarder.Serve()
		} else {
			log.Println(err.Error())
		}
	} else {
		log.Println("Redirecting tcp", *TargetHost+":"+*TargetPort, "to i2p")
		forwarder, err := i2ptunconf.NewSAMForwarderFromConf(config)
		if err == nil {
			forwarder.Serve()
		} else {
			log.Println(err.Error())
		}
	}
}
