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
