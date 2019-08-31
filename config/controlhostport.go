package i2ptunconf

import (
	"fmt"
	"strconv"
)

import "github.com/eyedeekay/sam-forwarder/interface"

// GetHost takes an argument and a default. If the argument differs from the
// default, the argument is always returned. If the argument and default are
// the same and the key exists, the key is returned. If the key is absent, the
// default is returned.
func (c *Conf) GetControlHost(arg, def string, label ...string) string {
	if arg != def {
		return arg
	}
	if c.Config == nil {
		return arg
	}
	if x, o := c.Get("controlhost", label...); o {
		return x
	}
	return arg
}

// GetPort takes an argument and a default. If the argument differs from the
// default, the argument is always returned. If the argument and default are
// the same and the key exists, the key is returned. If the key is absent, the
// default is returned.
func (c *Conf) GetControlPort(arg, def string, label ...string) string {
	if arg != def {
		return arg
	}
	if c.Config == nil {
		return arg
	}
	if x, o := c.Get("controlport", label...); o {
		return x
	}
	return arg
}

// SetHost sets the host to forward from the config file
func (c *Conf) SetControlHost(label ...string) {
	if v, ok := c.Get("controlhost", label...); ok {
		c.ControlHost = v
	} else {
		c.ControlHost = "127.0.0.1"
	}
}

// SetPort sets the port to forward from the config file
func (c *Conf) SetControlPort(label ...string) {
	if v, ok := c.Get("controlport", label...); ok {
		c.ControlPort = v
	} else {
		c.ControlPort = ""
	}
}

//SetControlHost sets the host of the service to forward
func SetControlHost(s string) func(samtunnel.SAMTunnel) error {
	return func(c samtunnel.SAMTunnel) error {
		c.(*Conf).ControlHost = s
		return nil
	}
}

//SetControlPort sets the port of the service to forward
func SetControlPort(s string) func(samtunnel.SAMTunnel) error {
	return func(c samtunnel.SAMTunnel) error {
		port, err := strconv.Atoi(s)
		if err != nil {
			return fmt.Errorf("Invalid TCP Server Target Port %s; non-number ", s)
		}
		if port < 65536 && port > -1 {
			c.(*Conf).ControlPort = s
			return nil
		}
		return fmt.Errorf("Invalid port")
	}
}
