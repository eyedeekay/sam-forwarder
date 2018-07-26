package samforwarder

import (
	"io"
	"log"
	"net"
)

import (
	"github.com/kpetku/sam3"
)

//SAMForwarder is a structure which automatically configured the forwarding of
//a local service to i2p over the SAM API.
type SAMForwarder struct {
    SamHost string
    SamPort string
    TunName string

    TargetHost string
    TargetPort string

    samConn *sam3.SAM
    samKeys sam3.I2PKeys
    publishStream *sam3.StreamSession
    publishListen *sam3.StreamListener
    publishConnection net.Conn

    //
    savePath string
    save bool
}

var err error

func (f *SAMForwarder) target() string {
    return f.TargetHost+":"+f.TargetPort
}

func (f *SAMForwarder) sam() string {
    return f.SamHost+":"+f.SamPort
}

func (f *SAMForwarder) forward(conn net.Conn) {
	client, err := net.Dial("tcp", f.target())
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

//Base32 returns the base32 address where the local service is being forwarded
func (f *SAMForwarder) Base32() string {
    return f.samKeys.Addr().Base32()
}

//Serve starts the SAM connection and and forwards the local host:port to i2p
func (f *SAMForwarder) Serve() error {
	if f.samConn, err = sam3.NewSAM(f.sam()); err != nil {
		return err
	}
	log.Println("SAM Bridge connection established.")
	if f.samKeys, err = f.samConn.NewKeys(); err != nil {
		return err
	}
	log.Println("Destination keys generated, tunnel name:", f.TunName, ".")
	if f.publishStream, err = f.samConn.NewStreamSession(f.TunName, f.samKeys,
        []string{"inbound.length=3", "outbound.length=3",
		"inbound.lengthVariance=1", "outbound.lengthVariance=1",
		"inbound.backupQuantity=3", "outbound.backupQuantity=3",
		"inbound.quantity=8", "outbound.quantity=8"}); err != nil {
		log.Println("Stream Creation error:", err.Error())
		return err
	}
	log.Println("SAM stream session established.")
	if f.publishListen, err = f.publishStream.Listen(); err != nil {
		return err
	}
	log.Println("Starting Listener.")
	b := string(f.samKeys.Addr().Base32())
	log.Println("SAM Listener created,", b+".b32.i2p")

	for {
		conn, err := f.publishListen.Accept()
		if err != nil {
			log.Fatalf("ERROR: failed to accept listener: %v", err)
		}
		log.Printf("Accepted connection %v\n", conn)
		go f.forward(conn)
	}
}

//NewSAMForwarder makes a new SAM forwarder with default options, accepts host:port arguments
func NewSAMForwarder(host, port string) (*SAMForwarder, error) {
    return NewSAMForwarderFromOptions(SetHost(host), SetPort(port))
}

//NewSAMForwarderFromOptions makes a new SAM forwarder with default options, accepts host:port arguments
func NewSAMForwarderFromOptions(opts ...func(*SAMForwarder) error) (*SAMForwarder, error) {
    var s SAMForwarder
    s.SamHost = "127.0.0.1"
    s.SamPort = "7656"
    s.TargetHost = "127.0.0.1"
    s.TargetPort = "8081"
    s.TunName = "samForwarder"
    for _, o := range opts {
		if err := o(&s); err != nil {
			return nil, err
		}
	}
    return &s, nil
}
