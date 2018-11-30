package i2ptunconf

// GetPort443 takes an argument and a default. If the argument differs from the
// default, the argument is always returned. If the argument and default are
// the same and the key exists, the key is returned. If the key is absent, the
// default is returned.
func (c *Conf) GetPort443(arg, def string, label ...string) string {
	if arg != def {
		return arg
	}
	if c.Config == nil {
		return arg
	}
	if x, o := c.Get("targetForPort.443", label...); o {
		return x
	}
	return arg
}

// SetTargetPort443 sets the port to forward from the config file
func (c *Conf) SetTargetPort443(label ...string) {
	if v, ok := c.Get("targetForPort.443", label...); ok {
		c.TargetForPort443 = v
	} else {
		c.TargetForPort443 = ""
	}
}
