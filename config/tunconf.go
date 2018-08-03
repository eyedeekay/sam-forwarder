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
	saveFile           bool
	TargetHost         string
	TargetPort         string
	SamHost            string
	SamPort            string
	TunName            string
	encryptLeaseSet    bool
	inAllowZeroHop     bool
	outAllowZeroHop    bool
	inLength           int
	outLength          int
	inQuantity         int
	outQuantity        int
	inVariance         int
	outVariance        int
	inBackupQuantity   int
	outBackupQuantity  int
	useCompression     bool
	reduceIdle         bool
	reduceIdleTime     int
	reduceIdleQuantity int
	accessListType     string
	accessList         []string
}

func (c *Conf) Print() {
	log.Println(
		"\n", c.saveFile,
		"\n", c.TargetHost,
		"\n", c.TargetPort,
		"\n", c.SamHost,
		"\n", c.SamPort,
		"\n", c.TunName,
		"\n", c.encryptLeaseSet,
		"\n", c.inAllowZeroHop,
		"\n", c.outAllowZeroHop,
		"\n", c.inLength,
		"\n", c.outLength,
		"\n", c.inQuantity,
		"\n", c.outQuantity,
		"\n", c.inVariance,
		"\n", c.outVariance,
		"\n", c.inBackupQuantity,
		"\n", c.outBackupQuantity,
		"\n", c.useCompression,
		"\n", c.reduceIdle,
		"\n", c.reduceIdleTime,
		"\n", c.reduceIdleQuantity,
		"\n", c.accessListType,
		"\n", c.accessList,
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
    for _, item := range c.accessList {
        if item == key {
            return
        }
    }
    c.accessList = append(c.accessList, key)
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
			c.saveFile = true
		} else {
			c.saveFile = false
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
			c.encryptLeaseSet = v
		} else {
			c.encryptLeaseSet = false
		}

		if v, ok := c.config.GetBool("inbound.allowZeroHop"); ok {
			c.inAllowZeroHop = v
		} else {
			c.inAllowZeroHop = false
		}
		if v, ok := c.config.GetBool("outbound.allowZeroHop"); ok {
			c.outAllowZeroHop = v
		} else {
			c.outAllowZeroHop = false
		}

		if v, ok := c.config.GetInt("inbound.length"); ok {
			c.inLength = v
		} else {
			c.inLength = 3
		}
		if v, ok := c.config.GetInt("outbound.length"); ok {
			c.outLength = v
		} else {
			c.outLength = 3
		}

		if v, ok := c.config.GetInt("inbound.quantity"); ok {
			c.inQuantity = v
		} else {
			c.inQuantity = 5
		}
		if v, ok := c.config.GetInt("outbound.quantity"); ok {
			c.outQuantity = v
		} else {
			c.outQuantity = 5
		}

		if v, ok := c.config.GetInt("inbound.variance"); ok {
			c.inVariance = v
		} else {
			c.inVariance = 0
		}
		if v, ok := c.config.GetInt("outbound.variance"); ok {
			c.outVariance = v
		} else {
			c.outVariance = 0
		}

		if v, ok := c.config.GetInt("inbound.backupQuantity"); ok {
			c.inBackupQuantity = v
		} else {
			c.inBackupQuantity = 2
		}
		if v, ok := c.config.GetInt("outbound.backupQuantity"); ok {
			c.outBackupQuantity = v
		} else {
			c.outBackupQuantity = 2
		}

		if v, ok := c.config.GetBool("gzip"); ok {
			c.useCompression = v
		} else {
			c.useCompression = true
		}

		if v, ok := c.config.GetBool("i2cp.reduceOnIdle"); ok {
			c.reduceIdle = v
		} else {
			c.reduceIdle = false
		}
		if v, ok := c.config.GetInt("i2cp.reduceIdleTime"); ok {
			c.reduceIdleTime = (v / 1000) / 60
		} else {
			c.reduceIdleTime = (6 * 60) * 1000
		}
		if v, ok := c.config.GetInt("i2cp.reduceQuantity"); ok {
			c.reduceIdleQuantity = v
		} else {
			c.reduceIdleQuantity = 3
		}

		if v, ok := c.config.GetBool("i2cp.enableBlackList"); ok {
			if v {
				c.accessListType = "blacklist"
			}
		}
		if v, ok := c.config.GetBool("i2cp.enableAccessList"); ok {
			if v {
				c.accessListType = "whitelist"
			}
		}
		if c.accessListType != "whitelist" && c.accessListType != "blacklist" {
			c.accessListType = "none"
		}
		if v, ok := c.config.Get("i2cp.accessList"); ok {
			csv := strings.Split(v, ",")
			for _, z := range csv {
				c.accessList = append(c.accessList, z)
			}
		}
		return &c, nil
	}
	return nil, nil
}

