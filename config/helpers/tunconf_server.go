package i2ptunhelper

import (
	"github.com/eyedeekay/eephttpd"
	"github.com/eyedeekay/sam-forwarder/config"
	"github.com/eyedeekay/sam-forwarder/options"
	"github.com/eyedeekay/sam-forwarder/tcp"
	"github.com/eyedeekay/sam-forwarder/udp"
)

// NewSAMForwarderFromConf generates a SAMforwarder from *i2ptunconf.Conf
func NewSAMForwarderFromConf(config *i2ptunconf.Conf) (*samforwarder.SAMForwarder, error) {
	if config != nil {
		return samforwarder.NewSAMForwarderFromOptions(
			samoptions.SetType(config.Type),
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
			samoptions.SetKeyFile(config.KeyFilePath),
			//samoptions.SetTargetForPort443(config.TargetForPort443),
		)
	}
	return nil, nil
}

// NewSAMForwarderFromConfig generates a new SAMForwarder from a config file
func NewSAMForwarderFromConfig(iniFile, SamHost, SamPort string, label ...string) (*samforwarder.SAMForwarder, error) {
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
		return NewSAMForwarderFromConf(config)
	}
	return nil, nil
}

// NewSAMDGForwarderFromConf generates a SAMSSUforwarder from *i2ptunconf.Conf
func NewSAMDGForwarderFromConf(config *i2ptunconf.Conf) (*samforwarderudp.SAMDGForwarder, error) {
	if config != nil {
		return samforwarderudp.NewSAMDGForwarderFromOptions(
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
			samoptions.SetKeyFile(config.KeyFilePath),
		)
	}
	return nil, nil
}

// NewSAMDGForwarderFromConfig generates a new SAMDGForwarder from a config file
func NewSAMDGForwarderFromConfig(iniFile, SamHost, SamPort string, label ...string) (*samforwarderudp.SAMDGForwarder, error) {
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
		return NewSAMDGForwarderFromConf(config)
	}
	return nil, nil
}

// NewEepHttpdFromConf generates a SAMforwarder from *i2ptunconf.Conf
func NewEepHttpdFromConf(config *i2ptunconf.Conf) (*eephttpd.EepHttpd, error) {
	if config != nil {
		return eephttpd.NewEepHttpdFromOptions(
			eephttpd.SetType(config.Type),
			eephttpd.SetSaveFile(config.SaveFile),
			eephttpd.SetFilePath(config.SaveDirectory),
			eephttpd.SetHost(config.TargetHost),
			eephttpd.SetPort(config.TargetPort),
			eephttpd.SetSAMHost(config.SamHost),
			eephttpd.SetSAMPort(config.SamPort),
			eephttpd.SetSigType(config.SigType),
			eephttpd.SetName(config.TunName),
			eephttpd.SetInLength(config.InLength),
			eephttpd.SetOutLength(config.OutLength),
			eephttpd.SetInVariance(config.InVariance),
			eephttpd.SetOutVariance(config.OutVariance),
			eephttpd.SetInQuantity(config.InQuantity),
			eephttpd.SetOutQuantity(config.OutQuantity),
			eephttpd.SetInBackups(config.InBackupQuantity),
			eephttpd.SetOutBackups(config.OutBackupQuantity),
			eephttpd.SetEncrypt(config.EncryptLeaseSet),
			eephttpd.SetLeaseSetKey(config.LeaseSetKey),
			eephttpd.SetLeaseSetPrivateKey(config.LeaseSetPrivateKey),
			eephttpd.SetLeaseSetPrivateSigningKey(config.LeaseSetPrivateSigningKey),
			eephttpd.SetAllowZeroIn(config.InAllowZeroHop),
			eephttpd.SetAllowZeroOut(config.OutAllowZeroHop),
			eephttpd.SetFastRecieve(config.FastRecieve),
			eephttpd.SetCompress(config.UseCompression),
			eephttpd.SetReduceIdle(config.ReduceIdle),
			eephttpd.SetReduceIdleTimeMs(config.ReduceIdleTime),
			eephttpd.SetReduceIdleQuantity(config.ReduceIdleQuantity),
			eephttpd.SetCloseIdle(config.CloseIdle),
			eephttpd.SetCloseIdleTimeMs(config.CloseIdleTime),
			eephttpd.SetAccessListType(config.AccessListType),
			eephttpd.SetAccessList(config.AccessList),
			eephttpd.SetMessageReliability(config.MessageReliability),
			eephttpd.SetKeyFile(config.KeyFilePath),
			eephttpd.SetServeDir(config.ServeDirectory),
			//eephttpd.SetTargetForPort443(config.TargetForPort443),
		)
	}
	return nil, nil
}

// NewEepHttpdFromConfig generates a new EepHttpd from a config file
func NewEepHttpdFromConfig(iniFile, SamHost, SamPort string, label ...string) (*eephttpd.EepHttpd, error) {
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
		return NewEepHttpdFromConf(config)
	}
	return nil, nil
}
