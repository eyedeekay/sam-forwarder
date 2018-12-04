Embedding i2p support in your Go application with samforwarder
==============================================================

One neat thing you can do with samforwarder is make eepWeb(?) services configure
themselves automatically by adding it to an existing Go application. To help
with this process, the samforwarder/config/ file has a bunch of helper
functions and a class for parsing configuration files directly. You can import
it, add a few flags(or however you configure your service) and fire off the
forwarder as a goroutne, all you have to do is configure it to forward the port
used by your service. This makes it extremely easy to do, but it should only be
used in this way for applications that would already be safe to host as services
in i2p or other overlay networks. In particular, it should only be used for
applications that don't require extensive login information and do not leak
information at the application layer.

So without further ado, a blatant copy-paste of information that shouldn't have
been in the README.md.

## Static eepsite in like no seconds

Using this port forwarder, it's possible to create an instant eepsite from a
folder full of html files(and the resources they call upon). Probably obviously
to everybody reading this right now, but maybe not obviously to everyone reading
this forever. An example of an application that works this way is available
[here at my eephttpd repo](https://github.com/eyedeekay/eephttpd).

## Quick-And-Dirty i2p-enabled golang web applications

Normal web applications can easily add the ability to serve itself over i2p by
importing and configuring this forwarding doodad. Wherever it takes the argument
for the web server's listening host and/or port, pass that same host and/or port
to a new instance of the "SAMForwarder" and then run the "Serve" function of the
SAMForwarder as a goroutine. This simply forwards the running service to the i2p
network, it doesn't do any filtering, and if your application establishes
out-of-band connections, those may escape. Also, if your application is
listening on all addresses, it will be visible from the local network.

Here's a simple example with a simple static file server:

```Diff
package main																		package main

import (																			import (
	"flag"																				"flag"
	"log"																				"log"
	"net/http"																			"net/http"
)																				    )

																			      >	import "github.com/eyedeekay/sam-forwarder"
																			      >
func main() {																			func main() {
	port := flag.String("p", "8100", "port to serve on")														port := flag.String("p", "8100", "port to serve on")
	directory := flag.String("d", ".", "the directory of static file to host")											directory := flag.String("d", ".", "the directory of static file to host")
	flag.Parse()																			flag.Parse()
																			      >
																			      >		forwarder, err := samforwarder.NewSAMForwarderFromOptions(
																			      >			samforwarder.SetHost("127.0.0.1"),
																			      >			samforwarder.SetPort(*port),
																			      >			samforwarder.SetSAMHost("127.0.0.1"),
																			      >			samforwarder.SetSAMPort("7656"),
																			      >			samforwarder.SetName("staticfiles"),
																			      >		)
																			      >		if err != nil {
																			      >			log.Fatal(err.Error())
																			      >		}
																			      >		go forwarder.Serve()

	http.Handle("/", http.FileServer(http.Dir(*directory)))														http.Handle("/", http.FileServer(http.Dir(*directory)))

	log.Printf("Serving %s on HTTP port: %s\n", *directory, *port)													log.Printf("Serving %s on HTTP port: %s\n", *directory, *port)
	log.Fatal(http.ListenAndServe("127.0.0.1:"+*port, nil))														log.Fatal(http.ListenAndServe("127.0.0.1:"+*port, nil))
}																				    }

```

[This tiny file server taken from here and used for this example](https://gist.github.com/paulmach/7271283)

## Integrating your Go web application with i2p using sam-forwarder
