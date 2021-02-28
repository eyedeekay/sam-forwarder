package i2ptunconf

import (
	"crypto/tls"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

import (
	//  "github.com/eyedeekay/sam3"
	"github.com/eyedeekay/sam3/i2pkeys"
	"github.com/zieckey/goini"
)

// Conf is a tructure containing an ini config, with some functions to help
// when you use it for in conjunction with command-line flags
type Conf struct {
	Config                    *goini.INI `default:&goini.INI{}`
	FilePath                  string     `default:"./"`
	KeyFilePath               string     `default:"./"`
	Labels                    []string   `default:{""}`
	Client                    bool       `default:true`
	ClientDest                string     `default:"idk.i2p"`
	SigType                   string     `default:"SIGNATURE_TYPE=EdDSA_SHA512_Ed25519"`
	Type                      string     `default:"client"`
	SaveDirectory             string     `default:"./"`
	ServeDirectory            string     `default:"./www"`
	SaveFile                  bool       `default:false`
	TargetHost                string     `default:"127.0.0.1"`
	TargetPort                string     `default:"7778"`
	SamHost                   string     `default:"127.0.0.1"`
	SamPort                   string     `default:"7656"`
	TunnelHost                string     `default:"127.0.0.1"`
	ControlHost               string     `default:"127.0.0.1"`
	ControlPort               string     `default:"7951"`
	TargetForPort443          string     `default:""`
	TunName                   string     `default:"goi2ptunnel"`
	EncryptLeaseSet           bool       `default:false`
	LeaseSetKey               string     `default:""`
	LeaseSetEncType           string     `default:"4,0"`
	LeaseSetPrivateKey        string     `default:""`
	LeaseSetPrivateSigningKey string     `default:""`
	InAllowZeroHop            bool       `default:false`
	OutAllowZeroHop           bool       `default:false`
	InLength                  int        `default:3`
	OutLength                 int        `default:3`
	InQuantity                int        `default:1`
	OutQuantity               int        `default:1`
	InVariance                int        `default:0`
	OutVariance               int        `default:0`
	InBackupQuantity          int        `default:1`
	OutBackupQuantity         int        `default:1`
	UseCompression            bool       `default:true`
	FastRecieve               bool       `default:true`
	ReduceIdle                bool       `default:false`
	ReduceIdleTime            int        `default:36000000`
	ReduceIdleQuantity        int        `default:1`
	CloseIdle                 bool       `default:false`
	CloseIdleTime             int        `default:36000000`
	AccessListType            string     `default:"none"`
	AccessList                []string   `default:{""}`
	MessageReliability        string     `default:""`
	exists                    bool       `default:false`
	UserName                  string     `default:""`
	Password                  string     `default:""`
	UseTLS                    bool       `default:false`
	TLSConf                   *tls.Config
	LoadedKeys                i2pkeys.I2PKeys
}

// PrintSlice returns and prints a formatted list of configured tunnel settings.
func (c *Conf) PrintSlice() []string {
	confstring := []string{
		c.SignatureType(),
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
		"i2cp.fastRecieve=" + strconv.FormatBool(c.FastRecieve),
		c.reliability(),
		c.accesslisttype(),
		c.accesslist(),
		c.lsk(),
		c.lspk(),
		c.lsspk(),
	}

	log.Println("Tunnel:", c.TunName, "using config:", confstring)
	return confstring
}

func (c *Conf) lsk() string {
	if c.LeaseSetKey != "" {
		return "i2cp.leaseSetKey=" + c.LeaseSetKey
	}
	return ""
}

func (c *Conf) lspk() string {
	if c.LeaseSetPrivateKey != "" {
		return "i2cp.leaseSetPrivateKey=" + c.LeaseSetPrivateKey
	}
	return ""
}

func (c *Conf) lsspk() string {
	if c.LeaseSetPrivateSigningKey != "" {
		return "i2cp.leaseSetSigningPrivateKey=" + c.LeaseSetPrivateSigningKey
	}
	return ""
}

func (c *Conf) SignatureType() string {
	if c.SigType == "" {
		return ""
	}
	return "SIGNATURE_TYPE=" + c.SigType
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
		c.SetEndpointHost(label...)
		c.SetTunName(label...)
		c.SetSigType(label...)
		c.SetEncryptLease(label...)
		c.SetLeaseSetEncType(label...)
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
		c.SetUserName(label...)
		c.SetPassword(label...)
		c.SetControlHost(label...)
		c.SetControlPort(label...)
		c.SetWWWDir(label...)
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
	//	var c Conf
	c := new(Conf)
	c.SamHost = "127.0.0.1"
	c.SamPort = "7656"
	c.TunName = "unksam"
	c.TargetHost = "127.0.0.1"
	c.TargetPort = "0"
	c.ClientDest = "idk.i2p"
	c.LeaseSetEncType = "4,0"
	c.Config = &goini.INI{}
	c.Config = goini.New()
	c.Config.Parse([]byte("[client]\nsamhost=\"127.0.0.1\"\nsamport=\"7656\"\n"), "\n", "=")
	return c
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
