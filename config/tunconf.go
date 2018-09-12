package i2ptunconf

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

import (
	"github.com/eyedeekay/sam-forwarder"
	"github.com/eyedeekay/sam-forwarder/udp"
	"github.com/zieckey/goini"
)

// Conf is a tructure containing an ini config, with some functions to help
// when you use it for in conjunction with command-line flags
type Conf struct {
	config             *goini.INI
	Labels             []string
	Client             bool
	Type               string
	SaveDirectory      string
	SaveFile           bool
	TargetHost         string
	TargetPort         string
	SamHost            string
	SamPort            string
	TargetForPort443   string
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
func (c *Conf) Get(key string, label ...string) (string, bool) {
	if len(c.Labels) > 0 {
		return c.config.SectionGet(c.Labels[0], key)
	} else {
		return c.config.Get(key)
	}
}

// GetBool passes directly through to goini.GetBool
func (c *Conf) GetBool(key string, label ...string) (bool, bool) {
	if len(c.Labels) > 0 {
		return c.config.SectionGetBool(c.Labels[0], key)
	} else {
		return c.config.GetBool(key)
	}
}

// GetInt passes directly through to goini.GetInt
func (c *Conf) GetInt(key string, label ...string) (int, bool) {
	if len(c.Labels) > 0 {
		return c.config.SectionGetInt(c.Labels[0], key)
	} else {
		return c.config.GetInt(key)
	}
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
func (c *Conf) GetHost(arg, def string, label ...string) string {
	if arg != def {
		return arg
	}
	if c.config == nil {
		return arg
	}
	if x, o := c.Get("host", label...); o {
		return x
	}
	return arg
}

// GetPort takes an argument and a default. If the argument differs from the
// default, the argument is always returned. If the argument and default are
// the same and the key exists, the key is returned. If the key is absent, the
// default is returned.
func (c *Conf) GetPort(arg, def string, label ...string) string {
	if arg != def {
		return arg
	}
	if c.config == nil {
		return arg
	}
	if x, o := c.Get("port", label...); o {
		return x
	}
	return arg
}

// GetPort443 takes an argument and a default. If the argument differs from the
// default, the argument is always returned. If the argument and default are
// the same and the key exists, the key is returned. If the key is absent, the
// default is returned.
func (c *Conf) GetPort443(arg, def string, label ...string) string {
	if arg != def {
		return arg
	}
	if c.config == nil {
		return arg
	}
	if x, o := c.Get("targetForPort.443", label...); o {
		return x
	}
	return arg
}

// GetType takes an argument and a default. If the argument differs from the
// default, the argument is always returned. If the argument and default are
// the same and the key exists, the key is returned. If the key is absent, the
// default is returned.
func (c *Conf) GetType(argc, argu, argh bool, def string, label ...string) string {
	var typ string
	if argu {
		typ += "udp"
	}
	if argc {
		typ += "client"
		c.Client = true
	} else {
		if argh == true {
			typ += "http"
		} else {
			typ += "server"
		}
	}
	if typ != def {
		return typ
	}
	if c.config == nil {
		return typ
	}
	if x, o := c.Get("type", label...); o {
		return x
	}
	return def
}

// GetSAMHost takes an argument and a default. If the argument differs from the
// default, the argument is always returned. If the argument and default are
// the same and the key exists, the key is returned. If the key is absent, the
// default is returned.
func (c *Conf) GetSAMHost(arg, def string, label ...string) string {
	if arg != def {
		return arg
	}
	if c.config == nil {
		return arg
	}
	if x, o := c.Get("samhost", label...); o {
		return x
	}
	return arg
}

// GetSAMPort takes an argument and a default. If the argument differs from the
// default, the argument is always returned. If the argument and default are
// the same and the key exists, the key is returned. If the key is absent, the
// default is returned.
func (c *Conf) GetSAMPort(arg, def string, label ...string) string {
	if arg != def {
		return arg
	}
	if c.config == nil {
		return arg
	}
	if x, o := c.Get("samport", label...); o {
		return x
	}
	return arg
}

// GetDir takes an argument and a default. If the argument differs from the
// default, the argument is always returned. If the argument and default are
// the same and the key exists, the key is returned. If the key is absent, the
// default is returned.
func (c *Conf) GetDir(arg, def string, label ...string) string {
	if arg != def {
		return arg
	}
	if c.config == nil {
		return arg
	}
	if x, o := c.Get("dir", label...); o {
		return x
	}
	return arg
}

// GetSaveFile takes an argument and a default. If the argument differs from the
// default, the argument is always returned. If the argument and default are
// the same and the key exists, the key is returned. If the key is absent, the
// default is returned.
func (c *Conf) GetSaveFile(arg, def bool, label ...string) bool {
	if arg != def {
		return arg
	}
	if c.config == nil {
		return arg
	}
	return c.SaveFile
}

// GetAccessListType takes an argument and a default. If the argument differs from the
// default, the argument is always returned. If the argument and default are
// the same and the key exists, the key is returned. If the key is absent, the
// default is returned.
func (c *Conf) GetAccessListType(arg, def string, label ...string) string {
	if arg != def {
		return arg
	}
	if c.config == nil {
		return arg
	}
	return c.AccessListType
}

// GetKeys takes an argument and a default. If the argument differs from the
// default, the argument is always returned. If the argument and default are
// the same and the key exists, the key is returned. If the key is absent, the
// default is returned.
func (c *Conf) GetKeys(arg, def string, label ...string) string {
	if arg != def {
		return arg
	}
	if c.config == nil {
		return arg
	}
	if x, o := c.Get("keys", label...); o {
		return x
	}
	return arg
}

// GetInLength takes an argument and a default. If the argument differs from the
// default, the argument is always returned. If the argument and default are
// the same and the key exists, the key is returned. If the key is absent, the
// default is returned.
func (c *Conf) GetInLength(arg, def int, label ...string) int {
	if arg != def {
		return arg
	}
	if c.config == nil {
		return arg
	}
	if x, o := c.GetInt("inbound.length", label...); o {
		return x
	}
	return arg
}

// GetOutLength takes an argument and a default. If the argument differs from the
// default, the argument is always returned. If the argument and default are
// the same and the key exists, the key is returned. If the key is absent, the
// default is returned.
func (c *Conf) GetOutLength(arg, def int, label ...string) int {
	if arg != def {
		return arg
	}
	if c.config == nil {
		return arg
	}
	if x, o := c.GetInt("outbound.length", label...); o {
		return x
	}
	return arg
}

// GetInVariance takes an argument and a default. If the argument differs from the
// default, the argument is always returned. If the argument and default are
// the same and the key exists, the key is returned. If the key is absent, the
// default is returned.
func (c *Conf) GetInVariance(arg, def int, label ...string) int {
	if arg != def {
		return arg
	}
	if c.config == nil {
		return arg
	}
	if x, o := c.GetInt("inbound.variance", label...); o {
		return x
	}
	return arg
}

// GetOutVariance takes an argument and a default. If the argument differs from the
// default, the argument is always returned. If the argument and default are
// the same and the key exists, the key is returned. If the key is absent, the
// default is returned.
func (c *Conf) GetOutVariance(arg, def int, label ...string) int {
	if arg != def {
		return arg
	}
	if c.config == nil {
		return arg
	}
	if x, o := c.GetInt("outbound.variance", label...); o {
		return x
	}
	return arg
}

// GetInQuantity takes an argument and a default. If the argument differs from the
// default, the argument is always returned. If the argument and default are
// the same and the key exists, the key is returned. If the key is absent, the
// default is returned.
func (c *Conf) GetInQuantity(arg, def int, label ...string) int {
	if arg != def {
		return arg
	}
	if c.config == nil {
		return arg
	}
	if x, o := c.GetInt("inbound.quantity", label...); o {
		return x
	}
	return arg
}

// GetOutQuantity takes an argument and a default. If the argument differs from the
// default, the argument is always returned. If the argument and default are
// the same and the key exists, the key is returned. If the key is absent, the
// default is returned.
func (c *Conf) GetOutQuantity(arg, def int, label ...string) int {
	if arg != def {
		return arg
	}
	if c.config == nil {
		return arg
	}
	if x, o := c.GetInt("outbound.quantity", label...); o {
		return x
	}
	return arg
}

// GetInBackups takes an argument and a default. If the argument differs from the
// default, the argument is always returned. If the argument and default are
// the same and the key exists, the key is returned. If the key is absent, the
// default is returned.
func (c *Conf) GetInBackups(arg, def int, label ...string) int {
	if arg != def {
		return arg
	}
	if c.config == nil {
		return arg
	}
	if x, o := c.GetInt("inbound.backupQuantity", label...); o {
		return x
	}
	return arg
}

// GetOutBackups takes an argument and a default. If the argument differs from the
// default, the argument is always returned. If the argument and default are
// the same and the key exists, the key is returned. If the key is absent, the
// default is returned.
func (c *Conf) GetOutBackups(arg, def int, label ...string) int {
	if arg != def {
		return arg
	}
	if c.config == nil {
		return arg
	}
	if x, o := c.GetInt("outbound.backupQuantity", label...); o {
		return x
	}
	return arg
}

// GetEncryptLeaseset takes an argument and a default. If the argument differs from the
// default, the argument is always returned. If the argument and default are
// the same and the key exists, the key is returned. If the key is absent, the
// default is returned.
func (c *Conf) GetEncryptLeaseset(arg, def bool, label ...string) bool {
	if arg != def {
		return arg
	}
	if c.config == nil {
		return arg
	}
	if x, o := c.GetBool("i2cp.encryptLeaseSet", label...); o {
		return x
	}
	return arg
}

// GetInAllowZeroHop takes an argument and a default. If the argument differs from the
// default, the argument is always returned. If the argument and default are
// the same and the key exists, the key is returned. If the key is absent, the
// default is returned.
func (c *Conf) GetInAllowZeroHop(arg, def bool, label ...string) bool {
	if arg != def {
		return arg
	}
	if c.config == nil {
		return arg
	}
	if x, o := c.GetBool("inbound.allowZeroHop", label...); o {
		return x
	}
	return arg
}

// GetOutAllowZeroHop takes an argument and a default. If the argument differs from the
// default, the argument is always returned. If the argument and default are
// the same and the key exists, the key is returned. If the key is absent, the
// default is returned.
func (c *Conf) GetOutAllowZeroHop(arg, def bool, label ...string) bool {
	if arg != def {
		return arg
	}
	if c.config == nil {
		return arg
	}
	if x, o := c.GetBool("outbound.allowZeroHop", label...); o {
		return x
	}
	return arg
}

// GetUseCompression takes an argument and a default. If the argument differs from the
// default, the argument is always returned. If the argument and default are
// the same and the key exists, the key is returned. If the key is absent, the
// default is returned.
func (c *Conf) GetUseCompression(arg, def bool, label ...string) bool {
	if arg != def {
		return arg
	}
	if c.config == nil {
		return arg
	}
	if x, o := c.GetBool("gzip", label...); o {
		return x
	}
	return arg
}

// GetCloseOnIdle takes an argument and a default. If the argument differs from the
// default, the argument is always returned. If the argument and default are
// the same and the key exists, the key is returned. If the key is absent, the
// default is returned.
func (c *Conf) GetCloseOnIdle(arg, def bool, label ...string) bool {
	if arg != def {
		return arg
	}
	if c.config == nil {
		return arg
	}
	if x, o := c.GetBool("i2cp.closeOnIdle", label...); o {
		return x
	}
	return arg
}

// GetCloseIdleTime takes an argument and a default. If the argument differs from the
// default, the argument is always returned. If the argument and default are
// the same and the key exists, the key is returned. If the key is absent, the
// default is returned.
func (c *Conf) GetCloseIdleTime(arg, def int, label ...string) int {
	if arg != def {
		return arg
	}
	if c.config == nil {
		return arg
	}
	if x, o := c.GetInt("i2cp.closeIdleTime", label...); o {
		return x
	}
	return arg
}

// GetReduceOnIdle takes an argument and a default. If the argument differs from the
// default, the argument is always returned. If the argument and default are
// the same and the key exists, the key is returned. If the key is absent, the
// default is returned.
func (c *Conf) GetReduceOnIdle(arg, def bool, label ...string) bool {
	if arg != def {
		return arg
	}
	if c.config == nil {
		return arg
	}
	if x, o := c.GetBool("i2cp.reduceOnIdle", label...); o {
		return x
	}
	return arg
}

// GetReduceIdleTime takes an argument and a default. If the argument differs from the
// default, the argument is always returned. If the argument and default are
// the same and the key exists, the key is returned. If the key is absent, the
// default is returned.
func (c *Conf) GetReduceIdleTime(arg, def int, label ...string) int {
	if arg != def {
		return arg
	}
	if c.config == nil {
		return arg
	}
	if x, o := c.GetInt("i2cp.reduceIdleTime", label...); o {
		return x
	}
	return arg
}

// GetReduceIdleQuantity takes an argument and a default. If the argument differs from the
// default, the argument is always returned. If the argument and default are
// the same and the key exists, the key is returned. If the key is absent, the
// default is returned.
func (c *Conf) GetReduceIdleQuantity(arg, def int, label ...string) int {
	if arg != def {
		return arg
	}
	if c.config == nil {
		return arg
	}
	if x, o := c.GetInt("i2cp.reduceIdleQuantity", label...); o {
		return x
	}
	return arg
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

// SetDir sets the key save directory from the config file
func (c *Conf) SetDir(label ...string) {
	if len(label) < 1 {
		if v, ok := c.Get("dir"); ok {
			c.SaveDirectory = v
		} else {
			c.SaveDirectory = "./"
		}
	}
}

// SetKeys sets the key name from the config file
func (c *Conf) SetKeys(label ...string) {
	if len(label) < 1 {
		if _, ok := c.Get("keys"); ok {
			c.SaveFile = true
		} else {
			c.SaveFile = false
		}
	}
}

// SetType sets the type of proxy to create from the config file
func (c *Conf) SetType(label ...string) {
	if len(label) < 1 {
		if v, ok := c.Get("type"); ok {
			if strings.Contains(v, "client") {
				c.Client = true
			}
			if c.Type == "server" || c.Type == "http" || c.Type == "client" || c.Type == "udpserver" || c.Type == "udpclient" {
				c.Type = v
			}
		} else {
			c.Type = "server"
		}
	}
}

// SetHost sets the host to forward from the config file
func (c *Conf) SetHost(label ...string) {
	if len(label) < 1 {
		if v, ok := c.Get("host"); ok {
			c.TargetHost = v
		} else {
			c.TargetHost = "127.0.0.1"
		}
	}
}

// SetPort sets the port to forward from the config file
func (c *Conf) SetPort(label ...string) {
	if len(label) < 1 {
		if v, ok := c.Get("port"); ok {
			c.TargetPort = v
		} else {
			c.TargetPort = "8081"
		}
	}
}

// SetTargetPort443 sets the port to forward from the config file
func (c *Conf) SetTargetPort443(label ...string) {
	if len(label) < 1 {
		if v, ok := c.Get("targetForPort.443"); ok {
			c.TargetForPort443 = v
		} else {
			c.TargetForPort443 = ""
		}
	}
}

// SetSAMHost sets the SAM host from the config file
func (c *Conf) SetSAMHost(label ...string) {
	if len(label) < 1 {
		if v, ok := c.Get("samhost"); ok {
			c.SamHost = v
		} else {
			c.SamHost = "127.0.0.1"
		}
	}
}

// SetSAMPort sets the SAM port from the config file
func (c *Conf) SetSAMPort(label ...string) {
	if len(label) < 1 {
		if v, ok := c.Get("samport"); ok {
			c.SamPort = v
		} else {
			c.SamPort = "7656"
		}
	}
}

// SetTunName sets the tunnel name from the config file
func (c *Conf) SetTunName(label ...string) {
	if len(label) < 1 {
		if v, ok := c.Get("keys"); ok {
			c.TunName = v
		} else {
			c.TunName = "fowarder"
		}
	}
}

// SetEncryptLease tells the conf to use encrypted leasesets the from the config file
func (c *Conf) SetEncryptLease(label ...string) {
	if len(label) < 1 {
		if v, ok := c.GetBool("i2cp.encryptLeaseSet"); ok {
			c.EncryptLeaseSet = v
		} else {
			c.EncryptLeaseSet = false
		}
	}
}

// SetAllowZeroHopIn sets the config to allow zero-hop tunnels
func (c *Conf) SetAllowZeroHopIn(label ...string) {
	if len(label) < 1 {
		if v, ok := c.GetBool("inbound.allowZeroHop"); ok {
			c.InAllowZeroHop = v
		} else {
			c.InAllowZeroHop = false
		}
	}
}

// SetAllowZeroHopOut sets the config to allow zero-hop tunnels
func (c *Conf) SetAllowZeroHopOut(label ...string) {
	if len(label) < 1 {
		if v, ok := c.GetBool("outbound.allowZeroHop"); ok {
			c.OutAllowZeroHop = v
		} else {
			c.OutAllowZeroHop = false
		}
	}
}

// SetInLength sets the inbound length from the config file
func (c *Conf) SetInLength(label ...string) {
	if len(label) < 1 {
		if v, ok := c.GetInt("outbound.length"); ok {
			c.OutLength = v
		} else {
			c.OutLength = 3
		}
	}
}

// SetOutLength sets the outbound lenth from the config file
func (c *Conf) SetOutLength(label ...string) {
	if len(label) < 1 {
		if v, ok := c.GetInt("inbound.length"); ok {
			c.InLength = v
		} else {
			c.InLength = 3
		}
	}
}

// SetInQuantity sets the inbound tunnel quantity from config file
func (c *Conf) SetInQuantity(label ...string) {
	if len(label) < 1 {
		if v, ok := c.GetInt("inbound.quantity"); ok {
			c.InQuantity = v
		} else {
			c.InQuantity = 5
		}
	}
}

// SetOutQuantity sets the outbound tunnel quantity from config file
func (c *Conf) SetOutQuantity(label ...string) {
	if len(label) < 1 {
		if v, ok := c.GetInt("outbound.quantity"); ok {
			c.OutQuantity = v
		} else {
			c.OutQuantity = 5
		}
	}
}

// SetInVariance sets the inbound tunnel variance from config file
func (c *Conf) SetInVariance(label ...string) {
	if len(label) < 1 {
		if v, ok := c.GetInt("inbound.variance"); ok {
			c.InVariance = v
		} else {
			c.InVariance = 0
		}
	}
}

// SetOutVariance sets the outbound tunnel variance from config file
func (c *Conf) SetOutVariance(label ...string) {
	if len(label) < 1 {
		if v, ok := c.GetInt("outbound.variance"); ok {
			c.OutVariance = v
		} else {
			c.OutVariance = 0
		}
	}
}

// SetInBackups sets the inbound tunnel backups from config file
func (c *Conf) SetInBackups(label ...string) {
	if len(label) < 1 {
		if v, ok := c.GetInt("inbound.backupQuantity"); ok {
			c.InBackupQuantity = v
		} else {
			c.InBackupQuantity = 2
		}
	}
}

// SetOutBackups sets the outbound tunnel backups from config file
func (c *Conf) SetOutBackups(label ...string) {
	if len(label) < 1 {
		if v, ok := c.GetInt("outbound.backupQuantity"); ok {
			c.OutBackupQuantity = v
		} else {
			c.OutBackupQuantity = 2
		}
	}
}

// SetCompressed sets the compression from the config file
func (c *Conf) SetCompressed(label ...string) {
	if len(label) < 1 {
		if v, ok := c.GetBool("gzip"); ok {
			c.UseCompression = v
		} else {
			c.UseCompression = true
		}
	}
}

// SetReduceIdle sets the config to reduce tunnels after idle time from config file
func (c *Conf) SetReduceIdle(label ...string) {
	if len(label) < 1 {
		if v, ok := c.GetBool("i2cp.reduceOnIdle"); ok {
			c.ReduceIdle = v
		} else {
			c.ReduceIdle = false
		}
	}
}

// SetReduceIdleTime sets the time to wait before reducing tunnels from config file
func (c *Conf) SetReduceIdleTime(label ...string) {
	if len(label) < 1 {
		if v, ok := c.GetInt("i2cp.reduceIdleTime"); ok {
			c.ReduceIdleTime = v
		} else {
			c.ReduceIdleTime = 300000
		}
	}
}

// SetReduceIdleQuantity sets the number of tunnels to reduce to from config file
func (c *Conf) SetReduceIdleQuantity(label ...string) {
	if len(label) < 1 {
		if v, ok := c.GetInt("i2cp.reduceQuantity"); ok {
			c.ReduceIdleQuantity = v
		} else {
			c.ReduceIdleQuantity = 3
		}
	}
}

// SetCloseIdle sets the tunnel to automatically close on idle from the config file
func (c *Conf) SetCloseIdle(label ...string) {
	if len(label) < 1 {
		if v, ok := c.GetBool("i2cp.closeOnIdle"); ok {
			c.CloseIdle = v
		} else {
			c.CloseIdle = false
		}
	}
}

// SetCloseIdleTime sets the time to wait before killing a tunnel from a config file
func (c *Conf) SetCloseIdleTime(label ...string) {
	if len(label) < 1 {
		if v, ok := c.GetInt("i2cp.closeIdleTime"); ok {
			c.CloseIdleTime = v
		} else {
			c.CloseIdleTime = 300000
		}
	}
}

// SetAccessListType sets the access list type from a config file
func (c *Conf) SetAccessListType(label ...string) {
	if len(label) < 1 {
		if v, ok := c.GetBool("i2cp.enableBlackList"); ok {
			if v {
				c.AccessListType = "blacklist"
			}
		}
		if v, ok := c.GetBool("i2cp.enableAccessList"); ok {
			if v {
				c.AccessListType = "whitelist"
			}
		}
		if c.AccessListType != "whitelist" && c.AccessListType != "blacklist" {
			c.AccessListType = "none"
		}
	}
}

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
	if iniFile != "none" {
		c.config = goini.New()
		err = c.config.ParseFile(iniFile)
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
		c.SetCompressed(label...)
		c.SetReduceIdle(label...)
		c.SetReduceIdleTime(label...)
		c.SetReduceIdleQuantity(label...)
		c.SetCloseIdle(label...)
		c.SetCloseIdleTime(label...)
		c.SetAccessListType(label...)
		c.SetTargetPort443(label...)

		if v, ok := c.Get("i2cp.accessList"); ok {
			csv := strings.Split(v, ",")
			for _, z := range csv {
				c.AccessList = append(c.AccessList, z)
			}
		}
		log.Println(c.Print())
	}
	return nil
}

// NewI2PBlankTunConf returns an empty but intialized tunconf
func NewI2PBlankTunConf() *Conf {
	var c Conf
	return &c
}

// NewI2PTunConf returns a Conf structure from an ini file, for modification
// before starting the tunnel
func NewI2PTunConf(iniFile string) (*Conf, error) {
	var err error
	var c Conf
	if err = c.I2PINILoad(iniFile); err != nil {
		return nil, err
	}
	return &c, nil
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
			samforwarder.SetAllowZeroIn(config.InAllowZeroHop),
			samforwarder.SetAllowZeroOut(config.OutAllowZeroHop),
			samforwarder.SetCompress(config.UseCompression),
			samforwarder.SetReduceIdle(config.ReduceIdle),
			samforwarder.SetReduceIdleTimeMs(config.ReduceIdleTime),
			samforwarder.SetReduceIdleQuantity(config.ReduceIdleQuantity),
			samforwarder.SetCloseIdle(config.CloseIdle),
			samforwarder.SetCloseIdleTimeMs(config.CloseIdleTime),
			samforwarder.SetAccessListType(config.AccessListType),
			samforwarder.SetAccessList(config.AccessList),
			//samforwarder.SetTargetForPort443(config.TargetForPort443),
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
			samforwarder.SetClientAllowZeroIn(config.InAllowZeroHop),
			samforwarder.SetClientAllowZeroOut(config.OutAllowZeroHop),
			samforwarder.SetClientCompress(config.UseCompression),
			samforwarder.SetClientReduceIdle(config.ReduceIdle),
			samforwarder.SetClientReduceIdleTimeMs(config.ReduceIdleTime),
			samforwarder.SetClientReduceIdleQuantity(config.ReduceIdleQuantity),
			samforwarder.SetClientCloseIdle(config.CloseIdle),
			samforwarder.SetClientCloseIdleTimeMs(config.CloseIdleTime),
			samforwarder.SetClientAccessListType(config.AccessListType),
			samforwarder.SetClientAccessList(config.AccessList),
		)
	}
	return nil, nil
}

