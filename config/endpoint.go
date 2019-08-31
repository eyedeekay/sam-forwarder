package i2ptunconf

// GetEndpointHost takes an argument and a default. If the argument differs from the
// default, the argument is always returned. If the argument and default are
// the same and the key exists, the key is returned. If the key is absent, the
// default is returned.
func (c *Conf) GetEndpointHost(arg, def string, label ...string) string {
	if arg != def {
		return arg
	}
	if c.Config == nil {
		return arg
	}
	if x, o := c.Get("tunhost", label...); o {
		return x
	}
	return arg
}

// SetEndpointHost sets the host to forward from the config file
func (c *Conf) SetEndpointHost(label ...string) {
	if v, ok := c.Get("tunhost", label...); ok {
		c.TunnelHost = v
	} else {
		c.TunnelHost = "10.79.0.1"
	}
}
