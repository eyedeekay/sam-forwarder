package i2ptunconf

// GetEncryptLeaseset takes an argument and a default. If the argument differs from the
// default, the argument is always returned. If the argument and default are
// the same and the key exists, the key is returned. If the key is absent, the
// default is returned.
func (c *Conf) GetEncryptLeaseset(arg, def bool, label ...string) bool {
	if arg != def {
		return arg
	}
	if c.Config == nil {
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

// GetLeasesetKey takes an argument and a default. If the argument differs from the
// default, the argument is always returned. If the argument and default are
// the same and the key exists, the key is returned. If the key is absent, the
// default is returned.
func (c *Conf) GetLeasesetKey(arg, def string, label ...string) string {
	if arg != def {
		return arg
	}
	if c.Config == nil {
		return arg
	}
	if x, o := c.Get("i2cp.leaseSetKey", label...); o {
		return x
	}
	return arg
}

// SetEncryptLease tells the conf to use encrypted leasesets the from the config file
func (c *Conf) SetLeasesetKey(label ...string) {
	if v, ok := c.Get("i2cp.leaseSetKey", label...); ok {
		c.LeaseSetKey = v
	} else {
		c.LeaseSetKey = ""
	}
}

// GetLeasesetPrivateKey takes an argument and a default. If the argument differs from the
// default, the argument is always returned. If the argument and default are
// the same and the key exists, the key is returned. If the key is absent, the
// default is returned.
func (c *Conf) GetLeasesetPrivateKey(arg, def string, label ...string) string {
	if arg != def {
		return arg
	}
	if c.Config == nil {
		return arg
	}
	if x, o := c.Get("i2cp.leaseSetPrivateKey", label...); o {
		return x
	}
	return arg
}

// SetLeasesetPrivateKey tells the conf to use encrypted leasesets the from the config file
func (c *Conf) SetLeasesetPrivateKey(label ...string) {
	if v, ok := c.Get("i2cp.leaseSetPrivateKey", label...); ok {
		c.LeaseSetPrivateKey = v
	} else {
		c.LeaseSetPrivateKey = ""
	}
}

// GetLeasesetPrivateSigningKey takes an argument and a default. If the argument differs from the
// default, the argument is always returned. If the argument and default are
// the same and the key exists, the key is returned. If the key is absent, the
// default is returned.
func (c *Conf) GetLeasesetPrivateSigningKey(arg, def string, label ...string) string {
	if arg != def {
		return arg
	}
	if c.Config == nil {
		return arg
	}
	if x, o := c.Get("i2cp.leaseSetPrivateSigningKey", label...); o {
		return x
	}
	return arg
}

// SetLeasesetPrivateSigningKey tells the conf to use encrypted leasesets the from the config file
func (c *Conf) SetLeasesetPrivateSigningKey(label ...string) {
	if v, ok := c.Get("i2cp.leaseSetPrivateKey", label...); ok {
		c.LeaseSetPrivateSigningKey = v
	} else {
		c.LeaseSetPrivateSigningKey = ""
	}
}

//SetEncrypt tells the router to use an encrypted leaseset
func SetEncrypt(b bool) func(*Conf) error {
	return func(c *Conf) error {
		if b {
			c.EncryptLeaseSet = b //"true"
			return nil
		}
		c.EncryptLeaseSet = b //"false"
		return nil
	}
}

//SetLeaseSetKey sets the host of the Conf's SAM bridge
func SetLeaseSetKey(s string) func(*Conf) error {
	return func(c *Conf) error {
		c.LeaseSetKey = s
		return nil
	}
}

//SetLeaseSetPrivateKey sets the host of the Conf's SAM bridge
func SetLeaseSetPrivateKey(s string) func(*Conf) error {
	return func(c *Conf) error {
		c.LeaseSetPrivateKey = s
		return nil
	}
}

//SetLeaseSetPrivateSigningKey sets the host of the Conf's SAM bridge
func SetLeaseSetPrivateSigningKey(s string) func(*Conf) error {
	return func(c *Conf) error {
		c.LeaseSetPrivateSigningKey = s
		return nil
	}
}
