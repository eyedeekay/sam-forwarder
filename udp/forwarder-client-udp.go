package samforwarderudp

import (
	"io"
	"log"
	"net"
	"strconv"
	"strings"
	"time"
)

import (
	"github.com/eyedeekay/sam-forwarder/config"
	"github.com/eyedeekay/sam-forwarder/hashhash"
	"github.com/eyedeekay/sam-forwarder/i2pkeys"
	"github.com/eyedeekay/sam-forwarder/interface"
	"github.com/eyedeekay/sam3"
	"github.com/eyedeekay/sam3/i2pkeys"
)

//SAMSSUClientForwarder is a structure which automatically configured the forwarding of
//a local port to i2p over the SAM API.
type SAMSSUClientForwarder struct {
	samConn           *sam3.SAM
	SamKeys           i2pkeys.I2PKeys
	Hasher            *hashhash.Hasher
	connectStream     *sam3.DatagramSession
	addr              i2pkeys.I2PAddr
	publishConnection net.PacketConn

	file io.ReadWriter
	up   bool

	// config
	Conf *i2ptunconf.Conf
}

func (f *SAMSSUClientForwarder) Config() *i2ptunconf.Conf {
	if f.Conf == nil {
		f.Conf = i2ptunconf.NewI2PBlankTunConf()
	}
	return f.Conf
}

func (f *SAMSSUClientForwarder) GetType() string {
	return f.Config().Type
}

func (f *SAMSSUClientForwarder) ID() string {
	return f.Config().TunName
}

func (f *SAMSSUClientForwarder) Keys() i2pkeys.I2PKeys {
	return f.SamKeys
}

func (f *SAMSSUClientForwarder) Cleanup() {
	f.publishConnection.Close()
	f.connectStream.Close()
	f.samConn.Close()
}

func (f *SAMSSUClientForwarder) Close() error {
	f.Cleanup()
	f.up = false
	return nil
}

func (f *SAMSSUClientForwarder) print() []string {
	return f.Config().PrintSlice()
}

