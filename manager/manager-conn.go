package sammanager

import (
	"context"
	"net"
	"os"
	"time"
)

func (s *SAMManager) Dial(network, address string) (net.Conn, error) {
	return nil, nil
}

func (s *SAMManager) DialContext(ctx context.Context, network, address string) (*net.Conn, error) {
	return nil, nil
}

func (s *SAMManager) DialTimeout(network, address string, timeout time.Duration) (net.Conn, error) {
	return nil, nil
}

func (s *SAMManager) FileConn(f *os.File) (c net.Conn, err error) {
	return nil, nil
}

func (s *SAMManager) Pipe() (net.Conn, net.Conn) {
	return nil, nil
}
