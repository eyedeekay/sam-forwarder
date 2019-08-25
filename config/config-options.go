package i2ptunconf

import (
	"fmt"
	"strconv"
)

//Option is a Conf Option
type Option func(*Conf) error

//SetFilePath sets the path to save the config file at.
func SetFilePath(s string) func(*Conf) error {
	return func(c *Conf) error {
		c.FilePath = s
		return nil
	}
}

//SetType sets the type of the forwarder server
func SetType(s string) func(*Conf) error {
	return func(c *Conf) error {
		if s == "http" {
			c.Type = s
			return nil
		} else {
			c.Type = "server"
			return nil
		}
	}
}

//SetSigType sets the type of the forwarder server
func SetSigType(s string) func(*Conf) error {
	return func(c *Conf) error {
		if s == "" {
			c.SigType = ""
		} else if s == "DSA_SHA1" {
			c.SigType = "DSA_SHA1"
		} else if s == "ECDSA_SHA256_P256" {
			c.SigType = "ECDSA_SHA256_P256"
		} else if s == "ECDSA_SHA384_P384" {
			c.SigType = "ECDSA_SHA384_P384"
		} else if s == "ECDSA_SHA512_P521" {
			c.SigType = "ECDSA_SHA512_P521"
		} else if s == "EdDSA_SHA512_Ed25519" {
			c.SigType = "EdDSA_SHA512_Ed25519"
		} else {
			c.SigType = "EdDSA_SHA512_Ed25519"
		}
		return nil
	}
}

//SetSaveFile tells the router to save the tunnel's keys long-term
func SetSaveFile(b bool) func(*Conf) error {
	return func(c *Conf) error {
		c.SaveFile = b
		return nil
	}
}

//SetHost sets the host of the service to forward
func SetHost(s string) func(*Conf) error {
	return func(c *Conf) error {
		c.TargetHost = s
		return nil
	}
}

//SetPort sets the port of the service to forward
func SetPort(s string) func(*Conf) error {
	return func(c *Conf) error {
		port, err := strconv.Atoi(s)
		if err != nil {
			return fmt.Errorf("Invalid TCP Server Target Port %s; non-number ", s)
		}
		if port < 65536 && port > -1 {
			c.TargetPort = s
			return nil
		}
		return fmt.Errorf("Invalid port")
	}
}

//SetSAMHost sets the host of the Conf's SAM bridge
func SetSAMHost(s string) func(*Conf) error {
	return func(c *Conf) error {
		c.SamHost = s
		return nil
	}
}

