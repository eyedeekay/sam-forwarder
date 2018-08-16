package samforwarderudp

import (
	"io"
	"log"
	"net"
	"os"
	"path/filepath"
	"strconv"
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
	SamKeys           sam3.I2PKeys
	publishConnection *sam3.DatagramSession
	clientConnection  net.PacketConn

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
	closeIdle          string
	closeIdleTime      string
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

// Target returns the host:port of the local service you want to forward to i2p
func (f *SAMSSUForwarder) Target() string {
	return f.TargetHost + ":" + f.TargetPort
}

func (f *SAMSSUForwarder) sam() string {
	return f.SamHost + ":" + f.SamPort
}

//func (f *SAMSSUForwarder) forward(conn net.Conn) {
func (f *SAMSSUForwarder) forward() {
	var err error
	p, _ := strconv.Atoi(f.TargetPort)
	f.clientConnection, err = net.DialUDP("udp", &net.UDPAddr{
		Port: p,
		IP:   net.ParseIP(f.TargetHost),
	}, nil)
	if err != nil {
		log.Fatalf("Dial failed: %v", err)
	}
	log.Printf("Connected to localhost %v\n", f.publishConnection)
	go func() {
		defer f.clientConnection.Close()
		defer f.publishConnection.Close()
		buffer := make([]byte, 1024)
		if size, addr, readerr := f.clientConnection.ReadFrom(buffer); readerr == nil {
			if size2, writeerr := f.publishConnection.WriteTo(buffer, addr); writeerr == nil {
				log.Printf("%q of %q bytes read", size, size2)
				log.Printf("%q of %q bytes written", size2, size)
			}
		}
	}()
	go func() {
		defer f.clientConnection.Close()
		defer f.publishConnection.Close()
		buffer := make([]byte, 1024)
		if size, addr, readerr := f.publishConnection.ReadFrom(buffer); readerr == nil {
			if size2, writeerr := f.clientConnection.WriteTo(buffer, addr); writeerr == nil {
				log.Printf("%q of %q bytes read", size, size2)
				log.Printf("%q of %q bytes written", size2, size)
			}
		}
	}()
}

//Base32 returns the base32 address where the local service is being forwarded
func (f *SAMSSUForwarder) Base32() string {
	return f.SamKeys.Addr().Base32()
}

//Base64 returns the base64 address where the local service is being forwarded
func (f *SAMSSUForwarder) Base64() string {
	return f.SamKeys.Addr().Base64()
}

//Serve starts the SAM connection and and forwards the local host:port to i2p
func (f *SAMSSUForwarder) Serve() error {
	sp, _ := strconv.Atoi(f.SamPort)
	if f.publishConnection, err = f.samConn.NewDatagramSession(f.TunName, f.SamKeys,
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
			"i2cp.closeOnIdle=" + f.closeIdle,
			"i2cp.closeIdleTime=" + f.closeIdleTime,
			f.accesslisttype(),
			f.accesslist(),
		}, sp); err != nil {
		log.Println("Stream Creation error:", err.Error())
		return err
	}
	log.Println("SAM stream session established.")
	log.Println("Starting Listener.")
	b := string(f.SamKeys.Addr().Base32())
	log.Println("SAM Keys created,", b)
	for {
		log.Printf("Accepted connection %v\n", f.publishConnection)
		go f.forward()
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
	s.encryptLeaseSet = "false"
	s.reduceIdle = "false"
	s.reduceIdleTime = "15"
	s.closeIdle = "false"
	s.closeIdleTime = "30"
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
	if s.SamKeys, err = s.samConn.NewKeys(); err != nil {
		return nil, err
	}
	log.Println("Destination keys generated, tunnel name:", s.TunName)
	if s.save {
		if _, err := os.Stat(filepath.Join(s.FilePath, s.TunName+".i2pkeys")); os.IsNotExist(err) {
			s.file, err = os.Create(filepath.Join(s.FilePath, s.TunName+".i2pkeys"))
			if err != nil {
				return nil, err
			}
			err = sam3.StoreKeysIncompat(s.SamKeys, s.file)
			if err != nil {
				return nil, err
			}
		}
		s.file, err = os.Open(filepath.Join(s.FilePath, s.TunName+".i2pkeys"))
		if err != nil {
			return nil, err
		}
		s.SamKeys, err = sam3.LoadKeysIncompat(s.file)
		if err != nil {
			return nil, err
		}
	}
	return &s, nil
}
