package command

import (
	"fmt"
	"strconv"
	"strings"
)

//Option is a SAMEmit Option
type Option func(*SAMEmit) error

//SetType sets the type of the forwarder server
func SetType(s string) func(*SAMEmit) error {
	return func(c *SAMEmit) error {
		if s == "STREAM" {
			c.Style = s
			return nil
		} else if s == "DATAGRAM" {
			c.Style = s
			return nil
		} else if s == "RAW" {
			c.Style = s
			return nil
		}
		return fmt.Errorf("Invalid session STYLE=%s, must be STREAM, DATAGRAM, or RAW")
	}
}

// SetSAMAddress sets the SAM address all-at-once
func SetSAMAddress(s string) func(*SAMEmit) error {
	return func(c *SAMEmit) error {
		sp := strings.Split(s, ":")
		if len(sp) > 2 {
			return fmt.Errorf("Invalid address string: %s", sp)
		}
		if len(sp) == 2 {
			c.I2PConfig.SamPort = sp[1]
		}
		c.I2PConfig.SamHost = sp[0]
		return nil
	}
}

//SetSAMHost sets the host of the SAMEmit's SAM bridge
func SetSAMHost(s string) func(*SAMEmit) error {
	return func(c *SAMEmit) error {
		c.I2PConfig.SamHost = s
		return nil
	}
}

//SetSAMPort sets the port of the SAMEmit's SAM bridge using a string
func SetSAMPort(s string) func(*SAMEmit) error {
	return func(c *SAMEmit) error {
		port, err := strconv.Atoi(s)
		if err != nil {
			return fmt.Errorf("Invalid SAM Port %s; non-number", s)
		}
		if port < 65536 && port > -1 {
			c.I2PConfig.SamPort = s
			return nil
		}
		return fmt.Errorf("Invalid port")
	}
}

//SetName sets the host of the SAMEmit's SAM bridge
func SetName(s string) func(*SAMEmit) error {
	return func(c *SAMEmit) error {
		c.I2PConfig.TunName = s
		return nil
	}
}

//SetInLength sets the number of hops inbound
func SetInLength(u int) func(*SAMEmit) error {
	return func(c *SAMEmit) error {
		if u < 7 && u >= 0 {
			c.I2PConfig.InLength = strconv.Itoa(u)
			return nil
		}
		return fmt.Errorf("Invalid inbound tunnel length")
	}
}

//SetOutLength sets the number of hops outbound
func SetOutLength(u int) func(*SAMEmit) error {
	return func(c *SAMEmit) error {
		if u < 7 && u >= 0 {
			c.I2PConfig.OutLength = strconv.Itoa(u)
			return nil
		}
		return fmt.Errorf("Invalid outbound tunnel length")
	}
}

//SetInVariance sets the variance of a number of hops inbound
func SetInVariance(i int) func(*SAMEmit) error {
	return func(c *SAMEmit) error {
		if i < 7 && i > -7 {
			c.I2PConfig.InVariance = strconv.Itoa(i)
			return nil
		}
		return fmt.Errorf("Invalid inbound tunnel length")
	}
}

//SetOutVariance sets the variance of a number of hops outbound
func SetOutVariance(i int) func(*SAMEmit) error {
	return func(c *SAMEmit) error {
		if i < 7 && i > -7 {
			c.I2PConfig.OutVariance = strconv.Itoa(i)
			return nil
		}
		return fmt.Errorf("Invalid outbound tunnel variance")
	}
}

//SetInQuantity sets the inbound tunnel quantity
func SetInQuantity(u int) func(*SAMEmit) error {
	return func(c *SAMEmit) error {
		if u <= 16 && u > 0 {
			c.I2PConfig.InQuantity = strconv.Itoa(u)
			return nil
		}
		return fmt.Errorf("Invalid inbound tunnel quantity")
	}
}

