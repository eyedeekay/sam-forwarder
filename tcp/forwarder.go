package samforwarder

import (
	"bufio"
	"crypto/tls"
	"io"
	"log"
	"net"
	"net/http"
	"net/http/httputil"

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

	i2pkeys "github.com/eyedeekay/sam3/i2pkeys"
)

//SAMForwarder is a structure which automatically configured the forwarding of
//a local service to i2p over the SAM API.
type SAMForwarder struct {
	*i2ptunconf.Conf
	samConn       *sam3.SAM
	SamKeys       i2pkeys.I2PKeys
	Hasher        *hashhash.Hasher
	publishStream *sam3.StreamSession
	publishListen net.Listener //*sam3.StreamListener
	//publishListen *sam3.StreamListener
	Bytes     map[string]int64
	ByteLimit int64

	file io.ReadWriter
	up   bool

	// conf
	//Conf *i2ptunconf.Conf

	clientLock bool
	connLock   bool

	connClientLock bool
	connConnLock   bool
}

var err error

func (f *SAMForwarder) Config() *i2ptunconf.Conf {
	if f.Conf == nil {
		f.Conf = i2ptunconf.NewI2PBlankTunConf()
	}
	return f.Conf
}

func (f *SAMForwarder) ID() string {
	return f.Config().TunName
}

func (f *SAMForwarder) Keys() i2pkeys.I2PKeys {
	return f.SamKeys
}

func (f *SAMForwarder) Cleanup() {
	f.publishStream.Close()
	f.publishListen.Close()
	f.samConn.Close()
}

func (f *SAMForwarder) GetType() string {
	return f.Conf.Type
}

func (f *SAMForwarder) print() []string {
	return f.Conf.PrintSlice()
}

