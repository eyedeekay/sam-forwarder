package i2ptunconf

// GetInAllowZeroHop takes an argument and a default. If the argument differs from the
// default, the argument is always returned. If the argument and default are
// the same and the key exists, the key is returned. If the key is absent, the
// default is returned.
func (c *Conf) GetInAllowZeroHop(arg, def bool, label ...string) bool {
	if arg != def {
		return arg
	}
	if c.Config == nil {
		return arg
	}
	if x, o := c.GetBool("inbound.allowZeroHop", label...); o {
		return x
	}
	return arg
}

// GetOutAllowZeroHop takes an argument and a default. If the argument differs from the
// default, the argument is always returned. If the argument and default are
// the same and the key exists, the key is returned. If the key is absent, the
// default is returned.
func (c *Conf) GetOutAllowZeroHop(arg, def bool, label ...string) bool {
	if arg != def {
		return arg
	}
	if c.Config == nil {
		return arg
	}
	if x, o := c.GetBool("outbound.allowZeroHop", label...); o {
		return x
	}
	return arg
}

// SetAllowZeroHopIn sets the config to allow zero-hop tunnels
func (c *Conf) SetAllowZeroHopIn(label ...string) {
	if v, ok := c.GetBool("inbound.allowZeroHop", label...); ok {
		c.InAllowZeroHop = v
	} else {
		c.InAllowZeroHop = false
	}
}

// SetAllowZeroHopOut sets the config to allow zero-hop tunnels
func (c *Conf) SetAllowZeroHopOut(label ...string) {
	if v, ok := c.GetBool("outbound.allowZeroHop", label...); ok {
		c.OutAllowZeroHop = v
	} else {
		c.OutAllowZeroHop = false
	}
}
