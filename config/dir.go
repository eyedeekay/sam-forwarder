package i2ptunconf

// GetDir takes an argument and a default. If the argument differs from the
// default, the argument is always returned. If the argument and default are
// the same and the key exists, the key is returned. If the key is absent, the
// default is returned.
func (c *Conf) GetDir(arg, def string, label ...string) string {
	if arg != def {
		return arg
	}
	if c.Config == nil {
		return arg
	}
	if x, o := c.Get("dir", label...); o {
		return x
	}
	return arg
}

// SetDir sets the key save directory from the config file
func (c *Conf) SetDir(label ...string) {
	if v, ok := c.Get("dir", label...); ok {
		c.SaveDirectory = v
	} else {
		c.SaveDirectory = "./"
	}
}
