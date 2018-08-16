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
		if *iniFile != "none" {
			forwarder, err = i2ptunconf.NewSAMSSUForwarderFromConfig(*iniFile, *SamHost, *SamPort)
		} else {
			forwarder, err = samforwarderudp.NewSAMSSUForwarderFromOptions(
				samforwarderudp.SetName(*TunName),
				samforwarderudp.SetFilePath(*TargetDir),
				samforwarderudp.SetSaveFile(*saveFile),
				samforwarderudp.SetHost(*TargetHost),
				samforwarderudp.SetPort(*TargetPort),
				samforwarderudp.SetSAMHost(*SamHost),
				samforwarderudp.SetSAMPort(*SamPort),
				samforwarderudp.SetName(*TunName),
				samforwarderudp.SetInLength(*inLength),
				samforwarderudp.SetOutLength(*outLength),
				samforwarderudp.SetInVariance(*inVariance),
				samforwarderudp.SetOutVariance(*outVariance),
				samforwarderudp.SetInQuantity(*inQuantity),
				samforwarderudp.SetOutQuantity(*outQuantity),
				samforwarderudp.SetInBackups(*inBackupQuantity),
				samforwarderudp.SetOutBackups(*outBackupQuantity),
				samforwarderudp.SetEncrypt(*encryptLeaseSet),
				samforwarderudp.SetAllowZeroIn(*inAllowZeroHop),
				samforwarderudp.SetAllowZeroOut(*outAllowZeroHop),
				samforwarderudp.SetCompress(*useCompression),
				samforwarderudp.SetReduceIdle(*reduceIdle),
				samforwarderudp.SetReduceIdleTime(*reduceIdleTime),
				samforwarderudp.SetReduceIdleQuantity(*reduceIdleQuantity),
				samforwarderudp.SetAccessListType(*accessListType),
				samforwarderudp.SetAccessList(accessList.StringSlice()),
			)
		}
		if err == nil {
			forwarder.Serve()
		} else {
			log.Println(err.Error())
		}
	} else {
		var forwarder *samforwarder.SAMForwarder
		log.Println("Redirecting tcp", *TargetHost+":"+*TargetPort, "to i2p")
		if *iniFile != "none" {
			forwarder, err = i2ptunconf.NewSAMForwarderFromConfig(*iniFile, *SamHost, *SamPort)
		} else {
			forwarder, err = samforwarder.NewSAMForwarderFromOptions(
				samforwarder.SetFilePath(*TargetDir),
				samforwarder.SetSaveFile(*saveFile),
				samforwarder.SetHost(*TargetHost),
				samforwarder.SetPort(*TargetPort),
				samforwarder.SetSAMHost(*SamHost),
				samforwarder.SetSAMPort(*SamPort),
				samforwarder.SetName(*TunName),
				samforwarder.SetInLength(*inLength),
				samforwarder.SetOutLength(*outLength),
				samforwarder.SetInVariance(*inVariance),
				samforwarder.SetOutVariance(*outVariance),
				samforwarder.SetInQuantity(*inQuantity),
				samforwarder.SetOutQuantity(*outQuantity),
				samforwarder.SetInBackups(*inBackupQuantity),
				samforwarder.SetOutBackups(*outBackupQuantity),
				samforwarder.SetEncrypt(*encryptLeaseSet),
				samforwarder.SetAllowZeroIn(*inAllowZeroHop),
				samforwarder.SetAllowZeroOut(*outAllowZeroHop),
				samforwarder.SetCompress(*useCompression),
				samforwarder.SetReduceIdle(*reduceIdle),
				samforwarder.SetReduceIdleTime(*reduceIdleTime),
				samforwarder.SetReduceIdleQuantity(*reduceIdleQuantity),
				samforwarder.SetAccessListType(*accessListType),
				samforwarder.SetAccessList(accessList.StringSlice()),
			)
		}
		if err == nil {
			forwarder.Serve()
		} else {
			log.Println(err.Error())
		}
	}
}
