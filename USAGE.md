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
    	(true or false)
  -host string
    	Target host(Host of service to forward to i2p) (default "127.0.0.1")
  -inback int
    	(0 to 5) (default 4)
  -incount int
    	(0 to 15) (default 8)
  -inlen int
    	(0 to 7) (default 3)
  -invar int
    	(-7 to 7)
  -name string
    	Tunnel name, this must be unique but can be anything. (default "forwarder")
  -outback int
    	(0 to 5) (default 4)
  -outcount int
    	(0 to 15) (default 8)
  -outlen int
    	(0 to 7) (default 3)
  -outvar int
    	(-7 to 7)
  -port string
    	Target port(Port of service to forward to i2p) (default "8081")
  -reduce
    	(true or false)
  -reducecount int
    	(0 to 5) (default 3)
  -reducetime int
    	(minutes) (default 3)
  -samhost string
    	SAM host (default "127.0.0.1")
  -samport string
    	SAM port (default "7656")
  -zeroin
    	(true or false)
  -zeroout
    	(true or false)
```
