package main

import (
	"flag"
	"log"
	"strings"
)

import (
	"github.com/eyedeekay/sam-forwarder"
	"github.com/eyedeekay/sam-forwarder/config"
	"github.com/eyedeekay/sam-forwarder/udp"
)

type flagOpts []string

func (f *flagOpts) String() string {
	r := ""
	for _, s := range *f {
		r += s + ","
	}
	return strings.TrimSuffix(r, ",")
}

func (f *flagOpts) Set(s string) error {
	*f = append(*f, s)
	return nil
}

func (f *flagOpts) StringSlice() []string {
	var r []string
	for _, s := range *f {
		r = append(r, s)
	}
	return r
}

func main() {
	var accessList flagOpts
	var err error

	saveFile := flag.Bool("save", false,
		"Use saved file and persist tunnel(If false, tunnel will not persist after program is stopped.")
	encryptLeaseSet := flag.Bool("encryptlease", true,
		"Use an encrypted leaseset(true or false)")
	inAllowZeroHop := flag.Bool("zeroin", false,
		"Allow zero-hop, non-anonymous tunnels in(true or false)")
	outAllowZeroHop := flag.Bool("zeroout", false,
		"Allow zero-hop, non-anonymous tunnels out(true or false)")
	useCompression := flag.Bool("gzip", false,
		"Uze gzip(true or false)")
	reduceIdle := flag.Bool("reduce", false,
		"Reduce tunnel quantity when idle(true or false)")
	udpMode := flag.Bool("udp", false,
		"UDP mode(true or false)")

	TargetDir := flag.String("dir", "",
		"Directory to save tunnel configuration file in.")
	iniFile := flag.String("ini", "none",
		"Use an ini file for configuration(config file options override passed arguments for now.)")
	TargetHost := flag.String("host", "127.0.0.1",
		"Target host(Host of service to forward to i2p)")
	TargetPort := flag.String("port", "8081",
		"Target port(Port of service to forward to i2p)")
	SamHost := flag.String("samhost", "127.0.0.1",
		"SAM host")
	SamPort := flag.String("samport", "7656",
		"SAM port")
	TunName := flag.String("name", "forwarder",
		"Tunnel name, this must be unique but can be anything.")
	accessListType := flag.String("access", "none",
		"Type of access list to use, can be \"whitelist\" \"blacklist\" or \"none\".")

	inLength := flag.Int("inlen", 3,
		"Set inbound tunnel length(0 to 7)")
	outLength := flag.Int("outlen", 3,
		"Set outbound tunnel length(0 to 7)")
	inQuantity := flag.Int("incount", 8,
		"Set inbound tunnel quantity(0 to 15)")
	outQuantity := flag.Int("outcount", 8,
		"Set outbound tunnel quantity(0 to 15)")
	inVariance := flag.Int("invar", 0,
		"Set inbound tunnel length variance(-7 to 7)")
	outVariance := flag.Int("outvar", 0,
		"Set outbound tunnel length variance(-7 to 7)")
	inBackupQuantity := flag.Int("inback", 4,
		"Set inbound tunnel backup quantity(0 to 5)")
	outBackupQuantity := flag.Int("outback", 4,
		"Set outbound tunnel backup quantity(0 to 5)")
	reduceIdleTime := flag.Int("reducetime", 10,
		"Reduce tunnel quantity after X (minutes)")
	reduceIdleQuantity := flag.Int("reducecount", 3,
		"Reduce idle tunnel quantity to X (0 to 5)")

	flag.Var(&accessList, "accesslist",
		"Specify an access list member(can be used multiple times)")

	flag.Parse()

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
				samforwarder.SetName(*TunName),
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
