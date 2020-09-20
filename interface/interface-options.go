package samtunnel

import (
	"fmt"
	"strconv"
)

//SetAccessListType tells the system to treat the accessList as a allowlist
func SetAccessListType(s string) func(SAMTunnel) error {
	return func(c SAMTunnel) error {
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

//SetAccessList tells the system to treat the accessList as a allowlist
func SetAccessList(s []string) func(SAMTunnel) error {
	return func(c SAMTunnel) error {
		if len(s) > 0 {
			for _, a := range s {
				c.Config().AccessList = append(c.Config().AccessList, a)
			}
			return nil
		}
		return nil
	}
}

//SetCompress tells clients to use compression
func SetCompress(b bool) func(SAMTunnel) error {
	return func(c SAMTunnel) error {
		c.Config().UseCompression = b // "false"
		return nil
	}
}

//SetFilePath sets the path to save the config file at.
func SetFilePath(s string) func(SAMTunnel) error {
	return func(c SAMTunnel) error {
		c.Config().FilePath = s
		return nil
	}
}

//SetControlHost sets the host of the service to present an API on
func SetControlHost(s string) func(SAMTunnel) error {
	return func(c SAMTunnel) error {
		c.Config().ControlHost = s
		return nil
	}
}

//SetControlPort sets the port of the service to present an API on
func SetControlPort(s string) func(SAMTunnel) error {
	return func(c SAMTunnel) error {
		port, err := strconv.Atoi(s)
		if err != nil {
			return fmt.Errorf("Invalid TCP Server Target Port %s; non-number ", s)
		}
		if port < 65536 && port > -1 {
			c.Config().ControlPort = s
			return nil
		}
		return fmt.Errorf("Invalid port")
	}
}

//SetKeyFile sets the path to a file containing a private key for decrypting
//locally-encrypted i2p keys.
func SetKeyFile(s string) func(SAMTunnel) error {
	return func(c SAMTunnel) error {
		c.Config().KeyFilePath = s
		return nil
	}
}

//SetDestination tells the
func SetDestination(b string) func(SAMTunnel) error {
	return func(c SAMTunnel) error {
		c.Config().ClientDest = b
		return nil
	}
}

//SetTunnelHost is used for VPN endpoints only.
func SetTunnelHost(s string) func(SAMTunnel) error {
	return func(c SAMTunnel) error {
		c.Config().TunnelHost = s
		return nil
	}
}

//SetFastRecieve tells clients to recieve all messages as quicky as possible
func SetFastRecieve(b bool) func(SAMTunnel) error {
	return func(c SAMTunnel) error {
		c.Config().FastRecieve = b
		return nil
	}
}

//SetHost sets the host of the service to forward
func SetHost(s string) func(SAMTunnel) error {
	return func(c SAMTunnel) error {
		c.Config().TargetHost = s
		return nil
	}
}

//SetPort sets the port of the service to forward
func SetPort(s string) func(SAMTunnel) error {
	return func(c SAMTunnel) error {
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

//SetEncrypt tells the outproxy.SetHttp to use an encrypted leaseset
func SetEncrypt(b bool) func(SAMTunnel) error {
	return func(c SAMTunnel) error {
		c.Config().EncryptLeaseSet = b //"false"
		return nil
	}
}

//SetLeaseSetKey sets key to use with the encrypted leaseset
func SetLeaseSetKey(s string) func(SAMTunnel) error {
	return func(c SAMTunnel) error {
		c.Config().LeaseSetKey = s
		return nil
	}
}

//SetLeaseSetPrivateKey sets the private key to use with the encrypted leaseset
func SetLeaseSetPrivateKey(s string) func(SAMTunnel) error {
	return func(c SAMTunnel) error {
		c.Config().LeaseSetPrivateKey = s
		return nil
	}
}

//SetLeaseSetPrivateSigningKey sets the private signing key to use with the encrypted leaseset
func SetLeaseSetPrivateSigningKey(s string) func(SAMTunnel) error {
	return func(c SAMTunnel) error {
		c.Config().LeaseSetPrivateSigningKey = s
		return nil
	}
}

//SetInLength sets the number of hops inbound
func SetInLength(u int) func(SAMTunnel) error {
	return func(c SAMTunnel) error {
		if u < 7 && u >= 0 {
			c.Config().InLength = u // strconv.Itoa(u)
			return nil
		}
		return fmt.Errorf("Invalid inbound tunnel length")
	}
}

//SetOutLength sets the number of hops outbound
func SetOutLength(u int) func(SAMTunnel) error {
	return func(c SAMTunnel) error {
		if u < 7 && u >= 0 {
			c.Config().OutLength = u // strconv.Itoa(u)
			return nil
		}
		return fmt.Errorf("Invalid outbound tunnel length")
	}
}

//SetName sets the host of the Conf's SAM bridge
func SetName(s string) func(SAMTunnel) error {
	return func(c SAMTunnel) error {
		c.Config().TunName = s
		return nil
	}
}

//SetSaveFile tells the application to save the tunnel's keys long-term
func SetSaveFile(b bool) func(SAMTunnel) error {
	return func(c SAMTunnel) error {
		c.Config().SaveFile = b
		return nil
	}
}

//SetPassword sets the host of the Conf's SAM bridge
func SetPassword(s string) func(SAMTunnel) error {
	return func(c SAMTunnel) error {
		c.Config().Password = s
		return nil
	}
}

//SetInQuantity sets the inbound tunnel quantity
func SetInQuantity(u int) func(SAMTunnel) error {
	return func(c SAMTunnel) error {
		if u <= 16 && u > 0 {
			c.Config().InQuantity = u //strconv.Itoa(u)
			return nil
		}
		return fmt.Errorf("Invalid inbound tunnel quantity")
	}
}

//SetOutQuantity sets the outbound tunnel quantity
func SetOutQuantity(u int) func(SAMTunnel) error {
	return func(c SAMTunnel) error {
		if u <= 16 && u > 0 {
			c.Config().OutQuantity = u // strconv.Itoa(u)
			return nil
		}
		return fmt.Errorf("Invalid outbound tunnel quantity")
	}
}

//SetReduceIdle tells the connection to reduce it's tunnels during extended idle time.
func SetReduceIdle(b bool) func(SAMTunnel) error {
	return func(c SAMTunnel) error {
		c.Config().ReduceIdle = b // "false"
		return nil
	}
}

//SetReduceIdleTime sets the time to wait before reducing tunnels to idle levels
func SetReduceIdleTime(u int) func(SAMTunnel) error {
	return func(c SAMTunnel) error {
		c.Config().ReduceIdleTime = 300000
		if u >= 6 {
			c.Config().ReduceIdleTime = (u * 60) * 1000 // strconv.Itoa((u * 60) * 1000)
			return nil
		}
		return fmt.Errorf("Invalid reduce idle timeout(Measured in minutes) %v", u)
	}
}

//SetReduceIdleTimeMs sets the time to wait before reducing tunnels to idle levels in milliseconds
func SetReduceIdleTimeMs(u int) func(SAMTunnel) error {
	return func(c SAMTunnel) error {
		c.Config().ReduceIdleTime = 300000
		if u >= 300000 {
			c.Config().ReduceIdleTime = u // strconv.Itoa(u)
			return nil
		}
		return fmt.Errorf("Invalid reduce idle timeout(Measured in milliseconds) %v", u)
	}
}

//SetReduceIdleQuantity sets minimum number of tunnels to reduce to during idle time
func SetReduceIdleQuantity(u int) func(SAMTunnel) error {
	return func(c SAMTunnel) error {
		if u < 5 {
			c.Config().ReduceIdleQuantity = u // strconv.Itoa(u)
			return nil
		}
		return fmt.Errorf("Invalid reduce tunnel quantity")
	}
}

//SetCloseIdle tells the connection to close it's tunnels during extended idle time.
func SetCloseIdle(b bool) func(SAMTunnel) error {
	return func(c SAMTunnel) error {
		c.Config().CloseIdle = b // "false"
		return nil
	}
}

//SetCloseIdleTime sets the time to wait before closing tunnels to idle levels
func SetCloseIdleTime(u int) func(SAMTunnel) error {
	return func(c SAMTunnel) error {
		c.Config().CloseIdleTime = 300000
		if u >= 6 {
			c.Config().CloseIdleTime = (u * 60) * 1000 // strconv.Itoa((u * 60) * 1000)
			return nil
		}
		return fmt.Errorf("Invalid close idle timeout(Measured in minutes) %v", u)
	}
}

//SetCloseIdleTimeMs sets the time to wait before closing tunnels to idle levels in milliseconds
func SetCloseIdleTimeMs(u int) func(SAMTunnel) error {
	return func(c SAMTunnel) error {
		c.Config().CloseIdleTime = 300000
		if u >= 300000 {
			c.Config().CloseIdleTime = u //strconv.Itoa(u)
			return nil
		}
		return fmt.Errorf("Invalid close idle timeout(Measured in milliseconds) %v", u)
	}
}

//SetMessageReliability sets the host of the Conf's SAM bridge
func SetMessageReliability(s string) func(SAMTunnel) error {
	return func(c SAMTunnel) error {
		c.Config().MessageReliability = s
		return nil
	}
}

//SetSAMHost sets the host of the Conf's SAM bridge
func SetSAMHost(s string) func(SAMTunnel) error {
	return func(c SAMTunnel) error {
		c.Config().SamHost = s
		return nil
	}
}

//SetSAMPort sets the port of the Conf's SAM bridge using a string
func SetSAMPort(s string) func(SAMTunnel) error {
	return func(c SAMTunnel) error {
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

//SetSigType sets the type of the forwarder server
func SetSigType(s string) func(SAMTunnel) error {
	return func(c SAMTunnel) error {
		switch s {
		case "DSA_SHA1":
			c.Config().SigType = "DSA_SHA1"
		case "ECDSA_SHA256_P256":
			c.Config().SigType = "ECDSA_SHA256_P256"
		case "ECDSA_SHA384_P384":
			c.Config().SigType = "ECDSA_SHA384_P384"
		case "ECDSA_SHA512_P521":
			c.Config().SigType = "ECDSA_SHA512_P521"
		case "EdDSA_SHA512_Ed25519":
			c.Config().SigType = "EdDSA_SHA512_Ed25519"
		default:
			c.Config().SigType = "EdDSA_SHA512_Ed25519"
		}
		return nil
	}
}

//SetType sets the type of the forwarder server
func SetType(s string) func(SAMTunnel) error {
	return func(c SAMTunnel) error {
		switch c.Config().Type {
		case "server":
			c.Config().Type = s
		case "http":
			c.Config().Type = s
		case "client":
			c.Config().Type = s
		case "httpclient":
			c.Config().Type = s
		case "browserclient":
			c.Config().Type = s
		case "udpserver":
			c.Config().Type = s
		case "udpclient":
			c.Config().Type = s
		case "vpnserver":
			c.Config().Type = s
		case "vpnclient":
			c.Config().Type = s
		case "kcpclient":
			c.Config().Type = s
		case "kcpserver":
			c.Config().Type = s
		default:
			c.Config().Type = "browserclient"
		}
		return nil
	}
}

//SetUserName sets username for authentication purposes
func SetUserName(s string) func(SAMTunnel) error {
	return func(c SAMTunnel) error {
		c.Config().UserName = s
		return nil
	}
}

//SetInVariance sets the variance of a number of hops inbound
func SetInVariance(i int) func(SAMTunnel) error {
	return func(c SAMTunnel) error {
		if i < 7 && i > -7 {
			c.Config().InVariance = i //strconv.Itoa(i)
			return nil
		}
		return fmt.Errorf("Invalid inbound tunnel length")
	}
}

//SetOutVariance sets the variance of a number of hops outbound
func SetOutVariance(i int) func(SAMTunnel) error {
	return func(c SAMTunnel) error {
		if i < 7 && i > -7 {
			c.Config().OutVariance = i //strconv.Itoa(i)
			return nil
		}
		return fmt.Errorf("Invalid outbound tunnel variance")
	}
}

//SetAllowZeroIn tells the tunnel to accept zero-hop peers
func SetAllowZeroIn(b bool) func(SAMTunnel) error {
	return func(c SAMTunnel) error {
		c.Config().InAllowZeroHop = b // "false"
		return nil
	}
}

//SetAllowZeroOut tells the tunnel to accept zero-hop peers
func SetAllowZeroOut(b bool) func(SAMTunnel) error {
	return func(c SAMTunnel) error {
		c.Config().OutAllowZeroHop = b // "false"
		return nil
	}
}

//SetInBackups sets the inbound tunnel backups
func SetInBackups(u int) func(SAMTunnel) error {
	return func(c SAMTunnel) error {
		if u < 6 && u >= 0 {
			c.Config().InBackupQuantity = u // strconv.Itoa(u)
			return nil
		}
		return fmt.Errorf("Invalid inbound tunnel backup quantity")
	}
}

//SetOutBackups sets the inbound tunnel backups
func SetOutBackups(u int) func(SAMTunnel) error {
	return func(c SAMTunnel) error {
		if u < 6 && u >= 0 {
			c.Config().OutBackupQuantity = u // strconv.Itoa(u)
			return nil
		}
		return fmt.Errorf("Invalid outbound tunnel backup quantity")
	}
}
