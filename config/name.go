package i2ptunconf

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
	if _, ok := c.Get("keys", label...); ok {
		c.SaveFile = true
	} else {
		c.SaveFile = false
	}
}

// SetTunName sets the tunnel name from the config file
func (c *Conf) SetTunName(label ...string) {
	if v, ok := c.Get("keys", label...); ok {
		c.TunName = v
	} else {
		c.TunName = "fowarder"
	}
}
