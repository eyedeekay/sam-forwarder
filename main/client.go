package main

import "log"

import (
	"github.com/eyedeekay/sam-forwarder"
	//"github.com/eyedeekay/sam-forwarder/config"
	//"github.com/eyedeekay/sam-forwarder/udp"
)

func ClientMode(){
    if *udpMode {

	} else {
		var forwarder *samforwarder.SAMClientForwarder
		log.Println("Proxying tcp", *TargetHost+":"+*TargetPort, "to i2p")
		if *iniFile != "none" {
			//forwarder, err = i2ptunconf.NewSAMForwarderFromConfig(*iniFile, *SamHost, *SamPort)
		} else {
			forwarder, err = samforwarder.NewSAMClientForwarderFromOptions(
				samforwarder.SetClientFilePath(*TargetDir),
				samforwarder.SetClientSaveFile(*saveFile),
				samforwarder.SetClientHost(*TargetHost),
				samforwarder.SetClientPort(*TargetPort),
				samforwarder.SetClientSAMHost(*SamHost),
				samforwarder.SetClientSAMPort(*SamPort),
				samforwarder.SetClientName(*TunName),
				samforwarder.SetClientInLength(*inLength),
				samforwarder.SetClientOutLength(*outLength),
				samforwarder.SetClientInVariance(*inVariance),
				samforwarder.SetClientOutVariance(*outVariance),
				samforwarder.SetClientInQuantity(*inQuantity),
				samforwarder.SetClientOutQuantity(*outQuantity),
				samforwarder.SetClientInBackups(*inBackupQuantity),
				samforwarder.SetClientOutBackups(*outBackupQuantity),
				samforwarder.SetClientEncrypt(*encryptLeaseSet),
				samforwarder.SetClientAllowZeroIn(*inAllowZeroHop),
				samforwarder.SetClientAllowZeroOut(*outAllowZeroHop),
				samforwarder.SetClientCompress(*useCompression),
				samforwarder.SetClientReduceIdle(*reduceIdle),
				samforwarder.SetClientReduceIdleTime(*reduceIdleTime),
				samforwarder.SetClientReduceIdleQuantity(*reduceIdleQuantity),
				samforwarder.SetClientAccessListType(*accessListType),
				samforwarder.SetClientAccessList(accessList.StringSlice()),
			)
		}
		if err == nil {
			forwarder.Serve()
		} else {
			log.Println(err.Error())
		}
	}
}
