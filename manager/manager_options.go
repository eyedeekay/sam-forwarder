package sammanager

import (
	"fmt"
	"strconv"
)

//ManagerOption is a SAMManager Option
type ManagerOption func(*SAMManager) error

//SetManagerFilePath sets the host of the SAMManager's SAM bridge
func SetManagerFilePath(s string) func(*SAMManager) error {
	return func(c *SAMManager) error {
		c.FilePath = s
		return nil
	}
}

//SetManagerSaveFile tells the router to use an encrypted leaseset
func SetManagerSaveFile(b bool) func(*SAMManager) error {
	return func(c *SAMManager) error {
		c.save = b
		return nil
	}
}

//SetManagerHost sets the host of the SAMManager's SAM bridge
func SetManagerHost(s string) func(*SAMManager) error {
	return func(c *SAMManager) error {
		c.TargetHost = s
		return nil
	}
}

//SetManagerDestination sets the destination to forwarder SAMManager's to
func SetManagerDestination(s string) func(*SAMManager) error {
	return func(c *SAMManager) error {
		c.dest = s
		return nil
	}
}

//SetManagerPort sets the port of the SAMManager's SAM bridge using a string
func SetManagerPort(s string) func(*SAMManager) error {
	return func(c *SAMManager) error {
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

//SetManagerSAMHost sets the host of the SAMManager's SAM bridge
func SetManagerSAMHost(s string) func(*SAMManager) error {
	return func(c *SAMManager) error {
		c.SamHost = s
		return nil
	}
}

//SetManagerSAMPort sets the port of the SAMManager's SAM bridge using a string
func SetManagerSAMPort(s string) func(*SAMManager) error {
	return func(c *SAMManager) error {
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

//SetManagerName sets the host of the SAMManager's SAM bridge
func SetManagerName(s string) func(*SAMManager) error {
	return func(c *SAMManager) error {
		c.TunName = s
		return nil
	}
}

//SetManagerInLength sets the number of hops inbound
func SetManagerInLength(u int) func(*SAMManager) error {
	return func(c *SAMManager) error {
		if u < 7 && u >= 0 {
			c.inLength = strconv.Itoa(u)
			return nil
		}
		return fmt.Errorf("Invalid inbound tunnel length")
	}
}

//SetManagerOutLength sets the number of hops outbound
func SetManagerOutLength(u int) func(*SAMManager) error {
	return func(c *SAMManager) error {
		if u < 7 && u >= 0 {
			c.outLength = strconv.Itoa(u)
			return nil
		}
		return fmt.Errorf("Invalid outbound tunnel length")
	}
}

//SetManagerInVariance sets the variance of a number of hops inbound
func SetManagerInVariance(i int) func(*SAMManager) error {
	return func(c *SAMManager) error {
		if i < 7 && i > -7 {
			c.inVariance = strconv.Itoa(i)
			return nil
		}
		return fmt.Errorf("Invalid inbound tunnel length")
	}
}

//SetManagerOutVariance sets the variance of a number of hops outbound
func SetManagerOutVariance(i int) func(*SAMManager) error {
	return func(c *SAMManager) error {
		if i < 7 && i > -7 {
			c.outVariance = strconv.Itoa(i)
			return nil
		}
		return fmt.Errorf("Invalid outbound tunnel variance")
	}
}

//SetManagerInQuantity sets the inbound tunnel quantity
func SetManagerInQuantity(u int) func(*SAMManager) error {
	return func(c *SAMManager) error {
		if u <= 16 && u > 0 {
			c.inQuantity = strconv.Itoa(u)
			return nil
		}
		return fmt.Errorf("Invalid inbound tunnel quantity")
	}
}

//SetManagerOutQuantity sets the outbound tunnel quantity
func SetManagerOutQuantity(u int) func(*SAMManager) error {
	return func(c *SAMManager) error {
		if u <= 16 && u > 0 {
			c.outQuantity = strconv.Itoa(u)
			return nil
		}
		return fmt.Errorf("Invalid outbound tunnel quantity")
	}
}

//SetManagerInBackups sets the inbound tunnel backups
func SetManagerInBackups(u int) func(*SAMManager) error {
	return func(c *SAMManager) error {
		if u < 6 && u >= 0 {
			c.inBackupQuantity = strconv.Itoa(u)
			return nil
		}
		return fmt.Errorf("Invalid inbound tunnel backup quantity")
	}
}

//SetManagerOutBackups sets the inbound tunnel backups
func SetManagerOutBackups(u int) func(*SAMManager) error {
	return func(c *SAMManager) error {
		if u < 6 && u >= 0 {
			c.outBackupQuantity = strconv.Itoa(u)
			return nil
		}
		return fmt.Errorf("Invalid outbound tunnel backup quantity")
	}
}

//SetManagerEncrypt tells the router to use an encrypted leaseset
func SetManagerEncrypt(b bool) func(*SAMManager) error {
	return func(c *SAMManager) error {
		if b {
			c.encryptLeaseSet = "true"
			return nil
		}
		c.encryptLeaseSet = "false"
		return nil
	}
}

//SetManagerAllowZeroIn tells the tunnel to accept zero-hop peers
func SetManagerAllowZeroIn(b bool) func(*SAMManager) error {
	return func(c *SAMManager) error {
		if b {
			c.inAllowZeroHop = "true"
			return nil
		}
		c.inAllowZeroHop = "false"
		return nil
	}
}

//SetManagerAllowZeroOut tells the tunnel to accept zero-hop peers
func SetManagerAllowZeroOut(b bool) func(*SAMManager) error {
	return func(c *SAMManager) error {
		if b {
			c.outAllowZeroHop = "true"
			return nil
		}
		c.outAllowZeroHop = "false"
		return nil
	}
}

//SetManagerCompress tells clients to use compression
func SetManagerCompress(b bool) func(*SAMManager) error {
	return func(c *SAMManager) error {
		if b {
			c.useCompression = "true"
			return nil
		}
		c.useCompression = "false"
		return nil
	}
}

//SetManagerReduceIdle tells the connection to reduce it's tunnels during extended idle time.
func SetManagerReduceIdle(b bool) func(*SAMManager) error {
	return func(c *SAMManager) error {
		if b {
			c.reduceIdle = "true"
			return nil
		}
		c.reduceIdle = "false"
		return nil
	}
}

//SetManagerReduceIdleTime sets the time to wait before reducing tunnels to idle levels
func SetManagerReduceIdleTime(u int) func(*SAMManager) error {
	return func(c *SAMManager) error {
		c.reduceIdleTime = strconv.Itoa(300000)
		if u >= 6 {
			c.reduceIdleTime = strconv.Itoa((u * 60) * 1000)
			return nil
		}
		return fmt.Errorf("Invalid reduce idle timeout(Measured in minutes) %v", u)
	}
}

//SetManagerReduceIdleTimeMs sets the time to wait before reducing tunnels to idle levels in milliseconds
func SetManagerReduceIdleTimeMs(u int) func(*SAMManager) error {
	return func(c *SAMManager) error {
		c.reduceIdleTime = strconv.Itoa(300000)
		if u >= 300000 {
			c.reduceIdleTime = strconv.Itoa(u)
			return nil
		}
		return fmt.Errorf("Invalid reduce idle timeout(Measured in milliseconds) %v", u)
	}
}

//SetManagerReduceIdleQuantity sets minimum number of tunnels to reduce to during idle time
func SetManagerReduceIdleQuantity(u int) func(*SAMManager) error {
	return func(c *SAMManager) error {
		if u < 5 {
			c.reduceIdleQuantity = strconv.Itoa(u)
			return nil
		}
		return fmt.Errorf("Invalid reduce tunnel quantity")
	}
}

//SetManagerCloseIdle tells the connection to close it's tunnels during extended idle time.
func SetManagerCloseIdle(b bool) func(*SAMManager) error {
	return func(c *SAMManager) error {
		if b {
			c.closeIdle = "true"
			return nil
		}
		c.closeIdle = "false"
		return nil
	}
}

//SetManagerCloseIdleTime sets the time to wait before closing tunnels to idle levels
func SetManagerCloseIdleTime(u int) func(*SAMManager) error {
	return func(c *SAMManager) error {
		c.closeIdleTime = strconv.Itoa(300000)
		if u >= 6 {
			c.closeIdleTime = strconv.Itoa((u * 60) * 1000)
			return nil
		}
		return fmt.Errorf("Invalid close idle timeout(Measured in minutes) %v", u)
	}
}

//SetManagerCloseIdleTimeMs sets the time to wait before closing tunnels to idle levels in milliseconds
func SetManagerCloseIdleTimeMs(u int) func(*SAMManager) error {
	return func(c *SAMManager) error {
		c.closeIdleTime = strconv.Itoa(300000)
		if u >= 300000 {
			c.closeIdleTime = strconv.Itoa(u)
			return nil
		}
		return fmt.Errorf("Invalid close idle timeout(Measured in milliseconds) %v", u)
	}
}
