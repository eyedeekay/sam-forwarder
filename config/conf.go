package i2ptunconf

import (
	"fmt"
	"strings"
)

import (
	sfi2pkeys "github.com/eyedeekay/sam-forwarder/i2pkeys"
	"github.com/eyedeekay/sam3/i2pkeys"
)

var err error

func (f *Conf) ID() string {
	return f.TunName
}

func (f *Conf) Keys() i2pkeys.I2PKeys {
	return f.LoadedKeys
}

func (f *Conf) Cleanup() {

}

func (f *Conf) GetType() string {
	return f.Type
}

/*func (f *Conf) targetForPort443() string {
	if f.TargetForPort443 != "" {
		return "targetForPort.4443=" + f.TargetHost + ":" + f.TargetForPort443
	}
	return ""
}*/

func (f *Conf) print() []string {
	lsk, lspk, lspsk := f.leasesetsettings()
	return []string{
		//f.targetForPort443(),
		"inbound.length=" + fmt.Sprintf("%d", f.InLength),
		"outbound.length=" + fmt.Sprintf("%d", f.OutLength),
		"inbound.lengthVariance=" + fmt.Sprintf("%d", f.InVariance),
		"outbound.lengthVariance=" + fmt.Sprintf("%d", f.OutVariance),
		"inbound.backupQuantity=" + fmt.Sprintf("%d", f.InBackupQuantity),
		"outbound.backupQuantity=" + fmt.Sprintf("%d", f.OutBackupQuantity),
		"inbound.quantity=" + fmt.Sprintf("%d", f.InQuantity),
		"outbound.quantity=" + fmt.Sprintf("%d", f.OutQuantity),
		"inbound.allowZeroHop=" + fmt.Sprintf("%t", f.InAllowZeroHop),
		"outbound.allowZeroHop=" + fmt.Sprintf("%t", f.OutAllowZeroHop),
		"i2cp.fastRecieve=" + fmt.Sprintf("%t", f.FastRecieve),
		"i2cp.gzip=" + fmt.Sprintf("%t", f.UseCompression),
		"i2cp.reduceOnIdle=" + fmt.Sprintf("%t", f.ReduceIdle),
		"i2cp.reduceIdleTime=" + fmt.Sprintf("%d", f.ReduceIdleTime),
		"i2cp.reduceQuantity=" + fmt.Sprintf("%d", f.ReduceIdleQuantity),
		"i2cp.closeOnIdle=" + fmt.Sprintf("%t", f.CloseIdle),
		"i2cp.closeIdleTime=" + fmt.Sprintf("%d", f.CloseIdleTime),
		"i2cp.messageReliability=" + f.MessageReliability,
		"i2cp.encryptLeaseSet=" + fmt.Sprintf("%t", f.EncryptLeaseSet),
		lsk, lspk, lspsk,
		f.accesslisttype(),
		f.accesslist(),
	}
}

func (f *Conf) Props() map[string]string {
	r := make(map[string]string)
	print := f.print()
	print = append(print, "base32="+f.Base32())
	print = append(print, "base64="+f.Base64())
	print = append(print, "base32words="+f.Base32Readable())
	for _, prop := range print {
		k, v := sfi2pkeys.Prop(prop)
		r[k] = v
	}
	return r
}

func (f *Conf) Print() string {
	var r string
	r += "name=" + f.TunName + "\n"
	r += "type=" + f.Type + "\n"
	if f.Type == "http" {
		r += "httpserver\n"
	} else {
		r += "ntcpserver\n"
	}
	for _, s := range f.print() {
		r += s + "\n"
	}
	return strings.Replace(r, "\n\n", "\n", -1)
}

func (f *Conf) Search(search string) string {
	terms := strings.Split(search, ",")
	if search == "" {
		return f.Print()
	}
	for _, value := range terms {
		if !strings.Contains(f.Print(), value) {
			return ""
		}
	}
	return f.Print()
}

/*
func (f *Conf) accesslisttype() string {
	if f.accessListType == "whitelist" {
		return "i2cp.enableAccessList=true"
	} else if f.accessListType == "blacklist" {
		return "i2cp.enableBlackList=true"
	} else if f.accessListType == "none" {
		return ""
	}
	return ""
}

func (f *Conf) accesslist() string {
	if f.accessListType != "" && len(f.accessList) > 0 {
		r := ""
		for _, s := range f.accessList {
			r += s + ","
		}
		return "i2cp.accessList=" + strings.TrimSuffix(r, ",")
	}
	return ""
}
*/
func (f *Conf) leasesetsettings() (string, string, string) {
	var r, s, t string
	if f.LeaseSetKey != "" {
		r = "i2cp.leaseSetKey=" + f.LeaseSetKey
	}
	if f.LeaseSetPrivateKey != "" {
		s = "i2cp.leaseSetPrivateKey=" + f.LeaseSetPrivateKey
	}
	if f.LeaseSetPrivateSigningKey != "" {
		t = "i2cp.leaseSetPrivateSigningKey=" + f.LeaseSetPrivateSigningKey
	}
	return r, s, t
}

// Target returns the host:port of the local service you want to forward to i2p
func (f *Conf) Target() string {
	return f.TargetHost + ":" + f.TargetPort
}

func (f *Conf) sam() string {
	return f.SamHost + ":" + f.SamPort
}

//Base32 returns the base32 address where the local service is being forwarded
func (f *Conf) Base32() string {
	return f.LoadedKeys.Addr().Base32()
}

//Base32Readable will always be an empty string when used here.
func (f *Conf) Base32Readable() string {
	return ""
}

//Base64 returns the base64 address where the local service is being forwarded
func (f *Conf) Base64() string {
	return f.LoadedKeys.Addr().Base64()
}

//Serve starts the SAM connection and and forwards the local host:port to i2p
func (f *Conf) Serve() error {
	return nil
}

func (f *Conf) Up() bool {
	return false
}

//Close shuts the whole thing down.
func (f *Conf) Close() error {
	return nil
}
