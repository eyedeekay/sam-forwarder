package samforwarder

import (
	"fmt"
	"strconv"

	"github.com/eyedeekay/sam-forwarder/interface"
)

//ClientOption is a SAMClientForwarder Option
type ClientOption func(samtunnel.SAMTunnel) error

//SetClientFilePath sets the host of the SAMClientForwarder's SAM bridge
func SetClientFilePath(s string) func(samtunnel.SAMTunnel) error {
	return func(c samtunnel.SAMTunnel) error {
		c.Config().FilePath = s
		return nil
	}
}

//SetClientSaveFile tells the router to save the tunnel keys long-term
func SetClientSaveFile(b bool) func(samtunnel.SAMTunnel) error {
	return func(c samtunnel.SAMTunnel) error {
		c.Config().SaveFile = b
		return nil
	}
}

//SetClientHost sets the host of the SAMClientForwarder's SAM bridge
func SetClientHost(s string) func(samtunnel.SAMTunnel) error {
	return func(c samtunnel.SAMTunnel) error {
		c.Config().TargetHost = s
		return nil
	}
}

//SetClientDestination sets the destination to forwarder SAMClientForwarder's to
func SetClientDestination(s string) func(samtunnel.SAMTunnel) error {
	return func(c samtunnel.SAMTunnel) error {
		c.Config().ClientDest = s
		return nil
	}
}

//SetClientPort sets the port of the SAMClientForwarder's SAM bridge using a string
func SetClientPort(s string) func(samtunnel.SAMTunnel) error {
	return func(c samtunnel.SAMTunnel) error {
		port, err := strconv.Atoi(s)
		if err != nil {
			return fmt.Errorf("Invalid TCP Client Target Port %s; non-number ", s)
		}
		if port < 65536 && port > -1 {
			c.Config().TargetPort = s
			return nil
		}
		return fmt.Errorf("Invalid port")
	}
}

//SetClientSAMHost sets the host of the SAMClientForwarder's SAM bridge
func SetClientSAMHost(s string) func(samtunnel.SAMTunnel) error {
	return func(c samtunnel.SAMTunnel) error {
		c.Config().SamHost = s
		return nil
	}
}

//SetClientSAMPort sets the port of the SAMClientForwarder's SAM bridge using a string
func SetClientSAMPort(s string) func(samtunnel.SAMTunnel) error {
	return func(c samtunnel.SAMTunnel) error {
		port, err := strconv.Atoi(s)
		if err != nil {
			return fmt.Errorf("Invalid SAM Port %s; non-number", s)
		}
		if port < 65536 && port > -1 {
			c.Config().SamPort = s
			return nil
		}
		return fmt.Errorf("Invalid port")
	}
}

//SetClientName sets the host of the SAMClientForwarder's SAM bridge
func SetClientName(s string) func(samtunnel.SAMTunnel) error {
	return func(c samtunnel.SAMTunnel) error {
		c.Config().TunName = s
		return nil
	}
}

//SetSigType sets the type of the forwarder server
func SetClientSigType(s string) func(samtunnel.SAMTunnel) error {
	return func(c samtunnel.SAMTunnel) error {
		if s == "" {
			c.Config().SigType = ""
		} else if s == "DSA_SHA1" {
			c.Config().SigType = "DSA_SHA1"
		} else if s == "ECDSA_SHA256_P256" {
			c.Config().SigType = "ECDSA_SHA256_P256"
		} else if s == "ECDSA_SHA384_P384" {
			c.Config().SigType = "ECDSA_SHA384_P384"
		} else if s == "ECDSA_SHA512_P521" {
			c.Config().SigType = "ECDSA_SHA512_P521"
		} else if s == "EdDSA_SHA512_Ed25519" {
			c.Config().SigType = "EdDSA_SHA512_Ed25519"
		} else {
			c.Config().SigType = "EdDSA_SHA512_Ed25519"
		}
		return nil
	}
}

//SetClientInLength sets the number of hops inbound
func SetClientInLength(u int) func(samtunnel.SAMTunnel) error {
	return func(c samtunnel.SAMTunnel) error {
		if u < 7 && u >= 0 {
			c.Config().InLength = u
			return nil
		}
		return fmt.Errorf("Invalid inbound tunnel length")
	}
}

//SetClientOutLength sets the number of hops outbound
func SetClientOutLength(u int) func(samtunnel.SAMTunnel) error {
	return func(c samtunnel.SAMTunnel) error {
		if u < 7 && u >= 0 {
			c.Config().OutLength = u
			return nil
		}
		return fmt.Errorf("Invalid outbound tunnel length")
	}
}

