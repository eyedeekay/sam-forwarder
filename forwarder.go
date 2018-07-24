package samlistener

import (
	"io"
	"log"
	"net"
)

import (
	"github.com/eyedeekay/i2pasta/convert"
	"github.com/kpetku/sam3"
)

var SamHost = "127.0.0.1"
var SamPort = "7656"
var TunName = "ephtun"

var samConn *sam3.SAM
var samKeys sam3.I2PKeys

var publishStream *sam3.StreamSession
var publishListen *sam3.StreamListener
var publishConnection net.Conn

var Target = "127.0.0.1:8081"

var err error

func forward(conn net.Conn) {
	client, err := net.Dial("tcp", Target)
	if err != nil {
		log.Fatalf("Dial failed: %v", err)
	}
	log.Printf("Connected to localhost %v\n", conn)
	go func() {
		defer client.Close()
		defer conn.Close()
		io.Copy(client, conn)
	}()
	go func() {
		defer client.Close()
		defer conn.Close()
		io.Copy(conn, client)
	}()
}

func Serve() error {
	if samConn, err = sam3.NewSAM(SamHost + ":" + SamPort); err != nil {
		return err
	}
	log.Println("SAM Bridge connection established.")
	if samKeys, err = samConn.NewKeys(); err != nil {
		return err
	}
	log.Println("Destination keys generated, tunnel name:", TunName, ".")
	if publishStream, err = samConn.NewStreamSession(TunName, samKeys, []string{"inbound.length=3", "outbound.length=3",
		"inbound.lengthVariance=1", "outbound.lengthVariance=1",
		"inbound.backupQuantity=3", "outbound.backupQuantity=3",
		"inbound.quantity=8", "outbound.quantity=8"}); err != nil {
		log.Println("Stream Creation error:", err.Error())
		return err
	}
	log.Println("SAM stream session established.")
	if publishListen, err = publishStream.Listen(); err != nil {
		return err
	}
	log.Println("Starting Listener.")
	I := i2pconv.I2pconv{}
	b, e := I.I2p64to32(string(samKeys.Addr()))
	if e != nil {
		return e
	}
	log.Println("SAM Listener created,", b+".b32.i2p")

	for {
		conn, err := publishListen.Accept()
		if err != nil {
			log.Fatalf("ERROR: failed to accept listener: %v", err)
		}
		log.Printf("Accepted connection %v\n", conn)
		go forward(conn)
	}
}
