package samtunnel

import (
	"github.com/eyedeekay/i2pkeys"
	"github.com/eyedeekay/sam-forwarder/config"
)

// SAMTunnel is an interface comprehensively representing an I2P tunnel over SAM
// in Go
type SAMTunnel interface {
	// Config returns the appropriate underlying config object all options, or
	// the common options passed into a compound tunnel.
	Config() *i2ptunconf.Conf
	// Tunnel Options
	// GetType Get the type of the tunnel in use(server, client, http, udp, etc)
	GetType() string
	// Print all the tunnel options as a string
	Print() string
	// Props Get a full list of tunnel properties as a map for user display/analysis
	Props() map[string]string
	//Search the Props for a common term
	Search(search string) string
	//Target The address of the local client or service to forward with a SAM tunnel
	Target() string
	//ID The user-chosen tunnel name
	ID() string
	//Destination() string

	// Key handling
	// Get the .b32.i2p address of your service
	Base32() string
	// Create a more-readable representation of the .b32.i2p address using English words
	Base32Readable() string
	// Get the public base64 address of your I2P service
	Base64() string
	// Get all the parts of the keys to your I2P service
	Keys() i2pkeys.I2PKeys

	// Service Management
	// Prepare tunnel keys and tunnel options
	Load() (SAMTunnel, error)
	// Start the tunnel
	Serve() error
	// Stop the tunnel and close all connections
	Close() error
	// Stop the tunnel but leave the sockets alone for now
	Cleanup()
	// Return "true" if the tunnel is ready to go up.
	Up() bool
}

// WebUI is an interface which is used to generate a minimal UI. Open to suggestions.
type WebUI interface {
	Title() string
	URL() string
	UseWebUI() bool
	Width() int
	Height() int
	Resizable() bool
}
