package i2ptunconf

//Option is a Conf Option
type Option func(*Conf) error

//SetTargetForPort sets the port of the Conf's SAM bridge using a string
/*func SetTargetForPort443(s string) func(samtunnel.SAMTunnel) error {
	return func(c samtunnel.SAMTunnel) error {
		port, err := strconv.Atoi(s)
		if err != nil {
			return fmt.Errorf("Invalid Target Port %s; non-number ", s)
		}
		if port < 65536 && port > -1 {
			c.TargetForPort443 = s
			return nil
		}
		return fmt.Errorf("Invalid port")
	}
}
*/
