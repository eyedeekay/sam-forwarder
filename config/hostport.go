package i2ptunconf

import (
	"fmt"
	"strconv"
)

// GetHost takes an argument and a default. If the argument differs from the
// default, the argument is always returned. If the argument and default are
// the same and the key exists, the key is returned. If the key is absent, the
// default is returned.
func (c *Conf) GetHost(arg, def string, label ...string) string {
	if arg != def {
		return arg
	}
	if c.Config == nil {
		return arg
	}
	if x, o := c.Get("host", label...); o {
		return x
	}
	return arg
}

// GetPort takes an argument and a default. If the argument differs from the
// default, the argument is always returned. If the argument and default are
// the same and the key exists, the key is returned. If the key is absent, the
// default is returned.
func (c *Conf) GetPort(arg, def string, label ...string) string {
	if arg != def {
		return arg
	}
	if c.Config == nil {
		return arg
	}
	if x, o := c.Get("port", label...); o {
		return x
	}
	return arg
}

// SetHost sets the host to forward from the config file
func (c *Conf) SetHost(label ...string) {
	if v, ok := c.Get("host", label...); ok {
		c.TargetHost = v
	} else {
		c.TargetHost = "127.0.0.1"
	}
}

// SetPort sets the port to forward from the config file
func (c *Conf) SetPort(label ...string) {
	if v, ok := c.Get("port", label...); ok {
		c.TargetPort = v
	} else {
		c.TargetPort = "8081"
	}
}

//SetHost sets the host of the service to forward
func SetHost(s string) func(*Conf) error {
	return func(c *Conf) error {
		c.TargetHost = s
		return nil
	}
}

//SetPort sets the port of the service to forward
func SetPort(s string) func(*Conf) error {
	return func(c *Conf) error {
		port, err := strconv.Atoi(s)
		if err != nil {
			return fmt.Errorf("Invalid TCP Server Target Port %s; non-number ", s)
		}
		if port < 65536 && port > -1 {
			c.TargetPort = s
			return nil
		}
		return fmt.Errorf("Invalid port")
	}
}
