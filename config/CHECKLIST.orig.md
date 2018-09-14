I2CP/Tunnel Options Implementation Checklist
============================================

key:

    - \[U\] - Undone/Unknoqn
    - \[C\] - Confirmed Working
    - \[W\] - Work in progress

                                            Version  Recommended Allowable            Default
        [U] - clientMessageTimeout                                8*1000 - 120*1000    60*1000       The timeout (ms) for all sent messages. Unused. See the protocol specification for per-message settings.
        [U] - crypto.lowTagThreshold          0.9.2               1-128                30            Minimum number of ElGamal/AES Session Tags before we send more. Recommended: approximately tagsToSend * 2/3
        [U] - crypto.tagsToSend               0.9.2               1-128                40            Number of ElGamal/AES Session Tags to send at a time. For clients with relatively low bandwidth per-client-pair (IRC, some UDP apps), this may be set lower.
        [U] - explicitPeers                                                            null          Comma-separated list of Base 64 Hashes of peers to build tunnels through; for debugging only
        [U] - i2cp.dontPublishLeaseSet                true,false                       false         Should generally be set to true for clients and false for servers
        [U] - i2cp.fastReceive                0.9.4   true,false                       false         If true, the router just sends the MessagePayload instead of sending a MessageStatus and awaiting a ReceiveMessageBegin.
        [U] - i2cp.messageReliability                             BestEffort, None     BestEffort    Guaranteed is disabled; None implemented in 0.8.1; the streaming lib default is None as of 0.8.1, the client side default is None as of 0.9.4
        [U] - i2cp.password                   0.8.2               string                             For authorization, if required by the router. If the client is running in the same JVM as a router, this option is not required. Warning - username and password are sent in the clear to the router, unless using SSL (i2cp.SSL=true). Authorization is only recommended when using SSL.
        [U] - i2cp.username                   0.8.2               string
        [U] - inbound.allowZeroHop                    true,false                       true          If incoming zero hop tunnel is allowed
        [U] - outbound.allowZeroHop                   true,false                       true          If outgoing zero hop tunnel is allowed
        [U] - inbound.backupQuantity                  0 to 3      No limit             0             Number of redundant fail-over for tunnels in
        [U] - outbound.backupQuantity                 0 to 3      No limit             0             Number of redundant fail-over for tunnels out
        [U] - inbound.IPRestriction                   0 to 4      0 to 4               2             Number of IP bytes to match to determine if two routers should not be in the same tunnel. 0 to disable.
        [U] - outbound.IPRestriction                  0 to 4      0 to 4               2             Number of IP bytes to match to determine if two routers should not be in the same tunnel. 0 to disable.
        [U] - inbound.length                          0 to 3      0 to 7               3             Length of tunnels in
        [U] - outbound.length                         0 to 3      0 to 7               3             Length of tunnels out
        [U] - inbound.lengthVariance                  -1 to 2    -7 to 7               0             Random amount to add or subtract to the length of tunnels in. A positive number x means add a random amount from 0 to x inclusive. A negative number -x means add a random amount from -x to x inclusive. The router will limit the total length of the tunnel to 0 to 7 inclusive. The default variance was 1 prior to release 0.7.6.
        [U] - outbound.lengthVariance                 -1 to 2    -7 to 7               0             Random amount to add or subtract to the length of tunnels out. A positive number x means add a random amount from 0 to x inclusive. A negative number -x means add a random amount from -x to x inclusive. The router will limit the total length of the tunnel to 0 to 7 inclusive. The default variance was 1 prior to release 0.7.6.
        [U] - inbound.nickname                        string                                         Name of tunnel - generally used in routerconsole, which will use the first few characters of the Base64 hash of the destination by default.
        [U] - outbound.nickname                       string                                         Name of tunnel - generally ignored unless inbound.nickname is unset.
        [U] - outbound.priority               0.9.4   -25 to 25  -25 to 25             0             Priority adjustment for outbound messages. Higher is higher priority.
        [U] - inbound.quantity                        1 to 3     1 to 16               2             Number of tunnels in. Limit was increased from 6 to 16 in release 0.9; however, numbers higher than 6 are incompatible with older releases.
        [U] - outbound.quantity                       1 to 3     No limit              2             Number of tunnels out
        [U] - inbound.randomKey               0.9.17             Base 64 encoding of 32 random bytes Used for consistent peer ordering across restarts.
        [U] - outbound.randomKey              0.9.17             Base 64 encoding of 32 random bytes Used for consistent peer ordering across restarts.
        [U] - inbound.*                                                                              Any other options prefixed with "inbound." are stored in the "unknown options" properties of the inbound tunnel pool's settings.
        [U] - outbound.*                                                                             Any other options prefixed with "outbound." are stored in the "unknown options" properties of the outbound tunnel pool's settings.
        [U] - shouldBundleReplyInfo           0.9.2   true,false                       true          Set to false to disable ever bundling a reply LeaseSet. For clients that do not publish their LeaseSet, this option must be true for any reply to be possible. "true" is also recommended for multihomed servers with long connection times. Setting to "false" may save significant outbound bandwidth, especially if the client is configured with a large number of inbound tunnels (Leases). If replies are still required, this may shift the bandwidth burden to the far-end client and the floodfill. There are several cases where "false" may be appropriate: Unidirectional communication, no reply required LeaseSet is published and higher reply latency is acceptable LeaseSet is published, client is a "server", all connections are inbound so the connecting far-end destination obviously has the leaseset already. Connections are either short, or it is acceptable for latency on a long-lived connection to temporarily increase while the other end re-fetches the LeaseSet after expiration. HTTP servers may fit these requirements.

        [U] - i2cp.closeIdleTime              0.7.1   1800000     300000 minimum                     (ms) Idle time required (default 30 minutes)
        [U] - i2cp.closeOnIdle                0.7.1   true,false                       false         Close I2P session when idle
        [U] - i2cp.encryptLeaseSet            0.7.1   true,false                       false         Encrypt the lease
        [U] - i2cp.fastReceive                0.9.4   true,false                       true          If true, the router just sends the MessagePayload instead of sending a MessageStatus and awaiting a ReceiveMessageBegin.
        [U] - i2cp.gzip                       0.6.5   true,false                       true          Gzip outbound data
        [U] - i2cp.leaseSetKey                0.7.1                                                  For encrypted leasesets. Base 64 SessionKey (44 characters)
        [U] - i2cp.leaseSetPrivateKey         0.9.18                                                 Base 64 private key for encryption. Optionally preceded by the key type and ':'. Only "ELGAMAL_2048:" is supported, which is the default. I2CP will generate the public key from the private key. Use for persistent leaseset keys across restarts.
        [U] - i2cp.leaseSetSigningPrivateKey  0.9.18                                                 Base 64 private key for signatures. Optionally preceded by the key type and ':'. DSA_SHA1 is the default. Key type must match the signature type in the destination. I2CP will generate the public key from the private key. Use for persistent leaseset keys across restarts.
        [U] - i2cp.messageReliability                             BestEffort, None     None          Guaranteed is disabled; None implemented in 0.8.1; None is the default as of 0.9.4
        [U] - i2cp.reduceIdleTime             0.7.1   1200000     300000 minimum                     (ms) Idle time required (default 20 minutes, minimum 5 minutes)
        [U] - i2cp.reduceOnIdle               0.7.1   true,false                       false         Reduce tunnel quantity when idle
        [U] - i2cp.reduceQuantity             0.7.1   1           1 to 5               1             Tunnel quantity when reduced (applies to both inbound and outbound)
        [U] - i2cp.SSL                        0.8.3   true,false                       false         Connect to the router using SSL. If the client is running in the same JVM as a router, this option is ignored, and the client connects to that router internally.
        [U] - i2cp.tcp.host                           127.0.0.1                                      Router hostname. If the client is running in the same JVM as a router, this option is ignored, and the client connects to that router internally.
        [U] - i2cp.tcp.port                           1-65535     7654                               Router I2CP port. If the client is running in the same JVM as a router, this option is ignored, and the client connects to that router internally.
