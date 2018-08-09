Embedding i2p support in your Go application with samforwarder
==============================================================

One neat thing you can do with samforwarder is make eepWeb(?) services configure
themselves automatically by adding it to an existing Go application. To help
with this process, the samforwarder/config/ file has a bunch of helper
functions and a class for parsing configuration files directly. You can import
it, add a few flags(or however you configure your service) and fire off the
forwarder as a goroutne, all you have to do is configure it to forward the port
used by your service. This makes it extremely easy to do, but in my opinion, it
should only be used in this way for applications that would already be safe to
host as services in i2p or other overlay networks. That means avoiding the risk
of out-of-band communication accidentally, such as by making the server retrieve
a resource from a clearnet service.

When I'm less tired and more interested, I'll try to finish this document.
