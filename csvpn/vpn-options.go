package samforwardervpn

import (
	"github.com/eyedeekay/sam-forwarder/config"
)

//Option is a SAMClientServerVPN Option
type Option func(*SAMClientServerVPN) error

func SetFilePath(s string) func(*SAMClientServerVPN) error {
	return func(c *SAMClientServerVPN) error {
		c.FilePath = s
		return nil
	}
}

func SetVPNConfig(s *i2ptunconf.Conf) func(*SAMClientServerVPN) error {
	return func(c *SAMClientServerVPN) error {
		c.Config = s
		return nil
	}
}
