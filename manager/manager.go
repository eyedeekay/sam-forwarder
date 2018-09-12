package sammanager

import (
	"context"
	"net"
)

import (
	"github.com/eyedeekay/sam-forwarder"
)

type SAMManager struct {
	FilePath string
	save     bool

	TargetHost string
	TargetPort string
	SamHost    string
	SamPort    string
	forwarders []samforwarder.SAMForwarder

	TunName           string
	inLength          string
	outLength         string
	inQuantity        string
	outQuantity       string
	inVariance        string
	outVariance       string
	inBackupQuantity  string
	outBackupQuantity string
	inAllowZeroHop    string
	outAllowZeroHop   string

	dontPublishedlient string
	encryptLeaseSet    string
	useCompression     string

	reduceIdle         string
	reduceIdleTime     string
	reduceIdleQuantity string
	closeIdle          string
	closeIdleTime      string
	dest               string
}

func (s *SAMManager) FindForwarder(lookup string) (bool, int) {
	for index, element := range s.forwarders {
		if element.TunName == lookup {
			return true, index
		}
	}
	return false, -1
}

func (s *SAMManager) Dial(ctx context.Context, network, address string) (*net.Conn, error) {
	return nil, nil
}

func (s *SAMManager) NewSAMMAnager(opts ...func(*SAMManager) error) (*SAMManager, error) {
	return nil, nil
}
