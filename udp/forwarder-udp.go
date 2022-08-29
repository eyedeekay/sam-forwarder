package samforwarderudp

import (
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"strings"
	"time"

	i2ptunconf "github.com/eyedeekay/sam-forwarder/config"
	"github.com/eyedeekay/sam-forwarder/hashhash"
	sfi2pkeys "github.com/eyedeekay/sam-forwarder/i2pkeys"
	samoptions "github.com/eyedeekay/sam-forwarder/options"

	"github.com/eyedeekay/i2pkeys"
	samtunnel "github.com/eyedeekay/sam-forwarder/interface"
	"github.com/eyedeekay/sam3"
)

//SAMDGForwarder is a structure which automatically configured the forwarding of
//a local service to i2p over the SAM API.
type SAMDGForwarder struct {
	*i2ptunconf.Conf

	samConn           *sam3.SAM
	SamKeys           i2pkeys.I2PKeys
	Hasher            *hashhash.Hasher
	publishConnection *sam3.DatagramSession
	clientConnection  net.PacketConn

	file io.ReadWriter
	up   bool
}

var err error

func (f *SAMDGForwarder) Config() *i2ptunconf.Conf {
	if f.Conf == nil {
		f.Conf = i2ptunconf.NewI2PBlankTunConf()
	}
	return f.Conf
}

func (f *SAMDGForwarder) GetType() string {
	return f.Config().Type
}

func (f *SAMDGForwarder) ID() string {
	return f.Config().TunName
}

func (f *SAMDGForwarder) Keys() i2pkeys.I2PKeys {
	return f.SamKeys
}

func (f *SAMDGForwarder) Cleanup() {
	f.publishConnection.Close()
	f.clientConnection.Close()
	f.samConn.Close()
}

func (f *SAMDGForwarder) Close() error {
	return nil
}

func (f *SAMDGForwarder) print() []string {
	/*lsk, lspk, lspsk := f.leasesetsettings()
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
	}*/
	return f.Config().PrintSlice()
}

