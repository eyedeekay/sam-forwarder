package i2ptunconf

import "fmt"

// GetInBackups takes an argument and a default. If the argument differs from the
// default, the argument is always returned. If the argument and default are
// the same and the key exists, the key is returned. If the key is absent, the
// default is returned.
func (c *Conf) GetInBackups(arg, def int, label ...string) int {
	if arg != def {
		return arg
	}
	if c.Config == nil {
		return arg
	}
	if x, o := c.GetInt("inbound.backupQuantity", label...); o {
		return x
	}
	return arg
}

// GetOutBackups takes an argument and a default. If the argument differs from the
// default, the argument is always returned. If the argument and default are
// the same and the key exists, the key is returned. If the key is absent, the
// default is returned.
func (c *Conf) GetOutBackups(arg, def int, label ...string) int {
	if arg != def {
		return arg
	}
	if c.Config == nil {
		return arg
	}
	if x, o := c.GetInt("outbound.backupQuantity", label...); o {
		return x
	}
	return arg
}

// SetInBackups sets the inbound tunnel backups from config file
func (c *Conf) SetInBackups(label ...string) {
	if v, ok := c.GetInt("inbound.backupQuantity", label...); ok {
		c.InBackupQuantity = v
	} else {
		c.InBackupQuantity = 2
	}
}

// SetOutBackups sets the outbound tunnel backups from config file
func (c *Conf) SetOutBackups(label ...string) {
	if v, ok := c.GetInt("outbound.backupQuantity", label...); ok {
		c.OutBackupQuantity = v
	} else {
		c.OutBackupQuantity = 2
	}
}

//SetInBackups sets the inbound tunnel backups
func SetInBackups(u int) func(*Conf) error {
	return func(c *Conf) error {
		if u < 6 && u >= 0 {
			c.InBackupQuantity = u // strconv.Itoa(u)
			return nil
		}
		return fmt.Errorf("Invalid inbound tunnel backup quantity")
	}
}

//SetOutBackups sets the inbound tunnel backups
func SetOutBackups(u int) func(*Conf) error {
	return func(c *Conf) error {
		if u < 6 && u >= 0 {
			c.OutBackupQuantity = u // strconv.Itoa(u)
			return nil
		}
		return fmt.Errorf("Invalid outbound tunnel backup quantity")
	}
}
