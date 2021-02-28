package i2ptunconf

import (
	"crypto/tls"
	"log"
)

// GetPort443 takes an argument and a default. If the argument differs from the
// default, the argument is always returned. If the argument and default are
// the same and the key exists, the key is returned. If the key is absent, the
// default is returned.
func (c *Conf) GetPort443(arg, def string, label ...string) string {
	if arg != def {
		return arg
	}
	if c.Config == nil {
		return arg
	}
	if x, o := c.Get("targetForPort.443", label...); o {
		return x
	}
	return arg
}

// SetTargetPort443 sets the port to forward from the config file
func (c *Conf) SetTargetPort443(label ...string) {
	if v, ok := c.Get("targetForPort.443", label...); ok {
		c.TargetForPort443 = v
	} else {
		c.TargetForPort443 = ""
	}
}

// Get
func (c *Conf) GetUseTLS(arg, def bool, label ...string) bool {
	if arg != def {
		return arg
	}
	if c.Config == nil {
		return arg
	}
	if x, o := c.GetBool("usetls", label...); o {
		return x
	}
	return arg
}

// SetAllowZeroHopOut sets the config to allow zero-hop tunnels
func (c *Conf) SetUseTLS(label ...string) {
	if v, ok := c.GetBool("usetls", label...); ok {
		c.UseTLS = v
	} else {
		c.UseTLS = false
	}
}

// GetTLSConfig
func (c *Conf) GetTLSConfig(arg, def string, label ...string) string {
	if arg != def {
		return arg
	}
	if c.Config == nil {
		return arg
	}
	if x, o := c.Get("cert", label...); o {
		return x
	}
	return arg
}

// SetClientDest sets the key name from the config file
func (c *Conf) SetTLSConfig(label ...string) {
	if v, ok := c.Get("cert", label...); ok {
		c.Cert = v
	} else {
		c.Cert = ""
	}
}

// GetTLSConfig
func (c *Conf) GetTLSConfigPem(arg, def string, label ...string) string {
	if arg != def {
		return arg
	}
	if c.Config == nil {
		return arg
	}
	if x, o := c.Get("pem", label...); o {
		return x
	}
	return arg
}

// SetClientDest sets the key name from the config file
func (c *Conf) SetTLSConfigPem(label ...string) {
	if v, ok := c.Get("pem", label...); ok {
		c.Pem = v
	} else {
		c.Pem = ""
	}
}

func (c *Conf) TLSConfig() *tls.Config {
	cert, err := tls.LoadX509KeyPair(c.Cert, c.Pem)
	if err != nil {
		log.Fatal(err)
	}
	return &tls.Config{Certificates: []tls.Certificate{cert}}
}
