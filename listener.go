package samlistener

import (
    "io"
    "log"
    "net"
    "os"
)

import (
	"github.com/kpetku/sam3"
)

var samhost = "127.0.0.1"
var samport = "7656"
var tunname = "ephtun"

var samConn *sam3.SAM
var samKeys sam3.I2PKeys

var publishStream     *sam3.StreamSession
var publishListen     *sam3.StreamListener
var publishConnection net.Conn

var err error

func forward(conn net.Conn) {
    client, err := net.Dial("tcp", os.Args[1])
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
    if samConn, err = sam3.NewSAM(samhost+":"+samport); err != nil {
		return err
	}
	log.Println("SAM Bridge connection established")
	if samKeys, err = samConn.NewKeys(); err != nil {
		return err
	}
	log.Println("Destination keys generated, tunnel name:", tunname)
	if publishStream, err = samConn.NewStreamSession(tunname, samKeys, sam3.Options_Small); err != nil {
		return err
	}
    log.Println("SAM stream session established")
	if publishListen, err = publishStream.Listen(); err != nil {
		return err
	}
	log.Println("SAM Listener created")

    for {
        conn, err := publishListen.Accept()
        if err != nil {
            log.Fatalf("ERROR: failed to accept listener: %v", err)
        }
        log.Printf("Accepted connection %v\n", conn)
        go forward(conn)
    }
}
