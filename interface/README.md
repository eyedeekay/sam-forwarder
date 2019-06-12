Implementing the sam-forwarder interface
========================================

The sam-forwrder interface(used int the Go sense of the word interface) is used
to create custom types of tunnels. It's kind of big, and maybe too complex, so
subject to change.

``` go
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
```
