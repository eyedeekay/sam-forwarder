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

type SAMClientVPN struct {
	// i2p tunnel
	I2PTunnel *samforwarderudp.SAMSSUClientForwarder
	Config    *i2ptunconf.Conf
	FilePath  string
	// VPN tunnel
	VPNTunnel  udptunnel.Tunnel
	ClientDest string
}

func (f *SAMClientVPN) sam() string {
	return f.Config.SamHost + ":" + f.Config.SamPort
}

// Target returns the host:port of the local service you want to forward to i2p
func (f *SAMClientVPN) Target() string {
	return f.Config.TargetHost + ":" + f.Config.TargetPort
}

func (f *SAMClientVPN) Base32() string {
	return f.I2PTunnel.Base32()
}

func NewSAMClientVPN(conf *i2ptunconf.Conf, destination ...string) (*SAMClientVPN, error) {
	if len(destination) == 0 {
		return NewSAMClientVPNFromOptions(SetClientVPNConfig(conf))
	} else if len(destination) == 1 {
		return NewSAMClientVPNFromOptions(SetClientVPNConfig(conf), SetClientDest(destination[0]))
	} else {
		return nil, fmt.Errorf("Error, argument for destination must be len==0 or len==1")
	}
}

func NewSAMClientVPNFromOptions(opts ...func(*SAMClientVPN) error) (*SAMClientVPN, error) {
	var s SAMClientVPN
	s.FilePath = ""
	s.ClientDest = ""
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
	if s.ClientDest != "" {
		s.Config.ClientDest = s.ClientDest
	}
	if s.Config.ClientDest == "" {
		return nil, fmt.Errorf("VPN Client destination cannot be empty")
	}
	s.I2PTunnel, err = i2ptunconf.NewSAMSSUClientForwarderFromConf(s.Config)
	if err != nil {
		return nil, err
	}
	go s.I2PTunnel.Serve()
	if !s.Config.Client {
		return nil, fmt.Errorf("Error, VPN client marked as server")
	} else {
		log.Println("VPN client tunnel destination", s.I2PTunnel.Base32())
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
	s.VPNTunnel = udptunnel.NewTunnel(!s.Config.Client, s.Config.TunName, "10.76.0.3", s.Target(), "", []uint16{},
		"i2pvpn", time.Duration(time.Second*300), Logger)
	go s.VPNTunnel.Run(ctx)
	return &s, nil
}

// NewSAMVPNForwarderFromConfig generates a new SAMVPNForwarder from a config file
func NewSAMVPNClientForwarderFromConfig(iniFile, SamHost, SamPort string, label ...string) (*SAMClientVPN, error) {
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
		return NewSAMClientVPN(config)
	}
	return nil, nil
}
