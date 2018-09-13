package sammanager

import (
	"context"
	"fmt"
	"log"
	"net"
)

import (
	"github.com/eyedeekay/sam-forwarder"
	"github.com/eyedeekay/sam-forwarder/config"
	"github.com/eyedeekay/sam-forwarder/udp"
)

type SAMManager struct {
	FilePath string
	save     bool
	start    bool
	config   *i2ptunconf.Conf

	ServerHost string
	ServerPort string
	SamHost    string
	SamPort    string

	forwarders          []samforwarder.SAMForwarder
	clientforwarders    []samforwarder.SAMClientForwarder
	udpforwarders       []samforwarderudp.SAMSSUForwarder
	udpclientforwarders []samforwarderudp.SAMSSUClientForwarder
}

func (s *SAMManager) FindForwarder(lookup string) (bool, int, string) {
	for index, element := range s.forwarders {
		if element.TunName == lookup {
			return true, index, element.Type
		}
	}
	for index, element := range s.clientforwarders {
		if element.TunName == lookup {
			return true, index, "client"
		}
	}
	for index, element := range s.udpforwarders {
		if element.TunName == lookup {
			return true, index, "udpserver"
		}
	}
	for index, element := range s.udpclientforwarders {
		if element.TunName == lookup {
			return true, index, "udpclient"
		}
	}
	return false, -1, ""
}

func (s *SAMManager) Serve() bool {
	log.Println("Starting tunnels")
	for _, element := range s.forwarders {
		log.Println("Starting NTCP service tunnel", element.TunName)
		go element.Serve()
	}
	for _, element := range s.clientforwarders {
		log.Println("Starting NTCP client tunnel", element.TunName)
		go element.Serve(element.Destination())
	}
	for _, element := range s.udpforwarders {
		log.Println("Starting SSU service tunnel", element.TunName)
		go element.Serve()
	}
	for _, element := range s.udpclientforwarders {
		log.Println("Starting SSU client tunnel", element.TunName)
		go element.Serve(element.Destination())
	}
	for true {
	}
	return false
}

func (s *SAMManager) LookupForwarder(lookup string, label ...string) (bool, string) {
	for _, l := range s.config.Labels {
		log.Println(l)
	}
	return false, ""
}

func (s *SAMManager) Dial(ctx context.Context, network, address string) (*net.Conn, error) {
	return nil, nil
}

