package i2ptunconf

import (
	"crypto/tls"
	"strings"

	i2ptls "github.com/eyedeekay/sam-forwarder/tls"
)

// GetPort443 takes an argument and a default. If the argument differs from the
// default, the argument is always returned. If the argument and default are
// the same and the key exists, the key is returned. If the key is absent, the
// default is returned.
func (c *Conf) GetPort443(arg, def string, label ...string) string {
	if arg != def {
		return arg
	}
	//if c.Config == nil {
	//	return arg
	//}
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
	//if c.Config == nil {
	//	return arg
	//}
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
func (c *Conf) GetTLSConfigCertPem(arg, def string, label ...string) string {
	if arg != def {
		return arg
	}
	//if c.Config == nil {
	//	return arg
	//}
	if x, o := c.Get("cert.pem", label...); o {
		return x
	}
	return arg
}

// SetClientDest sets the key name from the config file
func (c *Conf) SetTLSConfigCertPem(label ...string) {
	if v, ok := c.Get("cert.pem", label...); ok {
		c.Cert = v
	} else {
		c.Cert = ""
	}
}

// GetTLSConfig
func (c *Conf) GetTLSConfigKeyPem(arg, def string, label ...string) string {
	if arg != def {
		return arg
	}
	//if c.Config == nil {
	//	return arg
	//}
	if x, o := c.Get("key.pem", label...); o {
		return x
	}
	return arg
}

// SetClientDest sets the key name from the config file
func (c *Conf) SetTLSConfigKeyPem(label ...string) {
	if v, ok := c.Get("key.pem", label...); ok {
		c.Pem = v
	} else {
		c.Pem = ""
	}
}

func (c *Conf) TLSConfig() (*tls.Config, error) {
	names := []string{c.Base32()}
	if c.HostName != "" && strings.HasSuffix(c.HostName, ".i2p") {
		names = append(names, c.HostName)
	}
	if len(c.Cert) < 1 {
		c.Cert = "cert.pem"
	}
	if len(c.Pem) < 1 {
		c.Pem = "key.pem"
	}
	return i2ptls.TLSConfig(c.Cert, c.Pem, names)
}
