package samtunnel

type SAMTunnel interface {
	Cleanup()
	Print() string
	Search(search string) string
	Target() string
	Destination() string
	Base32() string
	Base64() string
	Serve() error
	Close() error
}
