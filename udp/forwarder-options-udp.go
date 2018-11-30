package samforwarderudp

import (
	"fmt"
	"strconv"
)

//Option is a SAMSSUForwarder Option
type Option func(*SAMSSUForwarder) error

//SetFilePath sets the host of the SAMSSUForwarder's SAM bridge
func SetFilePath(s string) func(*SAMSSUForwarder) error {
	return func(c *SAMSSUForwarder) error {
		c.FilePath = s
		return nil
	}
}

//SetSaveFile tells the router to use an encrypted leaseset
func SetSaveFile(b bool) func(*SAMSSUForwarder) error {
	return func(c *SAMSSUForwarder) error {
		c.save = b
		return nil
	}
}

//SetHost sets the host of the SAMSSUForwarder's SAM bridge
func SetHost(s string) func(*SAMSSUForwarder) error {
	return func(c *SAMSSUForwarder) error {
		c.TargetHost = s
		return nil
	}
}

//SetPort sets the port of the SAMSSUForwarder's SAM bridge using a string
func SetPort(s string) func(*SAMSSUForwarder) error {
	return func(c *SAMSSUForwarder) error {
		port, err := strconv.Atoi(s)
		if err != nil {
			return fmt.Errorf("Invalid SSU Server Target Port %s; non-number ", s)
		}
		if port < 65536 && port > -1 {
			c.TargetPort = s
			return nil
		}
		return fmt.Errorf("Invalid port")
	}
}

//SetSAMHost sets the host of the SAMSSUForwarder's SAM bridge
func SetSAMHost(s string) func(*SAMSSUForwarder) error {
	return func(c *SAMSSUForwarder) error {
		c.SamHost = s
		return nil
	}
}

//SetSAMPort sets the port of the SAMSSUForwarder's SAM bridge using a string
func SetSAMPort(s string) func(*SAMSSUForwarder) error {
	return func(c *SAMSSUForwarder) error {
		port, err := strconv.Atoi(s)
		if err != nil {
			return fmt.Errorf("Invalid SAM Port %s; non-number", s)
		}
		if port < 65536 && port > -1 {
			c.SamPort = s
			return nil
		}
		return fmt.Errorf("Invalid port")
	}
}

//SetName sets the host of the SAMSSUForwarder's SAM bridge
func SetName(s string) func(*SAMSSUForwarder) error {
	return func(c *SAMSSUForwarder) error {
		c.TunName = s
		return nil
	}
}

//SetInLength sets the number of hops inbound
func SetInLength(u int) func(*SAMSSUForwarder) error {
	return func(c *SAMSSUForwarder) error {
		if u < 7 && u >= 0 {
			c.inLength = strconv.Itoa(u)
			return nil
		}
		return fmt.Errorf("Invalid inbound tunnel length")
	}
}

//SetOutLength sets the number of hops outbound
func SetOutLength(u int) func(*SAMSSUForwarder) error {
	return func(c *SAMSSUForwarder) error {
		if u < 7 && u >= 0 {
			c.outLength = strconv.Itoa(u)
			return nil
		}
		return fmt.Errorf("Invalid outbound tunnel length")
	}
}

//SetInVariance sets the variance of a number of hops inbound
func SetInVariance(i int) func(*SAMSSUForwarder) error {
	return func(c *SAMSSUForwarder) error {
		if i < 7 && i > -7 {
			c.inVariance = strconv.Itoa(i)
			return nil
		}
		return fmt.Errorf("Invalid inbound tunnel length")
	}
}

//SetOutVariance sets the variance of a number of hops outbound
func SetOutVariance(i int) func(*SAMSSUForwarder) error {
	return func(c *SAMSSUForwarder) error {
		if i < 7 && i > -7 {
			c.outVariance = strconv.Itoa(i)
			return nil
		}
		return fmt.Errorf("Invalid outbound tunnel variance")
	}
}

