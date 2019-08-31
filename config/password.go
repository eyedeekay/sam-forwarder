package i2ptunconf

import "github.com/eyedeekay/sam-forwarder/interface"

// GetPassword takes an argument and a default. If the argument differs from the
// default, the argument is always returned. If the argument and default are
// the same and the key exists, the key is returned. If the key is absent, the
// default is returned.
func (c *Conf) GetPassword(arg, def string, label ...string) string {
	if arg != def {
		return arg
	}
	if c.Config == nil {
		return arg
	}
	if x, o := c.Get("username", label...); o {
		return x
	}
	return arg
}

// SetKeys sets the key name from the config file
func (c *Conf) SetPassword(label ...string) {
	if v, ok := c.Get("username", label...); ok {
		c.Password = v
	} else {
		c.Password = "samcatd"
	}
}

//SetPassword sets the host of the Conf's SAM bridge
func SetPassword(s string) func(samtunnel.SAMTunnel) error {
	return func(c samtunnel.SAMTunnel) error {
		c.(*Conf).Password = s
		return nil
	}
}