func (f *SAMDGForwarder) Props() map[string]string {
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

func (f *SAMDGForwarder) Print() string {
	var r string
	r += "name=" + f.Config().TunName + "\n"
	r += "type=" + f.Config().Type + "\n"
	r += "base32=" + f.Base32() + "\n"
	r += "base64=" + f.Base64() + "\n"
	r += "ssuserver\n"
	for _, s := range f.print() {
		r += s + "\n"
	}
	return strings.Replace(r, "\n\n", "\n", -1)
}

func (f *SAMDGForwarder) Search(search string) string {
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

func (f *SAMDGForwarder) accesslisttype() string {
	if f.Config().AccessListType == "allowlist" {
		return "i2cp.enableAccessList=true"
	} else if f.Config().AccessListType == "blocklist" {
		return "i2cp.enableBlackList=true"
	} else if f.Config().AccessListType == "none" {
		return ""
	}
	return ""
}

func (f *SAMDGForwarder) accesslist() string {
	if f.Config().AccessListType != "" && len(f.Config().AccessList) > 0 {
		r := ""
		for _, s := range f.Config().AccessList {
			r += s + ","
		}
		return "i2cp.accessList=" + strings.TrimSuffix(r, ",")
	}
	return ""
}

func (f *SAMDGForwarder) leasesetsettings() (string, string, string) {
	var r, s, t string
	if f.Config().LeaseSetKey != "" {
		r = "i2cp.leaseSetKey=" + f.Config().LeaseSetKey
	}
	if f.Config().LeaseSetPrivateKey != "" {
		s = "i2cp.leaseSetPrivateKey=" + f.Config().LeaseSetPrivateKey
	}
	if f.Config().LeaseSetPrivateSigningKey != "" {
		t = "i2cp.leaseSetPrivateSigningKey=" + f.Config().LeaseSetPrivateSigningKey
	}
	return r, s, t
}

// Target returns the host:port of the local service you want to forward to i2p
func (f *SAMDGForwarder) Target() string {
	return f.Config().TargetHost + ":" + f.Config().TargetPort
}

func (f *SAMDGForwarder) sam() string {
	return f.Config().SamHost + ":" + f.Config().SamPort
}

//func (f *SAMDGForwarder) forward(conn net.Conn) {
func (f *SAMDGForwarder) forward() {
	Loop := false
	if f.clientConnection == nil {
		for !Loop {
			log.Println("Attempting to resolve local UDP Service to forward to I2P", f.Config().Target())
			addr, err := net.ResolveUDPAddr("udp", f.Config().Target())
			Loop = f.errSleep(err)
			log.Println("Attempting to dial resolved UDP Address", f.Config().Target())
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
func (f *SAMDGForwarder) Base32() string {
	return f.SamKeys.Addr().Base32()
}

//Base32Readable returns the base32 address where the local service is being forwarded
func (f *SAMDGForwarder) Base32Readable() string {
	b32 := strings.Replace(f.Base32(), ".b32.i2p", "", 1)
	rhash, _ := f.Hasher.Friendly(b32)
	return rhash + " " + strconv.Itoa(len(b32))
}

//Base64 returns the base64 address where the local service is being forwarded
func (f *SAMDGForwarder) Base64() string {
	return f.SamKeys.Addr().Base64()
}

func (f *SAMDGForwarder) errSleep(err error) bool {
	if err != nil {
		log.Printf("Dial failed: %v, waiting 5 minutes to try again\n", err)
		time.Sleep(5 * time.Minute)
		return false
	} else {
		return true
	}
}

//Serve starts the SAM connection and and forwards the local host:port to i2p
func (f *SAMDGForwarder) Serve() error {
	var err error

	sp, err := strconv.Atoi(f.Config().SamPort)
	f.publishConnection, err = f.samConn.NewDatagramSession(
		f.Config().TunName,
		f.SamKeys,
		f.print(),
		sp-1,
	)
	if err != nil {
		return fmt.Errorf("Session creation Error%s", err)
	}

	log.Println("SAM datagram session established.")
	log.Println("Starting Forwarder.")
	log.Println("SAM Keys loaded,", f.Base32())
	log.Println("Human-readable hash:\n   ", f.Base32Readable())
	for {
		f.forward()
	}
}

func (s *SAMDGForwarder) Load() (samtunnel.SAMTunnel, error) {
	s.samConn, err = sam3.NewSAM(s.sam())
	if err != nil {
		return nil, fmt.Errorf("SAM connection error%s", err)
	}
	log.Println("SAM Bridge connection established.")
	if s.Config().SaveFile {
		log.Println("Saving i2p keys")
	}
	if s.SamKeys, err = sfi2pkeys.Load(s.Config().FilePath, s.Config().TunName, s.Config().KeyFilePath, s.samConn, s.Config().SaveFile); err != nil {
		return nil, fmt.Errorf("SAM key load/generate error%s", err)
	}
	log.Println("Destination keys generated, tunnel name:", s.Config().TunName)
	if s.Config().SaveFile {
		if err := sfi2pkeys.Save(s.Config().FilePath, s.Config().TunName, s.Config().KeyFilePath, s.SamKeys); err != nil {
			return nil, fmt.Errorf("SAM key save error%s", err)
		}
		log.Println("Saved tunnel keys for", s.Conf.TunName, "in", s.Conf.FilePath)
	}
	s.Hasher, err = hashhash.NewHasher(len(strings.Replace(s.Base32(), ".b32.i2p", "", 1)))
	if err != nil {
		return nil, fmt.Errorf("Word-hash generation error%s", err)
	}
	s.up = true
	return s, nil
}

func (f *SAMDGForwarder) Up() bool {
	return f.up
}

//NewSAMDGForwarder makes a new SAM forwarder with default options, accepts host:port arguments
func NewSAMDGForwarder(host, port string) (samtunnel.SAMTunnel, error) {
	return NewSAMDGForwarderFromOptions(samoptions.SetHost(host), samoptions.SetPort(port))
}

//NewSAMDGForwarderFromOptions makes a new SAM forwarder with default options, accepts host:port arguments
func NewSAMDGForwarderFromOptions(opts ...func(samtunnel.SAMTunnel) error) (*SAMDGForwarder, error) {
	var s SAMDGForwarder
	s.Conf = i2ptunconf.NewI2PBlankTunConf()
	s.Conf.Type = "udpserver"
	for _, o := range opts {
		if err := o(&s); err != nil {
			return nil, fmt.Errorf("Option Setting Error: %s", err)
		}
	}
	l, e := s.Load()
	if e != nil {
		return nil, fmt.Errorf("%s", e)
	}
	return l.(*SAMDGForwarder), nil
}
