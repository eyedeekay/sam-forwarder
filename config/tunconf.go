package i2ptunconf

import (
	"github.com/zieckey/goini"
	"strings"
)

type Conf struct {
	config             *goini.INI
	saveFile           bool
	TargetHost         string
	TargetPort         string
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

func (c *Conf) configParse(path string) (*goini.INI, error) {
	ini := goini.New()
	if err := ini.ParseFile(path); err != nil {
		return nil, err
	}
	return ini, nil
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

func NewI2PTunConf(iniFile string) (*Conf, error) {
	var c Conf
	var err error
	if iniFile != "none" {
		c.config, err = c.configParse(iniFile)
		if err != nil {
			return nil, err
		}
		if v, ok := c.config.GetBool("keys"); ok {
			c.saveFile = v
		}
		if v, ok := c.config.Get("host"); ok {
			c.TargetHost = v
		}
		if v, ok := c.config.Get("port"); ok {
			c.TargetPort = v
		}
		if v, ok := c.config.Get("keys"); ok {
			c.TunName = v
		}
		if v, ok := c.config.GetBool("i2cp.encryptLeaseSet"); ok {
			c.encryptLeaseSet = v
		}
		if v, ok := c.config.GetBool("inbound.allowZeroHop"); ok {
			c.inAllowZeroHop = v
		}
		if v, ok := c.config.GetBool("outbound.allowZeroHop"); ok {
			c.outAllowZeroHop = v
		}
		if v, ok := c.config.GetInt("inbound.length"); ok {
			c.inLength = v
		}
		if v, ok := c.config.GetInt("outbound.length"); ok {
			c.outLength = v
		}
		if v, ok := c.config.GetInt("inbound.quantity"); ok {
			c.inQuantity = v
		}
		if v, ok := c.config.GetInt("outbound.quantity"); ok {
			c.outQuantity = v
		}
		if v, ok := c.config.GetInt("inbound.variance"); ok {
			c.inVariance = v
		}
		if v, ok := c.config.GetInt("outbound.variance"); ok {
			c.outVariance = v
		}
		if v, ok := c.config.GetInt("inbound.backupQuantity"); ok {
			c.inBackupQuantity = v
		}
		if v, ok := c.config.GetInt("outbound.backupQuantity"); ok {
			c.outBackupQuantity = v
		}
		if v, ok := c.config.GetBool("gzip"); ok {
			c.useCompression = v
		}
		if v, ok := c.config.GetBool("i2cp.reduceOnIdle"); ok {
			c.reduceIdle = v
		}
		if v, ok := c.config.GetInt("i2cp.reduceIdleTime"); ok {
			c.reduceIdleTime = v
		}
		if v, ok := c.config.GetInt("i2cp.reduceQuantity"); ok {
			c.reduceIdleQuantity = v
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
