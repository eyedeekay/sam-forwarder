package main

import (
	"flag"
	"log"
	"strings"
)

import "github.com/eyedeekay/sam-forwarder/config"

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
	saveFile              = flag.Bool("save", false, "Use saved file and persist tunnel(If false, tunnel will not persist after program is stopped.")
	encryptLeaseSet       = flag.Bool("encryptlease", true, "Use an encrypted leaseset(true or false)")
	inAllowZeroHop        = flag.Bool("zeroin", false, "Allow zero-hop, non-anonymous tunnels in(true or false)")
	outAllowZeroHop       = flag.Bool("zeroout", false, "Allow zero-hop, non-anonymous tunnels out(true or false)")
	useCompression        = flag.Bool("gzip", false, "Uze gzip(true or false)")
	reduceIdle            = flag.Bool("reduce", false, "Reduce tunnel quantity when idle(true or false)")
	closeIdle             = flag.Bool("close", false, "Close tunnel idle(true or false)")
	udpMode               = flag.Bool("udp", false, "UDP mode(true or false)")
	client                = flag.Bool("client", false, "Client proxy mode(true or false)")
	injectHeaders         = flag.Bool("headers", false, "Inject X-I2P-DEST headers")
	encryptedLeasesetKeys = flag.String("lsk", "none", "path to saved encrypted leaseset keys")
	targetDir             = flag.String("dir", "", "Directory to save tunnel configuration file in.")
	iniFile               = flag.String("ini", "none", "Use an ini file for configuration(config file options override passed arguments for now.)")
	targetDestination     = flag.String("dest", "none", "Destination for client tunnels. Ignored for service tunnels.")
	targetHost            = flag.String("host", "127.0.0.1", "Target host(Host of service to forward to i2p)")
	targetPort            = flag.String("port", "8081", "Target port(Port of service to forward to i2p)")
	targetPort443         = flag.String("tlsport", "", "(Currently inoperative. Target TLS port(HTTPS Port of service to forward to i2p)")
	samHost               = flag.String("samhost", "127.0.0.1", "SAM host")
	samPort               = flag.String("samport", "7656", "SAM port")
	tunName               = flag.String("name", "forwarder", "Tunnel name, this must be unique but can be anything.")
	accessListType        = flag.String("access", "none", "Type of access list to use, can be \"whitelist\" \"blacklist\" or \"none\".")
	inLength              = flag.Int("inlen", 3, "Set inbound tunnel length(0 to 7)")
	outLength             = flag.Int("outlen", 3, "Set outbound tunnel length(0 to 7)")
	inQuantity            = flag.Int("incount", 6, "Set inbound tunnel quantity(0 to 15)")
	outQuantity           = flag.Int("outcount", 6, "Set outbound tunnel quantity(0 to 15)")
	inVariance            = flag.Int("invar", 0, "Set inbound tunnel length variance(-7 to 7)")
	outVariance           = flag.Int("outvar", 0, "Set outbound tunnel length variance(-7 to 7)")
	inBackupQuantity      = flag.Int("inback", 4, "Set inbound tunnel backup quantity(0 to 5)")
	outBackupQuantity     = flag.Int("outback", 4, "Set outbound tunnel backup quantity(0 to 5)")
	reduceIdleTime        = flag.Int("reducetime", 600000, "Reduce tunnel quantity after X (milliseconds)")
	closeIdleTime         = flag.Int("closetime", 600000, "Reduce tunnel quantity after X (milliseconds)")
	reduceIdleQuantity    = flag.Int("reducecount", 3, "Reduce idle tunnel quantity to X (0 to 5)")
)

var err error
var accessList flagOpts
var config *i2ptunconf.Conf

func main() {
	flag.Var(&accessList, "accesslist", "Specify an access list member(can be used multiple times)")
	flag.Parse()

	config = i2ptunconf.NewI2PBlankTunConf()
	if *iniFile != "none" {
		config, err = i2ptunconf.NewI2PTunConf(*iniFile)
	}
	config.TargetHost = config.GetHost(*targetHost, "127.0.0.1")
	config.TargetPort = config.GetPort(*targetPort, "8081")
	config.SaveFile = config.GetSaveFile(*saveFile, true)
	config.SaveDirectory = config.GetDir(*targetDir, "../")
	config.SamHost = config.GetSAMHost(*samHost, "127.0.0.1")
	config.SamPort = config.GetSAMPort(*samPort, "7656")
	config.TunName = config.GetKeys(*tunName, "forwarder")
	config.InLength = config.GetInLength(*inLength, 3)
	config.OutLength = config.GetOutLength(*outLength, 3)
	config.InVariance = config.GetInVariance(*inVariance, 0)
	config.OutVariance = config.GetOutVariance(*outVariance, 0)
	config.InQuantity = config.GetInQuantity(*inQuantity, 6)
	config.OutQuantity = config.GetOutQuantity(*outQuantity, 6)
	config.InBackupQuantity = config.GetInBackups(*inBackupQuantity, 5)
	config.OutBackupQuantity = config.GetOutBackups(*outBackupQuantity, 5)
	config.EncryptLeaseSet = config.GetEncryptLeaseset(*encryptLeaseSet, false)
	config.InAllowZeroHop = config.GetInAllowZeroHop(*inAllowZeroHop, false)
	config.OutAllowZeroHop = config.GetOutAllowZeroHop(*outAllowZeroHop, false)
	config.UseCompression = config.GetUseCompression(*useCompression, true)
	config.ReduceIdle = config.GetReduceOnIdle(*reduceIdle, true)
	config.ReduceIdleTime = config.GetReduceIdleTime(*reduceIdleTime, 600000)
	config.ReduceIdleQuantity = config.GetReduceIdleQuantity(*reduceIdleQuantity, 2)
	config.AccessListType = config.GetAccessListType(*accessListType, "none")
	config.CloseIdle = config.GetCloseOnIdle(*closeIdle, false)
	config.CloseIdleTime = config.GetCloseIdleTime(*closeIdleTime, 600000)
	config.Type = config.GetType(*client, *udpMode, *injectHeaders, "server")
	config.TargetForPort443 = config.GetPort443(*targetPort443, "")
	config.ClientDest = config.GetClientDest(*targetDestination, "", "")

	if config.Client {
		if *targetDestination == "none" {
			log.Fatal("Client mode requires you to specify a base32 or jump destination")
		} else {
			log.Println("Client mode is still experimental.")
			clientMode()
		}
	} else {
		serveMode()
	}
}
