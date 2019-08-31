package i2ptunconf

import "github.com/eyedeekay/sam-forwarder/interface"

// GetKeyFile takes an argument and a default. If the argument differs from the
// default, the argument is always returned. If the argument and default are
// the same and the key exists, the key is returned. If the key is absent, the
// default is returned.
func (c *Conf) GetKeyFile(arg, def string, label ...string) string {
	if arg != def {
		return arg
	}
	if c.Config == nil {
		return arg
	}
	if x, o := c.Get("keyfile", label...); o {
		return x
	}
	return arg
}

// SetKeyFile sets the key save directory from the config file
func (c *Conf) SetKeyFile(label ...string) {
	if v, ok := c.Get("keyfile", label...); ok {
		c.KeyFilePath = v
	} else {
		c.KeyFilePath = "./"
	}
}

//SetKeyFile sets
func SetKeyFile(s string) func(samtunnel.SAMTunnel) error {
	return func(c samtunnel.SAMTunnel) error {
		c.(*Conf).KeyFilePath = s
		return nil
	}
}
