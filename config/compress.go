package i2ptunconf

import "github.com/eyedeekay/sam-forwarder/interface"

// GetUseCompression takes an argument and a default. If the argument differs from the
// default, the argument is always returned. If the argument and default are
// the same and the key exists, the key is returned. If the key is absent, the
// default is returned.
func (c *Conf) GetUseCompression(arg, def bool, label ...string) bool {
	if arg != def {
		return arg
	}
	if c.Config == nil {
		return arg
	}
	if x, o := c.GetBool("gzip", label...); o {
		return x
	}
	return arg
}

// SetCompressed sets the compression from the config file
func (c *Conf) SetCompressed(label ...string) {
	if v, ok := c.GetBool("gzip", label...); ok {
		c.UseCompression = v
	} else {
		c.UseCompression = true
	}
}

//SetCompress tells clients to use compression
func SetCompress(b bool) func(samtunnel.SAMTunnel) error {
	return func(c samtunnel.SAMTunnel) error {
		if b {
			c.(*Conf).UseCompression = b // "true"
			return nil
		}
		c.(*Conf).UseCompression = b // "false"
		return nil
	}
}
