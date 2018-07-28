package main

import (
	"flag"
	"log"
	"strings"
)

import "github.com/eyedeekay/sam-forwarder"

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
	//
	TargetDir := *flag.String("dir", "",
		"Directory to save tunnel configuration file in.")
    saveFile := *flag.Bool("save", true,
		"Use saved file and persist tunnel(If false, tunnel will not persist after program is stopped.")
	TargetHost := *flag.String("host", "127.0.0.1",
		"Target host(Host of service to forward to i2p)")
	TargetPort := *flag.String("port", "8081",
		"Target port(Port of service to forward to i2p)")
	SamHost := *flag.String("samhost", "127.0.0.1",
		"SAM host")
	SamPort := *flag.String("samport", "7656",
		"SAM port")
	TunName := *flag.String("name", "forwarder",
		"Tunnel name, this must be unique but can be anything.")
	encryptLeaseSet := *flag.Bool("encryptlease", true,
		"Use an encrypted leaseset(true or false)")
	inAllowZeroHop := *flag.Bool("zeroin", false,
		"Allow zero-hop, non-anonymous tunnels in(true or false)")
	outAllowZeroHop := *flag.Bool("zeroout", false,
		"Allow zero-hop, non-anonymous tunnels out(true or false)")
	inLength := *flag.Int("inlen", 3,
		"Set inbound tunnel length(0 to 7)")
	outLength := *flag.Int("outlen", 3,
		"Set outbound tunnel length(0 to 7)")
	inQuantity := *flag.Int("incount", 8,
		"Set inbound tunnel quantity(0 to 15)")
	outQuantity := *flag.Int("outcount", 8,
		"Set outbound tunnel quantity(0 to 15)")
	inVariance := *flag.Int("invar", 0,
		"Set inbound tunnel length variance(-7 to 7)")
	outVariance := *flag.Int("outvar", 0,
		"Set outbound tunnel length variance(-7 to 7)")
	inBackupQuantity := *flag.Int("inback", 4,
		"Set inbound tunnel backup quantity(0 to 5)")
	outBackupQuantity := *flag.Int("outback", 4,
		"Set outbound tunnel backup quantity(0 to 5)")
	useCompression := *flag.Bool("gzip", false,
		"Uze gzip(true or false)")
	reduceIdle := *flag.Bool("reduce", false,
		"Reduce tunnel quantity when idle(true or false)")
	reduceIdleTime := *flag.Int("reducetime", 3,
		"Reduce tunnel quantity after X (minutes)")
	reduceIdleQuantity := *flag.Int("reducecount", 3,
		"Reduce idle tunnel quantity to X (0 to 5)")
	accessListType := *flag.String("access", "none",
		"Type of access list to use, can be \"whitelist\" \"blacklist\" or \"none\".")
	flag.Var(&accessList, "accesslist",
		"Specify an access list member(can be used multiple times)")

	flag.Parse()
	log.Println("Redirecting", TargetHost+":"+TargetPort, "to i2p")
	forwarder, err := samforwarder.NewSAMForwarderFromOptions(
		samforwarder.SetFilePath(TargetDir),
        samforwarder.SetSaveFile(saveFile),
		samforwarder.SetHost(TargetHost),
		samforwarder.SetPort(TargetPort),
		samforwarder.SetSAMHost(SamHost),
		samforwarder.SetSAMPort(SamPort),
		samforwarder.SetName(TunName),
		samforwarder.SetInLength(inLength),
		samforwarder.SetOutLength(outLength),
		samforwarder.SetInVariance(inVariance),
		samforwarder.SetOutVariance(outVariance),
		samforwarder.SetInQuantity(inQuantity),
		samforwarder.SetOutQuantity(outQuantity),
		samforwarder.SetInBackups(inBackupQuantity),
		samforwarder.SetOutBackups(outBackupQuantity),
		samforwarder.SetEncrypt(encryptLeaseSet),
		samforwarder.SetAllowZeroIn(inAllowZeroHop),
		samforwarder.SetAllowZeroOut(outAllowZeroHop),
		samforwarder.SetCompress(useCompression),
		samforwarder.SetReduceIdle(reduceIdle),
		samforwarder.SetReduceIdleTime(reduceIdleTime),
		samforwarder.SetReduceIdleQuantity(reduceIdleQuantity),
		samforwarder.SetAccessListType(accessListType),
		samforwarder.SetAccessList(accessList.StringSlice()),
	)
	if err == nil {
		forwarder.Serve()
	} else {
		log.Println(err.Error())
	}
}
