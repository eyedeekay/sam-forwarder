package samforwarder

//SetByteLimit sets the number of hops inbound
func SetByteLimit(u int64) func(*SAMForwarder) error {
	return func(c *SAMForwarder) error {
		c.ByteLimit = u
		return nil
	}
}
