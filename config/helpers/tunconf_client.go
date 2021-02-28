package i2ptunhelper

import (
	"github.com/eyedeekay/httptunnel"
	"github.com/eyedeekay/httptunnel/multiproxy"
	"github.com/eyedeekay/sam-forwarder/config"
	"github.com/eyedeekay/sam-forwarder/options"
	"github.com/eyedeekay/sam-forwarder/tcp"
	"github.com/eyedeekay/sam-forwarder/udp"
)

func NewSAMHTTPClientFromConf(config *i2ptunconf.Conf) (*i2phttpproxy.SAMHTTPProxy, error) {
	if config != nil {
		return i2phttpproxy.NewHttpProxy(
			i2phttpproxy.SetName(config.TunName),
			i2phttpproxy.SetKeysPath(config.KeyFilePath),
			i2phttpproxy.SetHost(config.SamHost),
			i2phttpproxy.SetPort(config.SamPort),
			i2phttpproxy.SetProxyAddr(config.TargetHost+":"+config.TargetPort),
			i2phttpproxy.SetControlHost(config.ControlHost),
			i2phttpproxy.SetControlPort(config.ControlPort),
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
		config, err := i2ptunconf.NewI2PTunConf(iniFile, label...)
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

func NewSAMBrowserClientFromConf(config *i2ptunconf.Conf) (*i2pbrowserproxy.SAMMultiProxy, error) {
	if config != nil {
		return i2pbrowserproxy.NewHttpProxy(
			i2pbrowserproxy.SetName(config.TunName),
			i2pbrowserproxy.SetKeysPath(config.KeyFilePath),
			i2pbrowserproxy.SetHost(config.SamHost),
			i2pbrowserproxy.SetPort(config.SamPort),
			i2pbrowserproxy.SetProxyAddr(config.TargetHost+":"+config.TargetPort),
			i2pbrowserproxy.SetControlHost(config.ControlHost),
			i2pbrowserproxy.SetControlPort(config.ControlPort),
			i2pbrowserproxy.SetInLength(uint(config.InLength)),
			i2pbrowserproxy.SetOutLength(uint(config.OutLength)),
			i2pbrowserproxy.SetInQuantity(uint(config.InQuantity)),
			i2pbrowserproxy.SetOutQuantity(uint(config.OutQuantity)),
			i2pbrowserproxy.SetInBackups(uint(config.InBackupQuantity)),
			i2pbrowserproxy.SetOutBackups(uint(config.OutBackupQuantity)),
			i2pbrowserproxy.SetInVariance(config.InVariance),
			i2pbrowserproxy.SetOutVariance(config.OutVariance),
			i2pbrowserproxy.SetUnpublished(config.Client),
			i2pbrowserproxy.SetReduceIdle(config.ReduceIdle),
			i2pbrowserproxy.SetCompression(config.UseCompression),
			i2pbrowserproxy.SetReduceIdleTime(uint(config.ReduceIdleTime)),
			i2pbrowserproxy.SetReduceIdleQuantity(uint(config.ReduceIdleQuantity)),
			//i2pbrowserproxy.SetCloseIdle(config.CloseIdle),
			//i2pbrowserproxy.SetCloseIdleTime(uint(config.CloseIdleTime)),
		)
	}
	return nil, nil
}

func NewSAMBrowserClientFromConfig(iniFile, SamHost, SamPort string, label ...string) (*i2pbrowserproxy.SAMMultiProxy, error) {
	if iniFile != "none" {
		config, err := i2ptunconf.NewI2PTunConf(iniFile, label...)
		if err != nil {
			return nil, err
		}
		if SamHost != "" && SamHost != "127.0.0.1" && SamHost != "localhost" {
			config.SamHost = config.GetSAMHost(SamHost, config.SamHost)
		}
		if SamPort != "" && SamPort != "7656" {
			config.SamPort = config.GetSAMPort(SamPort, config.SamPort)
		}
		return NewSAMBrowserClientFromConf(config)
	}
	return nil, nil
}

// NewSAMClientForwarderFromConf generates a SAMforwarder from *i2ptunconf.Conf
func NewSAMClientForwarderFromConf(config *i2ptunconf.Conf) (*samforwarder.SAMClientForwarder, error) {
	if config != nil {
		return samforwarder.NewSAMClientForwarderFromOptions(
			samoptions.SetSaveFile(config.SaveFile),
			samoptions.SetFilePath(config.SaveDirectory),
			samoptions.SetHost(config.TargetHost),
			samoptions.SetPort(config.TargetPort),
			samoptions.SetSAMHost(config.SamHost),
			samoptions.SetSAMPort(config.SamPort),
			samoptions.SetSigType(config.SigType),
			samoptions.SetName(config.TunName),
			samoptions.SetInLength(config.InLength),
			samoptions.SetOutLength(config.OutLength),
			samoptions.SetInVariance(config.InVariance),
			samoptions.SetOutVariance(config.OutVariance),
			samoptions.SetInQuantity(config.InQuantity),
			samoptions.SetOutQuantity(config.OutQuantity),
			samoptions.SetInBackups(config.InBackupQuantity),
			samoptions.SetOutBackups(config.OutBackupQuantity),
			samoptions.SetEncrypt(config.EncryptLeaseSet),
			samoptions.SetLeaseSetKey(config.LeaseSetKey),
			samoptions.SetLeaseSetPrivateKey(config.LeaseSetPrivateKey),
			samoptions.SetLeaseSetPrivateSigningKey(config.LeaseSetPrivateSigningKey),
			samoptions.SetAllowZeroIn(config.InAllowZeroHop),
			samoptions.SetAllowZeroOut(config.OutAllowZeroHop),
			samoptions.SetFastRecieve(config.FastRecieve),
			samoptions.SetCompress(config.UseCompression),
			samoptions.SetReduceIdle(config.ReduceIdle),
			samoptions.SetReduceIdleTimeMs(config.ReduceIdleTime),
			samoptions.SetReduceIdleQuantity(config.ReduceIdleQuantity),
			samoptions.SetCloseIdle(config.CloseIdle),
			samoptions.SetCloseIdleTimeMs(config.CloseIdleTime),
			samoptions.SetAccessListType(config.AccessListType),
			samoptions.SetAccessList(config.AccessList),
			samoptions.SetMessageReliability(config.MessageReliability),
			samoptions.SetPassword(config.KeyFilePath),
			samoptions.SetDestination(config.ClientDest),
		)
	}
	return nil, nil
}

// NewSAMClientForwarderFromConfig generates a new SAMForwarder from a config file
func NewSAMClientForwarderFromConfig(iniFile, SamHost, SamPort string, label ...string) (*samforwarder.SAMClientForwarder, error) {
	if iniFile != "none" {
		config, err := i2ptunconf.NewI2PTunConf(iniFile, label...)
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

// NewSAMDGClientForwarderFromConf generates a SAMSSUforwarder from *i2ptunconf.Conf
func NewSAMDGClientForwarderFromConf(config *i2ptunconf.Conf) (*samforwarderudp.SAMDGClientForwarder, error) {
	if config != nil {
		return samforwarderudp.NewSAMDGClientForwarderFromOptions(
			samoptions.SetSaveFile(config.SaveFile),
			samoptions.SetFilePath(config.SaveDirectory),
			samoptions.SetHost(config.TargetHost),
			samoptions.SetPort(config.TargetPort),
			samoptions.SetSAMHost(config.SamHost),
			samoptions.SetSAMPort(config.SamPort),
			samoptions.SetSigType(config.SigType),
			samoptions.SetName(config.TunName),
			samoptions.SetInLength(config.InLength),
			samoptions.SetOutLength(config.OutLength),
			samoptions.SetInVariance(config.InVariance),
			samoptions.SetOutVariance(config.OutVariance),
			samoptions.SetInQuantity(config.InQuantity),
			samoptions.SetOutQuantity(config.OutQuantity),
			samoptions.SetInBackups(config.InBackupQuantity),
			samoptions.SetOutBackups(config.OutBackupQuantity),
			samoptions.SetEncrypt(config.EncryptLeaseSet),
			samoptions.SetLeaseSetKey(config.LeaseSetKey),
			samoptions.SetLeaseSetPrivateKey(config.LeaseSetPrivateKey),
			samoptions.SetLeaseSetPrivateSigningKey(config.LeaseSetPrivateSigningKey),
			samoptions.SetAllowZeroIn(config.InAllowZeroHop),
			samoptions.SetAllowZeroOut(config.OutAllowZeroHop),
			samoptions.SetFastRecieve(config.FastRecieve),
			samoptions.SetCompress(config.UseCompression),
			samoptions.SetReduceIdle(config.ReduceIdle),
			samoptions.SetReduceIdleTimeMs(config.ReduceIdleTime),
			samoptions.SetReduceIdleQuantity(config.ReduceIdleQuantity),
			samoptions.SetCloseIdle(config.CloseIdle),
			samoptions.SetCloseIdleTimeMs(config.CloseIdleTime),
			samoptions.SetAccessListType(config.AccessListType),
			samoptions.SetAccessList(config.AccessList),
			samoptions.SetMessageReliability(config.MessageReliability),
			samoptions.SetPassword(config.KeyFilePath),
			samoptions.SetDestination(config.ClientDest),
		)
	}
	return nil, nil
}

func NewSAMDGClientForwarderFromConfig(iniFile, SamHost, SamPort string, label ...string) (*samforwarderudp.SAMDGClientForwarder, error) {
	if iniFile != "none" {
		config, err := i2ptunconf.NewI2PTunConf(iniFile, label...)
		if err != nil {
			return nil, err
		}
		if SamHost != "" && SamHost != "127.0.0.1" && SamHost != "localhost" {
			config.SamHost = config.GetSAMHost(SamHost, config.SamHost)
		}
		if SamPort != "" && SamPort != "7656" {
			config.SamPort = config.GetSAMPort(SamPort, config.SamPort)
		}
		return NewSAMDGClientForwarderFromConf(config)
	}
	return nil, nil
}
