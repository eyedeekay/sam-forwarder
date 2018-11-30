package i2ptunconf

//i2cp.messageReliability
// GetMessageReliability takes an argument and a default. If the argument differs from the
// default, the argument is always returned. If the argument and default are
// the same and the key exists, the key is returned. If the key is absent, the
// default is returned.
func (c *Conf) GetMessageReliability(arg, def string, label ...string) string {
	if arg != def {
		return arg
	}
	if c.Config == nil {
		return arg
	}
	return c.MessageReliability
}

// SetMessageReliability sets the access list type from a config file
func (c *Conf) SetMessageReliability(label ...string) {
	if v, ok := c.Get("i2cp.messageReliability", label...); ok {
		c.MessageReliability = v
	}
	if c.MessageReliability != "BestEffort" && c.MessageReliability != "none" {
		c.MessageReliability = "none"
	}
}

func (c *Conf) reliability() string {
	return "i2cp.messageReliability" + c.MessageReliability
}
