package i2ptunconf

import (
	"github.com/eyedeekay/sam-forwarder"
	"github.com/eyedeekay/sam-forwarder/udp"
	"github.com/zieckey/goini"
	"log"
	"strconv"
	"strings"
)

// Conf is a tructure containing an ini config, with some functions to help
// when you use it for in conjunction with command-line flags
type Conf struct {
	config             *goini.INI
	SaveDirectory      string
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

func (c *Conf) accesslisttype() string {
	if c.AccessListType == "whitelist" {
		return "i2cp.enableAccessList=true"
	} else if c.AccessListType == "blacklist" {
		return "i2cp.enableBlackList=true"
	} else if c.AccessListType == "none" {
		return ""
	}
	return ""
}

func (c *Conf) accesslist() string {
	if c.AccessListType != "" && len(c.AccessList) > 0 {
		r := ""
		for _, s := range c.AccessList {
			r += s + ","
		}
		return "i2cp.accessList=" + strings.TrimSuffix(r, ",")
	}
	return ""
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

// AddAccessListMember adds a member to either the blacklist or the whitelist
func (c *Conf) AddAccessListMember(key string) {
	for _, item := range c.AccessList {
		if item == key {
			return
		}
	}
	c.AccessList = append(c.AccessList, key)
}

// GetHost takes an argument and a default. If the argument differs from the
// default, the argument is always returned. If the argument and default are
// the same and the key exists, the key is returned. If the key is absent, the
// default is returned.
func (c *Conf) GetHost(arg, def string) string {
	if arg != def {
		return arg
	}
	if x, o := c.Get("host"); o {
		return x
	}
	return arg
}

// GetPort takes an argument and a default. If the argument differs from the
// default, the argument is always returned. If the argument and default are
// the same and the key exists, the key is returned. If the key is absent, the
// default is returned.
func (c *Conf) GetPort(arg, def string) string {
	if arg != def {
		return arg
	}
	if x, o := c.Get("port"); o {
		return x
	}
	return arg
}

// GetSAMHost takes an argument and a default. If the argument differs from the
// default, the argument is always returned. If the argument and default are
// the same and the key exists, the key is returned. If the key is absent, the
// default is returned.
func (c *Conf) GetSAMHost(arg, def string) string {
	if arg != def {
		return arg
	}
	if x, o := c.Get("samhost"); o {
		return x
	}
	return arg
}

// GetSAMPort takes an argument and a default. If the argument differs from the
// default, the argument is always returned. If the argument and default are
// the same and the key exists, the key is returned. If the key is absent, the
// default is returned.
func (c *Conf) GetSAMPort(arg, def string) string {
	if arg != def {
		return arg
	}
	if x, o := c.Get("samport"); o {
		return x
	}
	return arg
}

// GetDir takes an argument and a default. If the argument differs from the
// default, the argument is always returned. If the argument and default are
// the same and the key exists, the key is returned. If the key is absent, the
// default is returned.
func (c *Conf) GetDir(arg, def string) string {
	if arg != def {
		return arg
	}
	if x, o := c.Get("dir"); o {
		return x
	}
	return arg
}

// GetKeys takes an argument and a default. If the argument differs from the
// default, the argument is always returned. If the argument and default are
// the same and the key exists, the key is returned. If the key is absent, the
// default is returned.
func (c *Conf) GetKeys(arg, def string) string {
	if arg != def {
		return arg
	}
	if x, o := c.Get("keys"); o {
		return x
	}
	return arg
}

// GetInLength takes an argument and a default. If the argument differs from the
// default, the argument is always returned. If the argument and default are
// the same and the key exists, the key is returned. If the key is absent, the
// default is returned.
func (c *Conf) GetInLength(arg, def int) int {
	if arg != def {
		return arg
	}
	if x, o := c.GetInt("inbound.length"); o {
		return x
	}
	return arg
}

// GetOutLength takes an argument and a default. If the argument differs from the
// default, the argument is always returned. If the argument and default are
// the same and the key exists, the key is returned. If the key is absent, the
// default is returned.
func (c *Conf) GetOutLength(arg, def int) int {
	if arg != def {
		return arg
	}
	if x, o := c.GetInt("outbound.length"); o {
		return x
	}
	return arg
}

// GetInVariance takes an argument and a default. If the argument differs from the
// default, the argument is always returned. If the argument and default are
// the same and the key exists, the key is returned. If the key is absent, the
// default is returned.
func (c *Conf) GetInVariance(arg, def int) int {
	if arg != def {
		return arg
	}
	if x, o := c.GetInt("inbound.variance"); o {
		return x
	}
	return arg
}

// GetOutVariance takes an argument and a default. If the argument differs from the
// default, the argument is always returned. If the argument and default are
// the same and the key exists, the key is returned. If the key is absent, the
// default is returned.
func (c *Conf) GetOutVariance(arg, def int) int {
	if arg != def {
		return arg
	}
	if x, o := c.GetInt("outbound.variance"); o {
		return x
	}
	return arg
}

// GetInQuantity takes an argument and a default. If the argument differs from the
// default, the argument is always returned. If the argument and default are
// the same and the key exists, the key is returned. If the key is absent, the
// default is returned.
func (c *Conf) GetInQuantity(arg, def int) int {
	if arg != def {
		return arg
	}
	if x, o := c.GetInt("inbound.quantity"); o {
		return x
	}
	return arg
}

// GetOutQuantity takes an argument and a default. If the argument differs from the
// default, the argument is always returned. If the argument and default are
// the same and the key exists, the key is returned. If the key is absent, the
// default is returned.
func (c *Conf) GetOutQuantity(arg, def int) int {
	if arg != def {
		return arg
	}
	if x, o := c.GetInt("outbound.quantity"); o {
		return x
	}
	return arg
}

// GetInBackups takes an argument and a default. If the argument differs from the
// default, the argument is always returned. If the argument and default are
// the same and the key exists, the key is returned. If the key is absent, the
// default is returned.
func (c *Conf) GetInBackups(arg, def int) int {
	if arg != def {
		return arg
	}
	if x, o := c.GetInt("inbound.backupQuantity"); o {
		return x
	}
	return arg
}

// GetOutBackups takes an argument and a default. If the argument differs from the
// default, the argument is always returned. If the argument and default are
// the same and the key exists, the key is returned. If the key is absent, the
// default is returned.
func (c *Conf) GetOutBackups(arg, def int) int {
	if arg != def {
		return arg
	}
	if x, o := c.GetInt("outbound.backupQuantity"); o {
		return x
	}
	return arg
}

// GetEncryptLeaseset takes an argument and a default. If the argument differs from the
// default, the argument is always returned. If the argument and default are
// the same and the key exists, the key is returned. If the key is absent, the
// default is returned.
func (c *Conf) GetEncryptLeaseset(arg, def bool) bool {
	if arg != def {
		return arg
	}
	if x, o := c.GetBool("i2cp.encryptLeaseSet"); o {
		return x
	}
	return arg
}

// GetInAllowZeroHop takes an argument and a default. If the argument differs from the
// default, the argument is always returned. If the argument and default are
// the same and the key exists, the key is returned. If the key is absent, the
// default is returned.
func (c *Conf) GetInAllowZeroHop(arg, def bool) bool {
	if arg != def {
		return arg
	}
	if x, o := c.GetBool("inbound.allowZeroHop"); o {
		return x
	}
	return arg
}

// GetOutAllowZeroHop takes an argument and a default. If the argument differs from the
// default, the argument is always returned. If the argument and default are
// the same and the key exists, the key is returned. If the key is absent, the
// default is returned.
func (c *Conf) GetOutAllowZeroHop(arg, def bool) bool {
	if arg != def {
		return arg
	}
	if x, o := c.GetBool("outbound.allowZeroHop"); o {
		return x
	}
	return arg
}

// GetUseCompression takes an argument and a default. If the argument differs from the
// default, the argument is always returned. If the argument and default are
// the same and the key exists, the key is returned. If the key is absent, the
// default is returned.
func (c *Conf) GetUseCompression(arg, def bool) bool {
	if arg != def {
		return arg
	}
	if x, o := c.GetBool("gzip"); o {
		return x
	}
	return arg
}

// GetReduceOnIdle takes an argument and a default. If the argument differs from the
// default, the argument is always returned. If the argument and default are
// the same and the key exists, the key is returned. If the key is absent, the
// default is returned.
func (c *Conf) GetReduceOnIdle(arg, def bool) bool {
	if arg != def {
		return arg
	}
	if x, o := c.GetBool("i2cp.reduceOnIdle"); o {
		return x
	}
	return arg
}

// GetReduceIdleTime takes an argument and a default. If the argument differs from the
// default, the argument is always returned. If the argument and default are
// the same and the key exists, the key is returned. If the key is absent, the
// default is returned.
func (c *Conf) GetReduceIdleTime(arg, def int) int {
	if arg != def {
		return arg
	}
	if x, o := c.GetInt("i2cp.reduceIdleTime"); o {
		return x
	}
	return arg
}

// GetReduceIdleQuantity takes an argument and a default. If the argument differs from the
// default, the argument is always returned. If the argument and default are
// the same and the key exists, the key is returned. If the key is absent, the
// default is returned.
func (c *Conf) GetReduceIdleQuantity(arg, def int) int {
	if arg != def {
		return arg
	}
	if x, o := c.GetInt("i2cp.reduceIdleQuantity"); o {
		return x
	}
	return arg
}

// NewI2PTunConf returns a Conf structure from an ini file, for modification
// before starting the tunnel
func NewI2PTunConf(iniFile string) (*Conf, error) {
	var err error
	var c Conf
	c.config = goini.New()
	if iniFile != "none" {
		err = c.config.ParseFile(iniFile)
		if err != nil {
			return nil, err
		}

		if v, ok := c.config.Get("dir"); ok {
			c.SaveDirectory = v
		} else {
			c.SaveDirectory = "./"
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

// NewSAMForwarderFromConf generates a SAMforwarder from *i2ptunconf.Conf
func NewSAMForwarderFromConf(config *Conf) (*samforwarder.SAMForwarder, error) {
	if config != nil {
		return samforwarder.NewSAMForwarderFromOptions(
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

// NewSAMForwarderFromConfig generates a new SAMForwarder from a config file
func NewSAMForwarderFromConfig(iniFile, SamHost, SamPort string) (*samforwarder.SAMForwarder, error) {
	if iniFile != "none" {
		config, err := NewI2PTunConf(iniFile)
		if err != nil {
			return nil, err
		}
		return samforwarder.NewSAMForwarderFromOptions(
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

// NewSAMSSUForwarderFromConfig generates a new SAMSSUForwarder from a config file
func NewSAMSSUForwarderFromConfig(iniFile, SamHost, SamPort string) (*samforwarderudp.SAMSSUForwarder, error) {
	if iniFile != "none" {
		config, err := NewI2PTunConf(iniFile)
		if err != nil {
			return nil, err
		}
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
