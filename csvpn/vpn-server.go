package i2pvpn

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/eyedeekay/sam-forwarder/config"
	"github.com/eyedeekay/sam-forwarder/udp"
	"github.com/eyedeekay/udptunnel/tunnel"
)

type SAMClientServerVPN struct {
	// i2p tunnel
	I2PTunnel *samforwarderudp.SAMSSUForwarder
	Config    *i2ptunconf.Conf
	FilePath  string
	// VPN tunnel
	VPNTunnel udptunnel.Tunnel
}

func (f *SAMClientServerVPN) sam() string {
	return f.Config.SamHost + ":" + f.Config.SamPort
}

// Target returns the host:port of the local service you want to forward to i2p
func (f *SAMClientServerVPN) Target() string {
	return f.Config.TargetHost + ":" + f.Config.TargetPort
}

func (f *SAMClientServerVPN) Base32() string {
	return f.I2PTunnel.Base32()
}

func NewSAMClientServerVPN(conf *i2ptunconf.Conf) (*SAMClientServerVPN, error) {
	return NewSAMClientServerVPNFromOptions(SetVPNConfig(conf))
}

func NewSAMClientServerVPNFromOptions(opts ...func(*SAMClientServerVPN) error) (*SAMClientServerVPN, error) {
	var s SAMClientServerVPN
	s.FilePath = ""
	for _, o := range opts {
		if err := o(&s); err != nil {
			return nil, err
		}
	}
	var err error
	if s.Config == nil {
		if s.FilePath != "" {
			s.Config, err = i2ptunconf.NewI2PTunConf(s.FilePath)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, fmt.Errorf("No VPN configuration provided")
		}
	}
	s.I2PTunnel, err = i2ptunconf.NewSAMSSUForwarderFromConf(s.Config)
	if err != nil {
		return nil, err
	}
	go s.I2PTunnel.Serve()
	if !s.Config.Client {
		log.Println("VPN server tunnel listening on", s.I2PTunnel.Base32())
	} else {
		return nil, fmt.Errorf("Error, VPN server marked as client")
	}
	var logBuf bytes.Buffer
	Logger := log.New(io.MultiWriter(os.Stderr, &logBuf), "", log.Ldate|log.Ltime|log.Lshortfile)
	s.VPNTunnel = udptunnel.NewTunnel(!s.Config.Client, s.Config.TunName, "10.76.0.2", s.Target(), "", []uint16{},
		"i2pvpn", time.Duration(time.Second*300), Logger)
	return &s, nil
}