//SetOutQuantity sets the outbound tunnel quantity
func SetOutQuantity(u int) func(*SAMEmit) error {
	return func(c *SAMEmit) error {
		if u <= 16 && u > 0 {
			c.I2PConfig.OutQuantity = strconv.Itoa(u)
			return nil
		}
		return fmt.Errorf("Invalid outbound tunnel quantity")
	}
}

//SetInBackups sets the inbound tunnel backups
func SetInBackups(u int) func(*SAMEmit) error {
	return func(c *SAMEmit) error {
		if u < 6 && u >= 0 {
			c.I2PConfig.InBackupQuantity = strconv.Itoa(u)
			return nil
		}
		return fmt.Errorf("Invalid inbound tunnel backup quantity")
	}
}

//SetOutBackups sets the inbound tunnel backups
func SetOutBackups(u int) func(*SAMEmit) error {
	return func(c *SAMEmit) error {
		if u < 6 && u >= 0 {
			c.I2PConfig.OutBackupQuantity = strconv.Itoa(u)
			return nil
		}
		return fmt.Errorf("Invalid outbound tunnel backup quantity")
	}
}

//SetEncrypt tells the router to use an encrypted leaseset
func SetEncrypt(b bool) func(*SAMEmit) error {
	return func(c *SAMEmit) error {
		if b {
			c.I2PConfig.EncryptLeaseSet = "true"
			return nil
		}
		c.I2PConfig.EncryptLeaseSet = "false"
		return nil
	}
}

//SetLeaseSetKey sets the host of the SAMEmit's SAM bridge
func SetLeaseSetKey(s string) func(*SAMEmit) error {
	return func(c *SAMEmit) error {
		c.I2PConfig.LeaseSetKey = s
		return nil
	}
}

//SetLeaseSetPrivateKey sets the host of the SAMEmit's SAM bridge
func SetLeaseSetPrivateKey(s string) func(*SAMEmit) error {
	return func(c *SAMEmit) error {
		c.I2PConfig.LeaseSetPrivateKey = s
		return nil
	}
}

//SetLeaseSetPrivateSigningKey sets the host of the SAMEmit's SAM bridge
func SetLeaseSetPrivateSigningKey(s string) func(*SAMEmit) error {
	return func(c *SAMEmit) error {
		c.I2PConfig.LeaseSetPrivateSigningKey = s
		return nil
	}
}

//SetMessageReliability sets the host of the SAMEmit's SAM bridge
func SetMessageReliability(s string) func(*SAMEmit) error {
	return func(c *SAMEmit) error {
		c.I2PConfig.MessageReliability = s
		return nil
	}
}

//SetAllowZeroIn tells the tunnel to accept zero-hop peers
func SetAllowZeroIn(b bool) func(*SAMEmit) error {
	return func(c *SAMEmit) error {
		if b {
			c.I2PConfig.InAllowZeroHop = "true"
			return nil
		}
		c.I2PConfig.InAllowZeroHop = "false"
		return nil
	}
}

//SetAllowZeroOut tells the tunnel to accept zero-hop peers
func SetAllowZeroOut(b bool) func(*SAMEmit) error {
	return func(c *SAMEmit) error {
		if b {
			c.I2PConfig.OutAllowZeroHop = "true"
			return nil
		}
		c.I2PConfig.OutAllowZeroHop = "false"
		return nil
	}
}

//SetCompress tells clients to use compression
func SetCompress(b bool) func(*SAMEmit) error {
	return func(c *SAMEmit) error {
		if b {
			c.I2PConfig.UseCompression = "true"
			return nil
		}
		c.I2PConfig.UseCompression = "false"
		return nil
	}
}

//SetFastRecieve tells clients to use compression
func SetFastRecieve(b bool) func(*SAMEmit) error {
	return func(c *SAMEmit) error {
		if b {
			c.I2PConfig.FastRecieve = "true"
			return nil
		}
		c.I2PConfig.FastRecieve = "false"
		return nil
	}
}

//SetReduceIdle tells the connection to reduce it's tunnels during extended idle time.
func SetReduceIdle(b bool) func(*SAMEmit) error {
	return func(c *SAMEmit) error {
		if b {
			c.I2PConfig.ReduceIdle = "true"
			return nil
		}
		c.I2PConfig.ReduceIdle = "false"
		return nil
	}
}

