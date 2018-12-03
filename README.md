# sam-forwarder

*think of it as a shell-scriptable re-implementation of i2ptunnel. That's*
*pretty much it.*

Forward a local port to i2p over the SAM API, or proxy a destination to a port
on the local host. This is a work-in-progress, but the basic functionality is,
there and it's already pretty useful. Everything TCP works, but UDP forwarding
has much less real use than TCP. Turns out UDP was less broken than I thought
though. Yay.

## building
Just:

        make deps build

and it will be in the folder ./bin/

[![Build Status](https://travis-ci.org/eyedeekay/sam-forwarder.svg?branch=master)](https://travis-ci.org/eyedeekay/sam-forwarder)

## [usage/configuration](USAGE.md)

## binaries

Two binaries are produced by this repo. The first, ephsite, is only capable
of running one tunnel at a time. The second, samcatd, is more advanced. It can
start multiple tunnels with their own settings, or be used to start tunnels on
the fly like ephsite by passing the -s option. Eventually I'm probably just
going to use this to configure all of my tunnels.

Current limitations:
====================

I need to document it better.
[Besides fixing up the comments, this should help for now.](USAGE.md). I also
need to control output verbosity better.

I need samcatd to accept a configuration folder identical to
/etc/i2pd/tunnels.conf.d, since part of the point of this is to be compatible
with i2pd's tunnels configuration.

It doesn't encrypt the .i2pkeys file by default, so if someone can steal them,
then they can use them to construct tunnels to impersonate you. Experimental
support for encrypted saves has been added. The idea is that only the person
with the key will be able to decrypt and start the tunnels. It is up to the user
to determine how to go about managing these keys.

TCP and UDP are both working now. Additional functionality might be added by
adding other kinds of protocols overtop the TCP and UDP tunnels as a primitive.
A very basic UDP based VPN will be added soon. Obviously these won't be i2pd
compatible. Not sure what to do about that, except maybe make a "convert" tool
that will cull samcatd-specific options.

I've only enabled the use of a subset of the i2cp and tunnel configuration
options, the ones I use the most and for no other real reason assume other
people use the most. They're pretty easy to add, it's just boring. *If you*
*want an i2cp or tunnel option that isn't available, bring it to my attention*
*please.* I'm pretty responsive when people actually contact me, it'll probably
be added within 24 hours. I intend to have configuration options for all
relevant i2cp and tunnel options, which I'm keeping track of
[here](config/CHECKLIST.md).

I should probably have some options that are available in other general network
utilities. I've started to do this with samcatd.

I want it to be able to save ini files based on the settings used for a running
forwarder. Should be easy, I just need to decide how I want to do it. Also to
focus a bit more. I've got more of a plan here now. tunconf has the loaded ini
file inside it, and variables to track the state of the config options while
running, and they can be switched to save options that might be changed via some
interface or another.

Example tools built using this are being broken off into their own repos. Use
the other repos where appropriate, so I can leave the examples un-messed with.

It would be really awesome if I could make this run on Android. So I'll make
that happen eventually. I started a daemon for managing multiple tunnels and I
figure I give it a web interface to configure stuff with. I'll probably put that
in a different repo though. This is looking a little cluttered.

TLS configuration is experimental.

I'm eventually going to make the manager implement net.Conn. This won't be
exposed in the default application probably though, but rather as a library.
