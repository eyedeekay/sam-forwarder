package i2ptunhelper

import (
	"github.com/eyedeekay/outproxy"
	"github.com/eyedeekay/sam-forwarder/config"
)

// NewOutProxyFromConf generates a SAMforwarder from *i2ptunconf.Conf
func NewOutProxyFromConf(config *i2ptunconf.Conf) (*outproxy.OutProxy, error) {
	if config != nil {
		return outproxy.NewOutProxyFromOptions(
			outproxy.SetType(config.Type),
			outproxy.SetSaveFile(config.SaveFile),
			outproxy.SetFilePath(config.SaveDirectory),
			outproxy.SetHost(config.TargetHost),
			outproxy.SetPort(config.TargetPort),
			outproxy.SetSAMHost(config.SamHost),
			outproxy.SetSAMPort(config.SamPort),
			outproxy.SetSigType(config.SigType),
			outproxy.SetName(config.TunName),
			outproxy.SetInLength(config.InLength),
			outproxy.SetOutLength(config.OutLength),
			outproxy.SetInVariance(config.InVariance),
			outproxy.SetOutVariance(config.OutVariance),
			outproxy.SetInQuantity(config.InQuantity),
			outproxy.SetOutQuantity(config.OutQuantity),
			outproxy.SetInBackups(config.InBackupQuantity),
			outproxy.SetOutBackups(config.OutBackupQuantity),
			outproxy.SetEncrypt(config.EncryptLeaseSet),
			outproxy.SetLeaseSetKey(config.LeaseSetKey),
			outproxy.SetLeaseSetPrivateKey(config.LeaseSetPrivateKey),
			outproxy.SetLeaseSetPrivateSigningKey(config.LeaseSetPrivateSigningKey),
			outproxy.SetAllowZeroIn(config.InAllowZeroHop),
			outproxy.SetAllowZeroOut(config.OutAllowZeroHop),
			outproxy.SetFastRecieve(config.FastRecieve),
			outproxy.SetCompress(config.UseCompression),
			outproxy.SetReduceIdle(config.ReduceIdle),
			outproxy.SetReduceIdleTimeMs(config.ReduceIdleTime),
			outproxy.SetReduceIdleQuantity(config.ReduceIdleQuantity),
			outproxy.SetCloseIdle(config.CloseIdle),
			outproxy.SetCloseIdleTimeMs(config.CloseIdleTime),
			outproxy.SetAccessListType(config.AccessListType),
			outproxy.SetAccessList(config.AccessList),
			outproxy.SetMessageReliability(config.MessageReliability),
			outproxy.SetKeyFile(config.KeyFilePath),
			//outproxy.SetTargetForPort443(config.TargetForPort443),
		)
	}
	return nil, nil
}

// NewOutProxyFromConfig generates a new OutProxy from a config file
func NewOutProxyFromConfig(iniFile, SamHost, SamPort string, label ...string) (*outproxy.HttpOutProxy, error) {
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
		return NewHttpOutProxyFromConf(config)
	}
	return nil, nil
}

// NewOutProxyFromConf generates a SAMforwarder from *i2ptunconf.Conf
func NewHttpOutProxyFromConf(config *i2ptunconf.Conf) (*outproxy.HttpOutProxy, error) {
	if config != nil {
		return outproxy.NewHttpOutProxydFromOptions(
			outproxy.SetHttpType(config.Type),
			outproxy.SetHttpSaveFile(config.SaveFile),
			outproxy.SetHttpFilePath(config.SaveDirectory),
			outproxy.SetHttpHost(config.TargetHost),
			outproxy.SetHttpPort(config.TargetPort),
			outproxy.SetHttpSAMHost(config.SamHost),
			outproxy.SetHttpSAMPort(config.SamPort),
			outproxy.SetHttpSigType(config.SigType),
			outproxy.SetHttpName(config.TunName),
			outproxy.SetHttpInLength(config.InLength),
			outproxy.SetHttpOutLength(config.OutLength),
			outproxy.SetHttpInVariance(config.InVariance),
			outproxy.SetHttpOutVariance(config.OutVariance),
			outproxy.SetHttpInQuantity(config.InQuantity),
			outproxy.SetHttpOutQuantity(config.OutQuantity),
			outproxy.SetHttpInBackups(config.InBackupQuantity),
			outproxy.SetHttpOutBackups(config.OutBackupQuantity),
			outproxy.SetHttpEncrypt(config.EncryptLeaseSet),
			outproxy.SetHttpLeaseSetKey(config.LeaseSetKey),
			outproxy.SetHttpLeaseSetPrivateKey(config.LeaseSetPrivateKey),
			outproxy.SetHttpLeaseSetPrivateSigningKey(config.LeaseSetPrivateSigningKey),
			outproxy.SetHttpAllowZeroIn(config.InAllowZeroHop),
			outproxy.SetHttpAllowZeroOut(config.OutAllowZeroHop),
			outproxy.SetHttpFastRecieve(config.FastRecieve),
			outproxy.SetHttpCompress(config.UseCompression),
			outproxy.SetHttpReduceIdle(config.ReduceIdle),
			outproxy.SetHttpReduceIdleTimeMs(config.ReduceIdleTime),
			outproxy.SetHttpReduceIdleQuantity(config.ReduceIdleQuantity),
			outproxy.SetHttpCloseIdle(config.CloseIdle),
			outproxy.SetHttpCloseIdleTimeMs(config.CloseIdleTime),
			outproxy.SetHttpAccessListType(config.AccessListType),
			outproxy.SetHttpAccessList(config.AccessList),
			outproxy.SetHttpMessageReliability(config.MessageReliability),
			outproxy.SetHttpKeyFile(config.KeyFilePath),
			//outproxy.SetHttpTargetForPort443(config.TargetForPort443),
		)
	}
	return nil, nil
}

// NewOutProxyFromConfig generates a new OutProxy from a config file
func NewHttpOutProxyFromConfig(iniFile, SamHost, SamPort string, label ...string) (*outproxy.OutProxy, error) {
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
		return NewOutProxyFromConf(config)
	}
	return nil, nil
}
