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
		c.Conf.FilePath = s
		return nil
	}
}

//SetClientSaveFile tells the router to save the tunnel keys long-term
func SetClientSaveFile(b bool) func(*SAMClientForwarder) error {
	return func(c *SAMClientForwarder) error {
		c.Conf.SaveFile = b
		return nil
	}
}

//SetClientHost sets the host of the SAMClientForwarder's SAM bridge
func SetClientHost(s string) func(*SAMClientForwarder) error {
	return func(c *SAMClientForwarder) error {
		c.Conf.TargetHost = s
		return nil
	}
}

//SetClientDestination sets the destination to forwarder SAMClientForwarder's to
func SetClientDestination(s string) func(*SAMClientForwarder) error {
	return func(c *SAMClientForwarder) error {
		c.Conf.ClientDest = s
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
			c.Conf.TargetPort = s
			return nil
		}
		return fmt.Errorf("Invalid port")
	}
}

//SetClientSAMHost sets the host of the SAMClientForwarder's SAM bridge
func SetClientSAMHost(s string) func(*SAMClientForwarder) error {
	return func(c *SAMClientForwarder) error {
		c.Conf.SamHost = s
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
			c.Conf.SamPort = s
			return nil
		}
		return fmt.Errorf("Invalid port")
	}
}

//SetClientName sets the host of the SAMClientForwarder's SAM bridge
func SetClientName(s string) func(*SAMClientForwarder) error {
	return func(c *SAMClientForwarder) error {
		c.Conf.TunName = s
		return nil
	}
}

//SetSigType sets the type of the forwarder server
func SetClientSigType(s string) func(*SAMClientForwarder) error {
	return func(c *SAMClientForwarder) error {
		if s == "" {
			c.Conf.SigType = ""
		} else if s == "DSA_SHA1" {
			c.Conf.SigType = "DSA_SHA1"
		} else if s == "ECDSA_SHA256_P256" {
			c.Conf.SigType = "ECDSA_SHA256_P256"
		} else if s == "ECDSA_SHA384_P384" {
			c.Conf.SigType = "ECDSA_SHA384_P384"
		} else if s == "ECDSA_SHA512_P521" {
			c.Conf.SigType = "ECDSA_SHA512_P521"
		} else if s == "EdDSA_SHA512_Ed25519" {
			c.Conf.SigType = "EdDSA_SHA512_Ed25519"
		} else {
			c.Conf.SigType = "EdDSA_SHA512_Ed25519"
		}
		return nil
	}
}

//SetClientInLength sets the number of hops inbound
func SetClientInLength(u int) func(*SAMClientForwarder) error {
	return func(c *SAMClientForwarder) error {
		if u < 7 && u >= 0 {
			c.Conf.InLength = u
			return nil
		}
		return fmt.Errorf("Invalid inbound tunnel length")
	}
}

//SetClientOutLength sets the number of hops outbound
func SetClientOutLength(u int) func(*SAMClientForwarder) error {
	return func(c *SAMClientForwarder) error {
		if u < 7 && u >= 0 {
			c.Conf.OutLength = u
			return nil
		}
		return fmt.Errorf("Invalid outbound tunnel length")
	}
}

//SetClientInVariance sets the variance of a number of hops inbound
func SetClientInVariance(i int) func(*SAMClientForwarder) error {
	return func(c *SAMClientForwarder) error {
		if i < 7 && i > -7 {
			c.Conf.InVariance = i
			return nil
		}
		return fmt.Errorf("Invalid inbound tunnel length")
	}
}

//SetClientOutVariance sets the variance of a number of hops outbound
func SetClientOutVariance(i int) func(*SAMClientForwarder) error {
	return func(c *SAMClientForwarder) error {
		if i < 7 && i > -7 {
			c.Conf.OutVariance = i
			return nil
		}
		return fmt.Errorf("Invalid outbound tunnel variance")
	}
}

//SetClientInQuantity sets the inbound tunnel quantity
func SetClientInQuantity(u int) func(*SAMClientForwarder) error {
	return func(c *SAMClientForwarder) error {
		if u <= 16 && u > 0 {
			c.Conf.InQuantity = u
			return nil
		}
		return fmt.Errorf("Invalid inbound tunnel quantity")
	}
}

