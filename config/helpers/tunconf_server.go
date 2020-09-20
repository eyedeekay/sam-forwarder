package i2ptunhelper

import (
	"github.com/eyedeekay/eephttpd"
	"github.com/eyedeekay/sam-forwarder/config"
	"github.com/eyedeekay/sam-forwarder/tcp"
	"github.com/eyedeekay/sam-forwarder/udp"
)

// NewSAMForwarderFromConf generates a SAMforwarder from *i2ptunconf.Conf
func NewSAMForwarderFromConf(config *i2ptunconf.Conf) (*samforwarder.SAMForwarder, error) {
	if config != nil {
		return samforwarder.NewSAMForwarderFromOptions(
			samforwarder.SetType(config.Type),
			samforwarder.SetSaveFile(config.SaveFile),
			samforwarder.SetFilePath(config.SaveDirectory),
			samforwarder.SetHost(config.TargetHost),
			samforwarder.SetPort(config.TargetPort),
			samforwarder.SetSAMHost(config.SamHost),
			samforwarder.SetSAMPort(config.SamPort),
			samforwarder.SetSigType(config.SigType),
			samforwarder.SetName(config.TunName),
			samforwarder.SetInLength(config.InLength),
			samforwarder.SetOutLength(config.OutLength),
			samforwarder.SetInVariance(config.InVariance),
			samforwarder.SetOutVariance(config.OutVariance),
			samforwarder.SetInQuantity(config.InQuantity),
			samforwarder.SetOutQuantity(config.OutQuantity),
			samforwarder.SetInBackups(config.InBackupQuantity),
			samforwarder.SetOutBackups(config.OutBackupQuantity),
			samforwarder.SetEncrypt(config.EncryptLeaseSet),
			samforwarder.SetLeaseSetKey(config.LeaseSetKey),
			samforwarder.SetLeaseSetPrivateKey(config.LeaseSetPrivateKey),
			samforwarder.SetLeaseSetPrivateSigningKey(config.LeaseSetPrivateSigningKey),
			samforwarder.SetAllowZeroIn(config.InAllowZeroHop),
			samforwarder.SetAllowZeroOut(config.OutAllowZeroHop),
			samforwarder.SetFastRecieve(config.FastRecieve),
			samforwarder.SetCompress(config.UseCompression),
			samforwarder.SetReduceIdle(config.ReduceIdle),
			samforwarder.SetReduceIdleTimeMs(config.ReduceIdleTime),
			samforwarder.SetReduceIdleQuantity(config.ReduceIdleQuantity),
			samforwarder.SetCloseIdle(config.CloseIdle),
			samforwarder.SetCloseIdleTimeMs(config.CloseIdleTime),
			samforwarder.SetAccessListType(config.AccessListType),
			samforwarder.SetAccessList(config.AccessList),
			samforwarder.SetMessageReliability(config.MessageReliability),
			samforwarder.SetKeyFile(config.KeyFilePath),
			//samforwarder.SetTargetForPort443(config.TargetForPort443),
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
			samforwarderudp.SetSaveFile(config.SaveFile),
			samforwarderudp.SetFilePath(config.SaveDirectory),
			samforwarderudp.SetHost(config.TargetHost),
			samforwarderudp.SetPort(config.TargetPort),
			samforwarderudp.SetSAMHost(config.SamHost),
			samforwarderudp.SetSAMPort(config.SamPort),
			samforwarderudp.SetSigType(config.SigType),
			samforwarderudp.SetName(config.TunName),
			samforwarderudp.SetInLength(config.InLength),
			samforwarderudp.SetOutLength(config.OutLength),
			samforwarderudp.SetInVariance(config.InVariance),
			samforwarderudp.SetOutVariance(config.OutVariance),
			samforwarderudp.SetInQuantity(config.InQuantity),
			samforwarderudp.SetOutQuantity(config.OutQuantity),
			samforwarderudp.SetInBackups(config.InBackupQuantity),
			samforwarderudp.SetOutBackups(config.OutBackupQuantity),
			samforwarderudp.SetEncrypt(config.EncryptLeaseSet),
			samforwarderudp.SetLeaseSetKey(config.LeaseSetKey),
			samforwarderudp.SetLeaseSetPrivateKey(config.LeaseSetPrivateKey),
			samforwarderudp.SetLeaseSetPrivateSigningKey(config.LeaseSetPrivateSigningKey),
			samforwarderudp.SetAllowZeroIn(config.InAllowZeroHop),
			samforwarderudp.SetAllowZeroOut(config.OutAllowZeroHop),
			samforwarderudp.SetFastRecieve(config.FastRecieve),
			samforwarderudp.SetCompress(config.UseCompression),
			samforwarderudp.SetReduceIdle(config.ReduceIdle),
			samforwarderudp.SetReduceIdleTimeMs(config.ReduceIdleTime),
			samforwarderudp.SetReduceIdleQuantity(config.ReduceIdleQuantity),
			samforwarderudp.SetCloseIdle(config.CloseIdle),
			samforwarderudp.SetCloseIdleTimeMs(config.CloseIdleTime),
			samforwarderudp.SetAccessListType(config.AccessListType),
			samforwarderudp.SetAccessList(config.AccessList),
			samforwarderudp.SetMessageReliability(config.MessageReliability),
			samforwarderudp.SetKeyFile(config.KeyFilePath),
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
