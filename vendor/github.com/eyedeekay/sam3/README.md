# README #

go library for the I2P [SAMv3.0](https://geti2p.net/en/docs/api/samv3) bridge, used to build anonymous/pseudonymous end-to-end encrypted sockets.

This library is much better than ccondom (that use BOB), much more stable and much easier to maintain.

## Support/TODO ##

**What works:**

* Utils
    * Resolving domain names to I2P destinations
    * .b32.i2p hashes
    * Generating keys/i2p destinations
* Streaming
    * DialI2P() - Connecting to stuff in I2P
    * Listen()/Accept() - Handling incomming connections
    * Implements net.Conn and net.Listener
* Datagrams
    * Implements net.PacketConn
* Raw datagrams
    * Like datagrams, but without addresses

**Does not work:**

* Everything works! :D
* Probably needs some real-world testing

## Documentation ##

* Latest version-documentation:
    * set your GOPATH
    * Enter `godoc -http=:8081` into your terminal and hit enter.
    * Goto http://localhost:8081, click packages, and navigate to sam3

## Examples ##
```go
package main

import (
	"github.com/majestrate/i2p-tools/sam3"
	"fmt"
)

const yoursam = "127.0.0.1:7656" // sam bridge

func client(server I2PAddr) {
	sam, _ := NewSAM(yoursam)
	keys, _ := sam.NewKeys()
	stream, _ := sam.NewStreamSession("clientTun", keys, Options_Small)
	fmt.Println("Client: Connecting to " + server.Base32())
	conn, _ := stream.DialI2P(server)
	conn.Write([]byte("Hello world!"))
	return
}

func main() {
	sam, _ := NewSAM(yoursam)
	keys, _ := sam.NewKeys()
	go client(keys.Addr())
	stream, _ := sam.NewStreamSession("serverTun", keys, Options_Medium)
	listener, _ := stream.Listen()
	conn, _ := listener.Accept()
	buf := make([]byte, 4096)
	n, _ := conn.Read(buf)
	fmt.Println("Server received: " + string(buf[:n]))
}
```

The above will write to the terminal:

```text
Client: Connecting to zjnvfh4hs3et5vtz35ogwzrws26zvwkcad5uo5esecvg4qpk5b4a.b32.i2p
Server received: Hello world!
```

Error handling was omitted in the above code for readability.

## Testing ##

* `go test -tags=nettest` runs the whole suite (takes 90+ sec to perform!)
* `go test -short` runs the shorter variant, does not connect to anything

## License ##

Public domain.

## Author ##

* Kalle Vedin `kalle.vedin@fripost.org`
* Unknown Name (majestrate)
