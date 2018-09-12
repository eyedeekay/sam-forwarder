package sammanager

import (
	"context"
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
		for _, label := range s.config.Labels {
			if t, e := s.config.Get("type", label); e {
				switch t {
				case "http":
					log.Println("found http under", label)
				case "server":
					log.Println("found server under", label)
				case "client":
					log.Println("found client under", label)
				case "udpserver":
					log.Println("found udpserver under", label)
				case "udpclient":
					log.Println("found udpclient under", label)
				}
			}
		}
	}
	return &s, nil
}

func NewSAMManager(inifile, servhost, servport, samhost, samport string) (*SAMManager, error) {
	return NewSAMManagerFromOptions(
		SetManagerFilePath(inifile),
		SetManagerHost(servhost),
		SetManagerPort(servport),
		SetManagerSAMHost(samhost),
		SetManagerSAMPort(samport),
	)
}
