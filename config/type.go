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
		switch c.Type {
		case "server":
			c.Type = v
		case "http":
			c.Type = v
		case "client":
			c.Type = v
		case "httpclient":
			c.Type = v
		case "browserclient":
			c.Type = v
		case "udpserver":
			c.Type = v
		case "udpclient":
			c.Type = v
		case "vpnserver":
			c.Type = v
		case "vpnclient":
			c.Type = v
		case "kcpclient":
			c.Type = v
		case "kcpserver":
			c.Type = v
		default:
			c.Type = "browserclient"
		}
	} else {
		c.Type = "browserclient"
	}
}

//SetType sets the type of the forwarder server
func SetType(s string) func(*Conf) error {
	return func(c *Conf) error {
		switch c.Type {
		case "server":
			c.Type = s
		case "http":
			c.Type = s
		case "client":
			c.Type = s
		case "httpclient":
			c.Type = s
		case "browserclient":
			c.Type = s
		case "udpserver":
			c.Type = s
		case "udpclient":
			c.Type = s
		case "vpnserver":
			c.Type = s
		case "vpnclient":
			c.Type = s
		case "kcpclient":
			c.Type = s
		case "kcpserver":
			c.Type = s
		default:
			c.Type = "browserclient"
		}
		return nil
	}
}
