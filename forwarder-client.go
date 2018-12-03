package samforwarder

import (
	"io"
	"log"
	"net"
	//"os"
	//"path/filepath"
	"strings"
)

import (
	"github.com/eyedeekay/sam-forwarder/i2pkeys"
	"github.com/eyedeekay/sam3"
)

// SAMClientForwarder is a tcp proxy that automatically forwards ports to i2p
type SAMClientForwarder struct {
	SamHost string
	SamPort string
	TunName string
	Type    string

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

	// samcatd options
	passfile string

	// I2CP options
	encryptLeaseSet           string
	leaseSetKey               string
	leaseSetPrivateKey        string
	leaseSetPrivateSigningKey string
	LeaseSetKeys              *sam3.I2PKeys
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

func (f SAMClientForwarder) print() []string {
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

func (f SAMClientForwarder) Cleanup() {
	f.connectStream.Close()
	f.publishConnection.Close()
	f.samConn.Close()
}

func (f SAMClientForwarder) Print() string {
	var r string
	r += "name=" + f.TunName + "\n"
	r += "type=" + f.Type + "\n"
	r += "base32=" + f.Base32() + "\n"
	r += "base64=" + f.Base64() + "\n"
	r += "dest=" + f.dest + "\n"
	r += "ntcpclient\n"
	for _, s := range f.print() {
		r += s + "\n"
	}
	return strings.Replace(r, "\n\n", "\n", -1)
	return r
}

func (f SAMClientForwarder) Search(search string) string {
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

func (f SAMClientForwarder) accesslisttype() string {
	if f.accessListType == "whitelist" {
		return "i2cp.enableAccessList=true"
	} else if f.accessListType == "blacklist" {
		return "i2cp.enableBlackList=true"
	} else if f.accessListType == "none" {
		return ""
	}
	return ""
}

func (f SAMClientForwarder) accesslist() string {
	if f.accessListType != "" && len(f.accessList) > 0 {
		r := ""
		for _, s := range f.accessList {
			r += s + ","
		}
		return "i2cp.accessList=" + strings.TrimSuffix(r, ",")
	}
	return ""
}

func (f SAMClientForwarder) leasesetsettings() (string, string, string) {
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
func (f SAMClientForwarder) Target() string {
	return f.TargetHost + ":" + f.TargetPort
}

// Destination returns the destination of the i2p service you want to forward locally
func (f SAMClientForwarder) Destination() string {
	return f.addr.Base32()
}

func (f SAMClientForwarder) sam() string {
	return f.SamHost + ":" + f.SamPort
}

//Base32 returns the base32 address of the local destination
func (f SAMClientForwarder) Base32() string {
	return f.SamKeys.Addr().Base32()
}

//Base64 returns the base64 address of the local destiantion
func (f SAMClientForwarder) Base64() string {
	return f.SamKeys.Addr().Base64()
}

func (f SAMClientForwarder) forward(conn net.Conn) {
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
func (f SAMClientForwarder) Serve() error {
	if f.addr, err = f.samConn.Lookup(f.dest); err != nil {
		return err
	}
	if f.connectStream, err = f.samConn.NewStreamSession(f.TunName, f.SamKeys, f.print()); err != nil {
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

//Close shuts the whole thing down.
func (f SAMClientForwarder) Close() error {
	var err error
	err = f.samConn.Close()
	err = f.connectStream.Close()
	err = f.publishConnection.Close()
	return err
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
	s.inQuantity = "2"
	s.outQuantity = "2"
	s.inVariance = "1"
	s.outVariance = "1"
	s.inBackupQuantity = "3"
	s.outBackupQuantity = "3"
	s.inAllowZeroHop = "false"
	s.outAllowZeroHop = "false"
	s.encryptLeaseSet = "false"
	s.leaseSetKey = ""
	s.leaseSetPrivateKey = ""
	s.leaseSetPrivateSigningKey = ""
	s.fastRecieve = "false"
	s.useCompression = "true"
	s.reduceIdle = "false"
	s.reduceIdleTime = "300000"
	s.reduceIdleQuantity = "4"
	s.closeIdle = "false"
	s.closeIdleTime = "300000"
	s.dest = "none"
	s.Type = "client"
	s.messageReliability = "none"
	s.passfile = ""
	s.dest = "i2p-projekt.i2p"
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
		if err := i2pkeys.Save(s.FilePath, s.TunName, s.passfile, &s.SamKeys); err != nil {
			return nil, err
		}
	}
	return &s, nil
}
