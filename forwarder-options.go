package samforwarder

import (
	"fmt"
	"strconv"
)

//Option is a SAMForwarder Option
type Option func(*SAMForwarder) error

//SetHost sets the host of the SAMForwarder's SAM bridge
func SetHost(s string) func(*SAMForwarder) error {
	return func(c *SAMForwarder) error {
		c.TargetHost = s
		return nil
	}
}

//SetPort sets the port of the SAMForwarder's SAM bridge using a string
func SetPort(s string) func(*SAMForwarder) error {
	return func(c *SAMForwarder) error {
		port, err := strconv.Atoi(s)
		if err != nil {
			return fmt.Errorf("Invalid port; non-number")
		}
		if port < 65536 && port > -1 {
			c.TargetPort = s
			return nil
		}
		return fmt.Errorf("Invalid port")
	}
}

//SetSAMHost sets the host of the SAMForwarder's SAM bridge
func SetSAMHost(s string) func(*SAMForwarder) error {
	return func(c *SAMForwarder) error {
		c.SamHost = s
		return nil
	}
}

//SetSAMPort sets the port of the SAMForwarder's SAM bridge using a string
func SetSAMPort(s string) func(*SAMForwarder) error {
	return func(c *SAMForwarder) error {
		port, err := strconv.Atoi(s)
		if err != nil {
			return fmt.Errorf("Invalid port; non-number")
		}
		if port < 65536 && port > -1 {
			c.SamPort = s
			return nil
		}
		return fmt.Errorf("Invalid port")
	}
}

//SetName sets the host of the SAMForwarder's SAM bridge
func SetName(s string) func(*SAMForwarder) error {
	return func(c *SAMForwarder) error {
		c.TunName = s
		return nil
	}
}
