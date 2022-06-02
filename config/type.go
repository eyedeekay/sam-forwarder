package i2ptunconf

import (
	"strings"
)

// GetType takes an argument and a default. If the argument differs from the
// default, the argument is always returned. If the argument and default are
// the same and the key exists, the key is returned. If the key is absent, the
// default is returned.
func (c *Conf) GetTypes(argc, argu, argh bool, def string, label ...string) string {
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
		if typ != def {
			return typ
		}
	}
	if def == "kcpclient" {
		return def
	}
	if def == "kcpserver" {
		return def
	}
	if def == "eephttpd" {
		return def
	}
	if def == "vpnclient" {
		return def
	}
	if def == "vpnserver" {
		return def
	}
	if def == "outproxy" {
		return def
	}
	if def == "outproxyhttp" {
		return def
	}
	if def == "browserclient" {
		return def
	}
	//if c.Config == nil {
	//	return typ
	//}
	if x, o := c.Get("type", label...); o {
		return x
	}
	return def
}

func (c *Conf) GetOtherType(typ, def string, label ...string) string {
	if typ != def {
		return typ
	}
	//if c.Config == nil {
	//	return typ
	//}
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
		case "eephttpd":
			c.Type = v
		case "outproxy":
			c.Type = v
		case "outproxyhttp":
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
