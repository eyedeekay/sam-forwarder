package i2ptunconf

import (
	"fmt"
)

// GetReduceOnIdle takes an argument and a default. If the argument differs from the
// default, the argument is always returned. If the argument and default are
// the same and the key exists, the key is returned. If the key is absent, the
// default is returned.
func (c *Conf) GetReduceOnIdle(arg, def bool, label ...string) bool {
	if arg != def {
		return arg
	}
	if c.Config == nil {
		return arg
	}
	if x, o := c.GetBool("i2cp.reduceOnIdle", label...); o {
		return x
	}
	return arg
}

// GetReduceIdleTime takes an argument and a default. If the argument differs from the
// default, the argument is always returned. If the argument and default are
// the same and the key exists, the key is returned. If the key is absent, the
// default is returned.
func (c *Conf) GetReduceIdleTime(arg, def int, label ...string) int {
	if arg != def {
		return arg
	}
	if c.Config == nil {
		return arg
	}
	if x, o := c.GetInt("i2cp.reduceIdleTime", label...); o {
		return x
	}
	return arg
}

// GetReduceIdleQuantity takes an argument and a default. If the argument differs from the
// default, the argument is always returned. If the argument and default are
// the same and the key exists, the key is returned. If the key is absent, the
// default is returned.
func (c *Conf) GetReduceIdleQuantity(arg, def int, label ...string) int {
	if arg != def {
		return arg
	}
	if c.Config == nil {
		return arg
	}
	if x, o := c.GetInt("i2cp.reduceIdleQuantity", label...); o {
		return x
	}
	return arg
}

// GetCloseOnIdle takes an argument and a default. If the argument differs from the
// default, the argument is always returned. If the argument and default are
// the same and the key exists, the key is returned. If the key is absent, the
// default is returned.
func (c *Conf) GetCloseOnIdle(arg, def bool, label ...string) bool {
	if arg != def {
		return arg
	}
	if c.Config == nil {
		return arg
	}
	if x, o := c.GetBool("i2cp.closeOnIdle", label...); o {
		return x
	}
	return arg
}

// GetCloseIdleTime takes an argument and a default. If the argument differs from the
// default, the argument is always returned. If the argument and default are
// the same and the key exists, the key is returned. If the key is absent, the
// default is returned.
func (c *Conf) GetCloseIdleTime(arg, def int, label ...string) int {
	if arg != def {
		return arg
	}
	if c.Config == nil {
		return arg
	}
	if x, o := c.GetInt("i2cp.closeIdleTime", label...); o {
		return x
	}
	return arg
}

// SetReduceIdle sets the config to reduce tunnels after idle time from config file
func (c *Conf) SetReduceIdle(label ...string) {
	if v, ok := c.GetBool("i2cp.reduceOnIdle", label...); ok {
		c.ReduceIdle = v
	} else {
		c.ReduceIdle = false
	}
}

// SetReduceIdleTime sets the time to wait before reducing tunnels from config file
func (c *Conf) SetReduceIdleTime(label ...string) {
	if v, ok := c.GetInt("i2cp.reduceIdleTime", label...); ok {
		c.ReduceIdleTime = v
	} else {
		c.ReduceIdleTime = 300000
	}
}

// SetReduceIdleQuantity sets the number of tunnels to reduce to from config file
func (c *Conf) SetReduceIdleQuantity(label ...string) {
	if v, ok := c.GetInt("i2cp.reduceQuantity", label...); ok {
		c.ReduceIdleQuantity = v
	} else {
		c.ReduceIdleQuantity = 3
	}
}

// SetCloseIdle sets the tunnel to automatically close on idle from the config file
func (c *Conf) SetCloseIdle(label ...string) {
	if v, ok := c.GetBool("i2cp.closeOnIdle", label...); ok {
		c.CloseIdle = v
	} else {
		c.CloseIdle = false
	}
}

// SetCloseIdleTime sets the time to wait before killing a tunnel from a config file
func (c *Conf) SetCloseIdleTime(label ...string) {
	if v, ok := c.GetInt("i2cp.closeIdleTime", label...); ok {
		c.CloseIdleTime = v
	} else {
		c.CloseIdleTime = 300000
	}
}

//SetReduceIdle tells the connection to reduce it's tunnels during extended idle time.
func SetReduceIdle(b bool) func(*Conf) error {
	return func(c *Conf) error {
		if b {
			c.ReduceIdle = b // "true"
			return nil
		}
		c.ReduceIdle = b // "false"
		return nil
	}
}

//SetReduceIdleTime sets the time to wait before reducing tunnels to idle levels
func SetReduceIdleTime(u int) func(*Conf) error {
	return func(c *Conf) error {
		c.ReduceIdleTime = 300000
		if u >= 6 {
			c.ReduceIdleTime = (u * 60) * 1000 // strconv.Itoa((u * 60) * 1000)
			return nil
		}
		return fmt.Errorf("Invalid reduce idle timeout(Measured in minutes) %v", u)
	}
}

//SetReduceIdleTimeMs sets the time to wait before reducing tunnels to idle levels in milliseconds
func SetReduceIdleTimeMs(u int) func(*Conf) error {
	return func(c *Conf) error {
		c.ReduceIdleTime = 300000
		if u >= 300000 {
			c.ReduceIdleTime = u // strconv.Itoa(u)
			return nil
		}
		return fmt.Errorf("Invalid reduce idle timeout(Measured in milliseconds) %v", u)
	}
}

//SetReduceIdleQuantity sets minimum number of tunnels to reduce to during idle time
func SetReduceIdleQuantity(u int) func(*Conf) error {
	return func(c *Conf) error {
		if u < 5 {
			c.ReduceIdleQuantity = u // strconv.Itoa(u)
			return nil
		}
		return fmt.Errorf("Invalid reduce tunnel quantity")
	}
}

//SetCloseIdle tells the connection to close it's tunnels during extended idle time.
func SetCloseIdle(b bool) func(*Conf) error {
	return func(c *Conf) error {
		if b {
			c.CloseIdle = b // "true"
			return nil
		}
		c.CloseIdle = b // "false"
		return nil
	}
}

//SetCloseIdleTime sets the time to wait before closing tunnels to idle levels
func SetCloseIdleTime(u int) func(*Conf) error {
	return func(c *Conf) error {
		c.CloseIdleTime = 300000
		if u >= 6 {
			c.CloseIdleTime = (u * 60) * 1000 // strconv.Itoa((u * 60) * 1000)
			return nil
		}
		return fmt.Errorf("Invalid close idle timeout(Measured in minutes) %v", u)
	}
}

//SetCloseIdleTimeMs sets the time to wait before closing tunnels to idle levels in milliseconds
func SetCloseIdleTimeMs(u int) func(*Conf) error {
	return func(c *Conf) error {
		c.CloseIdleTime = 300000
		if u >= 300000 {
			c.CloseIdleTime = u //strconv.Itoa(u)
			return nil
		}
		return fmt.Errorf("Invalid close idle timeout(Measured in milliseconds) %v", u)
	}
}
