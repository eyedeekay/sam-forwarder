package samforwarder

import (
	"fmt"
	"strconv"
)

//ClientOption is a SAMClientForwarder Option
type ClientOption func(*SAMClientForwarder) error

//SetClientFilePath sets the host of the SAMClientForwarder's SAM bridge
func SetClientFilePath(s string) func(*SAMClientForwarder) error {
	return func(c *SAMClientForwarder) error {
		c.FilePath = s
		return nil
	}
}

//SetClientSaveFile tells the router to save the tunnel keys long-term
func SetClientSaveFile(b bool) func(*SAMClientForwarder) error {
	return func(c *SAMClientForwarder) error {
		c.save = b
		return nil
	}
}

//SetClientHost sets the host of the SAMClientForwarder's SAM bridge
func SetClientHost(s string) func(*SAMClientForwarder) error {
	return func(c *SAMClientForwarder) error {
		c.TargetHost = s
		return nil
	}
}

//SetClientDestination sets the destination to forwarder SAMClientForwarder's to
func SetClientDestination(s string) func(*SAMClientForwarder) error {
	return func(c *SAMClientForwarder) error {
		c.dest = s
		return nil
	}
}

//SetClientPort sets the port of the SAMClientForwarder's SAM bridge using a string
func SetClientPort(s string) func(*SAMClientForwarder) error {
	return func(c *SAMClientForwarder) error {
		port, err := strconv.Atoi(s)
		if err != nil {
			return fmt.Errorf("Invalid TCP Client Target Port %s; non-number ", s)
		}
		if port < 65536 && port > -1 {
			c.TargetPort = s
			return nil
		}
		return fmt.Errorf("Invalid port")
	}
}

//SetClientSAMHost sets the host of the SAMClientForwarder's SAM bridge
func SetClientSAMHost(s string) func(*SAMClientForwarder) error {
	return func(c *SAMClientForwarder) error {
		c.SamHost = s
		return nil
	}
}

