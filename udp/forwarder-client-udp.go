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

//SAMSSUClientForwarder is a structure which automatically configured the forwarding of
//a local port to i2p over the SAM API.
type SAMSSUClientForwarder struct {
	SamHost string
	SamPort string
	TunName string

	TargetHost string
	TargetPort string

	samConn           *sam3.SAM
	SamKeys           sam3.I2PKeys
	connectStream     *sam3.DatagramSession
	dest              string
	addr              sam3.I2PAddr
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
	closeIdle          string
	closeIdleTime      string
	reduceIdle         string
	reduceIdleTime     string
	reduceIdleQuantity string

	//Streaming Library options
	accessListType string
	accessList     []string
}

func (f *SAMSSUClientForwarder) accesslisttype() string {
	if f.accessListType == "whitelist" {
		return "i2cp.enableAccessList=true"
	} else if f.accessListType == "blacklist" {
		return "i2cp.enableBlackList=true"
	} else if f.accessListType == "none" {
		return ""
	}
	return ""
}

func (f *SAMSSUClientForwarder) accesslist() string {
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
func (f *SAMSSUClientForwarder) Target() string {
	return f.TargetHost + ":" + f.TargetPort
}

func (f *SAMSSUClientForwarder) sam() string {
	return f.SamHost + ":" + f.SamPort
}

//Base32 returns the base32 address of the local destination
func (f *SAMSSUClientForwarder) Base32() string {
	return f.SamKeys.Addr().Base32()
}

//Base64 returns the base64 address of the local destination
func (f *SAMSSUClientForwarder) Base64() string {
	return f.SamKeys.Addr().Base64()
}

func (f *SAMSSUClientForwarder) forward(conn net.PacketConn) {
	var err error
	//p, _ := strconv.Atoi(f.TargetPort)
	sp, _ := strconv.Atoi(f.SamPort)
	if f.connectStream, err = f.samConn.NewDatagramSession(f.TunName, f.SamKeys,
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
		log.Fatal("Stream Creation error:", err.Error())
	}
	log.Println("SAM stream session established.")
	log.Printf("Connected to localhost %v\n", f.publishConnection)
	go func() {
		defer f.connectStream.Close()
		defer f.publishConnection.Close()
		buffer := make([]byte, 1024)
		if size, addr, readerr := f.publishConnection.ReadFrom(buffer); readerr == nil {
			if size2, writeerr := f.connectStream.WriteTo(buffer, addr); writeerr == nil {
				log.Printf("%q of %q bytes read", size, size2)
				log.Printf("%q of %q bytes written", size2, size)
			}
		}
	}()
	go func() {
		defer f.connectStream.Close()
		defer f.publishConnection.Close()
		buffer := make([]byte, 1024)
		if size, addr, readerr := f.connectStream.ReadFrom(buffer); readerr == nil {
			if size2, writeerr := f.publishConnection.WriteTo(buffer, addr); writeerr == nil {
				log.Printf("%q of %q bytes read", size, size2)
				log.Printf("%q of %q bytes written", size2, size)
			}
		}
	}()
}

//Serve starts the SAM connection and and forwards the local host:port to i2p
func (f *SAMSSUClientForwarder) Serve(dest string) error {
	f.dest = dest
	if f.addr, err = f.samConn.Lookup(f.dest); err != nil {
		return err
	}

	for {
		//f.packetConn, err := f.publishConnection.Accept()
		if err != nil {
			return err
		}
		log.Println("Forwarding client to i2p address:", f.addr.Base32())
		go f.forward(f.publishConnection)
	}
}

//NewSAMSSUClientForwarderFromOptions makes a new SAM forwarder with default options, accepts host:port arguments
func NewSAMSSUClientForwarderFromOptions(opts ...func(*SAMSSUClientForwarder) error) (*SAMSSUClientForwarder, error) {
	var s SAMSSUClientForwarder
	s.SamHost = "127.0.0.1"
	s.SamPort = "7656"
	s.FilePath = ""
	s.save = false
	s.TargetHost = "127.0.0.1"
	s.TargetPort = "0"
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
	s.dest = "none"
	for _, o := range opts {
		if err := o(&s); err != nil {
			return nil, err
		}
	}
	if s.publishConnection, err = net.ListenPacket("udp", s.TargetHost+":"+s.TargetPort); err != nil {
		return nil, err
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
