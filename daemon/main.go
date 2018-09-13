package main

import (
	"flag"
	"log"
	"strings"
)

import (
	"github.com/eyedeekay/sam-forwarder/config"
	"github.com/eyedeekay/sam-forwarder/manager"
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
	saveFile = flag.Bool("t", false,
		"Use saved file and persist tunnel(If false, tunnel will not persist after program is stopped.")
	startUp = flag.Bool("s", false,
		"Start a tunnel with the passed parameters(Otherwise, they will be treated as default values.)")
	encryptLeaseSet = flag.Bool("l", true,
		"Use an encrypted leaseset(true or false)")
	inAllowZeroHop = flag.Bool("zi", false,
		"Allow zero-hop, non-anonymous tunnels in(true or false)")
	outAllowZeroHop = flag.Bool("zo", false,
		"Allow zero-hop, non-anonymous tunnels out(true or false)")
	useCompression = flag.Bool("z", false,
		"Uze gzip(true or false)")
	reduceIdle = flag.Bool("r", false,
		"Reduce tunnel quantity when idle(true or false)")
	closeIdle = flag.Bool("x", false,
		"Close tunnel idle(true or false)")
	udpMode = flag.Bool("u", false,
		"UDP mode(true or false)")
	client = flag.Bool("c", false,
		"Client proxy mode(true or false)")
	injectHeaders = flag.Bool("ih", false,
		"Inject X-I2P-DEST headers")
	encryptedLeasesetKeys = flag.String("k", "none",
		"path to saved encrypted leaseset keys")
	targetDir = flag.String("d", "",
		"Directory to save tunnel configuration file in.")
	iniFile = flag.String("f", "none",
		"Use an ini file for configuration(config file options override passed arguments for now.)")
	targetDestination = flag.String("i", "none",
		"Destination for client tunnels. Ignored for service tunnels.")
	targetHost = flag.String("h", "127.0.0.1",
		"Target host(Host of service to forward to i2p)")
	targetPort = flag.String("p", "8081",
		"Target port(Port of service to forward to i2p)")
	targetPort443 = flag.String("tls", "",
		"(Currently inoperative. Target TLS port(HTTPS Port of service to forward to i2p)")
	samHost = flag.String("sh", "127.0.0.1",
		"SAM host")
	samPort = flag.String("sp", "7656",
		"SAM port")
	tunName = flag.String("n", "forwarder",
		"Tunnel name, this must be unique but can be anything.")
	accessListType = flag.String("a", "none",
		"Type of access list to use, can be \"whitelist\" \"blacklist\" or \"none\".")
	inLength = flag.Int("il", 3,
		"Set inbound tunnel length(0 to 7)")
	outLength = flag.Int("ol", 3,
		"Set outbound tunnel length(0 to 7)")
	inQuantity = flag.Int("ic", 6,
		"Set inbound tunnel quantity(0 to 15)")
	outQuantity = flag.Int("oc", 6,
		"Set outbound tunnel quantity(0 to 15)")
	inVariance = flag.Int("iv", 0,
		"Set inbound tunnel length variance(-7 to 7)")
	outVariance = flag.Int("ov", 0,
		"Set outbound tunnel length variance(-7 to 7)")
	inBackupQuantity = flag.Int("ib", 4,
		"Set inbound tunnel backup quantity(0 to 5)")
	outBackupQuantity = flag.Int("ob", 4,
		"Set outbound tunnel backup quantity(0 to 5)")
	reduceIdleTime = flag.Int("rt", 600000,
		"Reduce tunnel quantity after X (milliseconds)")
	closeIdleTime = flag.Int("ct", 600000,
		"Reduce tunnel quantity after X (milliseconds)")
	reduceIdleQuantity = flag.Int("rc", 3,
		"Reduce idle tunnel quantity to X (0 to 5)")
)

var err error
var accessList flagOpts
var config *i2ptunconf.Conf

func main() {
	flag.Var(&accessList, "accesslist", "Specify an access list member(can be used multiple times)")
	flag.Parse()

	config = i2ptunconf.NewI2PBlankTunConf()
	if *iniFile != "none" && *iniFile != "" {
		config, err = i2ptunconf.NewI2PTunConf(*iniFile)
	} else {
		*startUp = true
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

	if manager, err := sammanager.NewSAMManagerFromConf(
		config,
		config.TargetHost,
		config.TargetPort,
		config.SamHost,
		config.SamPort,
		*startUp,
	); err == nil {
		manager.Serve()
	} else {
		log.Fatal(err)
	}
}