//SetSAMPort sets the port of the Conf's SAM bridge using a string
func SetSAMPort(s string) func(*Conf) error {
	return func(c *Conf) error {
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

//SetName sets the host of the Conf's SAM bridge
func SetName(s string) func(*Conf) error {
	return func(c *Conf) error {
		c.TunName = s
		return nil
	}
}

//SetInLength sets the number of hops inbound
func SetInLength(u int) func(*Conf) error {
	return func(c *Conf) error {
		if u < 7 && u >= 0 {
			c.InLength = u // strconv.Itoa(u)
			return nil
		}
		return fmt.Errorf("Invalid inbound tunnel length")
	}
}

//SetOutLength sets the number of hops outbound
func SetOutLength(u int) func(*Conf) error {
	return func(c *Conf) error {
		if u < 7 && u >= 0 {
			c.OutLength = u // strconv.Itoa(u)
			return nil
		}
		return fmt.Errorf("Invalid outbound tunnel length")
	}
}

//SetInVariance sets the variance of a number of hops inbound
func SetInVariance(i int) func(*Conf) error {
	return func(c *Conf) error {
		if i < 7 && i > -7 {
			c.InVariance = i //strconv.Itoa(i)
			return nil
		}
		return fmt.Errorf("Invalid inbound tunnel length")
	}
}

//SetOutVariance sets the variance of a number of hops outbound
func SetOutVariance(i int) func(*Conf) error {
	return func(c *Conf) error {
		if i < 7 && i > -7 {
			c.OutVariance = i //strconv.Itoa(i)
			return nil
		}
		return fmt.Errorf("Invalid outbound tunnel variance")
	}
}

//SetInQuantity sets the inbound tunnel quantity
func SetInQuantity(u int) func(*Conf) error {
	return func(c *Conf) error {
		if u <= 16 && u > 0 {
			c.InQuantity = u //strconv.Itoa(u)
			return nil
		}
		return fmt.Errorf("Invalid inbound tunnel quantity")
	}
}

//SetOutQuantity sets the outbound tunnel quantity
func SetOutQuantity(u int) func(*Conf) error {
	return func(c *Conf) error {
		if u <= 16 && u > 0 {
			c.OutQuantity = u // strconv.Itoa(u)
			return nil
		}
		return fmt.Errorf("Invalid outbound tunnel quantity")
	}
}

//SetInBackups sets the inbound tunnel backups
func SetInBackups(u int) func(*Conf) error {
	return func(c *Conf) error {
		if u < 6 && u >= 0 {
			c.InBackupQuantity = u // strconv.Itoa(u)
			return nil
		}
		return fmt.Errorf("Invalid inbound tunnel backup quantity")
	}
}

//SetOutBackups sets the inbound tunnel backups
func SetOutBackups(u int) func(*Conf) error {
	return func(c *Conf) error {
		if u < 6 && u >= 0 {
			c.OutBackupQuantity = u // strconv.Itoa(u)
			return nil
		}
		return fmt.Errorf("Invalid outbound tunnel backup quantity")
	}
}

//SetEncrypt tells the router to use an encrypted leaseset
func SetEncrypt(b bool) func(*Conf) error {
	return func(c *Conf) error {
		if b {
			c.EncryptLeaseSet = b //"true"
			return nil
		}
		c.EncryptLeaseSet = b //"false"
		return nil
	}
}

//SetLeaseSetKey sets the host of the Conf's SAM bridge
func SetLeaseSetKey(s string) func(*Conf) error {
	return func(c *Conf) error {
		c.LeaseSetKey = s
		return nil
	}
}

//SetLeaseSetPrivateKey sets the host of the Conf's SAM bridge
func SetLeaseSetPrivateKey(s string) func(*Conf) error {
	return func(c *Conf) error {
		c.LeaseSetPrivateKey = s
		return nil
	}
}

//SetLeaseSetPrivateSigningKey sets the host of the Conf's SAM bridge
func SetLeaseSetPrivateSigningKey(s string) func(*Conf) error {
	return func(c *Conf) error {
		c.LeaseSetPrivateSigningKey = s
		return nil
	}
}

//SetMessageReliability sets the host of the Conf's SAM bridge
func SetMessageReliability(s string) func(*Conf) error {
	return func(c *Conf) error {
		c.MessageReliability = s
		return nil
	}
}

//SetAllowZeroIn tells the tunnel to accept zero-hop peers
func SetAllowZeroIn(b bool) func(*Conf) error {
	return func(c *Conf) error {
		if b {
			c.InAllowZeroHop = b // "true"
			return nil
		}
		c.InAllowZeroHop = b // "false"
		return nil
	}
}

//SetAllowZeroOut tells the tunnel to accept zero-hop peers
func SetAllowZeroOut(b bool) func(*Conf) error {
	return func(c *Conf) error {
		if b {
			c.OutAllowZeroHop = b // "true"
			return nil
		}
		c.OutAllowZeroHop = b // "false"
		return nil
	}
}

//SetCompress tells clients to use compression
func SetCompress(b bool) func(*Conf) error {
	return func(c *Conf) error {
		if b {
			c.UseCompression = b // "true"
			return nil
		}
		c.UseCompression = b // "false"
		return nil
	}
}

//SetFastRecieve tells clients to use compression
func SetFastRecieve(b bool) func(*Conf) error {
	return func(c *Conf) error {
		if b {
			c.FastRecieve = b // "true"
			return nil
		}
		c.FastRecieve = b // "false"
		return nil
	}
}

//SetReduceIdle tells the connection to reduce it's tunnels during extended idle time.
func SetReduceIdle(b bool) func(*Conf) error {
	return func(c *Conf) error {
		if b {
			c.ReduceIdle = b // "true"
			return nil
		}
		c.ReduceIdle = b // "false"
		return nil
	}
}

//SetReduceIdleTime sets the time to wait before reducing tunnels to idle levels
func SetReduceIdleTime(u int) func(*Conf) error {
	return func(c *Conf) error {
		c.ReduceIdleTime = 300000
		if u >= 6 {
			c.ReduceIdleTime = (u * 60) * 1000 // strconv.Itoa((u * 60) * 1000)
			return nil
		}
		return fmt.Errorf("Invalid reduce idle timeout(Measured in minutes) %v", u)
	}
}

//SetReduceIdleTimeMs sets the time to wait before reducing tunnels to idle levels in milliseconds
func SetReduceIdleTimeMs(u int) func(*Conf) error {
	return func(c *Conf) error {
		c.ReduceIdleTime = 300000
		if u >= 300000 {
			c.ReduceIdleTime = u // strconv.Itoa(u)
			return nil
		}
		return fmt.Errorf("Invalid reduce idle timeout(Measured in milliseconds) %v", u)
	}
}

//SetReduceIdleQuantity sets minimum number of tunnels to reduce to during idle time
func SetReduceIdleQuantity(u int) func(*Conf) error {
	return func(c *Conf) error {
		if u < 5 {
			c.ReduceIdleQuantity = u // strconv.Itoa(u)
			return nil
		}
		return fmt.Errorf("Invalid reduce tunnel quantity")
	}
}

//SetCloseIdle tells the connection to close it's tunnels during extended idle time.
func SetCloseIdle(b bool) func(*Conf) error {
	return func(c *Conf) error {
		if b {
			c.CloseIdle = b // "true"
			return nil
		}
		c.CloseIdle = b // "false"
		return nil
	}
}

//SetCloseIdleTime sets the time to wait before closing tunnels to idle levels
func SetCloseIdleTime(u int) func(*Conf) error {
	return func(c *Conf) error {
		c.CloseIdleTime = 300000
		if u >= 6 {
			c.CloseIdleTime = (u * 60) * 1000 // strconv.Itoa((u * 60) * 1000)
			return nil
		}
		return fmt.Errorf("Invalid close idle timeout(Measured in minutes) %v", u)
	}
}

//SetCloseIdleTimeMs sets the time to wait before closing tunnels to idle levels in milliseconds
func SetCloseIdleTimeMs(u int) func(*Conf) error {
	return func(c *Conf) error {
		c.CloseIdleTime = 300000
		if u >= 300000 {
			c.CloseIdleTime = u //strconv.Itoa(u)
			return nil
		}
		return fmt.Errorf("Invalid close idle timeout(Measured in milliseconds) %v", u)
	}
}

//SetAccessListType tells the system to treat the accessList as a whitelist
func SetAccessListType(s string) func(*Conf) error {
	return func(c *Conf) error {
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

//SetAccessList tells the system to treat the accessList as a whitelist
func SetAccessList(s []string) func(*Conf) error {
	return func(c *Conf) error {
		if len(s) > 0 {
			for _, a := range s {
				c.AccessList = append(c.AccessList, a)
			}
			return nil
		}
		return nil
	}
}

//SetTargetForPort sets the port of the Conf's SAM bridge using a string
/*func SetTargetForPort443(s string) func(*Conf) error {
	return func(c *Conf) error {
		port, err := strconv.Atoi(s)
		if err != nil {
			return fmt.Errorf("Invalid Target Port %s; non-number ", s)
		}
		if port < 65536 && port > -1 {
			c.TargetForPort443 = s
			return nil
		}
		return fmt.Errorf("Invalid port")
	}
}
*/

//SetKeyFile sets
func SetKeyFile(s string) func(*Conf) error {
	return func(c *Conf) error {
		c.CryptFile = s
		return nil
	}
}
