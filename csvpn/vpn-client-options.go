package samforwardervpn

import (
	"github.com/eyedeekay/sam-forwarder/config"
)

//ClientOption is a SAMClientServerVPN Option
type ClientOption func(*SAMClientVPN) error

func SetClientFilePath(s string) func(*SAMClientVPN) error {
	return func(c *SAMClientVPN) error {
		c.FilePath = s
		return nil
	}
}

func SetClientVPNConfig(s *i2ptunconf.Conf) func(*SAMClientVPN) error {
	return func(c *SAMClientVPN) error {
		c.Config = s
		return nil
	}
}

func SetClientDest(s string) func(*SAMClientVPN) error {
	return func(c *SAMClientVPN) error {
		c.ClientDest = s
		return nil
	}
}
