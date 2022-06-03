module github.com/eyedeekay/sam-forwarder

go 1.12

require (
	crawshaw.io/littleboss v0.0.0-20190317185602-8957d0aedcce
	github.com/boreq/friendlyhash v0.0.0-20190522010448-1ca64b3ca69e
	github.com/eyedeekay/eephttpd v0.0.9996-0.20210919031443-3b354b7839bd
	github.com/eyedeekay/httptunnel v0.0.0-20220603041627-a064adf0ae4b
	github.com/eyedeekay/i2pkeys v0.0.0-20220310055120-b97558c06ac8
	github.com/eyedeekay/outproxy v0.0.0-20220603040929-b24e1e503f1f
	github.com/eyedeekay/portcheck v0.0.0-20190218044454-bb8718669680
	github.com/eyedeekay/sam3 v0.33.3-0.20220601222524-ee9930813dc1
	github.com/gtank/cryptopasta v0.0.0-20170601214702-1f550f6f2f69
	github.com/justinas/nosurf v0.0.0-20190416172904-05988550ea18
	github.com/zieckey/goini v0.0.0-20180118150432-0da17d361d26
	github.com/zserge/lorca v0.1.9
	github.com/zserge/webview v0.0.0-20190123072648-16c93bcaeaeb
)

replace gopkg.in/russross/blackfriday.v2 v2.1.0 => github.com/russross/blackfriday/v2 v2.1.0
