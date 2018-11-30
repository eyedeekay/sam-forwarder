package i2ptunconf

// GetFastRecieve takes an argument and a default. If the argument differs from the
// default, the argument is always returned. If the argument and default are
// the same and the key exists, the key is returned. If the key is absent, the
// default is returned.
func (c *Conf) GetFastRecieve(arg, def bool, label ...string) bool {
	if arg != def {
		return arg
	}
	if c.Config == nil {
		return arg
	}
	if x, o := c.GetBool("i2cp.fastRecieve", label...); o {
		return x
	}
	return arg
}

// SetFastRecieve sets the compression from the config file
func (c *Conf) SetFastRecieve(label ...string) {
	if v, ok := c.GetBool("i2cp.fastRecieve", label...); ok {
		c.FastRecieve = v
	} else {
		c.FastRecieve = false
	}
}