//SetClientInVariance sets the variance of a number of hops inbound
func SetClientInVariance(i int) func(samtunnel.SAMTunnel) error {
	return func(c samtunnel.SAMTunnel) error {
		if i < 7 && i > -7 {
			c.Config().InVariance = i
			return nil
		}
		return fmt.Errorf("Invalid inbound tunnel length")
	}
}

//SetClientOutVariance sets the variance of a number of hops outbound
func SetClientOutVariance(i int) func(samtunnel.SAMTunnel) error {
	return func(c samtunnel.SAMTunnel) error {
		if i < 7 && i > -7 {
			c.Config().OutVariance = i
			return nil
		}
		return fmt.Errorf("Invalid outbound tunnel variance")
	}
}

//SetClientInQuantity sets the inbound tunnel quantity
func SetClientInQuantity(u int) func(samtunnel.SAMTunnel) error {
	return func(c samtunnel.SAMTunnel) error {
		if u <= 16 && u > 0 {
			c.Config().InQuantity = u
			return nil
		}
		return fmt.Errorf("Invalid inbound tunnel quantity")
	}
}

//SetClientOutQuantity sets the outbound tunnel quantity
func SetClientOutQuantity(u int) func(samtunnel.SAMTunnel) error {
	return func(c samtunnel.SAMTunnel) error {
		if u <= 16 && u > 0 {
			c.Config().OutQuantity = u
			return nil
		}
		return fmt.Errorf("Invalid outbound tunnel quantity")
	}
}

//SetClientInBackups sets the inbound tunnel backups
func SetClientInBackups(u int) func(samtunnel.SAMTunnel) error {
	return func(c samtunnel.SAMTunnel) error {
		if u < 6 && u >= 0 {
			c.Config().InBackupQuantity = u
			return nil
		}
		return fmt.Errorf("Invalid inbound tunnel backup quantity")
	}
}

//SetClientOutBackups sets the inbound tunnel backups
func SetClientOutBackups(u int) func(samtunnel.SAMTunnel) error {
	return func(c samtunnel.SAMTunnel) error {
		if u < 6 && u >= 0 {
			c.Config().OutBackupQuantity = u
			return nil
		}
		return fmt.Errorf("Invalid outbound tunnel backup quantity")
	}
}

//SetClientEncrypt tells the router to use an encrypted leaseset
func SetClientEncrypt(b bool) func(samtunnel.SAMTunnel) error {
	return func(c samtunnel.SAMTunnel) error {
		if b {
			c.Config().EncryptLeaseSet = true
			return nil
		}
		c.Config().EncryptLeaseSet = false
		return nil
	}
}

//SetClientLeaseSetKey sets
func SetClientLeaseSetKey(s string) func(samtunnel.SAMTunnel) error {
	return func(c samtunnel.SAMTunnel) error {
		c.Config().LeaseSetKey = s
		return nil
	}
}

//SetClientLeaseSetPrivateKey sets
func SetClientLeaseSetPrivateKey(s string) func(samtunnel.SAMTunnel) error {
	return func(c samtunnel.SAMTunnel) error {
		c.Config().LeaseSetPrivateKey = s
		return nil
	}
}

//SetClientLeaseSetPrivateSigningKey sets
func SetClientLeaseSetPrivateSigningKey(s string) func(samtunnel.SAMTunnel) error {
	return func(c samtunnel.SAMTunnel) error {
		c.Config().LeaseSetPrivateSigningKey = s
		return nil
	}
}

//SetClientMessageReliability sets
func SetClientMessageReliability(s string) func(samtunnel.SAMTunnel) error {
	return func(c samtunnel.SAMTunnel) error {
		c.Config().MessageReliability = s
		return nil
	}
}

//SetClientAllowZeroIn tells the tunnel to accept zero-hop peers
func SetClientAllowZeroIn(b bool) func(samtunnel.SAMTunnel) error {
	return func(c samtunnel.SAMTunnel) error {
		if b {
			c.Config().InAllowZeroHop = true
			return nil
		}
		c.Config().InAllowZeroHop = false
		return nil
	}
}

//SetClientAllowZeroOut tells the tunnel to accept zero-hop peers
func SetClientAllowZeroOut(b bool) func(samtunnel.SAMTunnel) error {
	return func(c samtunnel.SAMTunnel) error {
		if b {
			c.Config().OutAllowZeroHop = true
			return nil
		}
		c.Config().OutAllowZeroHop = false
		return nil
	}
}

//SetClientFastRecieve tells clients use the i2cp.fastRecieve option
func SetClientFastRecieve(b bool) func(samtunnel.SAMTunnel) error {
	return func(c samtunnel.SAMTunnel) error {
		if b {
			c.Config().FastRecieve = true
			return nil
		}
		c.Config().FastRecieve = false
		return nil
	}
}

