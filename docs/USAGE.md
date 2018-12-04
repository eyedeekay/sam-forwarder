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
as a generalized i2p networking utility for power-users. It's
intended to be a Swiss-army knife for the SAM API.

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
  -conv string
    	Display the base32 and base64 values of a specified .i2pkeys file
  -cr string
    	Encrypt/decrypt the key files with a passfile
  -css string
    	custom CSS for web interface (default "css/styles.css")
  -ct int
    	Reduce tunnel quantity after X (milliseconds) (default 600000)
  -d string
    	Directory to save tunnel configuration file in.
  -de string
    	Destination to connect client's to by default.
  -f string
    	Use an ini file for configuration(config file options override passed arguments for now.) (default "none")
  -h string
    	Target host(Host of service to forward to i2p) (default "127.0.0.1")
  -i string
    	Destination for client tunnels. Ignored for service tunnels. (default "none")
  -ib int
    	Set inbound tunnel backup quantity(0 to 5) (default 2)
  -ih
    	Inject X-I2P-DEST headers
  -il int
    	Set inbound tunnel length(0 to 7) (default 3)
  -iq int
    	Set inbound tunnel quantity(0 to 15) (default 6)
  -iv int
    	Set inbound tunnel length variance(-7 to 7)
  -js string
    	custom JS for web interface (default "js/scripts.js")
  -k string
    	key for encrypted leaseset (default "none")
  -l	Use an encrypted leaseset(true or false) (default true)
  -n string
    	Tunnel name, this must be unique but can be anything. (default "forwarder")
  -ob int
    	Set outbound tunnel backup quantity(0 to 5) (default 2)
  -ol int
    	Set outbound tunnel length(0 to 7) (default 3)
  -oq int
    	Set outbound tunnel quantity(0 to 15) (default 6)
  -ov int
    	Set outbound tunnel length variance(-7 to 7)
  -p string
    	Target port(Port of service to forward to i2p) (default "8081")
  -pk string
    	private key for encrypted leaseset (default "none")
  -psk string
    	private signing key for encrypted leaseset (default "none")
  -r	Reduce tunnel quantity when idle(true or false)
  -rq int
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
  -w	Start web administration interface
  -wp string
    	Web port (default "7957")
  -x	Close tunnel idle(true or false)
  -z	Uze gzip(true or false)
  -zi
    	Allow zero-hop, non-anonymous tunnels in(true or false)
  -zo
    	Allow zero-hop, non-anonymous tunnels out(true or false)
```

managing samcatd save-encryption keys
=====================================

In order to keep from saving the .i2pkeys files in plaintext format, samcatd
can optionally generate a key and encrypt the .i2pkeys files securely. Of
course, to fully benefit from this arrangement, you need to move those keys
away from the machine where the tunnel keys(the .i2pkeys file) are located,
or protect them in some other way(sandboxing, etc). If you want to use
encrypted .i2pkeys files, you can specify a key file to use with the -cr
option on the terminal or with keyfile option in the .ini file.

example config - valid for both ephsite and samcat
==================================================
Options are still being added, pretty much as fast as I can put them
in. For up-to-the-minute options, see [the checklist](config/CHECKLIST.md)

(**ephsite** will only use top-level options, but they can be labeled or
unlabeled)

(**samcatd** treats the first set of options it sees as the default, and
does not start tunnels based on unlabeled options unless passed the
-s flag.)

``` ini

## Defaults, these are only invoked with the -start option or if labeled tunnels
## are not present(samcatd instructions). **THESE** are the correct config files
## to use as defaults, and not the ones in ../sam-forwarder/tunnels.ini, which
## are used for testing settings availability only.

inbound.length = 3
outbound.length = 3
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
keyfile = "/usr/share/samcatd/samcatd"

#[sam-forwarder-tcp-server]
#type = server
#host = 127.0.0.1
#port = 8081
#inbound.length = 3
#outbound.length = 3
#keys = forwarder

[sam-forwarder-tcp-client]
type = client
host = 127.0.0.1
port = 8082
inbound.length = 3
outbound.length = 3
destination = i2p-projekt.i2p
keys = forwarder-two

#[sam-forwarder-udp-server]
#type = udpserver
#host = 127.0.0.1
#port = 8084
#inbound.length = 6
#outbound.length = 3
#keys = forwarder-four

#[sam-forwarder-udp-client]
#type = udpclient
#host = 127.0.0.1
#port = 8083
#inbound.length = 3
#outbound.length = 3
#destination = i2p-projekt.i2p
#keys = forwarder-three

#[sam-forwarder-tcp-http-server]
#type = http
#host = 127.0.0.1
#port = 8085
#inbound.length = 3
#outbound.length = 3
#keys = forwarder-five

#[sam-forwarder-vpn-server]
#type = udpserver
#host = 127.0.0.1
#port = 8084
#inbound.length = 2
#outbound.length = 2
#inbound.backupQuantity = 3
#outbound.backupQuantity = 3
#inbound.quantity = 5
#outbound.quantity = 5
#i2cp.reduceOnIdle = true
#i2cp.reduceIdleTime = 3000000
#i2cp.reduceQuantity = 2
#i2cp.closeOnIdle = false
#keys = i2pvpnserver

#[sam-forwarder-vpn-client]
#type = udpclient
#host = 127.0.0.1
#port = 8085
#inbound.length = 2
#outbound.length = 2
#inbound.backupQuantity = 3
#outbound.backupQuantity = 3
#inbound.quantity = 5
#outbound.quantity = 5
#i2cp.reduceOnIdle = true
#i2cp.reduceIdleTime = 3000000
#i2cp.reduceQuantity = 2
#destination = adestinationisrequiredorbespecifiedatruntime.i2p
#keys = i2pvpnclient
```

