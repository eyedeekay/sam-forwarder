package i2ptunconf

import "fmt"

import "github.com/eyedeekay/sam-forwarder/interface"

// GetInVariance takes an argument and a default. If the argument differs from the
// default, the argument is always returned. If the argument and default are
// the same and the key exists, the key is returned. If the key is absent, the
// default is returned.
func (c *Conf) GetInVariance(arg, def int, label ...string) int {
	if arg != def {
		return arg
	}
	if c.Config == nil {
		return arg
	}
	if x, o := c.GetInt("inbound.variance", label...); o {
		return x
	}
	return arg
}

// GetOutVariance takes an argument and a default. If the argument differs from the
// default, the argument is always returned. If the argument and default are
// the same and the key exists, the key is returned. If the key is absent, the
// default is returned.
func (c *Conf) GetOutVariance(arg, def int, label ...string) int {
	if arg != def {
		return arg
	}
	if c.Config == nil {
		return arg
	}
	if x, o := c.GetInt("outbound.variance", label...); o {
		return x
	}
	return arg
}

// SetInVariance sets the inbound tunnel variance from config file
func (c *Conf) SetInVariance(label ...string) {
	if v, ok := c.GetInt("inbound.variance", label...); ok {
		c.InVariance = v
	} else {
		c.InVariance = 0
	}
}

// SetOutVariance sets the outbound tunnel variance from config file
func (c *Conf) SetOutVariance(label ...string) {
	if v, ok := c.GetInt("outbound.variance", label...); ok {
		c.OutVariance = v
	} else {
		c.OutVariance = 0
	}
}

//SetInVariance sets the variance of a number of hops inbound
func SetInVariance(i int) func(samtunnel.SAMTunnel) error {
	return func(c samtunnel.SAMTunnel) error {
		if i < 7 && i > -7 {
			c.(*Conf).InVariance = i //strconv.Itoa(i)
			return nil
		}
		return fmt.Errorf("Invalid inbound tunnel length")
	}
}

//SetOutVariance sets the variance of a number of hops outbound
func SetOutVariance(i int) func(samtunnel.SAMTunnel) error {
	return func(c samtunnel.SAMTunnel) error {
		if i < 7 && i > -7 {
			c.(*Conf).OutVariance = i //strconv.Itoa(i)
			return nil
		}
		return fmt.Errorf("Invalid outbound tunnel variance")
	}
}
