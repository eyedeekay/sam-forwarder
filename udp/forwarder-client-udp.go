package samforwarderudp

import (
	"io"
	"log"
	"net"
	//"os"
	//"path/filepath"
	"strconv"
	"strings"
)

import (
	"github.com/eyedeekay/sam-forwarder/i2pkeys"
	"github.com/eyedeekay/sam3"
)

//SAMSSUClientForwarder is a structure which automatically configured the forwarding of
//a local port to i2p over the SAM API.
type SAMSSUClientForwarder struct {
	SamHost string
	SamPort string
	TunName string
	Type    string

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

func (f SAMSSUClientForwarder) Cleanup() {
	f.publishConnection.Close()
	f.connectStream.Close()
	f.samConn.Close()
}

func (f SAMSSUClientForwarder) print() []string {
	lsk, lspk, lspsk := f.leasesetsettings()
	return []string{
		//f.targetForPort443(),
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
		"i2cp.fastRecieve=" + f.fastRecieve,
		"i2cp.gzip=" + f.useCompression,
		"i2cp.reduceOnIdle=" + f.reduceIdle,
		"i2cp.reduceIdleTime=" + f.reduceIdleTime,
		"i2cp.reduceQuantity=" + f.reduceIdleQuantity,
		"i2cp.closeOnIdle=" + f.closeIdle,
		"i2cp.closeIdleTime=" + f.closeIdleTime,
		"i2cp.messageReliability" + f.messageReliability,
		"i2cp.encryptLeaseSet=" + f.encryptLeaseSet,
		lsk, lspk, lspsk,
		f.accesslisttype(),
		f.accesslist(),
	}
}

func (f SAMSSUClientForwarder) Print() string {
	var r string
	r += "name=" + f.TunName + "\n"
	r += "type=" + f.Type + "\n"
	r += "base32=" + f.Base32() + "\n"
	r += "base64=" + f.Base64() + "\n"
	r += "dest=" + f.dest + "\n"
	r += "ssuclient\n"
	for _, s := range f.print() {
		r += s + "\n"
	}
	return strings.Replace(r, "\n\n", "\n", -1)
}

func (f SAMSSUClientForwarder) Search(search string) string {
	terms := strings.Split(search, ",")
	if search == "" {
		return f.Print()
	}
	for _, value := range terms {
		if !strings.Contains(f.Print(), value) {
			return ""
		}
	}
	return f.Print()
}

func (f SAMSSUClientForwarder) accesslisttype() string {
	if f.accessListType == "whitelist" {
		return "i2cp.enableAccessList=true"
	} else if f.accessListType == "blacklist" {
		return "i2cp.enableBlackList=true"
	} else if f.accessListType == "none" {
		return ""
	}
	return ""
}

func (f SAMSSUClientForwarder) accesslist() string {
	if f.accessListType != "" && len(f.accessList) > 0 {
		r := ""
		for _, s := range f.accessList {
			r += s + ","
		}
		return "i2cp.accessList=" + strings.TrimSuffix(r, ",")
	}
	return ""
}

func (f SAMSSUClientForwarder) leasesetsettings() (string, string, string) {
	var r, s, t string
	if f.leaseSetKey != "" {
		r = "i2cp.leaseSetKey=" + f.leaseSetKey
	}
	if f.leaseSetPrivateKey != "" {
		s = "i2cp.leaseSetPrivateKey=" + f.leaseSetPrivateKey
	}
	if f.leaseSetPrivateSigningKey != "" {
		t = "i2cp.leaseSetPrivateSigningKey=" + f.leaseSetPrivateSigningKey
	}
	return r, s, t
}

// Destination returns the destination of the i2p service you want to forward locally
func (f SAMSSUClientForwarder) Destination() string {
	return f.addr.Base32()
}

// Target returns the host:port of the local service you want to forward to i2p
func (f SAMSSUClientForwarder) Target() string {
	return f.TargetHost + ":" + f.TargetPort
}

func (f SAMSSUClientForwarder) sam() string {
	return f.SamHost + ":" + f.SamPort
}

//Base32 returns the base32 address of the local destination
func (f SAMSSUClientForwarder) Base32() string {
	return f.SamKeys.Addr().Base32()
}

//Base64 returns the base64 address of the local destination
func (f SAMSSUClientForwarder) Base64() string {
	return f.SamKeys.Addr().Base64()
}

func (f SAMSSUClientForwarder) forward(conn net.PacketConn) {
	var err error
	//p, _ := strconv.Atoi(f.TargetPort)
	sp, _ := strconv.Atoi(f.SamPort)
	if f.connectStream, err = f.samConn.NewDatagramSession(f.TunName, f.SamKeys,
		f.print(), sp); err != nil {
		log.Fatal("Stream Creation error:", err.Error())
	}
	log.Println("SAM stream session established.")
	log.Printf("Connected to localhost %v\n", f.publishConnection)
	go func() {
		defer f.connectStream.Close()
		defer f.publishConnection.Close()
		buffer := make([]byte, 1024)
		if size, _, readerr := f.publishConnection.ReadFrom(buffer); readerr == nil {
			if size2, writeerr := f.connectStream.WriteTo(buffer, f.addr); writeerr == nil {
				log.Printf("%q of %q bytes read", size, size2)
				log.Printf("%q of %q bytes written", size2, size)
			}
		}
	}()
	go func() {
		defer f.connectStream.Close()
		defer f.publishConnection.Close()
		buffer := make([]byte, 1024)
		if size, _, readerr := f.connectStream.ReadFrom(buffer); readerr == nil {
			if size2, writeerr := f.publishConnection.WriteTo(buffer, f.addr); writeerr == nil {
				log.Printf("%q of %q bytes read", size, size2)
				log.Printf("%q of %q bytes written", size2, size)
			}
		}
	}()
}

//Serve starts the SAM connection and and forwards the local host:port to i2p
func (f SAMSSUClientForwarder) Serve() error {
	if f.addr, err = f.samConn.Lookup(f.dest); err != nil {
		return err
	}

	for {
		p, _ := strconv.Atoi(f.TargetPort)
		f.publishConnection, err = net.DialUDP("udp", &net.UDPAddr{
			Port: p,
			IP:   net.ParseIP(f.TargetHost),
		}, nil)
		if err != nil {
			return err
		}
		log.Println("Forwarding client to i2p address:", f.addr.Base32())
		f.forward(f.publishConnection)
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
	s.inQuantity = "2"
	s.outQuantity = "2"
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
	s.reduceIdleTime = "15"
	s.closeIdle = "false"
	s.closeIdleTime = "30"
	s.reduceIdleQuantity = "4"
	s.dest = "none"
	s.Type = "udpclient"
	s.messageReliability = "none"
	s.passfile = ""
	s.dest = "i2p-projekt.i2p"
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
	if s.save {
		log.Println("Saving i2p keys")
	}
	if s.SamKeys, err = i2pkeys.Load(s.FilePath, s.TunName, s.passfile, s.samConn, s.save); err != nil {
		return nil, err
	}
	log.Println("Destination keys generated, tunnel name:", s.TunName)
	if s.save {
		if err := i2pkeys.Save(s.FilePath, s.TunName, s.passfile, s.SamKeys); err != nil {
			return nil, err
		}
		log.Println("Saved tunnel keys for", s.TunName)
	}
	return &s, nil
}
