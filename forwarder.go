package samforwarder

import (
	"bufio"
	"io"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"path/filepath"
	"strings"
)

import (
	"github.com/eyedeekay/sam3"
)

//SAMForwarder is a structure which automatically configured the forwarding of
//a local service to i2p over the SAM API.
type SAMForwarder struct {
	SamHost string
	SamPort string
	TunName string

	TargetHost string
	TargetPort string

	samConn           *sam3.SAM
	SamKeys           sam3.I2PKeys
	publishStream     *sam3.StreamSession
	publishListen     *sam3.StreamListener
	publishConnection net.Conn

	FilePath string
	file     io.ReadWriter
	save     bool

	Type string

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

var err error

/*func (f *SAMForwarder) targetForPort443() string {
	if f.TargetForPort443 != "" {
		return "targetForPort.4443=" + f.TargetHost + ":" + f.TargetForPort443
	}
	return ""
}*/

func (f *SAMForwarder) accesslisttype() string {
	if f.accessListType == "whitelist" {
		return "i2cp.enableAccessList=true"
	} else if f.accessListType == "blacklist" {
		return "i2cp.enableBlackList=true"
	} else if f.accessListType == "none" {
		return ""
	}
	return ""
}

func (f *SAMForwarder) accesslist() string {
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
func (f *SAMForwarder) Target() string {
	return f.TargetHost + ":" + f.TargetPort
}

func (f *SAMForwarder) sam() string {
	return f.SamHost + ":" + f.SamPort
}

func (f *SAMForwarder) HTTPRequestBytes(conn *sam3.SAMConn) ([]byte, error) {
	var request *http.Request
	var retrequest []byte
	var err error
	request, err = http.ReadRequest(bufio.NewReader(conn))
	if err != nil {
		return nil, err
	}
	dest := conn.RemoteAddr().(sam3.I2PAddr)
	log.Printf("Adding headers to http connection\n\tX-I2p-Dest-Base64=%s\n\tX-I2p-Dest-Base32=%s\n\tX-I2p-Dest-Hash=%s",
		dest.Base64(), dest.Base32(), dest.DestHash().String())
	request.Header.Add("X-I2p-Dest-Base64", dest.Base64())
	request.Header.Add("X-I2p-Dest-Base32", dest.Base32())
	request.Header.Add("X-I2p-Dest-Hash", dest.DestHash().String())
	if retrequest, err = httputil.DumpRequest(request, true); err != nil {
		return nil, err
	}
	return retrequest, nil
}

func (f *SAMForwarder) forward(conn *sam3.SAMConn) { //(conn net.Conn) {
	var err error
	client, err := net.Dial("tcp", f.Target())
	if err != nil {
		log.Fatalf("Dial failed: %v", err)
	}
	go func() {
		defer client.Close()
		defer conn.Close()
		if f.Type == "http" {
			if b, e := f.HTTPRequestBytes(conn); e != nil {
				client.Write(b)
			} else {
				log.Println(e)
			}
			//io.Copy(client, conn)
		} else {
			io.Copy(client, conn)
		}
	}()
	go func() {
		defer client.Close()
		defer conn.Close()
		if f.Type == "http" {
			io.Copy(conn, client)
		} else {
			io.Copy(conn, client)
		}
	}()
}

//Base32 returns the base32 address where the local service is being forwarded
func (f *SAMForwarder) Base32() string {
	return f.SamKeys.Addr().Base32()
}

//Base64 returns the base64 address where the local service is being forwarded
func (f *SAMForwarder) Base64() string {
	return f.SamKeys.Addr().Base64()
}

//Serve starts the SAM connection and and forwards the local host:port to i2p
func (f *SAMForwarder) Serve() error {
	if f.publishStream, err = f.samConn.NewStreamSession(f.TunName, f.SamKeys,
		[]string{
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
	if f.publishListen, err = f.publishStream.Listen(); err != nil {
		return err
	}
	log.Println("Starting Listener.")
	b := string(f.SamKeys.Addr().Base32())
	log.Println("SAM Listener created,", b)

	for {
		conn, err := f.publishListen.AcceptI2P()
		if err != nil {
			log.Fatalf("ERROR: failed to accept listener: %v", err)
		}
		log.Printf("Accepted connection %v\n", conn)
		go f.forward(conn)
	}
}

//NewSAMForwarder makes a new SAM forwarder with default options, accepts host:port arguments
func NewSAMForwarder(host, port string) (*SAMForwarder, error) {
	return NewSAMForwarderFromOptions(SetHost(host), SetPort(port))
}

//NewSAMForwarderFromOptions makes a new SAM forwarder with default options, accepts host:port arguments
func NewSAMForwarderFromOptions(opts ...func(*SAMForwarder) error) (*SAMForwarder, error) {
	var s SAMForwarder
	s.SamHost = "127.0.0.1"
	s.SamPort = "7656"
	s.FilePath = ""
	s.save = false
	s.TargetHost = "127.0.0.1"
	s.TargetPort = "8081"
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
	s.closeIdleTime = "300000"
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
