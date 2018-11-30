package i2ptunconf

// GetInLength takes an argument and a default. If the argument differs from the
// default, the argument is always returned. If the argument and default are
// the same and the key exists, the key is returned. If the key is absent, the
// default is returned.
func (c *Conf) GetInLength(arg, def int, label ...string) int {
	if arg != def {
		return arg
	}
	if c.Config == nil {
		return arg
	}
	if x, o := c.GetInt("inbound.length", label...); o {
		return x
	}
	return arg
}

// GetOutLength takes an argument and a default. If the argument differs from the
// default, the argument is always returned. If the argument and default are
// the same and the key exists, the key is returned. If the key is absent, the
// default is returned.
func (c *Conf) GetOutLength(arg, def int, label ...string) int {
	if arg != def {
		return arg
	}
	if c.Config == nil {
		return arg
	}
	if x, o := c.GetInt("outbound.length", label...); o {
		return x
	}
	return arg
}

// SetInLength sets the inbound length from the config file
func (c *Conf) SetInLength(label ...string) {
	if v, ok := c.GetInt("outbound.length", label...); ok {
		c.OutLength = v
	} else {
		c.OutLength = 3
	}
}

// SetOutLength sets the outbound lenth from the config file
func (c *Conf) SetOutLength(label ...string) {
	if v, ok := c.GetInt("inbound.length", label...); ok {
		c.InLength = v
	} else {
		c.InLength = 3
	}
}
