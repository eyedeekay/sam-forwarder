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
  -client
    	Client proxy mode(true or false)
  -close
    	Close tunnel idle(true or false)
  -closetime int
    	Reduce tunnel quantity after X (milliseconds) (default 600000)
  -dest string
    	Destination for client tunnels. Ignored for service tunnels. (default "none")
  -dir string
    	Directory to save tunnel configuration file in.
  -encryptlease
    	Use an encrypted leaseset(true or false) (default true)
  -gzip
    	Uze gzip(true or false)
  -headers
    	Inject X-I2P-DEST headers
  -host string
    	Target host(Host of service to forward to i2p) (default "127.0.0.1")
  -inback int
    	Set inbound tunnel backup quantity(0 to 5) (default 4)
  -incount int
    	Set inbound tunnel quantity(0 to 15) (default 6)
  -ini string
    	Use an ini file for configuration(config file options override passed arguments for now.) (default "none")
  -inlen int
    	Set inbound tunnel length(0 to 7) (default 3)
  -invar int
    	Set inbound tunnel length variance(-7 to 7)
  -lsk string
    	path to saved encrypted leaseset keys (default "none")
  -name string
    	Tunnel name, this must be unique but can be anything. (default "forwarder")
  -outback int
    	Set outbound tunnel backup quantity(0 to 5) (default 4)
  -outcount int
    	Set outbound tunnel quantity(0 to 15) (default 6)
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
    	Reduce tunnel quantity after X (milliseconds) (default 600000)
  -samhost string
    	SAM host (default "127.0.0.1")
  -samport string
    	SAM port (default "7656")
  -save
    	Use saved file and persist tunnel(If false, tunnel will not persist after program is stopped.
  -tlsport string
    	(Currently inoperative. Target TLS port(HTTPS Port of service to forward to i2p)
  -udp
    	UDP mode(true or false)
  -zeroin
    	Allow zero-hop, non-anonymous tunnels in(true or false)
  -zeroout
    	Allow zero-hop, non-anonymous tunnels out(true or false)
```

samcatd - Router-independent tunnel management for i2p
=========================================================

samcatd is a daemon which runs a group of forwarding proxies to
provide services over i2p independent of the router. It also serves
as a generalized i2p networking utility for power-users.

usage:
------

```
flag needs an argument: -h
Usage of ./bin/samcatd:
  -a string
    	Type of access list to use, can be "whitelist" "blacklist" or "none". (default "none")
  -accesslist value
    	Specify an access list member(can be used multiple times)
  -c	Client proxy mode(true or false)
  -ct int
    	Reduce tunnel quantity after X (milliseconds) (default 600000)
  -d string
    	Directory to save tunnel configuration file in.
  -f string
    	Use an ini file for configuration(config file options override passed arguments for now.) (default "none")
  -h string
    	Target host(Host of service to forward to i2p) (default "127.0.0.1")
  -i string
    	Destination for client tunnels. Ignored for service tunnels. (default "none")
  -ib int
    	Set inbound tunnel backup quantity(0 to 5) (default 4)
  -ic int
    	Set inbound tunnel quantity(0 to 15) (default 6)
  -ih
    	Inject X-I2P-DEST headers
  -il int
    	Set inbound tunnel length(0 to 7) (default 3)
  -iv int
    	Set inbound tunnel length variance(-7 to 7)
  -k string
    	path to saved encrypted leaseset keys (default "none")
  -l	Use an encrypted leaseset(true or false) (default true)
  -n string
    	Tunnel name, this must be unique but can be anything. (default "forwarder")
  -ob int
    	Set outbound tunnel backup quantity(0 to 5) (default 4)
  -oc int
    	Set outbound tunnel quantity(0 to 15) (default 6)
  -ol int
    	Set outbound tunnel length(0 to 7) (default 3)
  -ov int
    	Set outbound tunnel length variance(-7 to 7)
  -p string
    	Target port(Port of service to forward to i2p) (default "8081")
  -r	Reduce tunnel quantity when idle(true or false)
  -rc int
    	Reduce idle tunnel quantity to X (0 to 5) (default 3)
  -rt int
    	Reduce tunnel quantity after X (milliseconds) (default 600000)
  -s	Start a tunnel with the passed parameters(Otherwise, they will be treated as default values.)
  -sh string
    	SAM host (default "127.0.0.1")
  -sp string
    	SAM port (default "7656")
  -t	Use saved file and persist tunnel(If false, tunnel will not persist after program is stopped.
  -tls string
    	(Currently inoperative. Target TLS port(HTTPS Port of service to forward to i2p)
  -u	UDP mode(true or false)
  -x	Close tunnel idle(true or false)
  -z	Uze gzip(true or false)
  -zi
    	Allow zero-hop, non-anonymous tunnels in(true or false)
  -zo
    	Allow zero-hop, non-anonymous tunnels out(true or false)
```

example config - valid for both ephsite and samcat
==================================================

(ephsite will only use top-level options)

```

## Defaults, these are only invoked with the -start option or if labeled tunnels
## are not present(samcatd instructions)

inbound.length = 3
outbound.length = 6
inbound.lengthVariance = 0
outbound.lengthVariance = 0
inbound.backupQuantity = 3
outbound.backupQuantity = 3
inbound.quantity = 5
outbound.quantity = 5
inbound.allowZeroHop = false
outbound.allowZeroHop = false
i2cp.encryptLeaseSet = false
gzip = true
i2cp.reduceOnIdle = true
i2cp.reduceIdleTime = 3000000
i2cp.reduceQuantity = 2
i2cp.enableWhiteList = false
i2cp.enableBlackList = false

[sam-forwarder]
type = server
host = 127.0.0.1
port = 8081
inbound.length = 3
outbound.length = 6
keys = forwarder

[sam-forwarder-two]
type = client
host = 127.0.0.1
port = 8082
inbound.length = 6
outbound.length = 3
keys = forwarder-two

[sam-forwarder-three]
type = udpclient
host = 127.0.0.1
port = 8083
inbound.length = 3
outbound.length = 6
keys = forwarder-three

[sam-forwarder-four]
type = udpserver
host = 127.0.0.1
port = 8084
inbound.length = 6
outbound.length = 3
keys = forwarder-four

[sam-forwarder-five]
type = http
host = 127.0.0.1
port = 8085
inbound.length = 3
outbound.length = 6
keys = forwarder-five
```