func (f *SAMForwarder) Props() map[string]string {
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

func (f *SAMForwarder) Print() string {
	var r string
	r += "name=" + f.Config().TunName + "\n"
	r += "type=" + f.Conf.Type + "\n"
	r += "base32=" + f.Base32() + "\n"
	r += "base64=" + f.Base64() + "\n"
	if f.Conf.Type == "http" || f.Conf.Type == "https" {
		r += "httpserver\n"
	} else {
		r += "ntcpserver\n"
	}
	for _, s := range f.print() {
		r += s + "\n"
	}
	return strings.Replace(r, "\n\n", "\n", -1)
}

func (f *SAMForwarder) Search(search string) string {
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

func (f *SAMForwarder) accesslisttype() string {
	if f.Config().AccessListType == "allowlist" {
		return "i2cp.enableAccessList=true"
	} else if f.Config().AccessListType == "blocklist" {
		return "i2cp.enableBlackList=true"
	} else if f.Config().AccessListType == "none" {
		return ""
	}
	return ""
}

func (f *SAMForwarder) accesslist() string {
	if f.Config().AccessListType != "" && len(f.Config().AccessList) > 0 {
		r := ""
		for _, s := range f.Config().AccessList {
			r += s + ","
		}
		return "i2cp.accessList=" + strings.TrimSuffix(r, ",")
	}
	return ""
}

func (f *SAMForwarder) leasesetsettings() (string, string, string) {
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
func (f *SAMForwarder) Target() string {
	return f.Config().TargetHost + ":" + f.Config().TargetPort
}

func (f *SAMForwarder) sam() string {
	return f.Config().SamHost + ":" + f.Config().SamPort
}

func (f *SAMForwarder) ClientBase64(conn net.Conn) string {
	dest := conn.RemoteAddr().(i2pkeys.I2PAddr)
	return dest.Base32()
}

func (f *SAMForwarder) HTTPRequestBytes(conn net.Conn) ([]byte, *http.Request, error) {
	var request *http.Request
	var retrequest []byte
	var err error
	request, err = http.ReadRequest(bufio.NewReader(conn))
	if err != nil {
		return nil, nil, err
	}
	dest := conn.RemoteAddr().(i2pkeys.I2PAddr)
	request.Header.Add("X-I2p-Dest-Base64", dest.Base64())
	request.Header.Add("X-I2p-Dest-Base32", dest.Base32())
	request.Header.Add("X-I2p-Dest-Hash", dest.DestHash().Hash())
	if retrequest, err = httputil.DumpRequest(request, true); err != nil {
		return nil, nil, err
	}
	return retrequest, request, nil
}

func (f *SAMForwarder) HTTPResponseBytes(conn net.Conn, req *http.Request) ([]byte, error) {
	var response *http.Response
	var retresponse []byte
	var err error
	response, err = http.ReadResponse(bufio.NewReader(conn), req)
	if err != nil {
		return nil, err
	}
	if retresponse, err = httputil.DumpResponse(response, true); err != nil {
		return nil, err
	}
	return retresponse, nil
}

func (f *SAMForwarder) clientUnlockAndClose(cli, conn bool, client net.Conn) {
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

func (f *SAMForwarder) connUnlockAndClose(cli, conn bool, connection net.Conn) {
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

func (f *SAMForwarder) forward(conn net.Conn) { //(conn net.Conn) {
	if !f.Up() {
		return
	}
	var request *http.Request
	var requestbytes []byte
	var responsebytes []byte
	go func() {
		var err error
		var client net.Conn
		if client, err = net.Dial("tcp", f.Target()); err != nil {
			log.Fatalf("Dial failed: %v", err)
		}
		defer client.Close()
		if f.Conf.Type == "http" || f.Conf.Type == "https" {
			if requestbytes, request, err = f.HTTPRequestBytes(conn); err == nil {
				log.Printf("Forwarding modified request: \n\t %s", string(requestbytes))
				client.Write(requestbytes)
			} else {
				log.Println("Error: ", requestbytes, err)
			}
		} else {
			if f.ByteLimit > 0 {
				if val, ok := f.Bytes[f.ClientBase64(conn)]; ok == true {
					if val > f.ByteLimit {
						return
					}
				}
			}
			if count, err := io.Copy(client, conn); err == nil {
				if f.ByteLimit > 0 {
					f.Bytes[f.ClientBase64(conn)] += count
				}
			}
		}
	}()
	go func() {
		var err error
		var client net.Conn
		if client, err = net.Dial("tcp", f.Target()); err != nil {
			log.Fatalf("Dial failed: %v", err)
		}
		defer client.Close()
		if f.Conf.Type == "http" || f.Conf.Type == "https" {
			if responsebytes, err = f.HTTPResponseBytes(client, request); err == nil {
				log.Printf("Forwarding modified response: \n\t%s", string(responsebytes))
				conn.Write(responsebytes)
			} else {
				log.Println("Response Error: ", responsebytes, err)
			}
		} else {
			if val, ok := f.Bytes[f.ClientBase64(conn)]; ok == true {
				if val > f.ByteLimit {
					return
				}
			}
			io.Copy(conn, client)
		}
	}()
}

//Base32 returns the base32 address where the local service is being forwarded
func (f *SAMForwarder) Base32() string {
	return f.SamKeys.Addr().Base32()
}

//Base32Readable returns the base32 address where the local service is being forwarded
func (f *SAMForwarder) Base32Readable() string {
	b32 := strings.Replace(f.Base32(), ".b32.i2p", "", 1)
	rhash, _ := f.Hasher.Friendly(b32)
	return rhash + " " + strconv.Itoa(len(b32))
}

//Base64 returns the base64 address where the local service is being forwarded
func (f *SAMForwarder) Base64() string {
	return f.SamKeys.Addr().Base64()
}

//Serve starts the SAM connection and and forwards the local host:port to i2p
func (f *SAMForwarder) Serve() error {
	//lsk, lspk, lspsk := f.Config().Leasesetsettings()
	if f.Up() {
		if f.publishStream, err = f.samConn.NewStreamSession(f.Conf.TunName, f.SamKeys, f.print()); err != nil {
			log.Println("Stream Creation error:", err.Error())
			return err
		}
		log.Println("SAM stream session established.")
		publishListen, err := f.publishStream.Listen()
		if err != nil {
			return err
		}
		log.Println("Starting Listener.")
		log.Println("SAM Listener created,", f.Base32())
		log.Println("Human-readable hash:\n   ", f.Base32Readable())
		if f.Conf.UseTLS {
			tlsconf, err := f.Conf.TLSConfig()
			if err != nil {
				return err
			}
			f.publishListen = tls.NewListener(publishListen, tlsconf)
		} else {
			f.publishListen = publishListen
		}
		for {
			conn, err := f.publishListen.Accept()
			if err != nil {
				log.Printf("ERROR: failed to accept listener: %v", err)
				return nil
			}
			defer conn.Close()
			log.Printf("Accepted connection %v\n", conn)
			go f.forward(conn)
		}
	}
	return nil
}

func (f *SAMForwarder) Up() bool {
	return f.up
}

//Close shuts the whole thing down.
func (f *SAMForwarder) Close() error {
	var err error
	//err = f.samConn.Close()
	f.up = false
	err = f.publishStream.Close()
	//err = f.samConn.Close()
	//err = f.publishListen.Close()
	return err
}

func (s *SAMForwarder) Load() (samtunnel.SAMTunnel, error) {
	if s.samConn, err = sam3.NewSAM(s.sam()); err != nil {
		return nil, err
	}
	log.Println("SAM Bridge connection established.")
	if s.Conf.SaveFile {
		log.Println("Saving i2p keys")
	}
	if s.SamKeys, err = sfi2pkeys.Load(s.Conf.FilePath, s.Conf.TunName, s.Conf.KeyFilePath, s.samConn, s.Conf.SaveFile); err != nil {
		return nil, err
	}
	log.Println("Destination keys generated, tunnel name:", s.Conf.TunName)
	if s.Conf.SaveFile {
		if err := sfi2pkeys.Save(s.Conf.FilePath, s.Conf.TunName, s.Conf.KeyFilePath, s.SamKeys); err != nil {
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

//NewSAMForwarder makes a new SAM forwarder with default options, accepts host:port arguments
func NewSAMForwarder(host, port string) (samtunnel.SAMTunnel, error) {
	return NewSAMForwarderFromOptions(samoptions.SetHost(host), samoptions.SetPort(port))
}

//NewSAMForwarderFromOptions makes a new SAM forwarder with default options, accepts host:port arguments
func NewSAMForwarderFromOptions(opts ...func(samtunnel.SAMTunnel) error) (*SAMForwarder, error) {
	var s SAMForwarder
	s.Conf = i2ptunconf.NewI2PBlankTunConf()
	s.clientLock = false
	s.connLock = false
	s.ByteLimit = -1
	s.Bytes = make(map[string]int64)
	for _, o := range opts {
		if err := o(&s); err != nil {
			return nil, err
		}
	}
	l, e := s.Load()
	if e != nil {
		return nil, e
	}
	return l.(*SAMForwarder), nil
}
