package i2ptunconf

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

import (
	"github.com/eyedeekay/sam-forwarder"
    "github.com/eyedeekay/sam-forwarder/csvpn"
	"github.com/eyedeekay/sam-forwarder/udp"
	"github.com/zieckey/goini"
)

// Conf is a tructure containing an ini config, with some functions to help
// when you use it for in conjunction with command-line flags
type Conf struct {
	Config                    *goini.INI
	FilePath                  string
	KeyFilePath               string
	Labels                    []string
	Client                    bool
	ClientDest                string
	Type                      string
	SaveDirectory             string
	SaveFile                  bool
	TargetHost                string
	TargetPort                string
	SamHost                   string
	SamPort                   string
	TargetForPort443          string
	TunName                   string
	EncryptLeaseSet           bool
	LeaseSetKey               string
	LeaseSetPrivateKey        string
	LeaseSetPrivateSigningKey string
	InAllowZeroHop            bool
	OutAllowZeroHop           bool
	InLength                  int
	OutLength                 int
	InQuantity                int
	OutQuantity               int
	InVariance                int
	OutVariance               int
	InBackupQuantity          int
	OutBackupQuantity         int
	UseCompression            bool
	FastRecieve               bool
	ReduceIdle                bool
	ReduceIdleTime            int
	ReduceIdleQuantity        int
	CloseIdle                 bool
	CloseIdleTime             int
	AccessListType            string
	AccessList                []string
	MessageReliability        string
	exists                    bool
	VPN                       bool
}

// Print returns and prints a formatted list of configured tunnel settings.
func (c *Conf) Print() []string {
	confstring := []string{
		"inbound.length=" + strconv.Itoa(c.InLength),
		"outbound.length=" + strconv.Itoa(c.OutLength),
		"inbound.lengthVariance=" + strconv.Itoa(c.InVariance),
		"outbound.lengthVariance=" + strconv.Itoa(c.OutVariance),
		"inbound.backupQuantity=" + strconv.Itoa(c.InBackupQuantity),
		"outbound.backupQuantity=" + strconv.Itoa(c.OutBackupQuantity),
		"inbound.quantity=" + strconv.Itoa(c.InQuantity),
		"outbound.quantity=" + strconv.Itoa(c.OutQuantity),
		"inbound.allowZeroHop=" + strconv.FormatBool(c.InAllowZeroHop),
		"outbound.allowZeroHop=" + strconv.FormatBool(c.OutAllowZeroHop),
		"i2cp.encryptLeaseSet=" + strconv.FormatBool(c.EncryptLeaseSet),
		"i2cp.gzip=" + strconv.FormatBool(c.UseCompression),
		"i2cp.reduceOnIdle=" + strconv.FormatBool(c.ReduceIdle),
		"i2cp.reduceIdleTime=" + strconv.Itoa(c.ReduceIdleTime),
		"i2cp.reduceQuantity=" + strconv.Itoa(c.ReduceIdleQuantity),
		"i2cp.closeOnIdle=" + strconv.FormatBool(c.CloseIdle),
		"i2cp.closeIdleTime=" + strconv.Itoa(c.CloseIdleTime),
		c.accesslisttype(),
		c.accesslist(),
	}

	log.Println(confstring)
	return confstring
}

// Get passes directly through to goini.Get
func (c *Conf) Get(key string, label ...string) (string, bool) {
	if len(c.Labels) > 0 {
		if len(label) > 0 {
			return c.Config.SectionGet(label[0], key)
		}
		return c.Config.SectionGet(c.Labels[0], key)
	} else {
		if c.Config != nil {
			return c.Config.Get(key)
		} else {
			return "", false
		}
	}
	return "", false
}

// GetBool passes directly through to goini.GetBool
func (c *Conf) GetBool(key string, label ...string) (bool, bool) {
	if len(c.Labels) > 0 {
		if len(label) > 0 {
			return c.Config.SectionGetBool(label[0], key)
		}
		return c.Config.SectionGetBool(c.Labels[0], key)
	} else {
		if c.Config != nil {
			return c.Config.GetBool(key)
		} else {
			return false, false
		}
	}
	return false, false
}

// GetInt passes directly through to goini.GetInt
func (c *Conf) GetInt(key string, label ...string) (int, bool) {
	if len(c.Labels) > 0 {
		if len(label) > 0 {
			return c.Config.SectionGetInt(label[0], key)
		}
		return c.Config.SectionGetInt(c.Labels[0], key)
	} else {
		if c.Config != nil {
			return c.Config.GetInt(key)
		} else {
			return -1, false
		}
	}
	return -1, false
}

func (c *Conf) Write() error {
	if file, err := os.Open(filepath.Join(c.FilePath, c.TunName+".ini")); err != nil {
		defer file.Close()
		return err
	} else {
		defer file.Close()
		return c.Config.Write(file)
	}
}

/*
// Get takes an argument and a default. If the argument differs from the
// default, the argument is always returned. If the argument and default are
// the same and the key exists, the key is returned. If the key is absent, the
// default is returned.
func (c *Conf) Get(arg, def int, label ...string) int {
    return 0
}
*/

