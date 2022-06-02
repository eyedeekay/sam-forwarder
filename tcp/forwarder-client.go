package samforwarder

import (
	"io"
	"log"
	"net"

	//"os"
	//"path/filepath"
	"strconv"
	"strings"

	i2ptunconf "github.com/eyedeekay/sam-forwarder/config"
	"github.com/eyedeekay/sam-forwarder/hashhash"
	sfi2pkeys "github.com/eyedeekay/sam-forwarder/i2pkeys"
	samtunnel "github.com/eyedeekay/sam-forwarder/interface"

	samoptions "github.com/eyedeekay/sam-forwarder/options"
	"github.com/eyedeekay/sam3"
	"github.com/eyedeekay/sam3/i2pkeys"
)

// SAMClientForwarder is a tcp proxy that automatically forwards ports to i2p
type SAMClientForwarder struct {
	*i2ptunconf.Conf

	samConn           *sam3.SAM
	SamKeys           i2pkeys.I2PKeys
	Hasher            *hashhash.Hasher
	connectStream     *sam3.StreamSession
	addr              i2pkeys.I2PAddr
	publishConnection net.Listener
	Bytes             map[string]int
	ByteLimit         int

	file io.ReadWriter
	up   bool

	// config
	//Conf *i2ptunconf.Conf
}

func (f *SAMClientForwarder) Config() *i2ptunconf.Conf {
	if f.Conf == nil {
		f.Conf = i2ptunconf.NewI2PBlankTunConf()
	}
	return f.Conf
}

func (f *SAMClientForwarder) GetType() string {
	return f.Config().Type
}

func (f *SAMClientForwarder) ID() string {
	return f.Config().TunName
}

func (f *SAMClientForwarder) Keys() i2pkeys.I2PKeys {
	return f.SamKeys
}

func (f *SAMClientForwarder) print() []string {
	return f.Config().PrintSlice()
}

func (f *SAMClientForwarder) Props() map[string]string {
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

func (f *SAMClientForwarder) Cleanup() {
	f.connectStream.Close()
	f.publishConnection.Close()
	f.samConn.Close()
}

func (f *SAMClientForwarder) Print() string {
	var r string
	r += "name=" + f.Config().TunName + "\n"
	r += "type=" + f.Config().Type + "\n"
	r += "base32=" + f.Base32() + "\n"
	r += "base64=" + f.Base64() + "\n"
	r += "dest=" + f.Config().ClientDest + "\n"
	//r += "ntcpclient\n"
	for _, s := range f.print() {
		r += s + "\n"
	}
	return strings.Replace(r, "\n\n", "\n", -1)
	return r
}

func (f *SAMClientForwarder) Search(search string) string {
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

func (f *SAMClientForwarder) accesslisttype() string {
	if f.Config().AccessListType == "allowlist" {
		return "i2cp.enableAccessList=true"
	} else if f.Config().AccessListType == "blocklist" {
		return "i2cp.enableBlackList=true"
	} else if f.Config().AccessListType == "none" {
		return ""
	}
	return ""
}

func (f *SAMClientForwarder) accesslist() string {
	if f.Config().AccessListType != "" && len(f.Config().AccessList) > 0 {
		r := ""
		for _, s := range f.Config().AccessList {
			r += s + ","
		}
		return "i2cp.accessList=" + strings.TrimSuffix(r, ",")
	}
	return ""
}

func (f *SAMClientForwarder) leasesetsettings() (string, string, string) {
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
func (f *SAMClientForwarder) Target() string {
	return f.Config().TargetHost + ":" + f.Config().TargetPort
}

// Destination returns the destination of the i2p service you want to forward locally
func (f *SAMClientForwarder) Destination() string {
	return f.addr.Base32()
}

func (f *SAMClientForwarder) sam() string {
	return f.Config().SamHost + ":" + f.Config().SamPort
}

//Base32 returns the base32 address of the local destination
func (f *SAMClientForwarder) Base32() string {
	return f.SamKeys.Addr().Base32()
}

//Base32Readable returns the base32 address where the local service is being forwarded
func (f *SAMClientForwarder) Base32Readable() string {
	b32 := strings.Replace(f.Base32(), ".b32.i2p", "", 1)
	rhash, _ := f.Hasher.Friendly(b32)
	return rhash + " " + strconv.Itoa(len(b32))
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
func (f *SAMClientForwarder) Serve() error {
	if f.addr, err = f.samConn.Lookup(f.Config().ClientDest); err != nil {
		return err
	}
	if f.connectStream, err = f.samConn.NewStreamSession(f.Config().TunName, f.SamKeys, f.print()); err != nil {
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
		log.Println("Human-readable hash of Client:\n   ", f.Base32Readable())
		go f.forward(conn)
	}
}

func (f *SAMClientForwarder) Up() bool {
	return f.up
}

//Close shuts the whole thing down.
func (f *SAMClientForwarder) Close() error {
	var err error
	err = f.samConn.Close()
	f.up = false
	err = f.connectStream.Close()
	err = f.publishConnection.Close()
	return err
}

func (s *SAMClientForwarder) Load() (samtunnel.SAMTunnel, error) {
	if s.publishConnection, err = net.Listen("tcp", s.Config().TargetHost+":"+s.Config().TargetPort); err != nil {
		return nil, err
	}
	if s.samConn, err = sam3.NewSAM(s.sam()); err != nil {
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
		log.Println("Saved tunnel keys for", s.Conf.TunName, "in", s.Conf.FilePath)
	}
	s.Hasher, err = hashhash.NewHasher(len(strings.Replace(s.Base32(), ".b32.i2p", "", 1)))
	if err != nil {
		return nil, err
	}
	s.up = true
	return s, nil
}

//NewSAMClientForwarder makes a new SAM forwarder with default options, accepts host:port arguments
func NewSAMClientForwarder(host, port string) (samtunnel.SAMTunnel, error) {
	return NewSAMClientForwarderFromOptions(samoptions.SetHost(host), samoptions.SetPort(port))
}

//NewSAMClientForwarderFromOptions makes a new SAM forwarder with default options, accepts host:port arguments
func NewSAMClientForwarderFromOptions(opts ...func(samtunnel.SAMTunnel) error) (*SAMClientForwarder, error) {
	var s SAMClientForwarder
	s.Conf = i2ptunconf.NewI2PBlankTunConf()
	s.Conf.Type = "tcpclient"
	for _, o := range opts {
		if err := o(&s); err != nil {
			return nil, err
		}
	}
	l, e := s.Load()
	if e != nil {
		return nil, e
	}
	return l.(*SAMClientForwarder), nil
}
