package samforwarderudp

import (
	"io"
	"log"
	"net"
	"os"
	"path/filepath"
	"strings"
)

import (
	"github.com/kpetku/sam3"
)

//SAMSSUForwarder is a structure which automatically configured the forwarding of
//a local service to i2p over the SAM API.
type SAMSSUForwarder struct {
	SamHost string
	SamPort string
	TunName string

	TargetHost string
	TargetPort string

	samConn           *sam3.SAM
	samKeys           sam3.I2PKeys
	publishSession    *sam3.DatagramSession
	publishConnection net.PacketConn

	FilePath string
	file     io.ReadWriter
	save     bool

	// I2CP options
	encryptLeaseSet    string
	inAllowZeroHop     string
	outAllowZeroHop    string
	inLength           string
	outLength          string
	inQuantity         string
	outQuantity        string
	inVariance         string
	outVariance        string
	inBackupQuantity   string
	outBackupQuantity  string
	useCompression     string
	reduceIdle         string
	reduceIdleTime     string
	reduceIdleQuantity string

	//Streaming Library options
	accessListType string
	accessList     []string
}

var err error

func (f *SAMSSUForwarder) accesslisttype() string {
	if f.accessListType == "whitelist" {
		return "i2cp.enableAccessList=true"
	} else if f.accessListType == "blacklist" {
		return "i2cp.enableBlackList=true"
	} else if f.accessListType == "none" {
		return ""
	}
	return ""
}

func (f *SAMSSUForwarder) accesslist() string {
	if f.accessListType != "" && len(f.accessList) > 0 {
		r := ""
		for _, s := range f.accessList {
			r += s + ","
		}
		return "i2cp.accessList=" + strings.TrimSuffix(r, ",")
	}
	return ""
}

func (f *SAMSSUForwarder) Target() string {
	return f.TargetHost + ":" + f.TargetPort
}

func (f *SAMSSUForwarder) sam() string {
	return f.SamHost + ":" + f.SamPort
}

func (f *SAMSSUForwarder) forward(conn net.PacketConn) {
	client, err := net.Dial("udp", f.Target())
	if err != nil {
		log.Fatalf("Dial failed: %v", err)
	}
	log.Printf("Connected to localhost %v\n", conn)
	go func() {
		defer client.Close()
		defer conn.Close()
		//io.Copy(client, conn)
	}()
	go func() {
		defer client.Close()
		defer conn.Close()
		//io.Copy(conn, client)
	}()
}

//Base32 returns the base32 address where the local service is being forwarded
func (f *SAMSSUForwarder) Base32() string {
	return f.samKeys.Addr().Base32()
}

//Serve starts the SAM connection and and forwards the local host:port to i2p
func (f *SAMSSUForwarder) Serve() error {
	if f.publishSession, err = f.samConn.NewDatagramSession(f.TunName, f.samKeys,
		[]string{
			"inbound.length=" + f.inLength,
			"outbound.length=" + f.outLength,
			"inbound.lengthVariance=" + f.inVariance,
			"outbound.lengthVariance=" + f.outVariance,
			"inbound.backupQuantity=" + f.inBackupQuantity,
			"outbound.backupQuantity=" + f.outBackupQuantity,
			"inbound.quantity=" + f.inQuantity,
			"outbound.quantity=" + f.outQuantity,
			"inbound.allowZeroHop=" + f.inAllowZeroHop,
			"outbound.allowZeroHop=" + f.outAllowZeroHop,
			"i2cp.encryptLeaseSet=" + f.encryptLeaseSet,
			"i2cp.gzip=" + f.useCompression,
			"i2cp.reduceOnIdle=" + f.reduceIdle,
			"i2cp.reduceIdleTime=" + f.reduceIdleTime,
			"i2cp.reduceQuantity=" + f.reduceIdleQuantity,
			f.accesslisttype(),
			f.accesslist(),
		}, 0); err != nil {
		log.Println("Stream Creation error:", err.Error())
		return err
	}
	log.Println("SAM stream session established.")
	/*if f.publishListen, err = f.publishSession.Listen(); err != nil {
		return err
	}*/
	log.Println("Starting Listener.")
	b := string(f.samKeys.Addr().Base32())
	log.Println("SAM Listener created,", b+".b32.i2p")

	for {
		//conn, err := f.publishListen.Accept()
		if err != nil {
			log.Fatalf("ERROR: failed to accept listener: %v", err)
		}
		log.Printf("Accepted connection %v\n", f.publishSession)
		go f.forward(f.publishSession)
	}
}

//NewSAMSSUForwarder makes a new SAM forwarder with default options, accepts host:port arguments
func NewSAMSSUForwarder(host, port string) (*SAMSSUForwarder, error) {
	return NewSAMSSUForwarderFromOptions(SetHost(host), SetPort(port))
}

//NewSAMSSUForwarderFromOptions makes a new SAM forwarder with default options, accepts host:port arguments
func NewSAMSSUForwarderFromOptions(opts ...func(*SAMSSUForwarder) error) (*SAMSSUForwarder, error) {
	var s SAMSSUForwarder
	s.SamHost = "127.0.0.1"
	s.SamPort = "7656"
	s.FilePath = ""
	s.save = false
	s.TargetHost = "127.0.0.1"
	s.TargetPort = "8081"
	s.TunName = "samSSUForwarder"
	s.inLength = "3"
	s.outLength = "3"
	s.inQuantity = "8"
	s.outQuantity = "8"
	s.inVariance = "1"
	s.outVariance = "1"
	s.inBackupQuantity = "3"
	s.outBackupQuantity = "3"
	s.inAllowZeroHop = "false"
	s.outAllowZeroHop = "false"
	s.useCompression = "true"
	s.reduceIdle = "false"
	s.reduceIdleTime = "15"
	s.reduceIdleQuantity = "4"
	for _, o := range opts {
		if err := o(&s); err != nil {
			return nil, err
		}
	}
	if s.samConn, err = sam3.NewSAM(s.sam()); err != nil {
		return nil, err
	}
	log.Println("SAM Bridge connection established.")
	if s.samKeys, err = s.samConn.NewKeys(); err != nil {
		return nil, err
	}
	log.Println("Destination keys generated, tunnel name:", s.TunName)
	if s.save {
		if _, err := os.Stat(filepath.Join(s.FilePath, s.TunName+".i2pkeys")); os.IsNotExist(err) {
			s.file, err = os.Create(filepath.Join(s.FilePath, s.TunName+".i2pkeys"))
			if err != nil {
				return nil, err
			}
			err = sam3.StoreKeysIncompat(s.samKeys, s.file)
			if err != nil {
				return nil, err
			}
		}
		s.file, err = os.Open(filepath.Join(s.FilePath, s.TunName+".i2pkeys"))
		if err != nil {
			return nil, err
		}
		s.samKeys, err = sam3.LoadKeysIncompat(s.file)
		if err != nil {
			return nil, err
		}
	}
	return &s, nil
}