//SetClientSAMPort sets the port of the SAMClientForwarder's SAM bridge using a string
func SetClientSAMPort(s string) func(*SAMClientForwarder) error {
	return func(c *SAMClientForwarder) error {
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

//SetClientName sets the host of the SAMClientForwarder's SAM bridge
func SetClientName(s string) func(*SAMClientForwarder) error {
	return func(c *SAMClientForwarder) error {
		c.TunName = s
		return nil
	}
}

//SetClientInLength sets the number of hops inbound
func SetClientInLength(u int) func(*SAMClientForwarder) error {
	return func(c *SAMClientForwarder) error {
		if u < 7 && u >= 0 {
			c.inLength = strconv.Itoa(u)
			return nil
		}
		return fmt.Errorf("Invalid inbound tunnel length")
	}
}

//SetClientOutLength sets the number of hops outbound
func SetClientOutLength(u int) func(*SAMClientForwarder) error {
	return func(c *SAMClientForwarder) error {
		if u < 7 && u >= 0 {
			c.outLength = strconv.Itoa(u)
			return nil
		}
		return fmt.Errorf("Invalid outbound tunnel length")
	}
}

//SetClientInVariance sets the variance of a number of hops inbound
func SetClientInVariance(i int) func(*SAMClientForwarder) error {
	return func(c *SAMClientForwarder) error {
		if i < 7 && i > -7 {
			c.inVariance = strconv.Itoa(i)
			return nil
		}
		return fmt.Errorf("Invalid inbound tunnel length")
	}
}

//SetClientOutVariance sets the variance of a number of hops outbound
func SetClientOutVariance(i int) func(*SAMClientForwarder) error {
	return func(c *SAMClientForwarder) error {
		if i < 7 && i > -7 {
			c.outVariance = strconv.Itoa(i)
			return nil
		}
		return fmt.Errorf("Invalid outbound tunnel variance")
	}
}

//SetClientInQuantity sets the inbound tunnel quantity
func SetClientInQuantity(u int) func(*SAMClientForwarder) error {
	return func(c *SAMClientForwarder) error {
		if u <= 16 && u > 0 {
			c.inQuantity = strconv.Itoa(u)
			return nil
		}
		return fmt.Errorf("Invalid inbound tunnel quantity")
	}
}

//SetClientOutQuantity sets the outbound tunnel quantity
func SetClientOutQuantity(u int) func(*SAMClientForwarder) error {
	return func(c *SAMClientForwarder) error {
		if u <= 16 && u > 0 {
			c.outQuantity = strconv.Itoa(u)
			return nil
		}
		return fmt.Errorf("Invalid outbound tunnel quantity")
	}
}

//SetClientInBackups sets the inbound tunnel backups
func SetClientInBackups(u int) func(*SAMClientForwarder) error {
	return func(c *SAMClientForwarder) error {
		if u < 6 && u >= 0 {
			c.inBackupQuantity = strconv.Itoa(u)
			return nil
		}
		return fmt.Errorf("Invalid inbound tunnel backup quantity")
	}
}

//SetClientOutBackups sets the inbound tunnel backups
func SetClientOutBackups(u int) func(*SAMClientForwarder) error {
	return func(c *SAMClientForwarder) error {
		if u < 6 && u >= 0 {
			c.outBackupQuantity = strconv.Itoa(u)
			return nil
		}
		return fmt.Errorf("Invalid outbound tunnel backup quantity")
	}
}

//SetClientEncrypt tells the router to use an encrypted leaseset
func SetClientEncrypt(b bool) func(*SAMClientForwarder) error {
	return func(c *SAMClientForwarder) error {
		if b {
			c.encryptLeaseSet = "true"
			return nil
		}
		c.encryptLeaseSet = "false"
		return nil
	}
}

//SetClientLeaseSetKey sets
func SetClientLeaseSetKey(s string) func(*SAMClientForwarder) error {
	return func(c *SAMClientForwarder) error {
		c.leaseSetKey = s
		return nil
	}
}

//SetClientLeaseSetPrivateKey sets
func SetClientLeaseSetPrivateKey(s string) func(*SAMClientForwarder) error {
	return func(c *SAMClientForwarder) error {
		c.leaseSetPrivateKey = s
		return nil
	}
}

//SetClientLeaseSetPrivateSigningKey sets
func SetClientLeaseSetPrivateSigningKey(s string) func(*SAMClientForwarder) error {
	return func(c *SAMClientForwarder) error {
		c.leaseSetPrivateSigningKey = s
		return nil
	}
}

//SetClientMessageReliability sets
func SetClientMessageReliability(s string) func(*SAMClientForwarder) error {
	return func(c *SAMClientForwarder) error {
		c.messageReliability = s
		return nil
	}
}

//SetClientAllowZeroIn tells the tunnel to accept zero-hop peers
func SetClientAllowZeroIn(b bool) func(*SAMClientForwarder) error {
	return func(c *SAMClientForwarder) error {
		if b {
			c.inAllowZeroHop = "true"
			return nil
		}
		c.inAllowZeroHop = "false"
		return nil
	}
}

//SetClientAllowZeroOut tells the tunnel to accept zero-hop peers
func SetClientAllowZeroOut(b bool) func(*SAMClientForwarder) error {
	return func(c *SAMClientForwarder) error {
		if b {
			c.outAllowZeroHop = "true"
			return nil
		}
		c.outAllowZeroHop = "false"
		return nil
	}
}

//SetClientFastRecieve tells clients use the i2cp.fastRecieve option
func SetClientFastRecieve(b bool) func(*SAMClientForwarder) error {
	return func(c *SAMClientForwarder) error {
		if b {
			c.fastRecieve = "true"
			return nil
		}
		c.fastRecieve = "false"
		return nil
	}
}

//SetClientCompress tells clients to use compression
func SetClientCompress(b bool) func(*SAMClientForwarder) error {
	return func(c *SAMClientForwarder) error {
		if b {
			c.useCompression = "true"
			return nil
		}
		c.useCompression = "false"
		return nil
	}
}

//SetClientReduceIdle tells the connection to reduce it's tunnels during extended idle time.
func SetClientReduceIdle(b bool) func(*SAMClientForwarder) error {
	return func(c *SAMClientForwarder) error {
		if b {
			c.reduceIdle = "true"
			return nil
		}
		c.reduceIdle = "false"
		return nil
	}
}

//SetClientReduceIdleTime sets the time to wait before reducing tunnels to idle levels
func SetClientReduceIdleTime(u int) func(*SAMClientForwarder) error {
	return func(c *SAMClientForwarder) error {
		c.reduceIdleTime = strconv.Itoa(300000)
		if u >= 6 {
			c.reduceIdleTime = strconv.Itoa((u * 60) * 1000)
			return nil
		}
		return fmt.Errorf("Invalid reduce idle timeout(Measured in minutes) %v", u)
	}
}

//SetClientReduceIdleTimeMs sets the time to wait before reducing tunnels to idle levels in milliseconds
func SetClientReduceIdleTimeMs(u int) func(*SAMClientForwarder) error {
	return func(c *SAMClientForwarder) error {
		c.reduceIdleTime = strconv.Itoa(300000)
		if u >= 300000 {
			c.reduceIdleTime = strconv.Itoa(u)
			return nil
		}
		return fmt.Errorf("Invalid reduce idle timeout(Measured in milliseconds) %v", u)
	}
}

//SetClientReduceIdleQuantity sets minimum number of tunnels to reduce to during idle time
func SetClientReduceIdleQuantity(u int) func(*SAMClientForwarder) error {
	return func(c *SAMClientForwarder) error {
		if u < 5 {
			c.reduceIdleQuantity = strconv.Itoa(u)
			return nil
		}
		return fmt.Errorf("Invalid reduce tunnel quantity")
	}
}

//SetClientCloseIdle tells the connection to close it's tunnels during extended idle time.
func SetClientCloseIdle(b bool) func(*SAMClientForwarder) error {
	return func(c *SAMClientForwarder) error {
		if b {
			c.closeIdle = "true"
			return nil
		}
		c.closeIdle = "false"
		return nil
	}
}

//SetClientCloseIdleTime sets the time to wait before closing tunnels to idle levels
func SetClientCloseIdleTime(u int) func(*SAMClientForwarder) error {
	return func(c *SAMClientForwarder) error {
		c.closeIdleTime = "300000"
		if u >= 6 {
			c.closeIdleTime = strconv.Itoa((u * 60) * 1000)
			return nil
		}
		return fmt.Errorf("Invalid close idle timeout(Measured in minutes) %v", u)
	}
}

//SetClientCloseIdleTimeMs sets the time to wait before closing tunnels to idle levels in milliseconds
func SetClientCloseIdleTimeMs(u int) func(*SAMClientForwarder) error {
	return func(c *SAMClientForwarder) error {
		c.closeIdleTime = "300000"
		if u >= 300000 {
			c.closeIdleTime = strconv.Itoa(u)
			return nil
		}
		return fmt.Errorf("Invalid close idle timeout(Measured in milliseconds) %v", u)
	}
}

//SetClientAccessListType tells the system to treat the accessList as a whitelist
func SetClientAccessListType(s string) func(*SAMClientForwarder) error {
	return func(c *SAMClientForwarder) error {
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

//SetClientAccessList tells the system to treat the accessList as a whitelist
func SetClientAccessList(s []string) func(*SAMClientForwarder) error {
	return func(c *SAMClientForwarder) error {
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
func SetClientPassword(s string) func(*SAMClientForwarder) error {
	return func(c *SAMClientForwarder) error {
		c.passfile = s
		return nil
	}
}
