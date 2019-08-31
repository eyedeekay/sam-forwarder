package sam3

import (
	"fmt"
	"net"
	"strconv"

	"github.com/eyedeekay/sam3/i2pkeys"
)

// sam config

// options map
type Options map[string]string

// obtain sam options as list of strings
func (opts Options) AsList() (ls []string) {
	for k, v := range opts {
		ls = append(ls, fmt.Sprintf("%s=%s", k, v))
	}
	return
}

// Config is the config type for the sam connector api for i2p which allows applications to 'speak' with i2p
type Config struct {
	Addr    string
	Opts    Options
	Session string
	Keyfile string
}

// create new sam connector from config with a stream session
func (cfg *Config) StreamSession() (session *StreamSession, err error) {
	// connect
	var s *SAM
	s, err = NewSAM(cfg.Addr)
	if err == nil {
		// ensure keys exist
		var keys i2pkeys.I2PKeys
		keys, err = s.EnsureKeyfile(cfg.Keyfile)
		if err == nil {
			// create session
			session, err = s.NewStreamSession(cfg.Session, keys, cfg.Opts.AsList())
		}
	}
	return
}

// create new sam datagram session from config
func (cfg *Config) DatagramSession() (session *DatagramSession, err error) {
	// connect
	var s *SAM
	s, err = NewSAM(cfg.Addr)
	if err == nil {
		// ensure keys exist
		var keys i2pkeys.I2PKeys
		keys, err = s.EnsureKeyfile(cfg.Keyfile)
		if err == nil {
			// determine udp port
			var portstr string
			_, portstr, err = net.SplitHostPort(cfg.Addr)
			if err == nil {
				var port int
				port, err = strconv.Atoi(portstr)
				if err == nil && port > 0 {
					// udp port is 1 lower
					port--
					// create session
					session, err = s.NewDatagramSession(cfg.Session, keys, cfg.Opts.AsList(), port)
				}
			}
		}
	}
	return
}