//SetInQuantity sets the inbound tunnel quantity
func SetInQuantity(u int) func(*SAMSSUForwarder) error {
	return func(c *SAMSSUForwarder) error {
		if u <= 16 && u > 0 {
			c.inQuantity = strconv.Itoa(u)
			return nil
		}
		return fmt.Errorf("Invalid inbound tunnel quantity")
	}
}

//SetOutQuantity sets the outbound tunnel quantity
func SetOutQuantity(u int) func(*SAMSSUForwarder) error {
	return func(c *SAMSSUForwarder) error {
		if u <= 16 && u > 0 {
			c.outQuantity = strconv.Itoa(u)
			return nil
		}
		return fmt.Errorf("Invalid outbound tunnel quantity")
	}
}

//SetInBackups sets the inbound tunnel backups
func SetInBackups(u int) func(*SAMSSUForwarder) error {
	return func(c *SAMSSUForwarder) error {
		if u < 6 && u >= 0 {
			c.inBackupQuantity = strconv.Itoa(u)
			return nil
		}
		return fmt.Errorf("Invalid inbound tunnel backup quantity")
	}
}

//SetOutBackups sets the inbound tunnel backups
func SetOutBackups(u int) func(*SAMSSUForwarder) error {
	return func(c *SAMSSUForwarder) error {
		if u < 6 && u >= 0 {
			c.outBackupQuantity = strconv.Itoa(u)
			return nil
		}
		return fmt.Errorf("Invalid outbound tunnel backup quantity")
	}
}

//SetEncrypt tells the router to use an encrypted leaseset
func SetEncrypt(b bool) func(*SAMSSUForwarder) error {
	return func(c *SAMSSUForwarder) error {
		if b {
			c.encryptLeaseSet = "true"
			return nil
		}
		c.encryptLeaseSet = "false"
		return nil
	}
}

//SetLeaseSetKey sets
func SetLeaseSetKey(s string) func(*SAMSSUForwarder) error {
	return func(c *SAMSSUForwarder) error {
		c.leaseSetKey = s
		return nil
	}
}

//SetLeaseSetPrivateKey sets
func SetLeaseSetPrivateKey(s string) func(*SAMSSUForwarder) error {
	return func(c *SAMSSUForwarder) error {
		c.leaseSetPrivateKey = s
		return nil
	}
}

//SetLeaseSetPrivateSigningKey sets
func SetLeaseSetPrivateSigningKey(s string) func(*SAMSSUForwarder) error {
	return func(c *SAMSSUForwarder) error {
		c.leaseSetPrivateSigningKey = s
		return nil
	}
}

//SetMessageReliability sets
func SetMessageReliability(s string) func(*SAMSSUForwarder) error {
	return func(c *SAMSSUForwarder) error {
		c.messageReliability = s
		return nil
	}
}

//SetAllowZeroIn tells the tunnel to accept zero-hop peers
func SetAllowZeroIn(b bool) func(*SAMSSUForwarder) error {
	return func(c *SAMSSUForwarder) error {
		if b {
			c.inAllowZeroHop = "true"
			return nil
		}
		c.inAllowZeroHop = "false"
		return nil
	}
}

//SetAllowZeroOut tells the tunnel to accept zero-hop peers
func SetAllowZeroOut(b bool) func(*SAMSSUForwarder) error {
	return func(c *SAMSSUForwarder) error {
		if b {
			c.outAllowZeroHop = "true"
			return nil
		}
		c.outAllowZeroHop = "false"
		return nil
	}
}

//SetFastRecieve tells clients to use compression
func SetFastRecieve(b bool) func(*SAMSSUForwarder) error {
	return func(c *SAMSSUForwarder) error {
		if b {
			c.fastRecieve = "true"
			return nil
		}
		c.fastRecieve = "false"
		return nil
	}
}

//SetCompress tells clients to use compression
func SetCompress(b bool) func(*SAMSSUForwarder) error {
	return func(c *SAMSSUForwarder) error {
		if b {
			c.useCompression = "true"
			return nil
		}
		c.useCompression = "false"
		return nil
	}
}