//SetClientCompress tells clients to use compression
func SetClientCompress(b bool) func(samtunnel.SAMTunnel) error {
	return func(c samtunnel.SAMTunnel) error {
		if b {
			c.Config().UseCompression = true
			return nil
		}
		c.Config().UseCompression = false
		return nil
	}
}

//SetClientReduceIdle tells the connection to reduce it's tunnels during extended idle time.
func SetClientReduceIdle(b bool) func(samtunnel.SAMTunnel) error {
	return func(c samtunnel.SAMTunnel) error {
		if b {
			c.Config().ReduceIdle = true
			return nil
		}
		c.Config().ReduceIdle = false
		return nil
	}
}

//SetClientReduceIdleTime sets the time to wait before reducing tunnels to idle levels
func SetClientReduceIdleTime(u int) func(samtunnel.SAMTunnel) error {
	return func(c samtunnel.SAMTunnel) error {
		c.Config().ReduceIdleTime = 300000
		if u >= 6 {
			c.Config().ReduceIdleTime = (u * 60) * 1000
			return nil
		}
		return fmt.Errorf("Invalid reduce idle timeout(Measured in minutes) %v", u)
	}
}

//SetClientReduceIdleTimeMs sets the time to wait before reducing tunnels to idle levels in milliseconds
func SetClientReduceIdleTimeMs(u int) func(samtunnel.SAMTunnel) error {
	return func(c samtunnel.SAMTunnel) error {
		c.Config().ReduceIdleTime = 300000
		if u >= 300000 {
			c.Config().ReduceIdleTime = u
			return nil
		}
		return fmt.Errorf("Invalid reduce idle timeout(Measured in milliseconds) %v", u)
	}
}

//SetClientReduceIdleQuantity sets minimum number of tunnels to reduce to during idle time
func SetClientReduceIdleQuantity(u int) func(samtunnel.SAMTunnel) error {
	return func(c samtunnel.SAMTunnel) error {
		if u < 5 {
			c.Config().ReduceIdleQuantity = u
			return nil
		}
		return fmt.Errorf("Invalid reduce tunnel quantity")
	}
}

//SetClientCloseIdle tells the connection to close it's tunnels during extended idle time.
func SetClientCloseIdle(b bool) func(samtunnel.SAMTunnel) error {
	return func(c samtunnel.SAMTunnel) error {
		if b {
			c.Config().CloseIdle = true
			return nil
		}
		c.Config().CloseIdle = false
		return nil
	}
}

//SetClientCloseIdleTime sets the time to wait before closing tunnels to idle levels
func SetClientCloseIdleTime(u int) func(samtunnel.SAMTunnel) error {
	return func(c samtunnel.SAMTunnel) error {
		c.Config().CloseIdleTime = 300000
		if u >= 6 {
			c.Config().CloseIdleTime = (u * 60) * 1000
			return nil
		}
		return fmt.Errorf("Invalid close idle timeout(Measured in minutes) %v", u)
	}
}

//SetClientCloseIdleTimeMs sets the time to wait before closing tunnels to idle levels in milliseconds
func SetClientCloseIdleTimeMs(u int) func(samtunnel.SAMTunnel) error {
	return func(c samtunnel.SAMTunnel) error {
		c.Config().CloseIdleTime = 300000
		if u >= 300000 {
			c.Config().CloseIdleTime = u
			return nil
		}
		return fmt.Errorf("Invalid close idle timeout(Measured in milliseconds) %v", u)
	}
}

//SetClientAccessListType tells the system to treat the accessList as a allowlist
func SetClientAccessListType(s string) func(samtunnel.SAMTunnel) error {
	return func(c samtunnel.SAMTunnel) error {
		if s == "allowlist" {
			c.Config().AccessListType = "allowlist"
			return nil
		} else if s == "blocklist" {
			c.Config().AccessListType = "blocklist"
			return nil
		} else if s == "none" {
			c.Config().AccessListType = ""
			return nil
		} else if s == "" {
			c.Config().AccessListType = ""
			return nil
		}
		return fmt.Errorf("Invalid Access list type(allowlist, blocklist, none)")
	}
}

//SetClientAccessList tells the system to treat the accessList as a allowlist
func SetClientAccessList(s []string) func(samtunnel.SAMTunnel) error {
	return func(c samtunnel.SAMTunnel) error {
		if len(s) > 0 {
			for _, a := range s {
				c.Config().AccessList = append(c.Config().AccessList, a)
			}
			return nil
		}
		return nil
	}
}

//SetKeyFile sets
func SetClientPassword(s string) func(samtunnel.SAMTunnel) error {
	return func(c samtunnel.SAMTunnel) error {
		c.Config().KeyFilePath = s
		return nil
	}
}
