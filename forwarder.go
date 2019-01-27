package samforwarder

import (
	"bufio"
	"io"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	//"os"
	//"path/filepath"
	"strings"
)

import (
	"github.com/eyedeekay/sam-forwarder/i2pkeys"
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

	clientLock bool
	connLock   bool

	connClientLock bool
	connConnLock   bool
}

var err error

func (f SAMForwarder) Cleanup() {
	f.publishStream.Close()
	f.publishListen.Close()
	f.publishConnection.Close()
	f.samConn.Close()
}

/*func (f SAMForwarder) targetForPort443() string {
	if f.TargetForPort443 != "" {
		return "targetForPort.4443=" + f.TargetHost + ":" + f.TargetForPort443
	}
	return ""
}*/

func (f SAMForwarder) print() []string {
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

func (f SAMForwarder) Print() string {
	var r string
	r += "name=" + f.TunName + "\n"
	r += "type=" + f.Type + "\n"
	r += "base32=" + f.Base32() + "\n"
	r += "base64=" + f.Base64() + "\n"
	if f.Type == "http" {
		r += "httpserver\n"
	} else {
		r += "ntcpserver\n"
	}
	for _, s := range f.print() {
		r += s + "\n"
	}
	return strings.Replace(r, "\n\n", "\n", -1)
}

func (f SAMForwarder) Search(search string) string {
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

func (f SAMForwarder) accesslisttype() string {
	if f.accessListType == "whitelist" {
		return "i2cp.enableAccessList=true"
	} else if f.accessListType == "blacklist" {
		return "i2cp.enableBlackList=true"
	} else if f.accessListType == "none" {
		return ""
	}
	return ""
}

func (f SAMForwarder) accesslist() string {
	if f.accessListType != "" && len(f.accessList) > 0 {
		r := ""
		for _, s := range f.accessList {
			r += s + ","
		}
		return "i2cp.accessList=" + strings.TrimSuffix(r, ",")
	}
	return ""
}

func (f SAMForwarder) leasesetsettings() (string, string, string) {
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
func (f SAMForwarder) Target() string {
	return f.TargetHost + ":" + f.TargetPort
}

func (f SAMForwarder) sam() string {
	return f.SamHost + ":" + f.SamPort
}

func (f SAMForwarder) HTTPRequestBytes(conn *sam3.SAMConn) ([]byte, *http.Request, error) {
	var request *http.Request
	var retrequest []byte
	var err error
	request, err = http.ReadRequest(bufio.NewReader(conn))
	if err != nil {
		return nil, nil, err
	}
	dest := conn.RemoteAddr().(sam3.I2PAddr)
	request.Header.Add("X-I2p-Dest-Base64", dest.Base64())
	request.Header.Add("X-I2p-Dest-Base32", dest.Base32())
	request.Header.Add("X-I2p-Dest-Hash", dest.DestHash().String())
	if retrequest, err = httputil.DumpRequest(request, true); err != nil {
		return nil, nil, err
	}
	return retrequest, request, nil
}

func (f SAMForwarder) HTTPResponseBytes(conn net.Conn, req *http.Request) ([]byte, error) {
	var response *http.Response
	var retresponse []byte
	var err error
	response, err = http.ReadResponse(bufio.NewReader(conn), req)
	if err != nil {
		return nil, err
	}
	response.Header.Add("X-I2p-Dest-Base64", f.Base64())
	response.Header.Add("X-I2p-Dest-Base32", f.Base32())
	if retresponse, err = httputil.DumpResponse(response, true); err != nil {
		return nil, err
	}
	return retresponse, nil
}

func (f SAMForwarder) clientUnlockAndClose(cli, conn bool, client net.Conn) {
	if cli {
		f.clientLock = cli
	}
	if conn {
		f.connLock = conn
	}
	if f.clientLock && f.connLock {
		client.Close()
		f.clientLock = false
		f.connLock = false
	}
}

func (f SAMForwarder) connUnlockAndClose(cli, conn bool, connection *sam3.SAMConn) {
	if cli {
		f.connClientLock = cli
	}
	if conn {
		f.connConnLock = conn
	}
	if f.connClientLock && f.connConnLock {
		connection.Close()
		f.connClientLock = false
		f.connConnLock = false
	}
}

func (f SAMForwarder) forward(conn *sam3.SAMConn) { //(conn net.Conn) {
	var request *http.Request
	var requestbytes []byte
	var responsebytes []byte
	var err error
	var client net.Conn
	if client, err = net.Dial("tcp", f.Target()); err != nil {
		log.Fatalf("Dial failed: %v", err)
	}
	go func() {
		if f.Type == "http" {
			defer f.clientUnlockAndClose(true, false, client)
			defer f.connUnlockAndClose(false, true, conn)
			if requestbytes, request, err = f.HTTPRequestBytes(conn); err == nil {
				log.Printf("Forwarding modified request: \n\t %s", string(requestbytes))
				client.Write(requestbytes)
			} else {
				log.Println("Error: ", requestbytes, err)
			}
		} else {
			defer client.Close()
			defer conn.Close()
			io.Copy(client, conn)
		}
	}()
	go func() {
		if f.Type == "http" {
			defer f.clientUnlockAndClose(false, true, client)
			defer f.connUnlockAndClose(true, false, conn)
			if responsebytes, err = f.HTTPResponseBytes(client, request); err == nil {
				log.Printf("Forwarding modified response: \n\t%s", string(responsebytes))
				conn.Write(responsebytes)
			} else {
				log.Println("Response Error: ", responsebytes, err)
			}
		} else {
			defer client.Close()
			defer conn.Close()
			io.Copy(conn, client)
		}
	}()
}

//Base32 returns the base32 address where the local service is being forwarded
func (f SAMForwarder) Base32() string {
	return f.SamKeys.Addr().Base32()
}

//Base64 returns the base64 address where the local service is being forwarded
func (f SAMForwarder) Base64() string {
	return f.SamKeys.Addr().Base64()
}

//Serve starts the SAM connection and and forwards the local host:port to i2p
func (f SAMForwarder) Serve() error {
	//lsk, lspk, lspsk := f.leasesetsettings()
	if f.publishStream, err = f.samConn.NewStreamSession(f.TunName, f.SamKeys, f.print()); err != nil {
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

//Close shuts the whole thing down.
func (f SAMForwarder) Close() error {
	var err error
	err = f.samConn.Close()
	err = f.publishStream.Close()
	err = f.publishListen.Close()
	err = f.publishConnection.Close()
	return err
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
	s.Type = "server"
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
	s.reduceIdleTime = "15"
	s.reduceIdleQuantity = "4"
	s.closeIdle = "false"
	s.closeIdleTime = "300000"
	s.clientLock = false
	s.connLock = false
	s.messageReliability = "none"
	s.passfile = ""
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
		if err := i2pkeys.Save(s.FilePath, s.TunName, s.passfile, &s.SamKeys); err != nil {
			return nil, err
		}
	}
	return &s, nil
}
