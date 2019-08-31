package i2ptunconf

import "github.com/eyedeekay/sam-forwarder/interface"

// GetSaveFile takes an argument and a default. If the argument differs from the
// default, the argument is always returned. If the argument and default are
// the same and the key exists, the key is returned. If the key is absent, the
// default is returned.
func (c *Conf) GetSaveFile(arg, def bool, label ...string) bool {
	if arg != def {
		return arg
	}
	if c.Config == nil {
		return arg
	}
	return c.SaveFile
}

// GetKeys takes an argument and a default. If the argument differs from the
// default, the argument is always returned. If the argument and default are
// the same and the key exists, the key is returned. If the key is absent, the
// default is returned.
func (c *Conf) GetKeys(arg, def string, label ...string) string {
	if arg != def {
		return arg
	}
	if c.Config == nil {
		return arg
	}
	if x, o := c.Get("keys", label...); o {
		return x
	}
	return arg
}

// SetKeys sets the key name from the config file
func (c *Conf) SetKeys(label ...string) {
	if v, ok := c.Get("keys", label...); ok {
		c.TunName = v
		c.SaveFile = true
	} else {
		c.TunName = "forwarder"
		c.SaveFile = false
	}
}

// SetTunName sets the tunnel name from the config file
func (c *Conf) SetTunName(label ...string) {
	if v, ok := c.Get("keys", label...); ok {
		c.TunName = v
	} else {
		c.TunName = "forwarder"
	}
}

//SetName sets the host of the Conf's SAM bridge
func SetName(s string) func(samtunnel.SAMTunnel) error {
	return func(c samtunnel.SAMTunnel) error {
		c.(*Conf).TunName = s
		return nil
	}
}

//SetSaveFile tells the router to save the tunnel's keys long-term
func SetSaveFile(b bool) func(samtunnel.SAMTunnel) error {
	return func(c samtunnel.SAMTunnel) error {
		c.(*Conf).SaveFile = b
		return nil
	}
}
