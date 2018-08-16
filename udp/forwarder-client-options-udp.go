package samforwarderudp

import (
	"fmt"
	"strconv"
)

//ClientOption is a SAMSSUClientForwarder Option
type ClientOption func(*SAMSSUClientForwarder) error

//SetClientFilePath sets the host of the SAMSSUForwarder's SAM bridge
func SetClientFilePath(s string) func(*SAMSSUForwarder) error {
	return func(c *SAMSSUForwarder) error {
		c.FilePath = s
		return nil
	}
}

//SetClientSaveFile tells the router to use an encrypted leaseset
func SetClientSaveFile(b bool) func(*SAMSSUForwarder) error {
	return func(c *SAMSSUForwarder) error {
		c.save = b
		return nil
	}
}

//SetClientHost sets the host of the SAMSSUForwarder's SAM bridge
func SetClientHost(s string) func(*SAMSSUForwarder) error {
	return func(c *SAMSSUForwarder) error {
		c.TargetHost = s
		return nil
	}
}

//SetClientPort sets the port of the SAMSSUForwarder's SAM bridge using a string
func SetClientPort(s string) func(*SAMSSUForwarder) error {
	return func(c *SAMSSUForwarder) error {
		port, err := strconv.Atoi(s)
		if err != nil {
			return fmt.Errorf("Invalid Target Port %s; non-number ", s)
		}
		if port < 65536 && port > -1 {
			c.TargetPort = s
			return nil
		}
		return fmt.Errorf("Invalid port")
	}
}

//SetClientSAMHost sets the host of the SAMSSUForwarder's SAM bridge
func SetClientSAMHost(s string) func(*SAMSSUForwarder) error {
	return func(c *SAMSSUForwarder) error {
		c.SamHost = s
		return nil
	}
}

