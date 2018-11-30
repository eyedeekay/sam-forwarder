package sammanager

import (
	"fmt"
	"strconv"
)

import "github.com/eyedeekay/sam-forwarder/config"

//ManagerOption is a SAMManager Option
type ManagerOption func(*SAMManager) error

//SetManagerFilePath sets the host of the SAMManager's SAM bridge
func SetManagerFilePath(s string) func(*SAMManager) error {
	return func(c *SAMManager) error {
		c.FilePath = s
		return nil
	}
}

//SetManagerSaveFile tells the router to use an encrypted leaseset
func SetManagerSaveFile(b bool) func(*SAMManager) error {
	return func(c *SAMManager) error {
		c.save = b
		return nil
	}
}

//SetManagerHost sets the host of the SAMManager's SAM bridge
func SetManagerHost(s string) func(*SAMManager) error {
	return func(c *SAMManager) error {
		c.ServerHost = s
		return nil
	}
}

//SetManagerPort sets the port of the SAMManager's SAM bridge using a string
func SetManagerPort(s string) func(*SAMManager) error {
	return func(c *SAMManager) error {
		port, err := strconv.Atoi(s)
		if err != nil {
			return fmt.Errorf("Invalid Manager Server Port %s; non-number ", s)
		}
		if port < 65536 && port > -1 {
			c.ServerPort = s
			return nil
		}
		return fmt.Errorf("Invalid port")
	}
}

//SetManagerSAMHost sets the host of the SAMManager's SAM bridge
func SetManagerSAMHost(s string) func(*SAMManager) error {
	return func(c *SAMManager) error {
		c.SamHost = s
		return nil
	}
}

//SetManagerSAMPort sets the port of the SAMManager's SAM bridge using a string
func SetManagerSAMPort(s string) func(*SAMManager) error {
	return func(c *SAMManager) error {
		port, err := strconv.Atoi(s)
		if err != nil {
			return fmt.Errorf("Invalid Manager SAM Port %s; non-number", s)
		}
		if port < 65536 && port > -1 {
			c.SamPort = s
			return nil
		}
		return fmt.Errorf("Invalid port")
	}
}

//SetManagerWebHost sets the host of the SAMManager's SAM bridge
func SetManagerWebHost(s string) func(*SAMManager) error {
	return func(c *SAMManager) error {
		c.WebHost = s
		return nil
	}
}

//SetManagerWebPort sets the port of the SAMManager's SAM bridge using a string
func SetManagerWebPort(s string) func(*SAMManager) error {
	return func(c *SAMManager) error {
		port, err := strconv.Atoi(s)
		if err != nil {
			return fmt.Errorf("Invalid Manager SAM Port %s; non-number", s)
		}
		if port < 65536 && port > -1 {
			c.WebPort = s
			return nil
		}
		return fmt.Errorf("Invalid port")
	}
}

//SetManagerConf sets the host of the SAMManager's SAM bridge
func SetManagerConf(s *i2ptunconf.Conf) func(*SAMManager) error {
	return func(c *SAMManager) error {
		c.config = s
		return nil
	}
}

//SetManagerStart sets the host of the SAMManager's SAM bridge
func SetManagerStart(s bool) func(*SAMManager) error {
	return func(c *SAMManager) error {
		c.start = s
		return nil
	}
}

//SetTunName sets the host of the SAMManager's SAM bridge
func SetTunName(s string) func(*SAMManager) error {
	return func(c *SAMManager) error {
		c.tunName = s
		return nil
	}
}
