package i2pconfig

import (
	"strconv"
	"strings"
)

func ConstructEqualsConfigMap(options []string) map[string]string {
	configMap := make(map[string]string)
	for _, v := range options {
		val := strings.SplitN(v, "=", 2)
		if len(val) == 2 {
			configMap[val[0]] = val[1]
		}
	}
	return configMap
}

func StringOrNot(check map[string]string, key string, def string) string {
	if val, ok := check[key]; ok {
		return val
	}
	return def
}

func StringSliceOrNot(check map[string]string, key string, def string) []string {
	if val, ok := check[key]; ok {
		return strings.Split(val, ",")
	}
	return strings.Split(def, ",")
}

func IntOrNot(check map[string]string, key string, def int) int {
	if val, ok := check[key]; ok {
		if rval, err := strconv.Atoi(val); err == nil {
			return rval
		}
	}
	return def
}

func BoolOrNot(check map[string]string, key string, def bool) bool {
	if val, ok := check[key]; ok {
		if rval, err := strconv.ParseBool(val); err == nil {
			return rval
		}
	}
	return def
}

func ConstructEqualsConfig(options []string) (*I2PConfig, error) {
	confmap := ConstructEqualsConfigMap(options)
	accessListType := "none"
	if BoolOrNot(confmap, "i2cp.enableBlackList", false) {
		accessListType = "blacklist"
	}
	if BoolOrNot(confmap, "i2cp.enableWhiteList", false) {
		accessListType = "whitelist"
	}
	return NewConfig(
		SetTunType(StringOrNot(confmap, "type", "server")),
		SetConfigStyle(StringOrNot(confmap, "STYLE", "STREAM")),
		SetConfigSAMHost(StringOrNot(confmap, "samhost", "127.0.0.1")),
		SetConfigSAMPort(StringOrNot(confmap, "samport", "7657")),
		SetConfigFromPort(StringOrNot(confmap, "TO_PORT", "0")),
		SetConfigToPort(StringOrNot(confmap, "FROM_PORT", "0")),
		SetConfigName(StringOrNot(confmap, "keys", "unnamedunsafe")),
		SetConfigInLength(IntOrNot(confmap, "inbound.length", 2)),
		SetConfigOutLength(IntOrNot(confmap, "outbound.length", 2)),
		SetConfigInVariance(IntOrNot(confmap, "inbound.variance", 1)),
		SetConfigOutVariance(IntOrNot(confmap, "outbound.variance", 1)),
		SetConfigInQuantity(IntOrNot(confmap, "outbound.quantity", 1)),
		SetConfigOutQuantity(IntOrNot(confmap, "inbound.quantity", 1)),
		SetConfigInBackups(IntOrNot(confmap, "inbound.backups", 1)),
		SetConfigOutBackups(IntOrNot(confmap, "outbound.backups", 1)),
		SetConfigEncrypt(BoolOrNot(confmap, "i2cp.encryptLeaseSet", false)),
		SetConfigLeaseSetKey(StringOrNot(confmap, "i2cp.leaseSetKey", "")),
		SetConfigLeaseSetPrivateKey(StringOrNot(confmap, "i2cp.leaseSetPrivateKey", "")),
		SetConfigLeaseSetPrivateSigningKey(StringOrNot(confmap, "i2cp.leaseSetPrivateSigningKey", "")),
		SetConfigMessageReliability(StringOrNot(confmap, "i2cp.messageReliability", "")),
		SetConfigAllowZeroIn(BoolOrNot(confmap, "inbound.allowZeroHop", false)),
		SetConfigAllowZeroOut(BoolOrNot(confmap, "outbound.allowZeroHop", false)),
		SetConfigCompress(BoolOrNot(confmap, "gzip", true)),
		SetConfigFastRecieve(BoolOrNot(confmap, "i2cp.fastRecieve", false)),
		SetConfigReduceIdle(BoolOrNot(confmap, "i2cp.reduceIdle", false)),
		SetConfigReduceIdleTimeMs(IntOrNot(confmap, "i2cp.reduceIdleTime", 600000)),
		SetConfigReduceIdleQuantity(IntOrNot(confmap, "i2cp.reduceIdleQuantity", 1)),
		SetConfigCloseIdle(BoolOrNot(confmap, "i2cp.closeIdle", false)),
		SetConfigCloseIdleTimeMs(IntOrNot(confmap, "i2cp.closeIdleTime", 600000)),
		SetConfigAccessListType(accessListType),
		SetConfigAccessList(StringSliceOrNot(confmap, "i2cp.accessList", "")),
	)
}
