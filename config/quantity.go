package i2ptunconf

// GetInQuantity takes an argument and a default. If the argument differs from the
// default, the argument is always returned. If the argument and default are
// the same and the key exists, the key is returned. If the key is absent, the
// default is returned.
func (c *Conf) GetInQuantity(arg, def int, label ...string) int {
	if arg != def {
		return arg
	}
	if c.Config == nil {
		return arg
	}
	if x, o := c.GetInt("inbound.quantity", label...); o {
		return x
	}
	return arg
}

// GetOutQuantity takes an argument and a default. If the argument differs from the
// default, the argument is always returned. If the argument and default are
// the same and the key exists, the key is returned. If the key is absent, the
// default is returned.
func (c *Conf) GetOutQuantity(arg, def int, label ...string) int {
	if arg != def {
		return arg
	}
	if c.Config == nil {
		return arg
	}
	if x, o := c.GetInt("outbound.quantity", label...); o {
		return x
	}
	return arg
}

// SetInQuantity sets the inbound tunnel quantity from config file
func (c *Conf) SetInQuantity(label ...string) {
	if v, ok := c.GetInt("inbound.quantity", label...); ok {
		c.InQuantity = v
	} else {
		c.InQuantity = 5
	}
}

// SetOutQuantity sets the outbound tunnel quantity from config file
func (c *Conf) SetOutQuantity(label ...string) {
	if v, ok := c.GetInt("outbound.quantity", label...); ok {
		c.OutQuantity = v
	} else {
		c.OutQuantity = 5
	}
}
