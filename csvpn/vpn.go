package i2pvpn

import (
	"bytes"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/eyedeekay/sam-forwarder/udp"
	"github.com/eyedeekay/sam3"
	"github.com/eyedeekay/udptunnel/tunnel"
)

type SAMClientServerVPN struct {
	// i2p tunnel
	I2PTunnel samforwarderudp.SAMSSUForwarder
	// VPN tunnel
	VPNTunnel  udptunnel.Tunnel
	ServerMode bool
	samConn    *sam3.SAM
	SamKeys    sam3.I2PKeys
	SamHost    string
	SamPort    string
	TunName    string
	Type       string

	TargetHost string
	TargetPort string

	FilePath string
	file     io.ReadWriter
	save     bool

	// samcatd options
	passfile string

	// I2CP options
	encryptLeaseSet           string
	leaseSetKey               string
	leaseSetPrivateKey        string
	leaseSetPrivateSigningKey string
	inAllowZeroHop            string
	outAllowZeroHop           string
	inLength                  string
	outLength                 string
	inQuantity                string
	outQuantity               string
	inVariance                string
	outVariance               string
	inBackupQuantity          string
	outBackupQuantity         string
	fastRecieve               string
	useCompression            string
	messageReliability        string
	closeIdle                 string
	closeIdleTime             string
	reduceIdle                string
	reduceIdleTime            string
	reduceIdleQuantity        string

	//Streaming Library options
	accessListType string
	accessList     []string
}

func (f *SAMClientServerVPN) sam() string {
	return f.SamHost + ":" + f.SamPort
}

// Target returns the host:port of the local service you want to forward to i2p
func (f *SAMClientServerVPN) Target() string {
	return f.TargetHost + ":" + f.TargetPort
}

func NewSAMClientServerVPNFromOptions(opts ...func(*SAMClientServerVPN) error) (*SAMClientServerVPN, error) {
	var s SAMClientServerVPN
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
	s.fastRecieve = "false"
	s.useCompression = "true"
	s.encryptLeaseSet = "false"
	s.leaseSetKey = ""
	s.leaseSetPrivateKey = ""
	s.leaseSetPrivateSigningKey = ""
	s.reduceIdle = "false"
	s.reduceIdleTime = "300000"
	s.closeIdle = "false"
	s.closeIdleTime = "300000"
	s.reduceIdleQuantity = "4"
	s.Type = "udpserver"
	s.messageReliability = "none"
	s.passfile = ""
	for _, o := range opts {
		if err := o(&s); err != nil {
			return nil, err
		}
	}
	var err error
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
			err = Encrypt(filepath.Join(s.FilePath, s.TunName+".i2pkeys"), s.passfile)
			if err != nil {
				return nil, err
			}
		}
		s.file, err = os.Open(filepath.Join(s.FilePath, s.TunName+".i2pkeys"))
		if err != nil {
			return nil, err
		}
		err = Decrypt(filepath.Join(s.FilePath, s.TunName+".i2pkeys"), s.passfile)
		if err != nil {
			return nil, err
		}
		s.SamKeys, err = sam3.LoadKeysIncompat(s.file)
		if err != nil {
			return nil, err
		}
		err = Encrypt(filepath.Join(s.FilePath, s.TunName+".i2pkeys"), s.passfile)
		if err != nil {
			return nil, err
		}
	}
	var logBuf bytes.Buffer
	Logger = log.New(io.MultiWriter(os.Stderr, &logBuf), "", log.Ldate|log.Ltime|log.Lshortfile)
	s.VPNTunnel = udptunnel.NewTunnel(s.ServerMode, s.TunName, "10.76.0.2", s.Target(), "", []uint16{},
		"i2pvpn", time.Duration(time.Second*300), Logger)
	return &s, nil
}