//SetClientOutQuantity sets the outbound tunnel quantity
func SetClientOutQuantity(u int) func(*SAMClientForwarder) error {
	return func(c *SAMClientForwarder) error {
		if u <= 16 && u > 0 {
			c.Conf.OutQuantity = u
			return nil
		}
		return fmt.Errorf("Invalid outbound tunnel quantity")
	}
}

//SetClientInBackups sets the inbound tunnel backups
func SetClientInBackups(u int) func(*SAMClientForwarder) error {
	return func(c *SAMClientForwarder) error {
		if u < 6 && u >= 0 {
			c.Conf.InBackupQuantity = u
			return nil
		}
		return fmt.Errorf("Invalid inbound tunnel backup quantity")
	}
}

//SetClientOutBackups sets the inbound tunnel backups
func SetClientOutBackups(u int) func(*SAMClientForwarder) error {
	return func(c *SAMClientForwarder) error {
		if u < 6 && u >= 0 {
			c.Conf.OutBackupQuantity = u
			return nil
		}
		return fmt.Errorf("Invalid outbound tunnel backup quantity")
	}
}

//SetClientEncrypt tells the router to use an encrypted leaseset
func SetClientEncrypt(b bool) func(*SAMClientForwarder) error {
	return func(c *SAMClientForwarder) error {
		if b {
			c.Conf.EncryptLeaseSet = true
			return nil
		}
		c.Conf.EncryptLeaseSet = false
		return nil
	}
}

//SetClientLeaseSetKey sets
func SetClientLeaseSetKey(s string) func(*SAMClientForwarder) error {
	return func(c *SAMClientForwarder) error {
		c.Conf.LeaseSetKey = s
		return nil
	}
}

//SetClientLeaseSetPrivateKey sets
func SetClientLeaseSetPrivateKey(s string) func(*SAMClientForwarder) error {
	return func(c *SAMClientForwarder) error {
		c.Conf.LeaseSetPrivateKey = s
		return nil
	}
}

//SetClientLeaseSetPrivateSigningKey sets
func SetClientLeaseSetPrivateSigningKey(s string) func(*SAMClientForwarder) error {
	return func(c *SAMClientForwarder) error {
		c.Conf.LeaseSetPrivateSigningKey = s
		return nil
	}
}

//SetClientMessageReliability sets
func SetClientMessageReliability(s string) func(*SAMClientForwarder) error {
	return func(c *SAMClientForwarder) error {
		c.Conf.MessageReliability = s
		return nil
	}
}

//SetClientAllowZeroIn tells the tunnel to accept zero-hop peers
func SetClientAllowZeroIn(b bool) func(*SAMClientForwarder) error {
	return func(c *SAMClientForwarder) error {
		if b {
			c.Conf.InAllowZeroHop = true
			return nil
		}
		c.Conf.InAllowZeroHop = false
		return nil
	}
}

//SetClientAllowZeroOut tells the tunnel to accept zero-hop peers
func SetClientAllowZeroOut(b bool) func(*SAMClientForwarder) error {
	return func(c *SAMClientForwarder) error {
		if b {
			c.Conf.OutAllowZeroHop = true
			return nil
		}
		c.Conf.OutAllowZeroHop = false
		return nil
	}
}

//SetClientFastRecieve tells clients use the i2cp.fastRecieve option
func SetClientFastRecieve(b bool) func(*SAMClientForwarder) error {
	return func(c *SAMClientForwarder) error {
		if b {
			c.Conf.FastRecieve = true
			return nil
		}
		c.Conf.FastRecieve = false
		return nil
	}
}

//SetClientCompress tells clients to use compression
func SetClientCompress(b bool) func(*SAMClientForwarder) error {
	return func(c *SAMClientForwarder) error {
		if b {
			c.Conf.UseCompression = true
			return nil
		}
		c.Conf.UseCompression = false
		return nil
	}
}

