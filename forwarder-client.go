package samforwarder

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

// SAMClientForwarder is a tcp proxy that automatically forwards ports to i2p
type SAMClientForwarder struct {
	SamHost string
	SamPort string
	TunName string

	TargetHost string
	TargetPort string

	samConn           *sam3.SAM
	SamKeys           sam3.I2PKeys
	connectStream     *sam3.StreamSession
	dest              string
	addr              sam3.I2PAddr
	publishConnection net.Listener

	FilePath string
	file     io.ReadWriter
	save     bool

	// I2CP options
	encryptLeaseSet    string
	LeaseSetKeys       *sam3.I2PKeys
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

//var err error

func (f *SAMClientForwarder) accesslisttype() string {
	if f.accessListType == "whitelist" {
		return "i2cp.enableAccessList=true"
	} else if f.accessListType == "blacklist" {
		return "i2cp.enableBlackList=true"
	} else if f.accessListType == "none" {
		return ""
	}
	return ""
}

func (f *SAMClientForwarder) accesslist() string {
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
func (f *SAMClientForwarder) Target() string {
	return f.TargetHost + ":" + f.TargetPort
}

func (f *SAMClientForwarder) sam() string {
	return f.SamHost + ":" + f.SamPort
}

//Base32 returns the base32 address of the local destination
func (f *SAMClientForwarder) Base32() string {
	return f.SamKeys.Addr().Base32()
}

//Base64 returns the base64 address of the local destiantion
func (f *SAMClientForwarder) Base64() string {
	return f.SamKeys.Addr().Base64()
}

func (f *SAMClientForwarder) forward(conn net.Conn) {
	client, err := f.connectStream.DialI2P(f.addr)
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

//Serve starts the SAM connection and and forwards the local host:port to i2p
func (f *SAMClientForwarder) Serve(dest string) error {
	f.dest = dest
	if f.addr, err = f.samConn.Lookup(f.dest); err != nil {
		return err
	}
	if f.connectStream, err = f.samConn.NewStreamSession(f.TunName, f.SamKeys,
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
		}); err != nil {
		log.Println("Stream Creation error:", err.Error())
		return err
	}
	log.Println("SAM stream session established.")

	for {
		conn, err := f.publishConnection.Accept()
		if err != nil {
			return err
		}
		log.Println("Forwarding client to i2p address:", f.addr.Base32())
		go f.forward(conn)
	}
}

//NewSAMClientForwarder makes a new SAM forwarder with default options, accepts host:port arguments
func NewSAMClientForwarder(host, port string) (*SAMClientForwarder, error) {
	return NewSAMClientForwarderFromOptions(SetClientHost(host), SetClientPort(port))
}

//NewSAMClientForwarderFromOptions makes a new SAM forwarder with default options, accepts host:port arguments
func NewSAMClientForwarderFromOptions(opts ...func(*SAMClientForwarder) error) (*SAMClientForwarder, error) {
	var s SAMClientForwarder
	s.SamHost = "127.0.0.1"
	s.SamPort = "7656"
	s.FilePath = ""
	s.save = false
	s.TargetHost = "127.0.0.1"
	s.TargetPort = "0"
	s.TunName = "samForwarder"
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
	s.encryptLeaseSet = "false"
	s.useCompression = "true"
	s.reduceIdle = "false"
	s.reduceIdleTime = "15"
	s.reduceIdleQuantity = "4"
	s.closeIdle = "false"
	s.closeIdleTime = "30"
	s.dest = "none"
	for _, o := range opts {
		if err := o(&s); err != nil {
			return nil, err
		}
	}
	if s.publishConnection, err = net.Listen("tcp", s.TargetHost+":"+s.TargetPort); err != nil {
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