// SetLabels
func (c *Conf) SetLabels(iniFile string) {
	if tempfile, temperr := ioutil.ReadFile(iniFile); temperr == nil {
		tempstring := string(tempfile)
		tempslice := strings.Split(tempstring, "\n")
		for _, element := range tempslice {
			trimmedelement := strings.Trim(element, "\t\n ")
			if strings.HasPrefix(trimmedelement, "[") && strings.HasSuffix(trimmedelement, "]") {
				c.Labels = append(c.Labels, strings.Trim(trimmedelement, "[]"))
			}
		}
	}
	log.Println("Found Labels:", c.Labels)
}

/*
// Set
func (c *Conf) Set(label ...string) {

}
*/

// I2PINILoad loads variables from an ini file into the Conf data structure.
func (c *Conf) I2PINILoad(iniFile string, label ...string) error {
	var err error
	c.exists = true
	c.Config = goini.New()
	if iniFile != "none" && iniFile != "" {
		c.FilePath = iniFile
		err = c.Config.ParseFile(iniFile)
		if err != nil {
			return err
		}
		c.SetLabels(iniFile)
		c.SetDir(label...)
		c.SetType(label...)
		c.SetKeys(label...)
		c.SetHost(label...)
		c.SetPort(label...)
		c.SetSAMHost(label...)
		c.SetSAMPort(label...)
		c.SetTunName(label...)
		c.SetEncryptLease(label...)
		c.SetLeasesetKey(label...)
		c.SetLeasesetPrivateKey(label...)
		c.SetLeasesetPrivateSigningKey(label...)
		c.SetAllowZeroHopIn(label...)
		c.SetAllowZeroHopOut(label...)
		c.SetInLength(label...)
		c.SetOutLength(label...)
		c.SetInQuantity(label...)
		c.SetOutQuantity(label...)
		c.SetInVariance(label...)
		c.SetOutVariance(label...)
		c.SetInBackups(label...)
		c.SetOutBackups(label...)
		c.SetFastRecieve(label...)
		c.SetCompressed(label...)
		c.SetReduceIdle(label...)
		c.SetReduceIdleTime(label...)
		c.SetReduceIdleQuantity(label...)
		c.SetCloseIdle(label...)
		c.SetCloseIdleTime(label...)
		c.SetAccessListType(label...)
		c.SetTargetPort443(label...)
		c.SetMessageReliability(label...)
		c.SetClientDest(label...)
		c.SetKeyFile(label...)
		if v, ok := c.Get("i2cp.accessList", label...); ok {
			csv := strings.Split(v, ",")
			for _, z := range csv {
				c.AccessList = append(c.AccessList, z)
			}
		}
	}
	return nil
}

// NewI2PBlankTunConf returns an empty but intialized tunconf
func NewI2PBlankTunConf() *Conf {
	var c Conf
	c.Config = goini.New()
	return &c
}

// NewI2PTunConf returns a Conf structure from an ini file, for modification
// before starting the tunnel
func NewI2PTunConf(iniFile string, label ...string) (*Conf, error) {
	c := NewI2PBlankTunConf()
	if err := c.I2PINILoad(iniFile, label...); err != nil {
		return nil, err
	}
	return c, nil
}

// NewSAMForwarderFromConf generates a SAMforwarder from *i2ptunconf.Conf
func NewSAMForwarderFromConf(config *Conf) (*samforwarder.SAMForwarder, error) {
	if config != nil {
		return samforwarder.NewSAMForwarderFromOptions(
			samforwarder.SetType(config.Type),
			samforwarder.SetSaveFile(config.SaveFile),
			samforwarder.SetFilePath(config.SaveDirectory),
			samforwarder.SetHost(config.TargetHost),
			samforwarder.SetPort(config.TargetPort),
			samforwarder.SetSAMHost(config.SamHost),
			samforwarder.SetSAMPort(config.SamPort),
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
		return NewSAMForwarderFromConf(config)
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

// NewSAMSSUForwarderFromConf generates a SAMSSUforwarder from *i2ptunconf.Conf
func NewSAMSSUForwarderFromConf(config *Conf) (*samforwarderudp.SAMSSUForwarder, error) {
	if config != nil {
		return samforwarderudp.NewSAMSSUForwarderFromOptions(
			samforwarderudp.SetSaveFile(config.SaveFile),
			samforwarderudp.SetFilePath(config.SaveDirectory),
			samforwarderudp.SetHost(config.TargetHost),
			samforwarderudp.SetPort(config.TargetPort),
			samforwarderudp.SetSAMHost(config.SamHost),
			samforwarderudp.SetSAMPort(config.SamPort),
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

// NewSAMSSUForwarderFromConfig generates a new SAMSSUForwarder from a config file
func NewSAMSSUForwarderFromConfig(iniFile, SamHost, SamPort string, label ...string) (*samforwarderudp.SAMSSUForwarder, error) {
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
		return NewSAMSSUForwarderFromConf(config)
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

// NewSAMVPNClientForwarderFromConfig generates a new SAMVPNForwarder from a config file
func NewSAMVPNClientForwarderFromConfig(iniFile, SamHost, SamPort string, label ...string) (*samforwardervpn.SAMClientServerVPN, error) {
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
		return samforwardervpn.NewSAMClientServerVPN(config)
	}
	return nil, nil
}

// NewSAMVPNForwarderFromConfig generates a new SAMVPNForwarder from a config file
func NewSAMVPNForwarderFromConfig(iniFile, SamHost, SamPort string, label ...string) (*samforwardervpn.SAMClientVPN, error) {
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
		return samforwardervpn.NewSAMClientVPN(config)
	}
	return nil, nil
}