//SetReduceIdle tells the connection to reduce it's tunnels during extended idle time.
func SetReduceIdle(b bool) func(*SAMSSUForwarder) error {
	return func(c *SAMSSUForwarder) error {
		if b {
			c.reduceIdle = "true"
			return nil
		}
		c.reduceIdle = "false"
		return nil
	}
}

//SetReduceIdleTime sets the time to wait before reducing tunnels to idle levels
func SetReduceIdleTime(u int) func(*SAMSSUForwarder) error {
	return func(c *SAMSSUForwarder) error {
		c.reduceIdleTime = "300000"
		if u >= 6 {
			c.reduceIdleTime = strconv.Itoa((u * 60) * 1000)
			return nil
		}
		return fmt.Errorf("Invalid close idle timeout(Measured in minutes) %v", u)
	}
}

//SetReduceIdleTimeMs sets the time to wait before reducing tunnels to idle levels in milliseconds
func SetReduceIdleTimeMs(u int) func(*SAMSSUForwarder) error {
	return func(c *SAMSSUForwarder) error {
		c.reduceIdleTime = "300000"
		if u >= 300000 {
			c.reduceIdleTime = strconv.Itoa(u)
			return nil
		}
		return fmt.Errorf("Invalid close idle timeout(Measured in milliseconds) %v", u)
	}
}

//SetReduceIdleQuantity sets minimum number of tunnels to reduce to during idle time
func SetReduceIdleQuantity(u int) func(*SAMSSUForwarder) error {
	return func(c *SAMSSUForwarder) error {
		if u < 5 {
			c.reduceIdleQuantity = strconv.Itoa(u)
			return nil
		}
		return fmt.Errorf("Invalid reduce tunnel quantity")
	}
}

//SetCloseIdle tells the connection to close it's tunnels during extended idle time.
func SetCloseIdle(b bool) func(*SAMSSUForwarder) error {
	return func(c *SAMSSUForwarder) error {
		if b {
			c.closeIdle = "true"
			return nil
		}
		c.closeIdle = "false"
		return nil
	}
}

//SetCloseIdleTime sets the time to wait before closing tunnels to idle levels
func SetCloseIdleTime(u int) func(*SAMSSUForwarder) error {
	return func(c *SAMSSUForwarder) error {
		c.closeIdleTime = "300000"
		if u >= 6 {
			c.closeIdleTime = strconv.Itoa((u * 60) * 2000)
			return nil
		}
		return fmt.Errorf("Invalid reduce idle timeout(Measured in minutes)")
	}
}

//SetCloseIdleTimeMs sets the time to wait before closing tunnels to idle levels in milliseconds
func SetCloseIdleTimeMs(u int) func(*SAMSSUForwarder) error {
	return func(c *SAMSSUForwarder) error {
		c.closeIdleTime = "300000"
		if u >= 300000 {
			c.closeIdleTime = strconv.Itoa(u)
			return nil
		}
		return fmt.Errorf("Invalid reduce idle timeout(Measured in minutes)")
	}
}

//SetAccessListType tells the system to treat the accessList as a whitelist
func SetAccessListType(s string) func(*SAMSSUForwarder) error {
	return func(c *SAMSSUForwarder) error {
		if s == "whitelist" {
			c.accessListType = "whitelist"
			return nil
		} else if s == "blacklist" {
			c.accessListType = "blacklist"
			return nil
		} else if s == "none" {
			c.accessListType = ""
			return nil
		} else if s == "" {
			c.accessListType = ""
			return nil
		}
		return fmt.Errorf("Invalid Access list type(whitelist, blacklist, none)")
	}
}

//SetAccessList tells the system to treat the accessList as a whitelist
func SetAccessList(s []string) func(*SAMSSUForwarder) error {
	return func(c *SAMSSUForwarder) error {
		if len(s) > 0 {
			for _, a := range s {
				c.accessList = append(c.accessList, a)
			}
			return nil
		}
		return nil
	}
}

//SetKeyFile sets
func SetKeyFile(s string) func(*SAMSSUForwarder) error {
	return func(c *SAMSSUForwarder) error {
		c.passfile = s
		return nil
	}
}
