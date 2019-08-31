package command

import (
	"fmt"
	"log"
)

import (
	"github.com/eyedeekay/ramp/config"
)

type SAMEmit struct {
	i2pconfig.I2PConfig
}

func (e *SAMEmit) OptStr() string {
	optStr := ""
	for _, opt := range e.I2PConfig.Print() {
		optStr += opt + " "
	}
	return optStr
}

func (e *SAMEmit) Hello() string {
	return fmt.Sprintf("HELLO VERSION MIN=%s MAX=%s \n", e.I2PConfig.MinSAM(), e.I2PConfig.MaxSAM())
}

func (e *SAMEmit) HelloBytes() []byte {
	return []byte(e.Hello())
}

func (e *SAMEmit) GenerateDestination() string {
	return fmt.Sprintf("DEST GENERATE %s \n", e.I2PConfig.SignatureType())
}

func (e *SAMEmit) GenerateDestinationBytes() []byte {
	return []byte(e.GenerateDestination())
}

func (e *SAMEmit) Lookup(name string) string {
	return fmt.Sprintf("NAMING LOOKUP NAME=%s \n", name)
}

func (e *SAMEmit) LookupBytes(name string) []byte {
	return []byte(e.Lookup(name))
}

func (e *SAMEmit) Create() string {
	return fmt.Sprintf(
		//             //1 2 3 4 5 6 7
		"SESSION CREATE %s%s%s%s%s%s%s \n",
		e.I2PConfig.SessionStyle(),   //1
		e.I2PConfig.FromPort(),       //2
		e.I2PConfig.ToPort(),         //3
		e.I2PConfig.ID(),             //4
		e.I2PConfig.DestinationKey(), // 5
		e.I2PConfig.SignatureType(),  // 6
		e.OptStr(),                   // 7
	)
}

func (e *SAMEmit) CreateBytes() []byte {
	log.Println("sam command: " + e.Create())
	return []byte(e.Create())
}

func (e *SAMEmit) Connect(dest string) string {
	return fmt.Sprintf(
		"STREAM CONNECT ID=%s %s %s DESTINATION=%s \n",
		e.I2PConfig.ID(),
		e.I2PConfig.FromPort(),
		e.I2PConfig.ToPort(),
		dest,
	)
}

func (e *SAMEmit) ConnectBytes(dest string) []byte {
	return []byte(e.Connect(dest))
}

func (e *SAMEmit) Accept() string {
	return fmt.Sprintf(
		"STREAM ACCEPT ID=%s",
		e.I2PConfig.ID(),
		e.I2PConfig.FromPort(),
		e.I2PConfig.ToPort(),
	)
}

func (e *SAMEmit) AcceptBytes() []byte {
	return []byte(e.Accept())
}

func NewEmit(opts ...func(*SAMEmit) error) (*SAMEmit, error) {
	var emit SAMEmit
	for _, o := range opts {
		if err := o(&emit); err != nil {
			return nil, err
		}
	}
	return &emit, nil
}
