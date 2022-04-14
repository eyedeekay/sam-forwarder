package samoptions

import (
	"fmt"
	"strconv"

	samtunnel "github.com/eyedeekay/sam-forwarder/interface"
)

//Option is a SAMForwarder Option
type Option func(samtunnel.SAMTunnel) error

//SetFilePath sets the path to save the config file at.
func SetFilePath(s string) func(samtunnel.SAMTunnel) error {
	return func(c samtunnel.SAMTunnel) error {
		c.Config().FilePath = s
		return nil
	}
}

//SetType sets the type of the forwarder server
func SetType(s string) func(samtunnel.SAMTunnel) error {
	return func(c samtunnel.SAMTunnel) error {
		if s == "http" {
			c.Config().Type = s
			return nil
		} else {
			c.Config().Type = "server"
			return nil
		}
	}
}

//SetSigType sets the type of the forwarder server
func SetSigType(s string) func(samtunnel.SAMTunnel) error {
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

//SetSaveFile tells the router to save the tunnel's keys long-term
func SetSaveFile(b bool) func(samtunnel.SAMTunnel) error {
	return func(c samtunnel.SAMTunnel) error {
		c.Config().SaveFile = b
		return nil
	}
}

//SetHost sets the host of the service to forward
func SetHost(s string) func(samtunnel.SAMTunnel) error {
	return func(c samtunnel.SAMTunnel) error {
		c.Config().TargetHost = s
		return nil
	}
}

//SetPort sets the port of the service to forward
func SetPort(s string) func(samtunnel.SAMTunnel) error {
	return func(c samtunnel.SAMTunnel) error {
		port, err := strconv.Atoi(s)
		if err != nil {
			return fmt.Errorf("Invalid TCP Server Target Port %s; non-number ", s)
		}
		if port < 65536 && port > -1 {
			c.Config().TargetPort = s
			return nil
		}
		return fmt.Errorf("Invalid port")
	}
}

//SetSAMHost sets the host of the SAMForwarder's SAM bridge
func SetSAMHost(s string) func(samtunnel.SAMTunnel) error {
	return func(c samtunnel.SAMTunnel) error {
		c.Config().SamHost = s
		return nil
	}
}

//SetSAMPort sets the port of the SAMForwarder's SAM bridge using a string
func SetSAMPort(s string) func(samtunnel.SAMTunnel) error {
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

//SetName sets the host of the SAMForwarder's SAM bridge
func SetName(s string) func(samtunnel.SAMTunnel) error {
	return func(c samtunnel.SAMTunnel) error {
		c.Config().TunName = s
		return nil
	}
}

//SetInLength sets the number of hops inbound
func SetInLength(u int) func(samtunnel.SAMTunnel) error {
	return func(c samtunnel.SAMTunnel) error {
		if u < 7 && u >= 0 {
			c.Config().InLength = u
			return nil
		}
		return fmt.Errorf("Invalid inbound tunnel length")
	}
}

//SetOutLength sets the number of hops outbound
func SetOutLength(u int) func(samtunnel.SAMTunnel) error {
	return func(c samtunnel.SAMTunnel) error {
		if u < 7 && u >= 0 {
			c.Config().OutLength = u
			return nil
		}
		return fmt.Errorf("Invalid outbound tunnel length")
	}
}

//SetInVariance sets the variance of a number of hops inbound
func SetInVariance(i int) func(samtunnel.SAMTunnel) error {
	return func(c samtunnel.SAMTunnel) error {
		if i < 7 && i > -7 {
			c.Config().InVariance = i
			return nil
		}
		return fmt.Errorf("Invalid inbound tunnel length")
	}
}

//SetOutVariance sets the variance of a number of hops outbound
func SetOutVariance(i int) func(samtunnel.SAMTunnel) error {
	return func(c samtunnel.SAMTunnel) error {
		if i < 7 && i > -7 {
			c.Config().OutVariance = i
			return nil
		}
		return fmt.Errorf("Invalid outbound tunnel variance")
	}
}

//SetInQuantity sets the inbound tunnel quantity
func SetInQuantity(u int) func(samtunnel.SAMTunnel) error {
	return func(c samtunnel.SAMTunnel) error {
		if u <= 16 && u > 0 {
			c.Config().InQuantity = u
			return nil
		}
		return fmt.Errorf("Invalid inbound tunnel quantity")
	}
}

//SetOutQuantity sets the outbound tunnel quantity
func SetOutQuantity(u int) func(samtunnel.SAMTunnel) error {
	return func(c samtunnel.SAMTunnel) error {
		if u <= 16 && u > 0 {
			c.Config().OutQuantity = u
			return nil
		}
		return fmt.Errorf("Invalid outbound tunnel quantity")
	}
}

//SetInBackups sets the inbound tunnel backups
func SetInBackups(u int) func(samtunnel.SAMTunnel) error {
	return func(c samtunnel.SAMTunnel) error {
		if u < 6 && u >= 0 {
			c.Config().InBackupQuantity = u
			return nil
		}
		return fmt.Errorf("Invalid inbound tunnel backup quantity")
	}
}

//SetOutBackups sets the inbound tunnel backups
func SetOutBackups(u int) func(samtunnel.SAMTunnel) error {
	return func(c samtunnel.SAMTunnel) error {
		if u < 6 && u >= 0 {
			c.Config().OutBackupQuantity = u
			return nil
		}
		return fmt.Errorf("Invalid outbound tunnel backup quantity")
	}
}

//SetEncrypt tells the router to use an encrypted leaseset
func SetEncrypt(b bool) func(samtunnel.SAMTunnel) error {
	return func(c samtunnel.SAMTunnel) error {
		if b {
			c.Config().EncryptLeaseSet = true
			return nil
		}
		c.Config().EncryptLeaseSet = false
		return nil
	}
}

//SetLeaseSetKey sets the host of the SAMForwarder's SAM bridge
func SetLeaseSetKey(s string) func(samtunnel.SAMTunnel) error {
	return func(c samtunnel.SAMTunnel) error {
		c.Config().LeaseSetKey = s
		return nil
	}
}

//SetLeaseSetPrivateKey sets the host of the SAMForwarder's SAM bridge
func SetLeaseSetPrivateKey(s string) func(samtunnel.SAMTunnel) error {
	return func(c samtunnel.SAMTunnel) error {
		c.Config().LeaseSetPrivateKey = s
		return nil
	}
}

//SetLeaseSetPrivateSigningKey sets the host of the SAMForwarder's SAM bridge
func SetLeaseSetPrivateSigningKey(s string) func(samtunnel.SAMTunnel) error {
	return func(c samtunnel.SAMTunnel) error {
		c.Config().LeaseSetPrivateSigningKey = s
		return nil
	}
}

//SetMessageReliability sets the host of the SAMForwarder's SAM bridge
func SetMessageReliability(s string) func(samtunnel.SAMTunnel) error {
	return func(c samtunnel.SAMTunnel) error {
		c.Config().MessageReliability = s
		return nil
	}
}

//SetAllowZeroIn tells the tunnel to accept zero-hop peers
func SetAllowZeroIn(b bool) func(samtunnel.SAMTunnel) error {
	return func(c samtunnel.SAMTunnel) error {
		if b {
			c.Config().InAllowZeroHop = true
			return nil
		}
		c.Config().InAllowZeroHop = false
		return nil
	}
}

//SetAllowZeroOut tells the tunnel to accept zero-hop peers
func SetAllowZeroOut(b bool) func(samtunnel.SAMTunnel) error {
	return func(c samtunnel.SAMTunnel) error {
		if b {
			c.Config().OutAllowZeroHop = true
			return nil
		}
		c.Config().OutAllowZeroHop = false
		return nil
	}
}

//SetCompress tells clients to use compression
func SetCompress(b bool) func(samtunnel.SAMTunnel) error {
	return func(c samtunnel.SAMTunnel) error {
		if b {
			c.Config().UseCompression = true
			return nil
		}
		c.Config().UseCompression = false
		return nil
	}
}

//SetFastRecieve tells clients to use compression
func SetFastRecieve(b bool) func(samtunnel.SAMTunnel) error {
	return func(c samtunnel.SAMTunnel) error {
		if b {
			c.Config().FastRecieve = true
			return nil
		}
		c.Config().FastRecieve = false
		return nil
	}
}

//SetReduceIdle tells the connection to reduce it's tunnels during extended idle time.
func SetReduceIdle(b bool) func(samtunnel.SAMTunnel) error {
	return func(c samtunnel.SAMTunnel) error {
		if b {
			c.Config().ReduceIdle = true
			return nil
		}
		c.Config().ReduceIdle = false
		return nil
	}
}

//SetReduceIdleTime sets the time to wait before reducing tunnels to idle levels
func SetReduceIdleTime(u int) func(samtunnel.SAMTunnel) error {
	return func(c samtunnel.SAMTunnel) error {
		c.Config().ReduceIdleTime = 300000
		if u >= 6 {
			c.Config().ReduceIdleTime = (u * 60) * 1000
			return nil
		}
		return fmt.Errorf("Invalid reduce idle timeout(Measured in minutes) %v", u)
	}
}

//SetReduceIdleTimeMs sets the time to wait before reducing tunnels to idle levels in milliseconds
func SetReduceIdleTimeMs(u int) func(samtunnel.SAMTunnel) error {
	return func(c samtunnel.SAMTunnel) error {
		c.Config().ReduceIdleTime = 300000
		if u >= 300000 {
			c.Config().ReduceIdleTime = u
			return nil
		}
		return fmt.Errorf("Invalid reduce idle timeout(Measured in milliseconds) %v", u)
	}
}

//SetReduceIdleQuantity sets minimum number of tunnels to reduce to during idle time
func SetReduceIdleQuantity(u int) func(samtunnel.SAMTunnel) error {
	return func(c samtunnel.SAMTunnel) error {
		if u < 5 {
			c.Config().ReduceIdleQuantity = u
			return nil
		}
		return fmt.Errorf("Invalid reduce tunnel quantity")
	}
}

//SetCloseIdle tells the connection to close it's tunnels during extended idle time.
func SetCloseIdle(b bool) func(samtunnel.SAMTunnel) error {
	return func(c samtunnel.SAMTunnel) error {
		if b {
			c.Config().CloseIdle = true
			return nil
		}
		c.Config().CloseIdle = false
		return nil
	}
}

//SetCloseIdleTime sets the time to wait before closing tunnels to idle levels
func SetCloseIdleTime(u int) func(samtunnel.SAMTunnel) error {
	return func(c samtunnel.SAMTunnel) error {
		c.Config().CloseIdleTime = 300000
		if u >= 6 {
			c.Config().CloseIdleTime = (u * 60) * 1000
			return nil
		}
		return fmt.Errorf("Invalid close idle timeout(Measured in minutes) %v", u)
	}
}

//SetCloseIdleTimeMs sets the time to wait before closing tunnels to idle levels in milliseconds
func SetCloseIdleTimeMs(u int) func(samtunnel.SAMTunnel) error {
	return func(c samtunnel.SAMTunnel) error {
		c.Config().CloseIdleTime = 300000
		if u >= 300000 {
			c.Config().CloseIdleTime = u
			return nil
		}
		return fmt.Errorf("Invalid close idle timeout(Measured in milliseconds) %v", u)
	}
}

//SetAccessListType tells the system to treat the AccessList as a allowlist
func SetAccessListType(s string) func(samtunnel.SAMTunnel) error {
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

//SetAccessList tells the system to treat the AccessList as a allowlist
func SetAccessList(s []string) func(samtunnel.SAMTunnel) error {
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

func SetUseTLS(b bool) func(samtunnel.SAMTunnel) error {
	return func(c samtunnel.SAMTunnel) error {
		if b {
			c.Config().UseTLS = true
			return nil
		}
		return nil
	}
}

func SetTLSConfigCertPem(s string) func(samtunnel.SAMTunnel) error {
	return func(c samtunnel.SAMTunnel) error {
		if len(s) > 0 {
			c.Config().Cert = s
		}
		return nil
	}
}

func SetTLSConfigKeysPem(s string) func(samtunnel.SAMTunnel) error {
	return func(c samtunnel.SAMTunnel) error {
		if len(s) > 0 {
			c.Config().Pem = s
		}
		return nil
	}
}

//SetTargetForPort sets the port of the SAMForwarder's SAM bridge using a string
/*func SetTargetForPort443(s string) func(samtunnel.SAMTunnel) error {
	return func(c samtunnel.SAMTunnel) error {
		port, err := strconv.Atoi(s)
		if err != nil {
			return fmt.Errorf("Invalid Target Port %s; non-number ", s)
		}
		if port < 65536 && port > -1 {
			c.Config().TargetForPort443 = s
			return nil
		}
		return fmt.Errorf("Invalid port")
	}
}
*/

//SetKeyFile sets
func SetKeyFile(s string) func(samtunnel.SAMTunnel) error {
	return func(c samtunnel.SAMTunnel) error {
		c.Config().KeyFilePath = s
		return nil
	}
}

func SetPassword(s string) func(samtunnel.SAMTunnel) error {
	return func(c samtunnel.SAMTunnel) error {
		c.Config().KeyFilePath = s
		return nil
	}
}

func SetDestination(s string) func(samtunnel.SAMTunnel) error {
	return func(c samtunnel.SAMTunnel) error {
		c.Config().ClientDest = s
		return nil
	}
}
