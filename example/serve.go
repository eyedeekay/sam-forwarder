package main

import (
	"flag"
	"log"
	"net/http"
)

import (
	"github.com/eyedeekay/sam-forwarder"
	"github.com/eyedeekay/sam-forwarder/config"
)

func main() {
	host := flag.String("h", "127.0.0.1",
		"hostname to serve on")
	port := flag.String("p", "8100",
		"port to serve locally on")
	samhost := flag.String("sh", "127.0.0.1",
		"sam host to connect to")
	samport := flag.String("sp", "7656",
		"sam port to connect to")
	directory := flag.String("d", "./www",
		"the directory of static files to host(default ./www)")
	sdirectory := flag.String("s", ".",
		"the directory to save the keys in(default ./)")
	usei2p := flag.Bool("i", true,
		"save i2p keys(and thus destinations) across reboots")
	servicename := flag.String("n", "static-eepSite",
		"name to give the tunnel(default static-eepSite)")
	useCompression := flag.Bool("g", true,
		"Uze gzip(true or false)")
	encryptLeaseSet := flag.Bool("c", false,
		"Use an encrypted leaseset(true or false)")

	allowZeroHop := flag.Bool("z", false,
		"Allow zero-hop, non-anonymous tunnels(true or false)")

	reduceIdle := flag.Bool("r", false,
		"Reduce tunnel quantity when idle(true or false)")
	reduceIdleTime := flag.Int("rt", 10,
		"Reduce tunnel quantity after X (minutes)")
	reduceIdleQuantity := flag.Int("rc", 3,
		"Reduce idle tunnel quantity to X (0 to 5)")

	inLength := flag.Int("il", 3,
		"Set inbound tunnel length(0 to 7)")
	outLength := flag.Int("ol", 3,
		"Set outbound tunnel length(0 to 7)")
	inQuantity := flag.Int("iq", 8,
		"Set inbound tunnel quantity(0 to 15)")
	outQuantity := flag.Int("oq", 8,
		"Set outbound tunnel quantity(0 to 15)")
	inVariance := flag.Int("iv", 0,
		"Set inbound tunnel length variance(-7 to 7)")
	outVariance := flag.Int("ov", 0,
		"Set outbound tunnel length variance(-7 to 7)")
	inBackupQuantity := flag.Int("ib", 4,
		"Set inbound tunnel backup quantity(0 to 5)")
	outBackupQuantity := flag.Int("ob", 4,
		"Set outbound tunnel backup quantity(0 to 5)")
	iniFile := flag.String("f", "none",
		"Use an ini file for configuration")

	flag.Parse()

	var forwarder *samforwarder.SAMForwarder
	var err error
	if *iniFile != "none" {
		forwarder, err = i2ptunconf.NewSAMForwarderFromConfig(*iniFile, *samhost, *samport)
	} else {
		forwarder, err = samforwarder.NewSAMForwarderFromOptions(
			samforwarder.SetFilePath(*sdirectory),
			samforwarder.SetSaveFile(*usei2p),
			samforwarder.SetHost(*host),
			samforwarder.SetPort(*port),
			samforwarder.SetSAMHost(*samhost),
			samforwarder.SetSAMPort(*samport),
			samforwarder.SetName(*servicename),
			samforwarder.SetInLength(*inLength),
			samforwarder.SetOutLength(*outLength),
			samforwarder.SetInVariance(*inVariance),
			samforwarder.SetOutVariance(*outVariance),
			samforwarder.SetInQuantity(*inQuantity),
			samforwarder.SetOutQuantity(*outQuantity),
			samforwarder.SetInBackups(*inBackupQuantity),
			samforwarder.SetOutBackups(*outBackupQuantity),
			samforwarder.SetEncrypt(*encryptLeaseSet),
			samforwarder.SetAllowZeroIn(*allowZeroHop),
			samforwarder.SetAllowZeroOut(*allowZeroHop),
			samforwarder.SetCompress(*useCompression),
			samforwarder.SetReduceIdle(*reduceIdle),
			samforwarder.SetReduceIdleTime(*reduceIdleTime),
			samforwarder.SetReduceIdleQuantity(*reduceIdleQuantity),
		)
	}
	if err != nil {
		log.Fatal(err.Error())
	}
	go forwarder.Serve()

	http.Handle("/", http.FileServer(http.Dir(*directory)))

	log.Printf("Serving %s on HTTP port: %s\n", *directory, *port, "and on",
		forwarder.Base32()+".b32.i2p")
	log.Fatal(http.ListenAndServe(*host+":"+*port, nil))
}