// NewSAMClientForwarderFromConfig generates a new SAMForwarder from a config file
func NewSAMClientForwarderFromConfig(iniFile, SamHost, SamPort string) (*samforwarder.SAMClientForwarder, error) {
	if iniFile != "none" {
		config, err := NewI2PTunConf(iniFile)
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
			samforwarderudp.SetAllowZeroIn(config.InAllowZeroHop),
			samforwarderudp.SetAllowZeroOut(config.OutAllowZeroHop),
			samforwarderudp.SetCompress(config.UseCompression),
			samforwarderudp.SetReduceIdle(config.ReduceIdle),
			samforwarderudp.SetReduceIdleTimeMs(config.ReduceIdleTime),
			samforwarderudp.SetReduceIdleQuantity(config.ReduceIdleQuantity),
			samforwarderudp.SetCloseIdle(config.CloseIdle),
			samforwarderudp.SetCloseIdleTimeMs(config.CloseIdleTime),
			samforwarderudp.SetAccessListType(config.AccessListType),
			samforwarderudp.SetAccessList(config.AccessList),
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
			samforwarderudp.SetClientAllowZeroIn(config.InAllowZeroHop),
			samforwarderudp.SetClientAllowZeroOut(config.OutAllowZeroHop),
			samforwarderudp.SetClientCompress(config.UseCompression),
			samforwarderudp.SetClientReduceIdle(config.ReduceIdle),
			samforwarderudp.SetClientReduceIdleTimeMs(config.ReduceIdleTime),
			samforwarderudp.SetClientReduceIdleQuantity(config.ReduceIdleQuantity),
			samforwarderudp.SetClientCloseIdle(config.CloseIdle),
			samforwarderudp.SetClientCloseIdleTimeMs(config.CloseIdleTime),
			samforwarderudp.SetClientAccessListType(config.AccessListType),
			samforwarderudp.SetClientAccessList(config.AccessList),
		)
	}
	return nil, nil
}

// NewSAMSSUClientForwarderFromConfig generates a new SAMSSUForwarder from a config file
func NewSAMSSUClientForwarderFromConfig(iniFile, SamHost, SamPort string) (*samforwarderudp.SAMSSUClientForwarder, error) {
	if iniFile != "none" {
		config, err := NewI2PTunConf(iniFile)
		if err != nil {
			return nil, err
		}
		if SamHost != "" && SamHost != "127.0.0.1" && SamHost != "localhost" {
			config.SamHost = config.GetSAMHost(SamHost, config.SamHost)
		}
		if SamPort != "" && SamPort != "7656" {
			config.SamPort = config.GetSAMPort(SamPort, config.SamPort)
		}
		return NewSAMSSUClientForwarderFromConf(config)
	}
	return nil, nil
}
