i2phttpproxy
============

**This is *way* more useful than it is tested. Please be careful.**

On Unixes, always run it with the flag -littleboss=start to ensure that the
proxy always re-starts itself.

This is a very simple standalone HTTP Proxy for I2P based on the SAM Bridge. It
has a few advantages in certain situations, especially for adapting applications
that speak HTTP to the I2P network. It allows applications to start their own
HTTP proxies, with their own identities within I2P and their own discrete
configurations. It also has some disadvantages, it cannot add new readable
names to your I2P address book nor is it able to use an outproxy. It's new, but
it should be stable enough to experiment with a Tor Browser or a hardened
Firefox configuration.

It is not, and is not intended to be, and will not be intended for use by
multiple clients at the same time. It might be more-or-less OK as part of an
inproxy but you should only use it for one client at a time. A multi-client
solution will also be available soon([eeProxy](https://github.com/eyedeekay/eeProxy)).

This is not something you should use without understanding. It is also not
terribly hard to understand. It runs two services. One is an HTTP proxy that
uses the SAM API to retrieve information from I2P. The other is a controller
for that http proxy which, for now, only has one function that re-starts it
entirely. Doing this closes your SAM connection and your existing tunnels,
guaranteeing that you browse new sites with a fresh identity. The reset function
on UNIX-Like operating systems can optionally re-set an application's
configuration back to the state it was in when the proxy was started.

This is useful because I2P tunnel identities can be tracked across eepSites
during contemporary visits, and when combined with a long-term identifier like
a user account, can become "Retroactively Linkable." This enables an
easy-to-use, user-controlled interface for manually creating a new identity
before and after logging into an eepSite with a user account. An example
application which uses this is [i2psetproxy.js](https://github.com/eyedeekay/i2psetproxy.js)
a WebExtension which assures I2P Proxy settings and provides a "Reset Identity"
button on the Firefox toolbar.

Features: Done
--------------

  * Self-supervising, Self-restarting on Unixes
  * CONNECT support
  * "New Ident" signaling interface(Unix-only for now)(I guess I might have done
  for Windows too now but I haven't tried it out yet).

Features: Planned
-----------------

  * Outproxy Support
  * Traffic Shaping
  * Authentication
