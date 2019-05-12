package i2ptunconf

import (
	"github.com/eyedeekay/httptunnel"
	"github.com/eyedeekay/sam-forwarder/tcp"
	"github.com/eyedeekay/sam-forwarder/udp"
)

func NewSAMHTTPClientFromConf(config *Conf) (*i2phttpproxy.SAMHTTPProxy, error) {
	if config != nil {
		return i2phttpproxy.NewHttpProxy(
			i2phttpproxy.SetKeysPath(config.KeyFilePath),
			i2phttpproxy.SetHost(config.SamHost),
			i2phttpproxy.SetPort(config.SamPort),
			i2phttpproxy.SetProxyAddr(config.TargetHost+":"+config.TargetPort),
			//i2phttpproxy.SetControlAddr(),
			i2phttpproxy.SetInLength(uint(config.InLength)),
			i2phttpproxy.SetOutLength(uint(config.OutLength)),
			i2phttpproxy.SetInQuantity(uint(config.InQuantity)),
			i2phttpproxy.SetOutQuantity(uint(config.OutQuantity)),
			i2phttpproxy.SetInBackups(uint(config.InBackupQuantity)),
			i2phttpproxy.SetOutBackups(uint(config.OutBackupQuantity)),
			i2phttpproxy.SetInVariance(config.InVariance),
			i2phttpproxy.SetOutVariance(config.OutVariance),
			i2phttpproxy.SetUnpublished(config.Client),
			i2phttpproxy.SetReduceIdle(config.ReduceIdle),
			i2phttpproxy.SetCompression(config.UseCompression),
			i2phttpproxy.SetReduceIdleTime(uint(config.ReduceIdleTime)),
			i2phttpproxy.SetReduceIdleQuantity(uint(config.ReduceIdleQuantity)),
			i2phttpproxy.SetCloseIdle(config.CloseIdle),
			i2phttpproxy.SetCloseIdleTime(uint(config.CloseIdleTime)),
		)
	}
	return nil, nil
}

// NewSAMClientForwarderFromConfig generates a new SAMForwarder from a config file
func NewSAMHTTPClientFromConfig(iniFile, SamHost, SamPort string, label ...string) (*i2phttpproxy.SAMHTTPProxy, error) {
	if iniFile != "none" {
		config, err := NewI2PTunConf(iniFile, label...)
		if err != nil {
			return nil, err
		}
		if SamHost != "" && SamHost != "127.0.0.1" && SamHost != "localhost" {
			config.SamHost = config.GetSAMHost(SamHost, config.SamHost)
		}
		if SamPort != "" && SamPort != "7656" {
			config.SamPort = config.GetSAMPort(SamPort, config.SamPort)
		}
		return NewSAMHTTPClientFromConf(config)
	}
	return nil, nil
}

// NewSAMClientForwarderFromConf generates a SAMforwarder from *i2ptunconf.Conf
func NewSAMClientForwarderFromConf(config *Conf) (*samforwarder.SAMClientForwarder, error) {
	if config != nil {
		return samforwarder.NewSAMClientForwarderFromOptions(
			samforwarder.SetClientSaveFile(config.SaveFile),
			samforwarder.SetClientFilePath(config.SaveDirectory),
			samforwarder.SetClientHost(config.TargetHost),
			samforwarder.SetClientPort(config.TargetPort),
			samforwarder.SetClientSAMHost(config.SamHost),
			samforwarder.SetClientSAMPort(config.SamPort),
			samforwarder.SetClientSigType(config.SigType),
			samforwarder.SetClientName(config.TunName),
			samforwarder.SetClientInLength(config.InLength),
			samforwarder.SetClientOutLength(config.OutLength),
			samforwarder.SetClientInVariance(config.InVariance),
			samforwarder.SetClientOutVariance(config.OutVariance),
			samforwarder.SetClientInQuantity(config.InQuantity),
			samforwarder.SetClientOutQuantity(config.OutQuantity),
			samforwarder.SetClientInBackups(config.InBackupQuantity),
			samforwarder.SetClientOutBackups(config.OutBackupQuantity),
			samforwarder.SetClientEncrypt(config.EncryptLeaseSet),
			samforwarder.SetClientLeaseSetKey(config.LeaseSetKey),
			samforwarder.SetClientLeaseSetPrivateKey(config.LeaseSetPrivateKey),
			samforwarder.SetClientLeaseSetPrivateSigningKey(config.LeaseSetPrivateSigningKey),
			samforwarder.SetClientAllowZeroIn(config.InAllowZeroHop),
			samforwarder.SetClientAllowZeroOut(config.OutAllowZeroHop),
			samforwarder.SetClientFastRecieve(config.FastRecieve),
			samforwarder.SetClientCompress(config.UseCompression),
			samforwarder.SetClientReduceIdle(config.ReduceIdle),
			samforwarder.SetClientReduceIdleTimeMs(config.ReduceIdleTime),
			samforwarder.SetClientReduceIdleQuantity(config.ReduceIdleQuantity),
			samforwarder.SetClientCloseIdle(config.CloseIdle),
			samforwarder.SetClientCloseIdleTimeMs(config.CloseIdleTime),
			samforwarder.SetClientAccessListType(config.AccessListType),
			samforwarder.SetClientAccessList(config.AccessList),
			samforwarder.SetClientMessageReliability(config.MessageReliability),
			samforwarder.SetClientPassword(config.KeyFilePath),
			samforwarder.SetClientDestination(config.ClientDest),
		)
	}
	return nil, nil
}

