package i2ptunconf



// GetUserName takes an argument and a default. If the argument differs from the
// default, the argument is always returned. If the argument and default are
// the same and the key exists, the key is returned. If the key is absent, the
// default is returned.
func (c *Conf) GetUserName(arg, def string, label ...string) string {
	if arg != def {
		return arg
	}
	if c.Config == nil {
		return arg
	}
	if x, o := c.Get("username", label...); o {
		return x
	}
	return arg
}

// SetKeys sets the key name from the config file
func (c *Conf) SetUserName(label ...string) {
	if v, ok := c.Get("username", label...); ok {
		c.UserName = v
	} else {
		c.UserName = "samcatd"
	}
}
