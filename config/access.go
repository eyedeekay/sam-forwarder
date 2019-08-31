package i2ptunconf

import (
	"fmt"
	"strings"
)

import "github.com/eyedeekay/sam-forwarder/interface"

// GetAccessListType takes an argument and a default. If the argument differs from the
// default, the argument is always returned. If the argument and default are
// the same and the key exists, the key is returned. If the key is absent, the
// default is returned.
func (c *Conf) GetAccessListType(arg, def string, label ...string) string {
	if arg != def {
		return arg
	}
	if c.Config == nil {
		return arg
	}
	return c.AccessListType
}

// SetAccessListType sets the access list type from a config file
func (c *Conf) SetAccessListType(label ...string) {
	if v, ok := c.GetBool("i2cp.enableBlackList", label...); ok {
		if v {
			c.AccessListType = "blacklist"
		}
	}
	if v, ok := c.GetBool("i2cp.enableAccessList", label...); ok {
		if v {
			c.AccessListType = "whitelist"
		}
	}
	if c.AccessListType != "whitelist" && c.AccessListType != "blacklist" {
		c.AccessListType = "none"
	}
}

// AddAccessListMember adds a member to either the blacklist or the whitelist
func (c *Conf) AddAccessListMember(key string) {
	for _, item := range c.AccessList {
		if item == key {
			return
		}
	}
	c.AccessList = append(c.AccessList, key)
}

func (c *Conf) accesslisttype() string {
	if c.AccessListType == "whitelist" {
		return "i2cp.enableAccessList=true"
	} else if c.AccessListType == "blacklist" {
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

//SetAccessListType tells the system to treat the accessList as a whitelist
func SetAccessListType(s string) func(samtunnel.SAMTunnel) error {
	return func(c samtunnel.SAMTunnel) error {
		if s == "whitelist" {
			c.(*Conf).AccessListType = "whitelist"
			return nil
		} else if s == "blacklist" {
			c.(*Conf).AccessListType = "blacklist"
			return nil
		} else if s == "none" {
			c.(*Conf).AccessListType = ""
			return nil
		} else if s == "" {
			c.(*Conf).AccessListType = ""
			return nil
		}
		return fmt.Errorf("Invalid Access list type(whitelist, blacklist, none)")
	}
}

//SetAccessList tells the system to treat the accessList as a whitelist
func SetAccessList(s []string) func(samtunnel.SAMTunnel) error {
	return func(c samtunnel.SAMTunnel) error {
		if len(s) > 0 {
			for _, a := range s {
				c.(*Conf).AccessList = append(c.(*Conf).AccessList, a)
			}
			return nil
		}
		return nil
	}
}
