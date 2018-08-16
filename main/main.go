package main

import (
	"flag"
	"log"
	"strings"
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

var (
	saveFile        = flag.Bool("save", false, "Use saved file and persist tunnel(If false, tunnel will not persist after program is stopped.")
	encryptLeaseSet = flag.Bool("encryptlease", true, "Use an encrypted leaseset(true or false)")
	inAllowZeroHop  = flag.Bool("zeroin", false, "Allow zero-hop, non-anonymous tunnels in(true or false)")
	outAllowZeroHop = flag.Bool("zeroout", false, "Allow zero-hop, non-anonymous tunnels out(true or false)")
	useCompression  = flag.Bool("gzip", false, "Uze gzip(true or false)")
	reduceIdle      = flag.Bool("reduce", false, "Reduce tunnel quantity when idle(true or false)")
	udpMode         = flag.Bool("udp", false, "UDP mode(true or false)")
	clientMode      = flag.Bool("client", false, "Client proxy mode(true or false)")
	//EncryptedLeasesetKeys = flag.String("lsk","none", "path to saved encrypted leaseset keys")
	TargetDir          = flag.String("dir", "", "Directory to save tunnel configuration file in.")
	iniFile            = flag.String("ini", "none", "Use an ini file for configuration(config file options override passed arguments for now.)")
	TargetDestination  = flag.String("dest", "none", "Destination for client tunnels. Ignored for service tunnels.")
	TargetHost         = flag.String("host", "127.0.0.1", "Target host(Host of service to forward to i2p)")
	TargetPort         = flag.String("port", "8081", "Target port(Port of service to forward to i2p)")
	SamHost            = flag.String("samhost", "127.0.0.1", "SAM host")
	SamPort            = flag.String("samport", "7656", "SAM port")
	TunName            = flag.String("name", "forwarder", "Tunnel name, this must be unique but can be anything.")
	accessListType     = flag.String("access", "none", "Type of access list to use, can be \"whitelist\" \"blacklist\" or \"none\".")
	inLength           = flag.Int("inlen", 3, "Set inbound tunnel length(0 to 7)")
	outLength          = flag.Int("outlen", 3, "Set outbound tunnel length(0 to 7)")
	inQuantity         = flag.Int("incount", 8, "Set inbound tunnel quantity(0 to 15)")
	outQuantity        = flag.Int("outcount", 8, "Set outbound tunnel quantity(0 to 15)")
	inVariance         = flag.Int("invar", 0, "Set inbound tunnel length variance(-7 to 7)")
	outVariance        = flag.Int("outvar", 0, "Set outbound tunnel length variance(-7 to 7)")
	inBackupQuantity   = flag.Int("inback", 4, "Set inbound tunnel backup quantity(0 to 5)")
	outBackupQuantity  = flag.Int("outback", 4, "Set outbound tunnel backup quantity(0 to 5)")
	reduceIdleTime     = flag.Int("reducetime", 10, "Reduce tunnel quantity after X (minutes)")
	reduceIdleQuantity = flag.Int("reducecount", 3, "Reduce idle tunnel quantity to X (0 to 5)")
)

var err error
var accessList flagOpts

func main() {
	flag.Var(&accessList, "accesslist", "Specify an access list member(can be used multiple times)")
	flag.Parse()
	if *clientMode {
		if *TargetDestination == "none" {
			log.Fatal("Client mode requires you to specify a base32 or jump destination")
		} else {
			log.Println("Client mode not implemented yet.")
		}
	} else {
		ServeMode()
	}
}
