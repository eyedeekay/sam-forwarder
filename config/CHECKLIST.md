I2CP/Tunnel Options Implementation Checklist
============================================

This version of this document is valid for sam-forwarder. If you'd like to use
it, the original is at [CHECKLIST.orig.md](CHECKLIST.orig.md).

key:

    - \[U\] - Undone/Unknown
    - \[C\] - Confirmed Working
    - \[W\] - Work in progress
    - \[N\] - Not applicable/Won't be implemented without good reason.
    - \[*\] - See also

                                            Version  Recommended Allowable            Default
        [U] - clientMessageTimeout                                8*1000 - 120*1000    60*1000       The timeout (ms) for all sent messages. Unused. See the protocol specification for per-message settings.
        [N] - crypto.lowTagThreshold          0.9.2               1-128                30            Minimum number of ElGamal/AES Session Tags before we send more. Recommended: approximately tagsToSend * 2/3
        [N] - crypto.tagsToSend               0.9.2               1-128                40            Number of ElGamal/AES Session Tags to send at a time. For clients with relatively low bandwidth per-client-pair (IRC, some UDP apps), this may be set lower.
        [U] - explicitPeers                                                            null          Comma-separated list of Base 64 Hashes of peers to build tunnels through; for debugging only
        [C] - i2cp.dontPublishLeaseSet                true,false                       false         Should generally be set to true for clients and false for servers
        [C] - i2cp.fastReceive                0.9.4   true,false                       false         If true, the router just sends the MessagePayload instead of sending a MessageStatus and awaiting a ReceiveMessageBegin.
        [C] - i2cp.messageReliability                             BestEffort, None     BestEffort    Guaranteed is disabled; None implemented in 0.8.1; the streaming lib default is None as of 0.8.1, the client side default is None as of 0.9.4
        [U] - i2cp.password                   0.8.2               string                             For authorization, if required by the router. If the client is running in the same JVM as a router, this option is not required. Warning - username and password are sent in the clear to the router, unless using SSL (i2cp.SSL=true). Authorization is only recommended when using SSL.
        [U] - i2cp.username                   0.8.2               string
        [C] - inbound.allowZeroHop                    true,false                       true          If incoming zero hop tunnel is allowed
        [C] - outbound.allowZeroHop                   true,false                       true          If outgoing zero hop tunnel is allowed
        [C] - inbound.backupQuantity                  0 to 3      No limit             0             Number of redundant fail-over for tunnels in
        [C] - outbound.backupQuantity                 0 to 3      No limit             0             Number of redundant fail-over for tunnels out
        [U] - inbound.IPRestriction                   0 to 4      0 to 4               2             Number of IP bytes to match to determine if two routers should not be in the same tunnel. 0 to disable.
        [U] - outbound.IPRestriction                  0 to 4      0 to 4               2             Number of IP bytes to match to determine if two routers should not be in the same tunnel. 0 to disable.
        [C] - inbound.length                          0 to 3      0 to 7               3             Length of tunnels in
        [C] - outbound.length                         0 to 3      0 to 7               3             Length of tunnels out
        [C] - inbound.lengthVariance                  -1 to 2    -7 to 7               0             Random amount to add or subtract to the length of tunnels in. A positive number x means add a random amount from 0 to x inclusive. A negative number -x means add a random amount from -x to x inclusive. The router will limit the total length of the tunnel to 0 to 7 inclusive. The default variance was 1 prior to release 0.7.6.
        [C] - outbound.lengthVariance                 -1 to 2    -7 to 7               0             Random amount to add or subtract to the length of tunnels out. A positive number x means add a random amount from 0 to x inclusive. A negative number -x means add a random amount from -x to x inclusive. The router will limit the total length of the tunnel to 0 to 7 inclusive. The default variance was 1 prior to release 0.7.6.
        [U] - inbound.nickname                        string                                         Name of tunnel - generally used in routerconsole, which will use the first few characters of the Base64 hash of the destination by default.
        [U] - outbound.nickname                       string                                         Name of tunnel - generally ignored unless inbound.nickname is unset.
        [U] - outbound.priority               0.9.4   -25 to 25  -25 to 25             0             Priority adjustment for outbound messages. Higher is higher priority.
        [C] - inbound.quantity                        1 to 3     1 to 16               2             Number of tunnels in. Limit was increased from 6 to 16 in release 0.9; however, numbers higher than 6 are incompatible with older releases.
        [C] - outbound.quantity                       1 to 3     No limit              2             Number of tunnels out
        [U] - inbound.randomKey               0.9.17             Base 64 encoding of 32 random bytes Used for consistent peer ordering across restarts.
        [U] - outbound.randomKey              0.9.17             Base 64 encoding of 32 random bytes Used for consistent peer ordering across restarts.
        [*] - inbound.*                                                                              Any other options prefixed with "inbound." are stored in the "unknown options" properties of the inbound tunnel pool's settings.
        [*] - outbound.*                                                                             Any other options prefixed with "outbound." are stored in the "unknown options" properties of the outbound tunnel pool's settings.
        [U] - shouldBundleReplyInfo           0.9.2   true,false                       true          Set to false to disable ever bundling a reply LeaseSet. For clients that do not publish their LeaseSet, this option must be true for any reply to be possible. "true" is also recommended for multihomed servers with long connection times. Setting to "false" may save significant outbound bandwidth, especially if the client is configured with a large number of inbound tunnels (Leases). If replies are still required, this may shift the bandwidth burden to the far-end client and the floodfill. There are several cases where "false" may be appropriate: Unidirectional communication, no reply required LeaseSet is published and higher reply latency is acceptable LeaseSet is published, client is a "server", all connections are inbound so the connecting far-end destination obviously has the leaseset already. Connections are either short, or it is acceptable for latency on a long-lived connection to temporarily increase while the other end re-fetches the LeaseSet after expiration. HTTP servers may fit these requirements.
        [C] - i2cp.closeIdleTime              0.7.1   1800000     300000 minimum                     (ms) Idle time required (default 30 minutes)
        [C] - i2cp.closeOnIdle                0.7.1   true,false                       false         Close I2P session when idle
        [C] - i2cp.encryptLeaseSet            0.7.1   true,false                       false         Encrypt the lease
        [C] - i2cp.fastReceive                0.9.4   true,false                       true          If true, the router just sends the MessagePayload instead of sending a MessageStatus and awaiting a ReceiveMessageBegin.
        [C] - i2cp.gzip                       0.6.5   true,false                       true          Gzip outbound data
        [C] - i2cp.leaseSetKey                0.7.1                                                  For encrypted leasesets. Base 64 SessionKey (44 characters)
        [C] - i2cp.leaseSetPrivateKey         0.9.18                                                 Base 64 private key for encryption. Optionally preceded by the key type and ':'. Only "ELGAMAL_2048:" is supported, which is the default. I2CP will generate the public key from the private key. Use for persistent leaseset keys across restarts.
        [C] - i2cp.leaseSetSigningPrivateKey  0.9.18                                                 Base 64 private key for signatures. Optionally preceded by the key type and ':'. DSA_SHA1 is the default. Key type must match the signature type in the destination. I2CP will generate the public key from the private key. Use for persistent leaseset keys across restarts.
        [C] - i2cp.reduceIdleTime             0.7.1   1200000     300000 minimum                     (ms) Idle time required (default 20 minutes, minimum 5 minutes)
        [C] - i2cp.reduceOnIdle               0.7.1   true,false                       false         Reduce tunnel quantity when idle
        [C] - i2cp.reduceQuantity             0.7.1   1           1 to 5               1             Tunnel quantity when reduced (applies to both inbound and outbound)
        [*] - i2cp.SSL                        0.8.3   true,false                       false         Connect to the router using SSL. If the client is running in the same JVM as a router, this option is ignored, and the client connects to that router internally.
        [*] - i2cp.tcp.host                           127.0.0.1                                      Router hostname. If the client is running in the same JVM as a router, this option is ignored, and the client connects to that router internally.
        [*] - i2cp.tcp.port                           1-65535     7654                               Router I2CP port. If the client is running in the same JVM as a router, this option is ignored, and the client connects to that router internally.

                                                                  Default     Description
        [C] - i2cp.accessList                                      null       Comma- or space-separated list of Base64 peer Hashes used for either access list or blacklist. As of release 0.7.13.
        [U] - i2cp.destination.sigType                             DSA_SHA1   Use the access list as a whitelist for incoming connections. The name or number of the signature type for a transient destination. As of release 0.9.12.
        [C] - i2cp.enableAccessList                                false      Use the access list as a whitelist for incoming connections. As of release 0.7.13.
        [C] - i2cp.enableBlackList                                 false      Use the access list as a blacklist for incoming connections. As of release 0.7.13.
        [U] - i2p.streaming.answerPings                            true       Whether to respond to incoming pings
        [U] - i2p.streaming.blacklist                              null       Comma- or space-separated list of Base64 peer Hashes to be blacklisted for incoming connections to ALL destinations in the context. This option must be set in the context properties, NOT in the createManager() options argument. Note that setting this in the router context will not affect clients outside the router in a separate JVM and context. As of release 0.9.3.
        [U] - i2p.streaming.bufferSize                             64K        How much transmit data (in bytes) will be accepted that hasn't been written out yet.
        [U] - i2p.streaming.congestionAvoidanceGrowthRateFactor    1          When we're in congestion avoidance, we grow the window size at the rate of 1/(windowSize*factor). In standard TCP, window sizes are in bytes, while in I2P, window sizes are in messages. A higher number means slower growth.
        [U] - i2p.streaming.connectDelay                           -1         How long to wait after instantiating a new con before actually attempting to connect. If this is <= 0, connect immediately with no initial data. If greater than 0, wait until the output stream is flushed, the buffer fills, or that many milliseconds pass, and include any initial data with the SYN.
        [U] - i2p.streaming.connectTimeout                         5*60*1000  How long to block on connect, in milliseconds. Negative means indefinitely. Default is 5 minutes.
        [U] - i2p.streaming.disableRejectLogging                   false      Whether to disable warnings in the logs when an incoming connection is rejected due to connection limits. As of release 0.9.4.
        [U] - i2p.streaming.dsalist                                null       Comma- or space-separated list of Base64 peer Hashes or host names to be contacted using an alternate DSA destination. Only applies if multisession is enabled and the primary session is non-DSA (generally for shared clients only). This option must be set in the context properties, NOT in the createManager() options argument. Note that setting this in the router context will not affect clients outside the router in a separate JVM and context. As of release 0.9.21.
        [U] - i2p.streaming.enforceProtocol                        true       Whether to listen only for the streaming protocol. Setting to true will prohibit communication with Destinations earlier than release 0.7.1 (released March 2009). Set to true if running multiple protocols on this Destination. As of release 0.9.1. Default true as of release 0.9.36.
        [U] - i2p.streaming.inactivityAction                       2 (send)   (0=noop, 1=disconnect) What to do on an inactivity timeout - do nothing, disconnect, or send a duplicate ack.
        [U] - i2p.streaming.inactivityTimeout                      90*1000    Idle time before sending a keepalive
        [U] - i2p.streaming.initialAckDelay                        750        Delay before sending an ack
        [U] - i2p.streaming.initialResendDelay                     1000       The initial value of the resend delay field in the packet header, times 1000. Not fully implemented; see below.
        [U] - i2p.streaming.initialRTO                             9000       Initial timeout (if no sharing data available). As of release 0.9.8.
        [U] - i2p.streaming.initialRTT                             8000       Initial round trip time estimate (if no sharing data available). Disabled as of release 0.9.8; uses actual RTT.
        [U] - i2p.streaming.initialWindowSize                      6          (if no sharing data available) In standard TCP, window sizes are in bytes, while in I2P, window sizes are in messages.
        [U] - i2p.streaming.limitAction                            reset      What action to take when an incoming connection exceeds limits. Valid values are: reset (reset the connection); drop (drop the connection); or http (send a hardcoded HTTP 429 response). Any other value is a custom response to be sent. backslash-r and backslash-n will be replaced with CR and LF. As of release 0.9.34.
        [U] - i2p.streaming.maxConcurrentStreams                   -1         (0 or negative value means unlimited) This is a total limit for incoming and outgoing combined.
        [U] - i2p.streaming.maxConnsPerMinute                      0          Incoming connection limit (per peer; 0 means disabled) As of release 0.7.14.
        [U] - i2p.streaming.maxConnsPerHour                        0          (per peer; 0 means disabled) As of release 0.7.14.
        [U] - i2p.streaming.maxConnsPerDay                         0          (per peer; 0 means disabled) As of release 0.7.14.
        [U] - i2p.streaming.maxMessageSize                         1730       The MTU in bytes.
        [U] - i2p.streaming.maxResends                             8          Maximum number of retransmissions before failure.
        [U] - i2p.streaming.maxTotalConnsPerMinute                 0          Incoming connection limit (all peers; 0 means disabled) As of release 0.7.14.
        [U] - i2p.streaming.maxTotalConnsPerHour                   0          (all peers; 0 means disabled) Use with caution as exceeding this will disable a server for a long time. As of release 0.7.14.
        [U] - i2p.streaming.maxTotalConnsPerDay                    0          (all peers; 0 means disabled) Use with caution as exceeding this will disable a server for a long time. As of release 0.7.14.
        [U] - i2p.streaming.maxWindowSize                          128
        [U] - i2p.streaming.profile                                1 (bulk)   (2=interactive not supported) This doesn't currently do anything, but setting it to a value other than 1 will cause an error.
        [U] - i2p.streaming.readTimeout                            -1         How long to block on read, in milliseconds. Negative means indefinitely.
        [U] - i2p.streaming.slowStartGrowthRateFactor              1          When we're in slow start, we grow the window size at the rate of 1/(factor). In standard TCP, window sizes are in bytes, while in I2P, window sizes are in messages. A higher number means slower growth.
        [U] - i2p.streaming.tcbcache.rttDampening                  0.75       Ref: RFC 2140. Floating point value. May be set only via context properties, not connection options. As of release 0.9.8.
        [U] - i2p.streaming.tcbcache.rttdevDampening               0.75       Ref: RFC 2140. Floating point value. May be set only via context properties, not connection options. As of release 0.9.8.
        [U] - i2p.streaming.tcbcache.wdwDampening                  0.75       Ref: RFC 2140. Floating point value. May be set only via context properties, not connection options. As of release 0.9.8.
        [U] - i2p.streaming.writeTimeout                          -1          How long to block on write/flush, in milliseconds. Negative means indefinitely.

        [C] - destination                                                     useful to consider adding to custom applications for client ocnfiguration

\* : I'd like to have something like this setting internal to samcatd, but it
might not always be relevant to pass it through to the real i2p router. Right
now, I'm leaning toward a samcatd specific setting, but maybe just alter the
behavior of this setting for use with samcatd instead? Probably just give
samcatd it's own thing.

