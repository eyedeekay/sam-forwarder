# sam-forwarder
Forward a local port to i2p over the SAM API, or proxy a destination to a port
on the local host. This is a work-in-progress, but the basic functionality is,
there and it's already pretty useful.

## building

Just:

        make deps build

and it will be in the folder ./bin/

[![Build Status](https://travis-ci.org/eyedeekay/sam-forwarder.svg?branch=master)](https://travis-ci.org/eyedeekay/sam-forwarder)

## usage

        ./bin/ephsite -host=host -port=port

So, to serve an eepSite version of a local service on port 8080 -

        ./bin/ephsite -host=127.0.0.1 -port=8080

For more information, [look here](USAGE.md)

## ini-like configuration

I made it parse INI-like configuration files, optionally, which allows it to
generate tunnels from snippets of i2pd tunnel configuration files. That's kinda
useful. It appears to be more-or-less compatible with i2pd's tunnels.conf
format, but it only supports the following options:

        type = server
        host = 127.0.0.1
        port = 8081
        dir = /path/to/save/data/in #This is not shared with i2pd tunnels.conf
        inbound.length = 6
        outbound.length = 6
        inbound.lengthVariance = 6
        outbound.lengthVariance = 6
        inbound.backupQuantity = 5
        outbound.backupQuantity = 5
        inbound.quantity = 15
        outbound.quantity = 15
        inbound.allowZeroHop = true
        outbound.allowZeroHop = true
        i2cp.encryptLeaseSet = true
        gzip = true
        i2cp.reduceOnIdle = true
        i2cp.reduceIdleTime = 3000000
        i2cp.reduceQuantity = 4
        i2cp.enableWhiteList = false
        i2cp.enableBlackList = true
        i2cp.accessList = BASE64KEYSSEPARATEDBY,COMMAS
        keys = forwarder

Also it doesn't support sections. Didn't realize that at first. Will address
soon.

Other options are added to the config structure, but have to be referenced
manually, there are no convenience functions for them.

## Static eepsite in like no seconds

Using this port forwarder, it's possible to create an instant eepsite from a
folder full of html files(and the resources they call upon). Probably obviously
to everybody reading this right now, but maybe not obviously to everyone reading
this forever. A go application that does this I call eephttpd can be built with
the command:

        make server

and run from ./bin/eephttpd. The default behavior is to look for the files to
serve under the current directory in ./www. It can be configured to behave
differently according to the rules in [USAGE.md](USAGE.md). A Dockerfile is also
available.

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

Current limitations:
====================

Datagrams are still a work-in-progress. They're enabled, but I don't know for
sure how well they'll work yet. TCP is pretty good though.

I'm in the process of adding client proxying to a specific i2p destination by
base32 or (pre-added)jump address.

I've only enabled the use of a subset of the i2cp and tunnel configuration
options, the ones I use the most and for no other real reason assume other
people use the most. They're pretty easy to add, it's just boring. *If you*
*want an i2cp or tunnel option that isn't available, bring it to my attention*
*please.* I'm pretty responsive when people actually contact me, it'll probably
be added within 24 hours.

Encrypted leasesets are only half-implemented. The option seems to do nothing at
the moment. Soon it will be configurable.

I should probably have some options that are available in other general network
utilities like netcat and socat(ephsite may have it's name changed to samcat at
that point). Configuring timeouts and the like. In order to do this, some of the
existing flags should also be aliased to be more familiar and netcat-like.

I want it to be able to use poorly formed ini files, in order to accomodate the
use of INI-like labels. For now, my workaround is to comment out the labels
until I deal with this. Basically I just want it to ignore the lables and treat
the whole thing as flat. Alternatively I guess I could just have it start a
multiple forwarders, one-per-label, without losing features.

I want it to be able to save ini files based on the settings used for a running
forwarder. Should be easy, I just need to decide how I want to do it. Also to
focus a bit more.

I've written a handful of example tools, but some of them might be better as
their own projects. An i2p-native static site generator in the style of jekyll
(but in go) could be cool.

Haha. Well shit. I migrated colluding\_sites\_attack to auto-configure using
the forwarder and the X-I2p-Dest* headers aren't passed through. Implies some
interesting arrangements, but also makes colluding\_sites\_attack useless in
it's present state. I mean I know what I did with si-i2p-plugin works, so it's
not that important. I'll have to look for a way to make this behavior
configurable though.

It would be really awesome if I could make this run on Android. So I'll make
that happen eventually.
