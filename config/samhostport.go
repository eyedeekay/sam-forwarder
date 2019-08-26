package i2ptunconf

import (
	"fmt"
	"strconv"
)

// GetSAMHost takes an argument and a default. If the argument differs from the
// default, the argument is always returned. If the argument and default are
// the same and the key exists, the key is returned. If the key is absent, the
// default is returned.
func (c *Conf) GetSAMHost(arg, def string, label ...string) string {
	if arg != def {
		return arg
	}
	if c.Config == nil {
		return arg
	}
	if x, o := c.Get("samhost", label...); o {
		return x
	}
	return arg
}

// GetSAMPort takes an argument and a default. If the argument differs from the
// default, the argument is always returned. If the argument and default are
// the same and the key exists, the key is returned. If the key is absent, the
// default is returned.
func (c *Conf) GetSAMPort(arg, def string, label ...string) string {
	if arg != def {
		return arg
	}
	if c.Config == nil {
		return arg
	}
	if x, o := c.Get("samport", label...); o {
		return x
	}
	return arg
}

// SetSAMHost sets the SAM host from the config file
func (c *Conf) SetSAMHost(label ...string) {
	if v, ok := c.Get("samhost", label...); ok {
		c.SamHost = v
	} else {
		c.SamHost = "127.0.0.1"
	}
}

// SetSAMPort sets the SAM port from the config file
func (c *Conf) SetSAMPort(label ...string) {
	if v, ok := c.Get("samport", label...); ok {
		c.SamPort = v
	} else {
		c.SamPort = "7656"
	}
}

//SetSAMHost sets the host of the Conf's SAM bridge
func SetSAMHost(s string) func(*Conf) error {
	return func(c *Conf) error {
		c.SamHost = s
		return nil
	}
}

//SetSAMPort sets the port of the Conf's SAM bridge using a string
func SetSAMPort(s string) func(*Conf) error {
	return func(c *Conf) error {
		port, err := strconv.Atoi(s)
		if err != nil {
			return fmt.Errorf("Invalid SAM Port %s; non-number", s)
		}
		if port < 65536 && port > -1 {
			c.SamPort = s
			return nil
		}
		return fmt.Errorf("Invalid port")
	}
}
