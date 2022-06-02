package i2ptunconf

import (
	"strings"
)

// GetAccessListType takes an argument and a default. If the argument differs from the
// default, the argument is always returned. If the argument and default are
// the same and the key exists, the key is returned. If the key is absent, the
// default is returned.
func (c *Conf) GetAccessListType(arg, def string, label ...string) string {
	if arg != def {
		return arg
	}
	//if c.Config == nil {
	//	return arg
	//}
	return c.AccessListType
}

// SetAccessListType sets the access list type from a config file
func (c *Conf) SetAccessListType(label ...string) {
	if v, ok := c.GetBool("i2cp.enableBlackList", label...); ok {
		if v {
			c.AccessListType = "blocklist"
		}
	}
	if v, ok := c.GetBool("i2cp.enableAccessList", label...); ok {
		if v {
			c.AccessListType = "allowlist"
		}
	}
	if c.AccessListType != "allowlist" && c.AccessListType != "blocklist" {
		c.AccessListType = "none"
	}
}

// AddAccessListMember adds a member to either the blocklist or the allowlist
func (c *Conf) AddAccessListMember(key string) {
	for _, item := range c.AccessList {
		if item == key {
			return
		}
	}
	c.AccessList = append(c.AccessList, key)
}

func (c *Conf) accesslisttype() string {
	if c.AccessListType == "allowlist" {
		return "i2cp.enableAccessList=true"
	} else if c.AccessListType == "blocklist" {
		return "i2cp.enableBlackList=true"
	} else if c.AccessListType == "none" {
		return ""
	}
	return ""
}

func (c *Conf) accesslist() string {
	if c.AccessListType != "" && len(c.AccessList) > 0 {
		r := ""
		for _, s := range c.AccessList {
			r += s + ","
		}
		return "i2cp.accessList=" + strings.TrimSuffix(r, ",")
	}
	return ""
}
