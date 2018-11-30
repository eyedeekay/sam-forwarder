package i2ptunconf

// GetHost takes an argument and a default. If the argument differs from the
// default, the argument is always returned. If the argument and default are
// the same and the key exists, the key is returned. If the key is absent, the
// default is returned.
func (c *Conf) GetHost(arg, def string, label ...string) string {
	if arg != def {
		return arg
	}
	if c.Config == nil {
		return arg
	}
	if x, o := c.Get("host", label...); o {
		return x
	}
	return arg
}

// GetPort takes an argument and a default. If the argument differs from the
// default, the argument is always returned. If the argument and default are
// the same and the key exists, the key is returned. If the key is absent, the
// default is returned.
func (c *Conf) GetPort(arg, def string, label ...string) string {
	if arg != def {
		return arg
	}
	if c.Config == nil {
		return arg
	}
	if x, o := c.Get("port", label...); o {
		return x
	}
	return arg
}

// SetHost sets the host to forward from the config file
func (c *Conf) SetHost(label ...string) {
	if v, ok := c.Get("host", label...); ok {
		c.TargetHost = v
	} else {
		c.TargetHost = "127.0.0.1"
	}
}

// SetPort sets the port to forward from the config file
func (c *Conf) SetPort(label ...string) {
	if v, ok := c.Get("port", label...); ok {
		c.TargetPort = v
	} else {
		c.TargetPort = "8081"
	}
}
