package i2ptunconf

import (
	"strings"
)

// GetType takes an argument and a default. If the argument differs from the
// default, the argument is always returned. If the argument and default are
// the same and the key exists, the key is returned. If the key is absent, the
// default is returned.
func (c *Conf) GetType(argc, argu, argh bool, def string, label ...string) string {
	var typ string
	if argu {
		typ += "udp"
	}
	if argc {
		if argh == true {
			typ += "http"
		}
		typ += "client"
		c.Client = true
	} else {
		if argh == true {
			typ += "http"
		} else {
			typ += "server"
		}
	}
	if def == "kcpclient" {
		typ = "kcpclient"
	}
	if def == "kcpserver" {
		typ = "kcpserver"
	}
	if typ != def {
		return typ
	}
	if c.Config == nil {
		return typ
	}
	if x, o := c.Get("type", label...); o {
		return x
	}
	return def
}

func (c *Conf) GetOtherType(typ, def string, label ...string) string {
	if typ != def {
		return typ
	}
	if c.Config == nil {
		return typ
	}
	if x, o := c.Get("type", label...); o {
		return x
	}
	return def
}

// SetType sets the type of proxy to create from the config file
func (c *Conf) SetType(label ...string) {
	if v, ok := c.Get("type", label...); ok {
		if strings.Contains(v, "client") {
			c.Client = true
		}
		if c.Type == "server" || c.Type == "http" || c.Type == "client" || c.Type == "httpclient" || c.Type == "udpserver" || c.Type == "udpclient" || c.Type == "kcpclient" || c.Type == "kcpserver" {
			c.Type = v
		}
	} else {
		c.Type = "server"
	}
}
