package samtunnel

import (
	"github.com/eyedeekay/sam3/i2pkeys"
)

type SAMTunnel interface {
	GetType() string
	Cleanup()
	Print() string
	Props() map[string]string
	Search(search string) string
	Target() string
	ID() string
	//Destination() string
	Base32() string
	Base64() string
	Keys() i2pkeys.I2PKeys
	Serve() error
	Close() error
	Up() bool
	Load() (SAMTunnel, error)
}
