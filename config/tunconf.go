package i2ptunconf

import (
	"github.com/eyedeekay/sam-forwarder"
	"github.com/eyedeekay/sam-forwarder/udp"
	"github.com/zieckey/goini"
	"log"
	"strings"
)

type Conf struct {
	config             *goini.INI
	SaveFile           bool
	TargetHost         string
	TargetPort         string
	SamHost            string
	SamPort            string
	TunName            string
	EncryptLeaseSet    bool
	InAllowZeroHop     bool
	OutAllowZeroHop    bool
	InLength           int
	OutLength          int
	InQuantity         int
	OutQuantity        int
	InVariance         int
	OutVariance        int
	InBackupQuantity   int
	OutBackupQuantity  int
	UseCompression     bool
	ReduceIdle         bool
	ReduceIdleTime     int
	ReduceIdleQuantity int
	CloseIdle          bool
	CloseIdleTime      int
	AccessListType     string
	AccessList         []string
}

func (c *Conf) Print() {
	log.Println(
		"\n", c.SaveFile,
		"\n", c.TargetHost,
		"\n", c.TargetPort,
		"\n", c.SamHost,
		"\n", c.SamPort,
		"\n", c.TunName,
		"\n", c.EncryptLeaseSet,
		"\n", c.InAllowZeroHop,
		"\n", c.OutAllowZeroHop,
		"\n", c.InLength,
		"\n", c.OutLength,
		"\n", c.InQuantity,
		"\n", c.OutQuantity,
		"\n", c.InVariance,
		"\n", c.OutVariance,
		"\n", c.InBackupQuantity,
		"\n", c.OutBackupQuantity,
		"\n", c.UseCompression,
		"\n", c.ReduceIdle,
		"\n", c.ReduceIdleTime,
		"\n", c.ReduceIdleQuantity,
		"\n", c.AccessListType,
		"\n", c.AccessList,
	)
}

// Get passes directly through to goini.Get
func (c *Conf) Get(key string) (string, bool) {
	return c.config.Get(key)
}

// GetBool passes directly through to goini.GetBool
func (c *Conf) GetBool(key string) (bool, bool) {
	return c.config.GetBool(key)
}

// GetInt passes directly through to goini.GetInt
func (c *Conf) GetInt(key string) (int, bool) {
	return c.config.GetInt(key)
}

func (c *Conf) AddAccessListMember(key string) {
	for _, item := range c.AccessList {
		if item == key {
			return
		}
	}
	c.AccessList = append(c.AccessList, key)
}
func (c *Conf) GetHost(arg string) string {
	if x, o := c.Get("host"); o {
		return x
	}
	return arg
}

func (c *Conf) GetPort(arg string) string {
	if x, o := c.Get("port"); o {
		return x
	}
	return arg
}
func (c *Conf) GetSAMHost(arg string) string {
	if x, o := c.Get("samhost"); o {
		return x
	}
	return arg
}

func (c *Conf) GetSAMPort(arg string) string {
	if x, o := c.Get("samport"); o {
		return x
	}
	return arg
}

func (c *Conf) GetKeys(arg string) string {
	if x, o := c.Get("keys"); o {
		return x
	}
	return arg
}

func (c *Conf) GetInLength(arg int) int {
	if x, o := c.GetInt("inbound.length"); o {
		return x
	}
	return arg
}

