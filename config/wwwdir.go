package i2ptunconf

//

// GetDir takes an argument and a default. If the argument differs from the
// default, the argument is always returned. If the argument and default are
// the same and the key exists, the key is returned. If the key is absent, the
// default is returned.
func (c *Conf) GetWWWDir(arg, def string, label ...string) string {
	if arg != def {
		return arg
	}
	if c.Config == nil {
		return arg
	}
	if x, o := c.Get("wwwdir", label...); o {
		return x
	}
	return arg
}

// SetDir sets the key save directory from the config file
func (c *Conf) SetWWWDir(label ...string) {
	if v, ok := c.Get("wwwdir", label...); ok {
		c.ServeDirectory = v
	} else {
		c.ServeDirectory = "./www"
	}
}
