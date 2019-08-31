package i2pconfig

import (
	"fmt"
	"strconv"
	"strings"
)

//Option is a I2PConfig Option
type ConfigOption func(*I2PConfig) error

func SetTunType(s string) func(*I2PConfig) error {
	return func(c *I2PConfig) error {
		c.TunType = s
		return nil
	}
}

//SetConfigType sets the type of the forwarder server
func SetConfigStyle(s string) func(*I2PConfig) error {
	return func(c *I2PConfig) error {
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

// SetSAMAddress
func SetSAMAddress(s string) func(*I2PConfig) error {
	return func(c *I2PConfig) error {
		sp := strings.Split(s, ":")
		if len(sp) > 2 {
			return fmt.Errorf("Invalid address string: %s", sp)
		}
		if len(sp) == 2 {
			c.SamPort = sp[1]
		}
		c.SamHost = sp[0]
		return nil
	}
}

//SetConfigSAMHost sets the host of the I2PConfig's SAM bridge
func SetConfigSAMHost(s string) func(*I2PConfig) error {
	return func(c *I2PConfig) error {
		c.SamHost = s
		return nil
	}
}

//SetConfigSAMPort sets the port of the I2PConfig's SAM bridge using a string
func SetConfigSAMPort(s string) func(*I2PConfig) error {
	return func(c *I2PConfig) error {
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

//SetConfigFromPort sets the FROM_PORT propert to pass to SAM
func SetConfigFromPort(s string) func(*I2PConfig) error {
	return func(c *I2PConfig) error {
		port, err := strconv.Atoi(s)
		if err != nil {
			return fmt.Errorf("Invalid FROM Port %s; non-number", s)
		}
		if port < 65536 && port > -1 {
			c.Fromport = s
			return nil
		}
		return fmt.Errorf("Invalid port")
	}
}

//SetConfigToPort sets the TO_PORT property to pass to SAM
func SetConfigToPort(s string) func(*I2PConfig) error {
	return func(c *I2PConfig) error {
		port, err := strconv.Atoi(s)
		if err != nil {
			return fmt.Errorf("Invalid SAM Port %s; non-number", s)
		}
		if port < 65536 && port > -1 {
			c.Toport = s
			return nil
		}
		return fmt.Errorf("Invalid port")
	}
}

//SetConfigName sets the host of the I2PConfig's SAM bridge
func SetConfigName(s string) func(*I2PConfig) error {
	return func(c *I2PConfig) error {
		c.TunName = s
		return nil
	}
}

//SetConfigInLength sets the number of hops inbound
func SetConfigInLength(u int) func(*I2PConfig) error {
	return func(c *I2PConfig) error {
		if u < 7 && u >= 0 {
			c.InLength = strconv.Itoa(u)
			return nil
		}
		return fmt.Errorf("Invalid inbound tunnel length")
	}
}

//SetConfigOutLength sets the number of hops outbound
func SetConfigOutLength(u int) func(*I2PConfig) error {
	return func(c *I2PConfig) error {
		if u < 7 && u >= 0 {
			c.OutLength = strconv.Itoa(u)
			return nil
		}
		return fmt.Errorf("Invalid outbound tunnel length")
	}
}

//SetConfigInVariance sets the variance of a number of hops inbound
func SetConfigInVariance(i int) func(*I2PConfig) error {
	return func(c *I2PConfig) error {
		if i < 7 && i > -7 {
			c.InVariance = strconv.Itoa(i)
			return nil
		}
		return fmt.Errorf("Invalid inbound tunnel length")
	}
}

//SetConfigOutVariance sets the variance of a number of hops outbound
func SetConfigOutVariance(i int) func(*I2PConfig) error {
	return func(c *I2PConfig) error {
		if i < 7 && i > -7 {
			c.OutVariance = strconv.Itoa(i)
			return nil
		}
		return fmt.Errorf("Invalid outbound tunnel variance")
	}
}

//SetConfigInQuantity sets the inbound tunnel quantity
func SetConfigInQuantity(u int) func(*I2PConfig) error {
	return func(c *I2PConfig) error {
		if u <= 16 && u > 0 {
			c.InQuantity = strconv.Itoa(u)
			return nil
		}
		return fmt.Errorf("Invalid inbound tunnel quantity")
	}
}

//SetConfigOutQuantity sets the outbound tunnel quantity
func SetConfigOutQuantity(u int) func(*I2PConfig) error {
	return func(c *I2PConfig) error {
		if u <= 16 && u > 0 {
			c.OutQuantity = strconv.Itoa(u)
			return nil
		}
		return fmt.Errorf("Invalid outbound tunnel quantity")
	}
}

//SetConfigInBackups sets the inbound tunnel backups
func SetConfigInBackups(u int) func(*I2PConfig) error {
	return func(c *I2PConfig) error {
		if u < 6 && u >= 0 {
			c.InBackupQuantity = strconv.Itoa(u)
			return nil
		}
		return fmt.Errorf("Invalid inbound tunnel backup quantity")
	}
}

//SetConfigOutBackups sets the inbound tunnel backups
func SetConfigOutBackups(u int) func(*I2PConfig) error {
	return func(c *I2PConfig) error {
		if u < 6 && u >= 0 {
			c.OutBackupQuantity = strconv.Itoa(u)
			return nil
		}
		return fmt.Errorf("Invalid outbound tunnel backup quantity")
	}
}

//SetConfigEncrypt tells the router to use an encrypted leaseset
func SetConfigEncrypt(b bool) func(*I2PConfig) error {
	return func(c *I2PConfig) error {
		if b {
			c.EncryptLeaseSet = "true"
			return nil
		}
		c.EncryptLeaseSet = "false"
		return nil
	}
}

//SetConfigLeaseSetKey sets the host of the I2PConfig's SAM bridge
func SetConfigLeaseSetKey(s string) func(*I2PConfig) error {
	return func(c *I2PConfig) error {
		c.LeaseSetKey = s
		return nil
	}
}

//SetConfigLeaseSetPrivateKey sets the host of the I2PConfig's SAM bridge
func SetConfigLeaseSetPrivateKey(s string) func(*I2PConfig) error {
	return func(c *I2PConfig) error {
		c.LeaseSetPrivateKey = s
		return nil
	}
}

//SetConfigLeaseSetPrivateSigningKey sets the host of the I2PConfig's SAM bridge
func SetConfigLeaseSetPrivateSigningKey(s string) func(*I2PConfig) error {
	return func(c *I2PConfig) error {
		c.LeaseSetPrivateSigningKey = s
		return nil
	}
}

//SetConfigMessageReliability sets the host of the I2PConfig's SAM bridge
func SetConfigMessageReliability(s string) func(*I2PConfig) error {
	return func(c *I2PConfig) error {
		c.MessageReliability = s
		return nil
	}
}

//SetConfigAllowZeroIn tells the tunnel to accept zero-hop peers
func SetConfigAllowZeroIn(b bool) func(*I2PConfig) error {
	return func(c *I2PConfig) error {
		if b {
			c.InAllowZeroHop = "true"
			return nil
		}
		c.InAllowZeroHop = "false"
		return nil
	}
}

//SetConfigAllowZeroOut tells the tunnel to accept zero-hop peers
func SetConfigAllowZeroOut(b bool) func(*I2PConfig) error {
	return func(c *I2PConfig) error {
		if b {
			c.OutAllowZeroHop = "true"
			return nil
		}
		c.OutAllowZeroHop = "false"
		return nil
	}
}

//SetConfigCompress tells clients to use compression
func SetConfigCompress(b bool) func(*I2PConfig) error {
	return func(c *I2PConfig) error {
		if b {
			c.UseCompression = "true"
			return nil
		}
		c.UseCompression = "false"
		return nil
	}
}

//SetConfigFastRecieve tells clients to use compression
func SetConfigFastRecieve(b bool) func(*I2PConfig) error {
	return func(c *I2PConfig) error {
		if b {
			c.FastRecieve = "true"
			return nil
		}
		c.FastRecieve = "false"
		return nil
	}
}

//SetConfigReduceIdle tells the connection to reduce it's tunnels during extended idle time.
func SetConfigReduceIdle(b bool) func(*I2PConfig) error {
	return func(c *I2PConfig) error {
		if b {
			c.ReduceIdle = "true"
			return nil
		}
		c.ReduceIdle = "false"
		return nil
	}
}

//SetConfigReduceIdleTime sets the time to wait before reducing tunnels to idle levels
func SetConfigReduceIdleTime(u int) func(*I2PConfig) error {
	return func(c *I2PConfig) error {
		c.ReduceIdleTime = "300000"
		if u >= 6 {
			c.ReduceIdleTime = strconv.Itoa((u * 60) * 1000)
			return nil
		}
		return fmt.Errorf("Invalid reduce idle timeout(Measured in minutes) %v", u)
	}
}

//SetConfigReduceIdleTimeMs sets the time to wait before reducing tunnels to idle levels in milliseconds
func SetConfigReduceIdleTimeMs(u int) func(*I2PConfig) error {
	return func(c *I2PConfig) error {
		c.ReduceIdleTime = "300000"
		if u >= 300000 {
			c.ReduceIdleTime = strconv.Itoa(u)
			return nil
		}
		return fmt.Errorf("Invalid reduce idle timeout(Measured in milliseconds) %v", u)
	}
}

//SetConfigReduceIdleQuantity sets minimum number of tunnels to reduce to during idle time
func SetConfigReduceIdleQuantity(u int) func(*I2PConfig) error {
	return func(c *I2PConfig) error {
		if u < 5 {
			c.ReduceIdleQuantity = strconv.Itoa(u)
			return nil
		}
		return fmt.Errorf("Invalid reduce tunnel quantity")
	}
}

//SetConfigCloseIdle tells the connection to close it's tunnels during extended idle time.
func SetConfigCloseIdle(b bool) func(*I2PConfig) error {
	return func(c *I2PConfig) error {
		if b {
			c.CloseIdle = "true"
			return nil
		}
		c.CloseIdle = "false"
		return nil
	}
}

//SetConfigCloseIdleTime sets the time to wait before closing tunnels to idle levels
func SetConfigCloseIdleTime(u int) func(*I2PConfig) error {
	return func(c *I2PConfig) error {
		c.CloseIdleTime = "300000"
		if u >= 6 {
			c.CloseIdleTime = strconv.Itoa((u * 60) * 1000)
			return nil
		}
		return fmt.Errorf("Invalid close idle timeout(Measured in minutes) %v", u)
	}
}

//SetConfigCloseIdleTimeMs sets the time to wait before closing tunnels to idle levels in milliseconds
func SetConfigCloseIdleTimeMs(u int) func(*I2PConfig) error {
	return func(c *I2PConfig) error {
		c.CloseIdleTime = "300000"
		if u >= 300000 {
			c.CloseIdleTime = strconv.Itoa(u)
			return nil
		}
		return fmt.Errorf("Invalid close idle timeout(Measured in milliseconds) %v", u)
	}
}

//SetConfigAccessListType tells the system to treat the AccessList as a whitelist
func SetConfigAccessListType(s string) func(*I2PConfig) error {
	return func(c *I2PConfig) error {
		if s == "whitelist" {
			c.AccessListType = "whitelist"
			return nil
		} else if s == "blacklist" {
			c.AccessListType = "blacklist"
			return nil
		} else if s == "none" {
			c.AccessListType = ""
			return nil
		} else if s == "" {
			c.AccessListType = ""
			return nil
		}
		return fmt.Errorf("Invalid Access list type(whitelist, blacklist, none)")
	}
}

//SetConfigAccessList tells the system to treat the AccessList as a whitelist
func SetConfigAccessList(s []string) func(*I2PConfig) error {
	return func(c *I2PConfig) error {
		if len(s) > 0 {
			for _, a := range s {
				c.AccessList = append(c.AccessList, a)
			}
			return nil
		}
		return nil
	}
}