func NewSAMManagerFromOptions(opts ...func(*SAMManager) error) (*SAMManager, error) {
	var s SAMManager
	s.FilePath = ""
	s.save = true
	s.start = false
	s.config = i2ptunconf.NewI2PBlankTunConf()
	s.ServerHost = "localhost"
	s.ServerPort = "7957"
	s.SamHost = "localhost"
	s.SamPort = "7656"
	for _, o := range opts {
		if err := o(&s); err != nil {
			return nil, err
		}
	}
	var err error
	if s.FilePath != "" {
		s.config, err = i2ptunconf.NewI2PTunConf(s.FilePath)
		if err != nil {
			return nil, err
		}
	} else {
		if s.config == nil {
			return nil, fmt.Errorf("Configuration not found")
		}
		s.FilePath = s.config.FilePath
	}
	if s.start {
		t, b := s.config.Get("type")
		if !b {
			return nil, fmt.Errorf("samcat was instructed to start a tunnel with insufficient default settings information.")
		}
		switch t {
		case "http":
			log.Println("found http under")
			if f, e := i2ptunconf.NewSAMForwarderFromConfig(s.FilePath, s.SamHost, s.SamPort); e == nil {
				s.forwarders = append(s.forwarders, *f)
			}
		case "server":
			log.Println("found server under")
			if f, e := i2ptunconf.NewSAMForwarderFromConfig(s.FilePath, s.SamHost, s.SamPort); e == nil {
				s.forwarders = append(s.forwarders, *f)
			}
		case "client":
			log.Println("found client under")
			if f, e := i2ptunconf.NewSAMClientForwarderFromConfig(s.FilePath, s.SamHost, s.SamPort); e == nil {
				s.clientforwarders = append(s.clientforwarders, *f)
			}
		case "udpserver":
			log.Println("found udpserver under")
			if f, e := i2ptunconf.NewSAMSSUForwarderFromConfig(s.FilePath, s.SamHost, s.SamPort); e == nil {
				s.udpforwarders = append(s.udpforwarders, *f)
			}
		case "udpclient":
			log.Println("found udpclient under")
			if f, e := i2ptunconf.NewSAMSSUClientForwarderFromConfig(s.FilePath, s.SamHost, s.SamPort); e == nil {
				s.udpclientforwarders = append(s.udpclientforwarders, *f)
			}
		}
	}
	for _, label := range s.config.Labels {
		if t, e := s.config.Get("type", label); e {
			switch t {
			case "http":
				log.Println("found http under", label)
				if f, e := i2ptunconf.NewSAMForwarderFromConfig(s.FilePath, s.SamHost, s.SamPort, label); e == nil {
					s.forwarders = append(s.forwarders, *f)
				}
			case "server":
				log.Println("found server under", label)
				if f, e := i2ptunconf.NewSAMForwarderFromConfig(s.FilePath, s.SamHost, s.SamPort, label); e == nil {
					s.forwarders = append(s.forwarders, *f)
				}
			case "client":
				log.Println("found client under", label)
				if f, e := i2ptunconf.NewSAMClientForwarderFromConfig(s.FilePath, s.SamHost, s.SamPort, label); e == nil {
					s.clientforwarders = append(s.clientforwarders, *f)
				}
			case "udpserver":
				log.Println("found udpserver under", label)
				if f, e := i2ptunconf.NewSAMSSUForwarderFromConfig(s.FilePath, s.SamHost, s.SamPort, label); e == nil {
					s.udpforwarders = append(s.udpforwarders, *f)
				}
			case "udpclient":
				log.Println("found udpclient under", label)
				if f, e := i2ptunconf.NewSAMSSUClientForwarderFromConfig(s.FilePath, s.SamHost, s.SamPort, label); e == nil {
					s.udpclientforwarders = append(s.udpclientforwarders, *f)
				}
			}
		}
	}
	if len(s.config.Labels) == 0 && !s.start {
		t, b := s.config.Get("type")
		if !b {
			return nil, fmt.Errorf("samcat was instructed to start a tunnel with insufficient default settings information.")
		}
		switch t {
		case "http":
			log.Println("found http under")
			if f, e := i2ptunconf.NewSAMForwarderFromConfig(s.FilePath, s.SamHost, s.SamPort); e == nil {
				s.forwarders = append(s.forwarders, *f)
			}
		case "server":
			log.Println("found server under")
			if f, e := i2ptunconf.NewSAMForwarderFromConfig(s.FilePath, s.SamHost, s.SamPort); e == nil {
				s.forwarders = append(s.forwarders, *f)
			}
		case "client":
			log.Println("found client under")
			if f, e := i2ptunconf.NewSAMClientForwarderFromConfig(s.FilePath, s.SamHost, s.SamPort); e == nil {
				s.clientforwarders = append(s.clientforwarders, *f)
			}
		case "udpserver":
			log.Println("found udpserver under")
			if f, e := i2ptunconf.NewSAMSSUForwarderFromConfig(s.FilePath, s.SamHost, s.SamPort); e == nil {
				s.udpforwarders = append(s.udpforwarders, *f)
			}
		case "udpclient":
			log.Println("found udpclient under")
			if f, e := i2ptunconf.NewSAMSSUClientForwarderFromConfig(s.FilePath, s.SamHost, s.SamPort); e == nil {
				s.udpclientforwarders = append(s.udpclientforwarders, *f)
			}
		}
	}
	return &s, nil
}

func NewSAMManager(inifile, servhost, servport, samhost, samport string, start bool) (*SAMManager, error) {
	return NewSAMManagerFromOptions(
		SetManagerFilePath(inifile),
		SetManagerHost(servhost),
		SetManagerPort(servport),
		SetManagerSAMHost(samhost),
		SetManagerSAMPort(samport),
		SetManagerStart(start),
	)
}

func NewSAMManagerFromConf(conf *i2ptunconf.Conf, servhost, servport, samhost, samport string, start bool) (*SAMManager, error) {
	return NewSAMManagerFromOptions(
		SetManagerConf(conf),
		SetManagerHost(servhost),
		SetManagerPort(servport),
		SetManagerSAMHost(samhost),
		SetManagerSAMPort(samport),
		SetManagerStart(start),
	)
}
