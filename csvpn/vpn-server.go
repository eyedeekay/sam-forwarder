package samforwardervpn

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"syscall"
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
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		sigc := make(chan os.Signal, 1)
		signal.Notify(sigc, syscall.SIGINT, syscall.SIGTERM)
		Logger.Printf("received %v - initiating shutdown", <-sigc)
		cancel()
	}()
	s.VPNTunnel = udptunnel.NewTunnel(!s.Config.Client, s.Config.TunName, "10.76.0.2", s.Target(), "", []uint16{},
		"i2pvpn", time.Duration(time.Second*300), Logger)
	go s.VPNTunnel.Run(ctx)
	return &s, nil
}

// NewSAMVPNClientForwarderFromConfig generates a new SAMVPNForwarder from a config file
func NewSAMVPNForwarderFromConfig(iniFile, SamHost, SamPort string, label ...string) (*SAMClientServerVPN, error) {
	if iniFile != "none" {
		config, err := i2ptunconf.NewI2PTunConf(iniFile, label...)
		if err != nil {
			return nil, err
		}
		if SamHost != "" && SamHost != "127.0.0.1" && SamHost != "localhost" {
			config.SamHost = config.GetSAMHost(SamHost, config.SamHost)
		}
		if SamPort != "" && SamPort != "7656" {
			config.SamPort = config.GetSAMPort(SamPort, config.SamPort)
		}
		return NewSAMClientServerVPN(config)
	}
	return nil, nil
}
