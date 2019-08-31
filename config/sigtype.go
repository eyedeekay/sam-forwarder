package i2ptunconf



// GetSigType takes an argument and a default. If the argument differs from the
// default, the argument is always returned. If the argument and default are
// the same and the key exists, the key is returned. If the key is absent, the
// default is returned.
func (c *Conf) GetSigType(argt, def string, label ...string) string {
	var typ string
	if argt == "" {
		typ = ""
	} else if argt == "DSA_SHA1" {
		typ = "DSA_SHA1"
	} else if argt == "ECDSA_SHA256_P256" {
		typ = "ECDSA_SHA256_P256"
	} else if argt == "ECDSA_SHA384_P384" {
		typ = "ECDSA_SHA384_P384"
	} else if argt == "ECDSA_SHA512_P521" {
		typ = "ECDSA_SHA512_P521"
	} else if argt == "EdDSA_SHA512_Ed25519" {
		typ = "EdDSA_SHA512_Ed25519"
	} else {
		typ = "EdDSA_SHA512_Ed25519"
	}
	if typ != def {
		return typ
	}
	if c.Config == nil {
		return typ
	}
	if x, o := c.Get("signaturetype", label...); o {
		return x
	}
	return def
}

// SetSigType sets the type of proxy to create from the config file
func (c *Conf) SetSigType(label ...string) {
	if v, ok := c.Get("signaturetype", label...); ok {
		if c.SigType == "" || c.SigType == "DSA_SHA1" || c.SigType == "ECDSA_SHA256_P256" || c.SigType == "ECDSA_SHA384_P384" || c.SigType == "ECDSA_SHA512_P521" || c.SigType == "EdDSA_SHA512_Ed25519" {
			c.SigType = v
		}
	} else {
		c.SigType = "EdDSA_SHA512_Ed25519"
	}
}