func (f *SAMSSUClientForwarder) Props() map[string]string {
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

func (f *SAMSSUClientForwarder) Print() string {
	var r string
	r += "name=" + f.Config().TunName + "\n"
	r += "type=" + f.Config().Type + "\n"
	r += "base32=" + f.Base32() + "\n"
	r += "base64=" + f.Base64() + "\n"
	r += "dest=" + f.Config().ClientDest + "\n"
	r += "ssuclient\n"
	for _, s := range f.print() {
		r += s + "\n"
	}
	return strings.Replace(r, "\n\n", "\n", -1)
}

func (f *SAMSSUClientForwarder) Search(search string) string {
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

func (f *SAMSSUClientForwarder) accesslisttype() string {
	if f.Config().AccessListType == "whitelist" {
		return "i2cp.enableAccessList=true"
	} else if f.Config().AccessListType == "blacklist" {
		return "i2cp.enableBlackList=true"
	} else if f.Config().AccessListType == "none" {
		return ""
	}
	return ""
}

func (f *SAMSSUClientForwarder) accesslist() string {
	if f.Config().AccessListType != "" && len(f.Config().AccessList) > 0 {
		r := ""
		for _, s := range f.Config().AccessList {
			r += s + ","
		}
		return "i2cp.accessList=" + strings.TrimSuffix(r, ",")
	}
	return ""
}

func (f *SAMSSUClientForwarder) leasesetsettings() (string, string, string) {
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

// Destination returns the destination of the i2p service you want to forward locally
func (f *SAMSSUClientForwarder) Destination() string {
	return f.addr.Base32()
}

// Target returns the host:port of the local service you want to forward to i2p
func (f *SAMSSUClientForwarder) Target() string {
	return f.Config().TargetHost + ":" + f.Config().TargetPort
}

func (f *SAMSSUClientForwarder) sam() string {
	return f.Config().SamHost + ":" + f.Config().SamPort
}

//Base32 returns the base32 address of the local destination
func (f *SAMSSUClientForwarder) Base32() string {
	return f.SamKeys.Addr().Base32()
}

//Base32Readable returns the base32 address where the local service is being forwarded
func (f *SAMSSUClientForwarder) Base32Readable() string {
	b32 := strings.Replace(f.Base32(), ".b32.i2p", "", 1)
	rhash, _ := f.Hasher.Friendly(b32)
	return rhash + " " + strconv.Itoa(len(b32))
}

//Base64 returns the base64 address of the local destination
func (f *SAMSSUClientForwarder) Base64() string {
	return f.SamKeys.Addr().Base64()
}

func (f *SAMSSUClientForwarder) forward(conn net.PacketConn) {
	//	go func() {
	//		defer f.connectStream.Close()
	//		defer f.publishConnection.Close()
	buffer := make([]byte, 1024)
	if size, _, readerr := f.publishConnection.ReadFrom(buffer); readerr == nil {
		if size2, writeerr := f.connectStream.WriteTo(buffer, f.addr); writeerr == nil {
			log.Printf("%q of %q bytes read", size, size2)
			log.Printf("%q of %q bytes written", size2, size)
		}
	}
	//	}()
	//	go func() {
	//		defer f.connectStream.Close()
	//		defer f.publishConnection.Close()
	//buffer := make([]byte, 1024)
	if size, _, readerr := f.connectStream.ReadFrom(buffer); readerr == nil {
		if size2, writeerr := f.publishConnection.WriteTo(buffer, f.addr); writeerr == nil {
			log.Printf("%q of %q bytes read", size, size2)
			log.Printf("%q of %q bytes written", size2, size)
		}
	}
	//	}()
}

func (f *SAMSSUClientForwarder) Up() bool {
	return f.up
}

func (f *SAMSSUClientForwarder) errSleep(err error) bool {
	if err != nil {
		log.Printf("Dial failed: %v, waiting 5 minutes to try again\n", err)
		time.Sleep(5 * time.Minute)
		return false
	} else {
		return true
	}
}

//Serve starts the SAM connection and and forwards the local host:port to i2p
func (f *SAMSSUClientForwarder) Serve() error {
	var err error
	log.Println("Establishing a SAM datagram session.")
	if f.publishConnection, err = net.ListenPacket("udp", f.Config().Target()); err != nil {
		return err
	}
	//p, _ := strconv.Atoi(f.Config().TargetPort)
	sp, _ := strconv.Atoi(f.Config().SamPort)
	f.connectStream, err = f.samConn.NewDatagramSession(
		f.Config().TunName,
		f.SamKeys,
		f.print(),
		sp-1,
	)
	if err != nil {
		log.Fatal("Datagram Session Creation error:", err.Error())
	}
	log.Println("SAM datagram session established.")
	log.Printf("Connected to localhost %v\n", f.publishConnection)
	log.Println("Human-readable hash of Client:\n   ", f.Base32Readable())
	/*	Close := false
		for !Close {
			//addr, err := net.ResolveUDPAddr("udp", f.Config().Target())
			Close = f.errSleep(err)
			//f.publishConnection, err = net.DialUDP("udp", nil, addr)
			Close = f.errSleep(err)
			log.Printf("Forwarding client to i2p address: %v\n", f.publishConnection)
		}*/
	for {
		f.forward(f.publishConnection)
	}
	return nil
}

func (s *SAMSSUClientForwarder) Load() (samtunnel.SAMTunnel, error) {
	if s.samConn, err = sam3.NewSAM(s.sam()); err != nil {
		return nil, err
	}
	if s.addr, err = s.samConn.Lookup(s.Config().ClientDest); err != nil {
		return nil, err
	}
	log.Println("SAM Bridge connection established.")
	if s.Config().SaveFile {
		log.Println("Saving i2p keys")
	}
	if s.SamKeys, err = sfi2pkeys.Load(s.Config().FilePath, s.Config().TunName, s.Config().KeyFilePath, s.samConn, s.Config().SaveFile); err != nil {
		return nil, err
	}
	log.Println("Destination keys generated, tunnel name:", s.Config().TunName)
	if s.Config().SaveFile {
		if err := sfi2pkeys.Save(s.Config().FilePath, s.Config().TunName, s.Config().KeyFilePath, s.SamKeys); err != nil {
			return nil, err
		}
		log.Println("Saved tunnel keys for", s.Config().TunName)
	}
	s.Hasher, err = hashhash.NewHasher(len(strings.Replace(s.Base32(), ".b32.i2p", "", 1)))
	if err != nil {
		return nil, err
	}
	s.up = true
	return s, nil
}

//NewSAMSSUClientForwarderFromOptions makes a new SAM forwarder with default options, accepts host:port arguments
func NewSAMSSUClientForwarderFromOptions(opts ...func(*SAMSSUClientForwarder) error) (*SAMSSUClientForwarder, error) {
	var s SAMSSUClientForwarder
	s.Conf = i2ptunconf.NewI2PBlankTunConf()
	s.Conf.Type = "udpclient"
	for _, o := range opts {
		if err := o(&s); err != nil {
			return nil, err
		}
	}
	l, e := s.Load()
	if e != nil {
		return nil, e
	}
	return l.(*SAMSSUClientForwarder), nil
}