//SetClientSAMPort sets the port of the SAMSSUForwarder's SAM bridge using a string
func SetClientSAMPort(s string) func(*SAMSSUForwarder) error {
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

//SetClientName sets the host of the SAMSSUForwarder's SAM bridge
func SetClientName(s string) func(*SAMSSUForwarder) error {
	return func(c *SAMSSUForwarder) error {
		c.TunName = s
		return nil
	}
}

//SetClientInLength sets the number of hops inbound
func SetClientInLength(u int) func(*SAMSSUForwarder) error {
	return func(c *SAMSSUForwarder) error {
		if u < 7 && u >= 0 {
			c.inLength = strconv.Itoa(u)
			return nil
		}
		return fmt.Errorf("Invalid inbound tunnel length")
	}
}

//SetClientOutLength sets the number of hops outbound
func SetClientOutLength(u int) func(*SAMSSUForwarder) error {
	return func(c *SAMSSUForwarder) error {
		if u < 7 && u >= 0 {
			c.outLength = strconv.Itoa(u)
			return nil
		}
		return fmt.Errorf("Invalid outbound tunnel length")
	}
}

//SetClientInVariance sets the variance of a number of hops inbound
func SetClientInVariance(i int) func(*SAMSSUForwarder) error {
	return func(c *SAMSSUForwarder) error {
		if i < 7 && i > -7 {
			c.inVariance = strconv.Itoa(i)
			return nil
		}
		return fmt.Errorf("Invalid inbound tunnel length")
	}
}

//SetClientOutVariance sets the variance of a number of hops outbound
func SetClientOutVariance(i int) func(*SAMSSUForwarder) error {
	return func(c *SAMSSUForwarder) error {
		if i < 7 && i > -7 {
			c.outVariance = strconv.Itoa(i)
			return nil
		}
		return fmt.Errorf("Invalid outbound tunnel variance")
	}
}

//SetClientInQuantity sets the inbound tunnel quantity
func SetClientInQuantity(u int) func(*SAMSSUForwarder) error {
	return func(c *SAMSSUForwarder) error {
		if u <= 16 && u > 0 {
			c.inQuantity = strconv.Itoa(u)
			return nil
		}
		return fmt.Errorf("Invalid inbound tunnel quantity")
	}
}

//SetClientOutQuantity sets the outbound tunnel quantity
func SetClientOutQuantity(u int) func(*SAMSSUForwarder) error {
	return func(c *SAMSSUForwarder) error {
		if u <= 16 && u > 0 {
			c.outQuantity = strconv.Itoa(u)
			return nil
		}
		return fmt.Errorf("Invalid outbound tunnel quantity")
	}
}

//SetClientInBackups sets the inbound tunnel backups
func SetClientInBackups(u int) func(*SAMSSUForwarder) error {
	return func(c *SAMSSUForwarder) error {
		if u < 6 && u >= 0 {
			c.inBackupQuantity = strconv.Itoa(u)
			return nil
		}
		return fmt.Errorf("Invalid inbound tunnel backup quantity")
	}
}

//SetClientOutBackups sets the inbound tunnel backups
func SetClientOutBackups(u int) func(*SAMSSUForwarder) error {
	return func(c *SAMSSUForwarder) error {
		if u < 6 && u >= 0 {
			c.outBackupQuantity = strconv.Itoa(u)
			return nil
		}
		return fmt.Errorf("Invalid outbound tunnel backup quantity")
	}
}

//SetClientEncrypt tells the router to use an encrypted leaseset
func SetClientEncrypt(b bool) func(*SAMSSUForwarder) error {
	return func(c *SAMSSUForwarder) error {
		if b {
			c.encryptLeaseSet = "true"
			return nil
		}
		c.encryptLeaseSet = "false"
		return nil
	}
}

//SetClientAllowZeroIn tells the tunnel to accept zero-hop peers
func SetClientAllowZeroIn(b bool) func(*SAMSSUForwarder) error {
	return func(c *SAMSSUForwarder) error {
		if b {
			c.inAllowZeroHop = "true"
			return nil
		}
		c.inAllowZeroHop = "false"
		return nil
	}
}

//SetClientAllowZeroOut tells the tunnel to accept zero-hop peers
func SetClientAllowZeroOut(b bool) func(*SAMSSUForwarder) error {
	return func(c *SAMSSUForwarder) error {
		if b {
			c.outAllowZeroHop = "true"
			return nil
		}
		c.outAllowZeroHop = "false"
		return nil
	}
}

//SetClientCompress tells clients to use compression
func SetClientCompress(b bool) func(*SAMSSUForwarder) error {
	return func(c *SAMSSUForwarder) error {
		if b {
			c.useCompression = "true"
			return nil
		}
		c.useCompression = "false"
		return nil
	}
}

//SetClientReduceIdle tells the connection to reduce it's tunnels during extended idle time.
func SetClientReduceIdle(b bool) func(*SAMSSUForwarder) error {
	return func(c *SAMSSUForwarder) error {
		if b {
			c.reduceIdle = "true"
			return nil
		}
		c.reduceIdle = "false"
		return nil
	}
}

//SetClientReduceIdleTime sets the time to wait before reducing tunnels to idle levels
func SetClientReduceIdleTime(u int) func(*SAMSSUForwarder) error {
	return func(c *SAMSSUForwarder) error {
		if u > 6 {
			c.reduceIdleTime = strconv.Itoa((u * 60) * 1000)
			return nil
		}
		return fmt.Errorf("Invalid reduce idle timeout(Measured in minutes)")
	}
}

//SetClientReduceIdleTimeMs sets the time to wait before reducing tunnels to idle levels in milliseconds
func SetClientReduceIdleTimeMs(u int) func(*SAMSSUForwarder) error {
	return func(c *SAMSSUForwarder) error {
		if u > 300000 {
			c.reduceIdleTime = strconv.Itoa(u)
			return nil
		}
		return fmt.Errorf("Invalid reduce idle timeout(Measured in minutes)")
	}
}

//SetClientReduceIdleQuantity sets minimum number of tunnels to reduce to during idle time
func SetClientReduceIdleQuantity(u int) func(*SAMSSUForwarder) error {
	return func(c *SAMSSUForwarder) error {
		if u < 5 {
			c.reduceIdleQuantity = strconv.Itoa(u)
			return nil
		}
		return fmt.Errorf("Invalid reduce tunnel quantity")
	}
}

//SetClientCloseIdle tells the connection to close it's tunnels during extended idle time.
func SetClientCloseIdle(b bool) func(*SAMSSUForwarder) error {
	return func(c *SAMSSUForwarder) error {
		if b {
			c.closeIdle = "true"
			return nil
		}
		c.closeIdle = "false"
		return nil
	}
}

//SetClientCloseIdleTime sets the time to wait before closing tunnels to idle levels
func SetClientCloseIdleTime(u int) func(*SAMSSUForwarder) error {
	return func(c *SAMSSUForwarder) error {
		if u > 6 {
			c.closeIdleTime = strconv.Itoa((u * 60) * 2000)
			return nil
		}
		return fmt.Errorf("Invalid reduce idle timeout(Measured in minutes)")
	}
}

//SetClientCloseIdleTimeMs sets the time to wait before closing tunnels to idle levels in milliseconds
func SetClientCloseIdleTimeMs(u int) func(*SAMSSUForwarder) error {
	return func(c *SAMSSUForwarder) error {
		if u > 600000 {
			c.closeIdleTime = strconv.Itoa(u)
			return nil
		}
		return fmt.Errorf("Invalid reduce idle timeout(Measured in minutes)")
	}
}

//SetClientAccessListType tells the system to treat the accessList as a whitelist
func SetClientAccessListType(s string) func(*SAMSSUForwarder) error {
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
		}
		return fmt.Errorf("Invalid Access list type(whitelist, blacklist, none)")
	}
}

//SetClientAccessList tells the system to treat the accessList as a whitelist
func SetClientAccessList(s []string) func(*SAMSSUForwarder) error {
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