func NewSAMForwarderFromConf(config *Conf) (*samforwarder.SAMForwarder, error) {
	if config != nil {
		return samforwarder.NewSAMForwarderFromOptions(
			samforwarder.SetSaveFile(config.saveFile),
			samforwarder.SetHost(config.TargetHost),
			samforwarder.SetPort(config.TargetPort),
			samforwarder.SetSAMHost(config.SamHost),
			samforwarder.SetSAMPort(config.SamPort),
			samforwarder.SetName(config.TunName),
			samforwarder.SetInLength(config.inLength),
			samforwarder.SetOutLength(config.outLength),
			samforwarder.SetInVariance(config.inVariance),
			samforwarder.SetOutVariance(config.outVariance),
			samforwarder.SetInQuantity(config.inQuantity),
			samforwarder.SetOutQuantity(config.outQuantity),
			samforwarder.SetInBackups(config.inBackupQuantity),
			samforwarder.SetOutBackups(config.outBackupQuantity),
			samforwarder.SetEncrypt(config.encryptLeaseSet),
			samforwarder.SetAllowZeroIn(config.inAllowZeroHop),
			samforwarder.SetAllowZeroOut(config.outAllowZeroHop),
			samforwarder.SetCompress(config.useCompression),
			samforwarder.SetReduceIdle(config.reduceIdle),
			samforwarder.SetReduceIdleTime(config.reduceIdleTime),
			samforwarder.SetReduceIdleQuantity(config.reduceIdleQuantity),
			samforwarder.SetAccessListType(config.accessListType),
			samforwarder.SetAccessList(config.accessList),
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
			samforwarder.SetSaveFile(config.saveFile),
			samforwarder.SetHost(config.TargetHost),
			samforwarder.SetPort(config.TargetPort),
			samforwarder.SetSAMHost(config.SamHost),
			samforwarder.SetSAMPort(config.SamPort),
			samforwarder.SetName(config.TunName),
			samforwarder.SetInLength(config.inLength),
			samforwarder.SetOutLength(config.outLength),
			samforwarder.SetInVariance(config.inVariance),
			samforwarder.SetOutVariance(config.outVariance),
			samforwarder.SetInQuantity(config.inQuantity),
			samforwarder.SetOutQuantity(config.outQuantity),
			samforwarder.SetInBackups(config.inBackupQuantity),
			samforwarder.SetOutBackups(config.outBackupQuantity),
			samforwarder.SetEncrypt(config.encryptLeaseSet),
			samforwarder.SetAllowZeroIn(config.inAllowZeroHop),
			samforwarder.SetAllowZeroOut(config.outAllowZeroHop),
			samforwarder.SetCompress(config.useCompression),
			samforwarder.SetReduceIdle(config.reduceIdle),
			samforwarder.SetReduceIdleTime(config.reduceIdleTime),
			samforwarder.SetReduceIdleQuantity(config.reduceIdleQuantity),
			samforwarder.SetAccessListType(config.accessListType),
			samforwarder.SetAccessList(config.accessList),
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
			samforwarderudp.SetSaveFile(config.saveFile),
			samforwarderudp.SetHost(config.TargetHost),
			samforwarderudp.SetPort(config.TargetPort),
			samforwarderudp.SetSAMHost(config.SamHost),
			samforwarderudp.SetSAMPort(config.SamPort),
			samforwarderudp.SetName(config.TunName),
			samforwarderudp.SetInLength(config.inLength),
			samforwarderudp.SetOutLength(config.outLength),
			samforwarderudp.SetInVariance(config.inVariance),
			samforwarderudp.SetOutVariance(config.outVariance),
			samforwarderudp.SetInQuantity(config.inQuantity),
			samforwarderudp.SetOutQuantity(config.outQuantity),
			samforwarderudp.SetInBackups(config.inBackupQuantity),
			samforwarderudp.SetOutBackups(config.outBackupQuantity),
			samforwarderudp.SetEncrypt(config.encryptLeaseSet),
			samforwarderudp.SetAllowZeroIn(config.inAllowZeroHop),
			samforwarderudp.SetAllowZeroOut(config.outAllowZeroHop),
			samforwarderudp.SetCompress(config.useCompression),
			samforwarderudp.SetReduceIdle(config.reduceIdle),
			samforwarderudp.SetReduceIdleTime(config.reduceIdleTime),
			samforwarderudp.SetReduceIdleQuantity(config.reduceIdleQuantity),
			samforwarderudp.SetAccessListType(config.accessListType),
			samforwarderudp.SetAccessList(config.accessList),
		)
	}
	return nil, nil
}

func NewSAMSSUForwarderFromConf(config *Conf) (*samforwarderudp.SAMSSUForwarder, error) {
	if config != nil {
		return samforwarderudp.NewSAMSSUForwarderFromOptions(
			samforwarderudp.SetSaveFile(config.saveFile),
			samforwarderudp.SetHost(config.TargetHost),
			samforwarderudp.SetPort(config.TargetPort),
			samforwarderudp.SetSAMHost(config.SamHost),
			samforwarderudp.SetSAMPort(config.SamPort),
			samforwarderudp.SetName(config.TunName),
			samforwarderudp.SetInLength(config.inLength),
			samforwarderudp.SetOutLength(config.outLength),
			samforwarderudp.SetInVariance(config.inVariance),
			samforwarderudp.SetOutVariance(config.outVariance),
			samforwarderudp.SetInQuantity(config.inQuantity),
			samforwarderudp.SetOutQuantity(config.outQuantity),
			samforwarderudp.SetInBackups(config.inBackupQuantity),
			samforwarderudp.SetOutBackups(config.outBackupQuantity),
			samforwarderudp.SetEncrypt(config.encryptLeaseSet),
			samforwarderudp.SetAllowZeroIn(config.inAllowZeroHop),
			samforwarderudp.SetAllowZeroOut(config.outAllowZeroHop),
			samforwarderudp.SetCompress(config.useCompression),
			samforwarderudp.SetReduceIdle(config.reduceIdle),
			samforwarderudp.SetReduceIdleTime(config.reduceIdleTime),
			samforwarderudp.SetReduceIdleQuantity(config.reduceIdleQuantity),
			samforwarderudp.SetAccessListType(config.accessListType),
			samforwarderudp.SetAccessList(config.accessList),
		)
	}
	return nil, nil
}
