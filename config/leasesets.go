package i2ptunconf

// GetEncryptLeaseset takes an argument and a default. If the argument differs from the
// default, the argument is always returned. If the argument and default are
// the same and the key exists, the key is returned. If the key is absent, the
// default is returned.
func (c *Conf) GetEncryptLeaseset(arg, def bool, label ...string) bool {
	if arg != def {
		return arg
	}
	if c.config == nil {
		return arg
	}
	if x, o := c.GetBool("i2cp.encryptLeaseSet", label...); o {
		return x
	}
	return arg
}

// SetEncryptLease tells the conf to use encrypted leasesets the from the config file
func (c *Conf) SetEncryptLease(label ...string) {
	if v, ok := c.GetBool("i2cp.encryptLeaseSet", label...); ok {
		c.EncryptLeaseSet = v
	} else {
		c.EncryptLeaseSet = false
	}
}
