package i2ptunconf

import "github.com/eyedeekay/sam-forwarder/interface"

// GetClientDest takes an argument and a default. If the argument differs from the
// default, the argument is always returned. If the argument and default are
// the same and the key exists, the key is returned. If the key is absent, the
// default is returned.
func (c *Conf) GetClientDest(arg, def string, label ...string) string {
	if arg != def {
		return arg
	}
	if c.Config == nil {
		return arg
	}
	if x, o := c.Get("destination", label...); o {
		return x
	}
	return arg
}

// SetClientDest sets the key name from the config file
func (c *Conf) SetClientDest(label ...string) {
	if v, ok := c.Get("destination", label...); ok {
		c.ClientDest = v
	} else {
		c.ClientDest = v
	}
}

//SetSaveFile tells the router to save the tunnel's keys long-term
func SetDestination(b string) func(samtunnel.SAMTunnel) error {
	return func(c samtunnel.SAMTunnel) error {
		c.(*Conf).ClientDest = b
		return nil
	}
}