// NewSAMClientForwarderFromConfig generates a new SAMForwarder from a config file
func NewSAMClientForwarderFromConfig(iniFile, SamHost, SamPort string, label ...string) (*samforwarder.SAMClientForwarder, error) {
	if iniFile != "none" {
		config, err := NewI2PTunConf(iniFile, label...)
		if err != nil {
			return nil, err
		}
		if SamHost != "" && SamHost != "127.0.0.1" && SamHost != "localhost" {
			config.SamHost = config.GetSAMHost(SamHost, config.SamHost)
		}
		if SamPort != "" && SamPort != "7656" {
			config.SamPort = config.GetSAMPort(SamPort, config.SamPort)
		}
		return NewSAMClientForwarderFromConf(config)
	}
	return nil, nil
}

// NewSAMSSUClientForwarderFromConf generates a SAMSSUforwarder from *i2ptunconf.Conf
func NewSAMSSUClientForwarderFromConf(config *Conf) (*samforwarderudp.SAMSSUClientForwarder, error) {
	if config != nil {
		return samforwarderudp.NewSAMSSUClientForwarderFromOptions(
			samforwarderudp.SetClientSaveFile(config.SaveFile),
			samforwarderudp.SetClientFilePath(config.SaveDirectory),
			samforwarderudp.SetClientHost(config.TargetHost),
			samforwarderudp.SetClientPort(config.TargetPort),
			samforwarderudp.SetClientSAMHost(config.SamHost),
			samforwarderudp.SetClientSAMPort(config.SamPort),
			samforwarderudp.SetClientSigType(config.SigType),
			samforwarderudp.SetClientName(config.TunName),
			samforwarderudp.SetClientInLength(config.InLength),
			samforwarderudp.SetClientOutLength(config.OutLength),
			samforwarderudp.SetClientInVariance(config.InVariance),
			samforwarderudp.SetClientOutVariance(config.OutVariance),
			samforwarderudp.SetClientInQuantity(config.InQuantity),
			samforwarderudp.SetClientOutQuantity(config.OutQuantity),
			samforwarderudp.SetClientInBackups(config.InBackupQuantity),
			samforwarderudp.SetClientOutBackups(config.OutBackupQuantity),
			samforwarderudp.SetClientEncrypt(config.EncryptLeaseSet),
			samforwarderudp.SetClientLeaseSetKey(config.LeaseSetKey),
			samforwarderudp.SetClientLeaseSetPrivateKey(config.LeaseSetPrivateKey),
			samforwarderudp.SetClientLeaseSetPrivateSigningKey(config.LeaseSetPrivateSigningKey),
			samforwarderudp.SetClientAllowZeroIn(config.InAllowZeroHop),
			samforwarderudp.SetClientAllowZeroOut(config.OutAllowZeroHop),
			samforwarderudp.SetClientFastRecieve(config.FastRecieve),
			samforwarderudp.SetClientCompress(config.UseCompression),
			samforwarderudp.SetClientReduceIdle(config.ReduceIdle),
			samforwarderudp.SetClientReduceIdleTimeMs(config.ReduceIdleTime),
			samforwarderudp.SetClientReduceIdleQuantity(config.ReduceIdleQuantity),
			samforwarderudp.SetClientCloseIdle(config.CloseIdle),
			samforwarderudp.SetClientCloseIdleTimeMs(config.CloseIdleTime),
			samforwarderudp.SetClientAccessListType(config.AccessListType),
			samforwarderudp.SetClientAccessList(config.AccessList),
			samforwarderudp.SetClientMessageReliability(config.MessageReliability),
			samforwarderudp.SetClientPassword(config.KeyFilePath),
			samforwarderudp.SetClientDestination(config.ClientDest),
		)
	}
	return nil, nil
}

func NewSAMSSUClientForwarderFromConfig(iniFile, SamHost, SamPort string, label ...string) (*samforwarderudp.SAMSSUClientForwarder, error) {
	if iniFile != "none" {
		config, err := NewI2PTunConf(iniFile, label...)
		if err != nil {
			return nil, err
		}
		if SamHost != "" && SamHost != "127.0.0.1" && SamHost != "localhost" {
			config.SamHost = config.GetSAMHost(SamHost, config.SamHost)
		}
		if SamPort != "" && SamPort != "7656" {
			config.SamPort = config.GetSAMPort(SamPort, config.SamPort)
		}
		return NewSAMSSUClientForwarderFromConf(config)
	}
	return nil, nil
}