func (c *Conf) GetOutLength(arg int) int {
	if x, o := c.GetInt("outbound.length"); o {
		return x
	}
	return arg
}
func (c *Conf) GetInVariance(arg int) int {
	if x, o := c.GetInt("inbound.variance"); o {
		return x
	}
	return arg
}
func (c *Conf) GetOutVariance(arg int) int {
	if x, o := c.GetInt("outbound.variance"); o {
		return x
	}
	return arg
}
func (c *Conf) GetInQuantity(arg int) int {
	if x, o := c.GetInt("inbound.quantity"); o {
		return x
	}
	return arg
}
func (c *Conf) GetOutQuantity(arg int) int {
	if x, o := c.GetInt("outbound.quantity"); o {
		return x
	}
	return arg
}
func (c *Conf) GetInBackups(arg int) int {
	if x, o := c.GetInt("inbound.backupQuantity"); o {
		return x
	}
	return arg
}
func (c *Conf) GetOutBackups(arg int) int {
	if x, o := c.GetInt("outbound.backupQuantity"); o {
		return x
	}
	return arg
}
func (c *Conf) GetEncryptLeaseset(arg bool) bool {
	if x, o := c.GetBool("i2cp.encryptLeaseSet"); o {
		return x
	}
	return arg
}
func (c *Conf) GetInAllowZeroHop(arg bool) bool {
	if x, o := c.GetBool("inbound.allowZeroHop"); o {
		return x
	}
	return arg
}
func (c *Conf) GetOutAllowZeroHop(arg bool) bool {
	if x, o := c.GetBool("outbound.allowZeroHop"); o {
		return x
	}
	return arg
}
func (c *Conf) GetUseCompression(arg bool) bool {
	if x, o := c.GetBool("gzip"); o {
		return x
	}
	return arg
}
func (c *Conf) GetReduceOnIdle(arg bool) bool {
	if x, o := c.GetBool("i2cp.reduceOnIdle"); o {
		return x
	}
	return arg
}
func (c *Conf) GetReduceIdleTime(arg int) int {
	if x, o := c.GetInt("i2cp.reduceIdleTime"); o {
		return x
	}
	return arg
}
func (c *Conf) GetReduceIdleQuantity(arg int) int {
	if x, o := c.GetInt("i2cp.reduceIdleQuantity"); o {
		return x
	}
	return arg
}
func NewI2PTunConf(iniFile string) (*Conf, error) {
	var err error
	var c Conf
	c.config = goini.New()
	if iniFile != "none" {
		err = c.config.ParseFile(iniFile)
		if err != nil {
			return nil, err
		}

		if _, ok := c.config.Get("keys"); ok {
			c.SaveFile = true
		} else {
			c.SaveFile = false
		}

		if v, ok := c.config.Get("host"); ok {
			c.TargetHost = strings.Replace(v, ":", "", -1)
		} else {
			c.TargetHost = "127.0.0.1"
		}
		if v, ok := c.config.Get("port"); ok {
			c.TargetPort = strings.Replace(v, ":", "", -1)
		} else {
			c.TargetPort = "8081"
		}

		if v, ok := c.config.Get("samhost"); ok {
			c.SamHost = strings.Replace(v, ":", "", -1)
		} else {
			c.SamHost = "127.0.0.1"
		}
		if v, ok := c.config.Get("samport"); ok {
			c.SamPort = strings.Replace(v, ":", "", -1)
		} else {
			c.SamPort = "7656"
		}

		if v, ok := c.config.Get("keys"); ok {
			c.TunName = v
		} else {
			c.TunName = "fowarder"
		}
		if v, ok := c.config.GetBool("i2cp.encryptLeaseSet"); ok {
			c.EncryptLeaseSet = v
		} else {
			c.EncryptLeaseSet = false
		}

		if v, ok := c.config.GetBool("inbound.allowZeroHop"); ok {
			c.InAllowZeroHop = v
		} else {
			c.InAllowZeroHop = false
		}
		if v, ok := c.config.GetBool("outbound.allowZeroHop"); ok {
			c.OutAllowZeroHop = v
		} else {
			c.OutAllowZeroHop = false
		}

		if v, ok := c.config.GetInt("inbound.length"); ok {
			c.InLength = v
		} else {
			c.InLength = 3
		}
		if v, ok := c.config.GetInt("outbound.length"); ok {
			c.OutLength = v
		} else {
			c.OutLength = 3
		}

		if v, ok := c.config.GetInt("inbound.quantity"); ok {
			c.InQuantity = v
		} else {
			c.InQuantity = 5
		}
		if v, ok := c.config.GetInt("outbound.quantity"); ok {
			c.OutQuantity = v
		} else {
			c.OutQuantity = 5
		}

		if v, ok := c.config.GetInt("inbound.variance"); ok {
			c.InVariance = v
		} else {
			c.InVariance = 0
		}
		if v, ok := c.config.GetInt("outbound.variance"); ok {
			c.OutVariance = v
		} else {
			c.OutVariance = 0
		}

		if v, ok := c.config.GetInt("inbound.backupQuantity"); ok {
			c.InBackupQuantity = v
		} else {
			c.InBackupQuantity = 2
		}
		if v, ok := c.config.GetInt("outbound.backupQuantity"); ok {
			c.OutBackupQuantity = v
		} else {
			c.OutBackupQuantity = 2
		}

		if v, ok := c.config.GetBool("gzip"); ok {
			c.UseCompression = v
		} else {
			c.UseCompression = true
		}

		if v, ok := c.config.GetBool("i2cp.reduceOnIdle"); ok {
			c.ReduceIdle = v
		} else {
			c.ReduceIdle = false
		}
		if v, ok := c.config.GetInt("i2cp.reduceIdleTime"); ok {
			c.ReduceIdleTime = (v / 1000) / 60
		} else {
			c.ReduceIdleTime = (6 * 60) * 1000
		}
		if v, ok := c.config.GetInt("i2cp.reduceQuantity"); ok {
			c.ReduceIdleQuantity = v
		} else {
			c.ReduceIdleQuantity = 3
		}

		if v, ok := c.config.GetBool("i2cp.closeOnIdle"); ok {
			c.CloseIdle = v
		} else {
			c.CloseIdle = false
		}
		if v, ok := c.config.GetInt("i2cp.closeIdleTime"); ok {
			c.CloseIdleTime = (v / 2000) / 60
		} else {
			c.CloseIdleTime = (6 * 60) * 2000
		}

		if v, ok := c.config.GetBool("i2cp.enableBlackList"); ok {
			if v {
				c.AccessListType = "blacklist"
			}
		}
		if v, ok := c.config.GetBool("i2cp.enableAccessList"); ok {
			if v {
				c.AccessListType = "whitelist"
			}
		}
		if c.AccessListType != "whitelist" && c.AccessListType != "blacklist" {
			c.AccessListType = "none"
		}
		if v, ok := c.config.Get("i2cp.accessList"); ok {
			csv := strings.Split(v, ",")
			for _, z := range csv {
				c.AccessList = append(c.AccessList, z)
			}
		}
		return &c, nil
	}
	return nil, nil
}

