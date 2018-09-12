package sammanager

import (
	"context"
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

	TargetHost string
	TargetPort string
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
	return false, ""
}

func (s *SAMManager) Dial(ctx context.Context, network, address string) (*net.Conn, error) {
	return nil, nil
}

func (s *SAMManager) NewSAMManager(opts ...func(*SAMManager) error) (*SAMManager, error) {
	return nil, nil
}

func (s *SAMManager) NewSAMManagerStrings(servhost, servpost samhost, samport string) (*SAMManager, error) {
	return nil, nil
}