//SetClientReduceIdle tells the connection to reduce it's tunnels during extended idle time.
func SetClientReduceIdle(b bool) func(*SAMClientForwarder) error {
	return func(c *SAMClientForwarder) error {
		if b {
			c.Conf.ReduceIdle = true
			return nil
		}
		c.Conf.ReduceIdle = false
		return nil
	}
}

//SetClientReduceIdleTime sets the time to wait before reducing tunnels to idle levels
func SetClientReduceIdleTime(u int) func(*SAMClientForwarder) error {
	return func(c *SAMClientForwarder) error {
		c.Conf.ReduceIdleTime = 300000
		if u >= 6 {
			c.Conf.ReduceIdleTime = (u * 60) * 1000
			return nil
		}
		return fmt.Errorf("Invalid reduce idle timeout(Measured in minutes) %v", u)
	}
}

//SetClientReduceIdleTimeMs sets the time to wait before reducing tunnels to idle levels in milliseconds
func SetClientReduceIdleTimeMs(u int) func(*SAMClientForwarder) error {
	return func(c *SAMClientForwarder) error {
		c.Conf.ReduceIdleTime = 300000
		if u >= 300000 {
			c.Conf.ReduceIdleTime = u
			return nil
		}
		return fmt.Errorf("Invalid reduce idle timeout(Measured in milliseconds) %v", u)
	}
}

//SetClientReduceIdleQuantity sets minimum number of tunnels to reduce to during idle time
func SetClientReduceIdleQuantity(u int) func(*SAMClientForwarder) error {
	return func(c *SAMClientForwarder) error {
		if u < 5 {
			c.Conf.ReduceIdleQuantity = u
			return nil
		}
		return fmt.Errorf("Invalid reduce tunnel quantity")
	}
}

//SetClientCloseIdle tells the connection to close it's tunnels during extended idle time.
func SetClientCloseIdle(b bool) func(*SAMClientForwarder) error {
	return func(c *SAMClientForwarder) error {
		if b {
			c.Conf.CloseIdle = true
			return nil
		}
		c.Conf.CloseIdle = false
		return nil
	}
}

//SetClientCloseIdleTime sets the time to wait before closing tunnels to idle levels
func SetClientCloseIdleTime(u int) func(*SAMClientForwarder) error {
	return func(c *SAMClientForwarder) error {
		c.Conf.CloseIdleTime = 300000
		if u >= 6 {
			c.Conf.CloseIdleTime = (u * 60) * 1000
			return nil
		}
		return fmt.Errorf("Invalid close idle timeout(Measured in minutes) %v", u)
	}
}

//SetClientCloseIdleTimeMs sets the time to wait before closing tunnels to idle levels in milliseconds
func SetClientCloseIdleTimeMs(u int) func(*SAMClientForwarder) error {
	return func(c *SAMClientForwarder) error {
		c.Conf.CloseIdleTime = 300000
		if u >= 300000 {
			c.Conf.CloseIdleTime = u
			return nil
		}
		return fmt.Errorf("Invalid close idle timeout(Measured in milliseconds) %v", u)
	}
}

//SetClientAccessListType tells the system to treat the accessList as a allowlist
func SetClientAccessListType(s string) func(*SAMClientForwarder) error {
	return func(c *SAMClientForwarder) error {
		if s == "allowlist" {
			c.Conf.AccessListType = "allowlist"
			return nil
		} else if s == "blocklist" {
			c.Conf.AccessListType = "blocklist"
			return nil
		} else if s == "none" {
			c.Conf.AccessListType = ""
			return nil
		} else if s == "" {
			c.Conf.AccessListType = ""
			return nil
		}
		return fmt.Errorf("Invalid Access list type(allowlist, blocklist, none)")
	}
}

//SetClientAccessList tells the system to treat the accessList as a allowlist
func SetClientAccessList(s []string) func(*SAMClientForwarder) error {
	return func(c *SAMClientForwarder) error {
		if len(s) > 0 {
			for _, a := range s {
				c.Conf.AccessList = append(c.Conf.AccessList, a)
			}
			return nil
		}
		return nil
	}
}

//SetKeyFile sets
func SetClientPassword(s string) func(*SAMClientForwarder) error {
	return func(c *SAMClientForwarder) error {
		c.Conf.KeyFilePath = s
		return nil
	}
}
