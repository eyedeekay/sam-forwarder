package i2ptunconf

// GetHost takes an argument and a default. If the argument differs from the
// default, the argument is always returned. If the argument and default are
// the same and the key exists, the key is returned. If the key is absent, the
// default is returned.
func (c *Conf) GetControlHost(arg, def string, label ...string) string {
	if arg != def {
		return arg
	}
	//if c.Config == nil {
	//	return arg
	//}
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
	//if c.Config == nil {
	//	return arg
	//}
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