func NewSAMForwarderFromConf(config *Conf) (*samforwarder.SAMForwarder, error) {
	if config != nil {
		return samforwarder.NewSAMForwarderFromOptions(
			samforwarder.SetSaveFile(config.SaveFile),
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
			samforwarder.SetAllowZeroIn(config.InAllowZeroHop),
			samforwarder.SetAllowZeroOut(config.OutAllowZeroHop),
			samforwarder.SetCompress(config.UseCompression),
			samforwarder.SetReduceIdle(config.ReduceIdle),
			samforwarder.SetReduceIdleTime(config.ReduceIdleTime),
			samforwarder.SetReduceIdleQuantity(config.ReduceIdleQuantity),
			samforwarder.SetCloseIdle(config.CloseIdle),
			samforwarder.SetCloseIdleTime(config.CloseIdleTime),
			samforwarder.SetAccessListType(config.AccessListType),
			samforwarder.SetAccessList(config.AccessList),
		)
	}
	return nil, nil
}

func NewSAMForwarderFromConfig(iniFile, SamHost, SamPort string) (*samforwarder.SAMForwarder, error) {
	if iniFile != "none" {
		config, err := NewI2PTunConf(iniFile)
		if err != nil {
			return nil, err
		}
		return samforwarder.NewSAMForwarderFromOptions(
			samforwarder.SetSaveFile(config.SaveFile),
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
			samforwarder.SetAllowZeroIn(config.InAllowZeroHop),
			samforwarder.SetAllowZeroOut(config.OutAllowZeroHop),
			samforwarder.SetCompress(config.UseCompression),
			samforwarder.SetReduceIdle(config.ReduceIdle),
			samforwarder.SetReduceIdleTime(config.ReduceIdleTime),
			samforwarder.SetReduceIdleQuantity(config.ReduceIdleQuantity),
			samforwarder.SetCloseIdle(config.CloseIdle),
			samforwarder.SetCloseIdleTime(config.CloseIdleTime),
			samforwarder.SetAccessListType(config.AccessListType),
			samforwarder.SetAccessList(config.AccessList),
		)
	}
	return nil, nil
}

func NewSAMSSUForwarderFromConfig(iniFile, SamHost, SamPort string) (*samforwarderudp.SAMSSUForwarder, error) {
	if iniFile != "none" {
		config, err := NewI2PTunConf(iniFile)
		if err != nil {
			return nil, err
		}
		return samforwarderudp.NewSAMSSUForwarderFromOptions(
			samforwarderudp.SetSaveFile(config.SaveFile),
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
			samforwarderudp.SetAllowZeroIn(config.InAllowZeroHop),
			samforwarderudp.SetAllowZeroOut(config.OutAllowZeroHop),
			samforwarderudp.SetCompress(config.UseCompression),
			samforwarderudp.SetReduceIdle(config.ReduceIdle),
			samforwarderudp.SetReduceIdleTime(config.ReduceIdleTime),
			samforwarderudp.SetReduceIdleQuantity(config.ReduceIdleQuantity),
			samforwarderudp.SetCloseIdle(config.CloseIdle),
			samforwarderudp.SetCloseIdleTime(config.CloseIdleTime),
			samforwarderudp.SetAccessListType(config.AccessListType),
			samforwarderudp.SetAccessList(config.AccessList),
		)
	}
	return nil, nil
}

func NewSAMSSUForwarderFromConf(config *Conf) (*samforwarderudp.SAMSSUForwarder, error) {
	if config != nil {
		return samforwarderudp.NewSAMSSUForwarderFromOptions(
			samforwarderudp.SetSaveFile(config.SaveFile),
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
			samforwarderudp.SetAllowZeroIn(config.InAllowZeroHop),
			samforwarderudp.SetAllowZeroOut(config.OutAllowZeroHop),
			samforwarderudp.SetCompress(config.UseCompression),
			samforwarderudp.SetReduceIdle(config.ReduceIdle),
			samforwarderudp.SetReduceIdleTime(config.ReduceIdleTime),
			samforwarderudp.SetReduceIdleQuantity(config.ReduceIdleQuantity),
			samforwarderudp.SetCloseIdle(config.CloseIdle),
			samforwarderudp.SetCloseIdleTime(config.CloseIdleTime),
			samforwarderudp.SetAccessListType(config.AccessListType),
			samforwarderudp.SetAccessList(config.AccessList),
		)
	}
	return nil, nil
}
