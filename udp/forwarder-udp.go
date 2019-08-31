package samforwarderudp

import (
	"io"
	"log"
	"net"
	//"os"
	//"path/filepath"
	"strconv"
	"strings"
	"time"
)

import (
	"github.com/eyedeekay/sam-forwarder/hashhash"
	"github.com/eyedeekay/sam-forwarder/i2pkeys"
	"github.com/eyedeekay/sam-forwarder/interface"
	"github.com/eyedeekay/sam3"
	"github.com/eyedeekay/sam3/i2pkeys"
    "github.com/eyedeekay/sam-forwarder/config"
)

//SAMSSUForwarder is a structure which automatically configured the forwarding of
//a local service to i2p over the SAM API.
type SAMSSUForwarder struct {
	SamHost string
	SamPort string
	TunName string
	Type    string

	TargetHost string
	TargetPort string

	samConn           *sam3.SAM
	SamKeys           i2pkeys.I2PKeys
	Hasher            *hashhash.Hasher
	publishConnection *sam3.DatagramSession
	clientConnection  net.PacketConn

	FilePath string
	file     io.ReadWriter
	save     bool
	up       bool

    // config
    Conf *i2ptunconf.Conf

	// samcatd options
	passfile string
	sigType  string

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

var err error

func (f *SAMSSUForwarder) Config() *i2ptunconf.Conf {
    return f.Conf
}

func (f *SAMSSUForwarder) GetType() string {
	return f.Type
}

func (f *SAMSSUForwarder) ID() string {
	return f.TunName
}

func (f *SAMSSUForwarder) Keys() i2pkeys.I2PKeys {
	return f.SamKeys
}

func (f *SAMSSUForwarder) Cleanup() {
	f.publishConnection.Close()
	f.clientConnection.Close()
	f.samConn.Close()
}

func (f *SAMSSUForwarder) Close() error {
	return nil
}

func (f *SAMSSUForwarder) print() []string {
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
		"i2cp.messageReliability=" + f.messageReliability,
		"i2cp.encryptLeaseSet=" + f.encryptLeaseSet,
		lsk, lspk, lspsk,
		f.accesslisttype(),
		f.accesslist(),
	}
}

func (f *SAMSSUForwarder) Props() map[string]string {
	r := make(map[string]string)
	print := f.print()
	print = append(print, "base32="+f.Base32())
	print = append(print, "base64="+f.Base64())
	print = append(print, "base32words="+f.Base32Readable())
	for _, prop := range print {
		k, v := sfi2pkeys.Prop(prop)
		r[k] = v
	}
	return r
}

func (f *SAMSSUForwarder) Print() string {
	var r string
	r += "name=" + f.TunName + "\n"
	r += "type=" + f.Type + "\n"
	r += "base32=" + f.Base32() + "\n"
	r += "base64=" + f.Base64() + "\n"
	r += "ssuserver\n"
	for _, s := range f.print() {
		r += s + "\n"
	}
	return strings.Replace(r, "\n\n", "\n", -1)
}

