package samtunnel

import (
	"github.com/eyedeekay/sam3/i2pkeys"
)

type SAMTunnel interface {

	// Tunnel Options
	GetType() string             // Get the type of the tunnel in use(server, client, http, udp, etc)
	Print() string               // Print all the tunnel options as a string
	Props() map[string]string    //Get a full list of tunnel properties as a map for user display/analysis
	Search(search string) string //Search the Props for a common term
	Target() string              //The address of the local client or service to forward with a SAM tunnel
	ID() string                  //The user-chosen tunnel ID
	//Destination() string

	// Key handling
	Base32() string         // Get the .b32.i2p address of your service
	Base32Readable() string // Create a more-readable representation of the .b32.i2p address using English words
	Base64() string         // Get the public base64 address of your I2P service
	Keys() i2pkeys.I2PKeys  // Get all the parts of the keys to your I2P service

	// Service Management
	Load() (SAMTunnel, error) // Prepare tunnel keys and tunnel options
	Serve() error             // Start the tunnel
	Close() error             // Stop the tunnel and close all connections
	Cleanup()                 // Stop the tunnel but leave the sockets alone for now
	Up() bool                 // Return "true" if the tunnel is ready to go up.
}

type WebUI interface {
	Title() string
	URL() string
	UseWebUI() bool
	Width() int
	Height() int
	Resizable() bool
}
