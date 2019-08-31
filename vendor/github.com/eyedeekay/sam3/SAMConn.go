package sam3

import (
	"github.com/eyedeekay/sam3/i2pkeys"
	"net"
	"time"
)

/*
import (
	. "github.com/eyedeekay/sam3/i2pkeys"
)
*/
// Implements net.Conn
type SAMConn struct {
	laddr i2pkeys.I2PAddr
	raddr i2pkeys.I2PAddr
	conn  net.Conn
}

// Implements net.Conn
func (sc *SAMConn) Read(buf []byte) (int, error) {
	n, err := sc.conn.Read(buf)
	return n, err
}

// Implements net.Conn
func (sc *SAMConn) Write(buf []byte) (int, error) {
	n, err := sc.conn.Write(buf)
	return n, err
}

// Implements net.Conn
func (sc *SAMConn) Close() error {
	return sc.conn.Close()
}

func (sc *SAMConn) LocalAddr() net.Addr {
	return sc.localAddr()
}

// Implements net.Conn
func (sc *SAMConn) localAddr() i2pkeys.I2PAddr {
	return sc.laddr
}

func (sc *SAMConn) RemoteAddr() net.Addr {
	return sc.remoteAddr()
}

// Implements net.Conn
func (sc *SAMConn) remoteAddr() i2pkeys.I2PAddr {
	return sc.raddr
}

// Implements net.Conn
func (sc *SAMConn) SetDeadline(t time.Time) error {
	return sc.conn.SetDeadline(t)
}

// Implements net.Conn
func (sc *SAMConn) SetReadDeadline(t time.Time) error {
	return sc.conn.SetReadDeadline(t)
}

// Implements net.Conn
func (sc *SAMConn) SetWriteDeadline(t time.Time) error {
	return sc.conn.SetWriteDeadline(t)
}
