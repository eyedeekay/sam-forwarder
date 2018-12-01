package main

/*
   WARNING: This is not the official version of eephttpd. It is an older
   verion I use to test new sam-forwarder features. It is not intended for
   use.
*/

import (
	"crypto/tls"
	"flag"
	"log"
	"net/http"
	"path/filepath"
)

import (
	"github.com/eyedeekay/sam-forwarder"
	"github.com/eyedeekay/sam-forwarder/config"
)

var cfg = &tls.Config{
	MinVersion:               tls.VersionTLS12,
	CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
	PreferServerCipherSuites: true,
	CipherSuites: []uint16{
		tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
		tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
		tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
		tls.TLS_RSA_WITH_AES_256_CBC_SHA,
	},
}

var (
	host               = flag.String("a", "127.0.0.1", "hostname to serve on")
	port               = flag.String("p", "7880", "port to serve locally on")
	samhost            = flag.String("sh", "127.0.0.1", "sam host to connect to")
	samport            = flag.String("sp", "7656", "sam port to connect to")
	directory          = flag.String("d", "./www", "the directory of static files to host(default ./www)")
	sdirectory         = flag.String("s", ".", "the directory to save the keys in(default ./)")
	usei2p             = flag.Bool("i", true, "save i2p keys(and thus destinations) across reboots")
	servicename        = flag.String("n", "static-eepSite", "name to give the tunnel(default static-eepSite)")
	useCompression     = flag.Bool("g", true, "Uze gzip(true or false)")
	injectHeaders      = flag.Bool("x", true, "Inject X-I2P-DEST headers")
	accessListType     = flag.String("l", "none", "Type of access list to use, can be \"whitelist\" \"blacklist\" or \"none\".")
	encryptLeaseSet    = flag.Bool("c", false, "Use an encrypted leaseset(true or false)")
	allowZeroHop       = flag.Bool("z", false, "Allow zero-hop, non-anonymous tunnels(true or false)")
	reduceIdle         = flag.Bool("r", false, "Reduce tunnel quantity when idle(true or false)")
	reduceIdleTime     = flag.Int("rt", 600000, "Reduce tunnel quantity after X (milliseconds)")
	reduceIdleQuantity = flag.Int("rc", 3, "Reduce idle tunnel quantity to X (0 to 5)")
	inLength           = flag.Int("il", 3, "Set inbound tunnel length(0 to 7)")
	outLength          = flag.Int("ol", 3, "Set outbound tunnel length(0 to 7)")
	inQuantity         = flag.Int("iq", 8, "Set inbound tunnel quantity(0 to 15)")
	outQuantity        = flag.Int("oq", 8, "Set outbound tunnel quantity(0 to 15)")
	inVariance         = flag.Int("iv", 0, "Set inbound tunnel length variance(-7 to 7)")
	outVariance        = flag.Int("ov", 0, "Set outbound tunnel length variance(-7 to 7)")
	inBackupQuantity   = flag.Int("ib", 4, "Set inbound tunnel backup quantity(0 to 5)")
	outBackupQuantity  = flag.Int("ob", 4, "Set outbound tunnel backup quantity(0 to 5)")
	iniFile            = flag.String("f", "none", "Use an ini file for configuration")
	useTLS             = flag.Bool("t", false, "Generate or use an existing TLS certificate")
	certFile           = flag.String("m", "cert", "Certificate name to use")
)

func main() {
	flag.Parse()
	var forwarder *samforwarder.SAMForwarder
	var err error
	config := i2ptunconf.NewI2PBlankTunConf()
	if *iniFile != "none" {
		config, err = i2ptunconf.NewI2PTunConf(*iniFile)
	}
	config.TargetHost = config.GetHost(*host, "127.0.0.1")
	config.TargetPort = config.GetPort(*port, "7880")
	config.SaveFile = config.GetSaveFile(*usei2p, true)
	config.SaveDirectory = config.GetDir(*sdirectory, "../")
	config.SamHost = config.GetSAMHost(*samhost, "127.0.0.1")
	config.SamPort = config.GetSAMPort(*samport, "7656")
	config.TunName = config.GetKeys(*servicename, "static-eepSite")
	config.InLength = config.GetInLength(*inLength, 3)
	config.OutLength = config.GetOutLength(*outLength, 3)
	config.InVariance = config.GetInVariance(*inVariance, 0)
	config.OutVariance = config.GetOutVariance(*outVariance, 0)
	config.InQuantity = config.GetInQuantity(*inQuantity, 6)
	config.OutQuantity = config.GetOutQuantity(*outQuantity, 6)
	config.InBackupQuantity = config.GetInBackups(*inBackupQuantity, 5)
	config.OutBackupQuantity = config.GetOutBackups(*outBackupQuantity, 5)
	config.EncryptLeaseSet = config.GetEncryptLeaseset(*encryptLeaseSet, false)
	config.InAllowZeroHop = config.GetInAllowZeroHop(*allowZeroHop, false)
	config.OutAllowZeroHop = config.GetOutAllowZeroHop(*allowZeroHop, false)
	config.UseCompression = config.GetUseCompression(*useCompression, true)
	config.ReduceIdle = config.GetReduceOnIdle(*reduceIdle, true)
	config.ReduceIdleTime = config.GetReduceIdleTime(*reduceIdleTime, 600000)
	config.ReduceIdleQuantity = config.GetReduceIdleQuantity(*reduceIdleQuantity, 2)
	config.CloseIdleTime = config.GetCloseIdleTime(*reduceIdleTime, 600000)
	config.AccessListType = config.GetAccessListType(*accessListType, "none")
	config.Type = config.GetType(false, false, *injectHeaders, "server")

	if forwarder, err = i2ptunconf.NewSAMForwarderFromConf(config); err != nil {
		log.Fatal(err.Error())
	}
	go forwarder.Serve()

	if *useTLS {
		srv := &http.Server{
			Addr:         *host + ":" + *port,
			Handler:      http.FileServer(http.Dir(*directory)),
			TLSConfig:    cfg,
			TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
		}
		log.Printf("Serving %s on HTTPS port: %s\n\t and on \n%s", *directory, *port, forwarder.Base32())
		log.Fatal(
			srv.ListenAndServeTLS(
				filepath.Join(*sdirectory+"/", *certFile+".crt"),
				filepath.Join(*sdirectory+"/", *certFile+".key"),
			),
		)
	} else {
		log.Printf("Serving %s on HTTP port: %s\n\t and on \n%s", *directory, *port, forwarder.Base32())
		log.Fatal(
			http.ListenAndServe(
				*host+":"+*port,
				http.FileServer(http.Dir(*directory)),
			),
		)
	}
}