func (f *SAMSSUForwarder) Search(search string) string {
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

func (f *SAMSSUForwarder) leasesetsettings() (string, string, string) {
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

// Target returns the host:port of the local service you want to forward to i2p
func (f *SAMSSUForwarder) Target() string {
	return f.TargetHost + ":" + f.TargetPort
}

func (f *SAMSSUForwarder) sam() string {
	return f.SamHost + ":" + f.SamPort
}

//func (f *SAMSSUForwarder) forward(conn net.Conn) {
func (f *SAMSSUForwarder) forward() {
	Loop := false
	if f.clientConnection == nil {
		for !Loop {
			log.Println("Attempting to resolve local UDP Service to forward to I2P", f.Target())
			addr, err := net.ResolveUDPAddr("udp", f.Target())
			Loop = f.errSleep(err)
			log.Println("Attempting to dial resolved UDP Address", f.Target())
			f.clientConnection, err = net.DialUDP("udp", nil, addr)
			Loop = f.errSleep(err)
			log.Printf("Connected %v to localhost %v\n", f.publishConnection, f.clientConnection)
		}
	}
	buffer := make([]byte, 1024)
	if size, addr, readerr := f.clientConnection.ReadFrom(buffer); readerr == nil {
		if size2, writeerr := f.publishConnection.WriteTo(buffer, addr); writeerr == nil {
			log.Printf("%q of %q bytes read", size, size2)
			log.Printf("%q of %q bytes written", size2, size)
		} else {
			log.Println(writeerr)
		}
	} else {
		log.Println(readerr)
	}
	if size, addr, readerr := f.publishConnection.ReadFrom(buffer); readerr == nil {
		if size2, writeerr := f.clientConnection.WriteTo(buffer, addr); writeerr == nil {
			log.Printf("%q of %q bytes read", size, size2)
			log.Printf("%q of %q bytes written", size2, size)
		} else {
			log.Println(writeerr)
		}
	} else {
		log.Println(readerr)
	}
}

//Base32 returns the base32 address where the local service is being forwarded
func (f *SAMSSUForwarder) Base32() string {
	return f.SamKeys.Addr().Base32()
}

//Base32Readable returns the base32 address where the local service is being forwarded
func (f *SAMSSUForwarder) Base32Readable() string {
	b32 := strings.Replace(f.Base32(), ".b32.i2p", "", 1)
	rhash, _ := f.Hasher.Friendly(b32)
	return rhash + " " + strconv.Itoa(len(b32))
}

//Base64 returns the base64 address where the local service is being forwarded
func (f *SAMSSUForwarder) Base64() string {
	return f.SamKeys.Addr().Base64()
}

func (f *SAMSSUForwarder) errSleep(err error) bool {
	if err != nil {
		log.Printf("Dial failed: %v, waiting 5 minutes to try again\n", err)
		time.Sleep(5 * time.Minute)
		return false
	} else {
		return true
	}
}

//Serve starts the SAM connection and and forwards the local host:port to i2p
func (f *SAMSSUForwarder) Serve() error {
	var err error

	sp, _ := strconv.Atoi(f.SamPort)
	f.publishConnection, err = f.samConn.NewDatagramSession(
		f.TunName,
		f.SamKeys,
		f.print(),
		sp-1,
	)
	if err != nil {
		log.Println("Session Creation error:", err.Error())
		return err
	}

	log.Println("SAM datagram session established.")
	log.Println("Starting Forwarder.")
	log.Println("SAM Keys loaded,", f.Base32())
	log.Println("Human-readable hash:\n   ", f.Base32Readable())
	for {
		f.forward()
	}
}

func (s *SAMSSUForwarder) Load() (samtunnel.SAMTunnel, error) {
	s.samConn, err = sam3.NewSAM(s.sam())
	if err != nil {
		return nil, err
	}
	log.Println("SAM Bridge connection established.")
	if s.save {
		log.Println("Saving i2p keys")
	}
	if s.SamKeys, err = sfi2pkeys.Load(s.FilePath, s.TunName, s.passfile, s.samConn, s.save); err != nil {
		return nil, err
	}
	log.Println("Destination keys generated, tunnel name:", s.TunName)
	if s.save {
		if err := sfi2pkeys.Save(s.FilePath, s.TunName, s.passfile, s.SamKeys); err != nil {
			return nil, err
		}
		log.Println("Saved tunnel keys for", s.TunName)
	}
	s.Hasher, err = hashhash.NewHasher(len(strings.Replace(s.Base32(), ".b32.i2p", "", 1)))
	if err != nil {
		return nil, err
	}
	s.up = true
	return s, nil
}

func (f *SAMSSUForwarder) Up() bool {
	return f.up
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
	l, e := s.Load()
	if e != nil {
		return nil, e
	}
	return l.(*SAMSSUForwarder), nil
}
