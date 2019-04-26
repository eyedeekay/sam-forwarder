package samtunnel

type SAMTunnel interface {
	//GetType() string
	Cleanup()
	Print() string
	Props() map[string]string
	Search(search string) string
	Target() string
	ID() string
	//Destination() string
	Base32() string
	Base64() string
	Serve() error
	Close() error
}
