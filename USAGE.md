ephsite - Easy forwarding of local services to i2p
==================================================

ephsite is a forwarding proxy designed to configure a tunnel for use
with i2p. It can be used to easily forward a local service to the
i2p network using i2p's SAM API instead of the tunnel interface.

usage:
------

```
Usage of ./bin/ephsite:
  -access string
    	Type of access list to use, can be "whitelist" "blacklist" or "none". (default "none")
  -accesslist value
    	Specify an access list member(can be used multiple times)
  -dir string
    	Directory to save tunnel configuration file in.
  -encryptlease
    	Use an encrypted leaseset(true or false) (default true)
  -gzip
    	Uze gzip(true or false)
  -host string
    	Target host(Host of service to forward to i2p) (default "127.0.0.1")
  -inback int
    	Set inbound tunnel backup quantity(0 to 5) (default 4)
  -incount int
    	Set inbound tunnel quantity(0 to 15) (default 8)
  -ini string
    	Use an ini file for configuration(config file options override passed arguments for now.) (default "none")
  -inlen int
    	Set inbound tunnel length(0 to 7) (default 3)
  -invar int
    	Set inbound tunnel length variance(-7 to 7)
  -name string
    	Tunnel name, this must be unique but can be anything. (default "forwarder")
  -outback int
    	Set outbound tunnel backup quantity(0 to 5) (default 4)
  -outcount int
    	Set outbound tunnel quantity(0 to 15) (default 8)
  -outlen int
    	Set outbound tunnel length(0 to 7) (default 3)
  -outvar int
    	Set outbound tunnel length variance(-7 to 7)
  -port string
    	Target port(Port of service to forward to i2p) (default "8081")
  -reduce
    	Reduce tunnel quantity when idle(true or false)
  -reducecount int
    	Reduce idle tunnel quantity to X (0 to 5) (default 3)
  -reducetime int
    	Reduce tunnel quantity after X (minutes) (default 3)
  -samhost string
    	SAM host (default "127.0.0.1")
  -samport string
    	SAM port (default "7656")
  -save
    	Use saved file and persist tunnel(If false, tunnel will not persist after program is stopped. (default true)
  -zeroin
    	Allow zero-hop, non-anonymous tunnels in(true or false)
  -zeroout
    	Allow zero-hop, non-anonymous tunnels out(true or false)
```