//SetReduceIdleTime sets the time to wait before reducing tunnels to idle levels
func SetReduceIdleTime(u int) func(*SAMEmit) error {
	return func(c *SAMEmit) error {
		c.I2PConfig.ReduceIdleTime = "300000"
		if u >= 6 {
			c.I2PConfig.ReduceIdleTime = strconv.Itoa((u * 60) * 1000)
			return nil
		}
		return fmt.Errorf("Invalid reduce idle timeout(Measured in minutes) %v", u)
	}
}

//SetReduceIdleTimeMs sets the time to wait before reducing tunnels to idle levels in milliseconds
func SetReduceIdleTimeMs(u int) func(*SAMEmit) error {
	return func(c *SAMEmit) error {
		c.I2PConfig.ReduceIdleTime = "300000"
		if u >= 300000 {
			c.I2PConfig.ReduceIdleTime = strconv.Itoa(u)
			return nil
		}
		return fmt.Errorf("Invalid reduce idle timeout(Measured in milliseconds) %v", u)
	}
}

//SetReduceIdleQuantity sets minimum number of tunnels to reduce to during idle time
func SetReduceIdleQuantity(u int) func(*SAMEmit) error {
	return func(c *SAMEmit) error {
		if u < 5 {
			c.I2PConfig.ReduceIdleQuantity = strconv.Itoa(u)
			return nil
		}
		return fmt.Errorf("Invalid reduce tunnel quantity")
	}
}

//SetCloseIdle tells the connection to close it's tunnels during extended idle time.
func SetCloseIdle(b bool) func(*SAMEmit) error {
	return func(c *SAMEmit) error {
		if b {
			c.I2PConfig.CloseIdle = "true"
			return nil
		}
		c.I2PConfig.CloseIdle = "false"
		return nil
	}
}

//SetCloseIdleTime sets the time to wait before closing tunnels to idle levels
func SetCloseIdleTime(u int) func(*SAMEmit) error {
	return func(c *SAMEmit) error {
		c.I2PConfig.CloseIdleTime = "300000"
		if u >= 6 {
			c.I2PConfig.CloseIdleTime = strconv.Itoa((u * 60) * 1000)
			return nil
		}
		return fmt.Errorf("Invalid close idle timeout(Measured in minutes) %v", u)
	}
}

//SetCloseIdleTimeMs sets the time to wait before closing tunnels to idle levels in milliseconds
func SetCloseIdleTimeMs(u int) func(*SAMEmit) error {
	return func(c *SAMEmit) error {
		c.I2PConfig.CloseIdleTime = "300000"
		if u >= 300000 {
			c.I2PConfig.CloseIdleTime = strconv.Itoa(u)
			return nil
		}
		return fmt.Errorf("Invalid close idle timeout(Measured in milliseconds) %v", u)
	}
}

//SetAccessListType tells the system to treat the AccessList as a whitelist
func SetAccessListType(s string) func(*SAMEmit) error {
	return func(c *SAMEmit) error {
		if s == "whitelist" {
			c.I2PConfig.AccessListType = "whitelist"
			return nil
		} else if s == "blacklist" {
			c.I2PConfig.AccessListType = "blacklist"
			return nil
		} else if s == "none" {
			c.I2PConfig.AccessListType = ""
			return nil
		} else if s == "" {
			c.I2PConfig.AccessListType = ""
			return nil
		}
		return fmt.Errorf("Invalid Access list type(whitelist, blacklist, none)")
	}
}

//SetAccessList tells the system to treat the AccessList as a whitelist
func SetAccessList(s []string) func(*SAMEmit) error {
	return func(c *SAMEmit) error {
		if len(s) > 0 {
			for _, a := range s {
				c.I2PConfig.AccessList = append(c.I2PConfig.AccessList, a)
			}
			return nil
		}
		return nil
	}
}
