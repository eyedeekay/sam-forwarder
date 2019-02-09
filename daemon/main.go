package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"strings"
)

import (
	"github.com/eyedeekay/littleboss"
	"github.com/eyedeekay/sam-forwarder/config"
	"github.com/eyedeekay/sam-forwarder/manager"
	"github.com/eyedeekay/samcatd-web"
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
	encryptKeyFiles = flag.String("cr", "",
		"Encrypt/decrypt the key files with a passfile")
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
	webAdmin = flag.Bool("w", false,
		"Start web administration interface")
	webPort = flag.String("wp", "7957",
		"Web port")
	webCSS = flag.String("css", "css/styles.css",
		"custom CSS for web interface")
	webJS = flag.String("js", "js/scripts.js",
		"custom JS for web interface")
	leaseSetKey = flag.String("k", "none",
		"key for encrypted leaseset")
	leaseSetPrivateKey = flag.String("pk", "none",
		"private key for encrypted leaseset")
	leaseSetPrivateSigningKey = flag.String("psk", "none",
		"private signing key for encrypted leaseset")
	targetDir = flag.String("d", "",
		"Directory to save tunnel configuration file in.")
	targetDest = flag.String("de", "",
		"Destination to connect client's to by default.")
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
	inQuantity = flag.Int("iq", 6,
		"Set inbound tunnel quantity(0 to 15)")
	outQuantity = flag.Int("oq", 6,
		"Set outbound tunnel quantity(0 to 15)")
	inVariance = flag.Int("iv", 0,
		"Set inbound tunnel length variance(-7 to 7)")
	outVariance = flag.Int("ov", 0,
		"Set outbound tunnel length variance(-7 to 7)")
	inBackupQuantity = flag.Int("ib", 2,
		"Set inbound tunnel backup quantity(0 to 5)")
	outBackupQuantity = flag.Int("ob", 2,
		"Set outbound tunnel backup quantity(0 to 5)")
	reduceIdleTime = flag.Int("rt", 600000,
		"Reduce tunnel quantity after X (milliseconds)")
	closeIdleTime = flag.Int("ct", 600000,
		"Reduce tunnel quantity after X (milliseconds)")
	reduceIdleQuantity = flag.Int("rq", 3,
		"Reduce idle tunnel quantity to X (0 to 5)")
	readKeys = flag.String("conv", "", "Display the base32 and base64 values of a specified .i2pkeys file")
)

var (
	webinterface    *samcatweb.SAMWebConfig
	webinterfaceerr error
	err             error
	accessList      flagOpts
	config          *i2ptunconf.Conf
)

func main() {
	lb := littleboss.New("service-name")
	lb.Run(func(ctx context.Context) {
		lbMain(ctx)
	})
}

func lbMain(ctx context.Context) {
	flag.Var(&accessList, "accesslist", "Specify an access list member(can be used multiple times)")
	flag.Parse()

	if *readKeys != "" {

	}

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
	config.LeaseSetKey = config.GetLeasesetKey(*leaseSetKey, "")
	config.LeaseSetPrivateKey = config.GetLeasesetPrivateKey(*leaseSetPrivateKey, "")
	config.LeaseSetPrivateSigningKey = config.GetLeasesetPrivateSigningKey(*leaseSetPrivateSigningKey, "")
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
	config.KeyFilePath = config.GetKeyFile(*encryptKeyFiles, "")
	config.ClientDest = config.GetClientDest(*targetDest, "", "")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	if manager, err := sammanager.NewSAMManagerFromConf(
		config,
		config.TargetHost,
		config.TargetPort,
		config.SamHost,
		config.SamPort,
		"localhost",
		*webPort,
		*startUp,
	); err == nil {
		go func() {
			for sig := range c {
				if sig == os.Interrupt {
					manager.Cleanup()
				}
			}
		}()
		if *webAdmin {
			go samcatweb.Serve(manager, "", "", manager.WebHost, manager.WebPort)
		}
		manager.Serve()
	} else {
		log.Fatal(err)
	}
	ctx.Done()
}
