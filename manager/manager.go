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
	s.FilePath = "tunnels.ini"
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
	return nil, nil
}

func (s *SAMManager) NewSAMManager(inifile, servhost, servport, samhost, samport string) (*SAMManager, error) {
	return NewSamManagerFromOptions(
		SetManagerFilePath(inifile),
		SetManagerHost(servhost),
		SetManagerPort(servport),
		SetManagerSAMHost(samhost),
		SetManagerSAMPort(samport),
	)
}
