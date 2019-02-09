
GOPATH=$(HOME)/go

appname = ephsite
packagename = sam-forwarder
eephttpd = eephttpd
samcatd = samcatd
network = si
samhost = sam-host
samport = 7656
args = -r

PREFIX := /
VAR := var/
RUN := run/
LIB := lib/
LOG := log/
ETC := etc/
USR := usr/
LOCAL := local/
VERSION := 0.1

WEB_INTERFACE = -tags "webface netgo"

echo:
	@echo "$(GOPATH)"
	find . -path ./.go -prune -o -name "*.go" -exec gofmt -w {} \;
	find . -path ./.go -prune -o -name "*.i2pkeys" -exec rm {} \;
	find . -path ./.go -prune -o -name "*.go" -exec cat {} \; | nl

## TODO: Remove this, replace with something right
fix-debian:
	find ./debian -type f -exec sed -i 's|lair repo key|eyedeekay|g' {} \;
	find ./debian -type f -exec sed -i 's|eyedeekay@safe-mail.net|hankhill19580@gmail.com|g' {} \;

test: test-keys test-ntcp test-ssu test-config test-manager

long-test: test-serve test

full-test: test-serve test-vpn test

test-serve:
	cd serve_test && go test

test-ntcp:
	go test

test-ssu:
	cd udp && go test

test-vpn:
	cd csvpn && go test

test-config:
	cd config && go test

test-manager:
	cd manager && go test

test-keys:
	cd i2pkeys && go test

try-web:
	cd bin && \
		./samcatd-web -w -f ../etc/samcatd/tunnels.ini

gdb-web:
	cd bin && \
		gdb ./samcatd-web -w -f ../etc/samcatd/tunnels.ini

refresh:

deps:
	go get -u github.com/songgao/water
	go get -u github.com/gtank/cryptopasta
	go get -u github.com/zieckey/goini
	go get -u github.com/eyedeekay/udptunnel
	go get -u github.com/eyedeekay/sam-forwarder
	go get -u github.com/eyedeekay/sam-forwarder/i2pkeys
	go get -u github.com/eyedeekay/sam-forwarder/i2pkeys/keys
	go get -u github.com/eyedeekay/sam-forwarder/udp
	go get -u github.com/eyedeekay/sam-forwarder/config
	go get -u github.com/eyedeekay/sam-forwarder/manager
	go get -u github.com/eyedeekay/sam3
	go get -u github.com/eyedeekay/littleboss
	go get -u github.com/eyedeekay/samcatd-web

mine:
	go get -u github.com/kpetku/sam3

webdep:
	go get -u github.com/eyedeekay/samcatd-web

build: clean bin/$(appname)

install:
	install -m755 ./bin/$(appname) $(PREFIX)$(USR)$(LOCAL)/bin/
	install -m755 ./bin/$(samcatd) $(PREFIX)$(USR)$(LOCAL)/bin/
	install -m755 ./bin/$(samcatd)-web $(PREFIX)$(USR)$(LOCAL)/bin/
	install -m644 ./etc/init.d/samcatd $(PREFIX)$(ETC)/init.d
	mkdir -p $(PREFIX)$(ETC)/samcatd/ $(PREFIX)$(ETC)/sam-forwarder/ $(PREFIX)$(ETC)/i2pvpn/
	install -m644 ./etc/samcatd/tunnels.ini $(PREFIX)$(ETC)/samcatd/
	install -m644 ./etc/sam-forwarder/tunnels.ini $(PREFIX)$(ETC)/sam-forwarder/
	install -m644 ./etc/i2pvpn/i2pvpn.ini $(PREFIX)$(ETC)/i2pvpn/
	install -m644 ./etc/i2pvpn/i2pvpnclient.ini $(PREFIX)$(ETC)/i2pvpn/

bin/$(appname):
	mkdir -p bin
	cd main && go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o ../bin/$(appname)

server: clean-server bin/$(eephttpd)

bin/$(eephttpd):
	mkdir -p bin
	go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o ./bin/$(eephttpd) ./example/serve.go

daemon: clean-daemon bin/$(samcatd)

bin/$(samcatd):
	mkdir -p bin
	go build -a -tags netgo \
		-ldflags '-w -extldflags "-static"' \
		-o ./bin/$(samcatd) \
		./daemon/*.go

daemon-web: clean-daemon-web bin/$(samcatd)-web

bin/$(samcatd)-web:
	mkdir -p bin
	go build -a $(WEB_INTERFACE) \
		-ldflags '-w -extldflags "-static"' \
		-o ./bin/$(samcatd)-web \
		./daemon/*.go

all: daemon daemon-web build server

clean-all: clean clean-server clean-daemon clean-daemon-web

clean:
	rm -f bin/$(appname)

clean-server:
	rm -f bin/$(eephttpd)

clean-daemon:
	rm -f bin/$(samcatd)

clean-daemon-web:
	rm -f bin/$(samcatd)-web

noopts: clean
	mkdir -p bin
	cd main && go build -o ../bin/$(appname)

install-forwarder:
	install -m755 bin/ephsite /usr/local/bin/ephsite

install-server:
	install -m755 bin/eephttpd /usr/local/bin/eephttpd

install-all: install install-server

remove:
	rm -rf /usr/local/bin/ephsite /usr/local/bin/eephttpd

gendoc:
	@echo "$(appname) - Easy forwarding of local services to i2p" > USAGE.md
	@echo "==================================================" >> USAGE.md
	@echo "" >> USAGE.md
	@echo "$(appname) is a forwarding proxy designed to configure a tunnel for use" >> USAGE.md
	@echo "with i2p. It can be used to easily forward a local service to the" >> USAGE.md
	@echo "i2p network using i2p's SAM API instead of the tunnel interface." >> USAGE.md
	@echo "" >> USAGE.md
	@echo "usage:" >> USAGE.md
	@echo "------" >> USAGE.md
	@echo "" >> USAGE.md
	@echo '```' >> USAGE.md
	./bin/$(appname) -help  2>> USAGE.md; true
	@echo '```' >> USAGE.md
	@echo "" >> USAGE.md
	@echo "$(samcatd) - Router-independent tunnel management for i2p" >> USAGE.md
	@echo "=========================================================" >> USAGE.md
	@echo "" >> USAGE.md
	@echo "$(samcatd) is a daemon which runs a group of forwarding proxies to" >> USAGE.md
	@echo "provide services over i2p independent of the router. It also serves" >> USAGE.md
	@echo "as a generalized i2p networking utility for power-users. It's" >> USAGE.md
	@echo "intended to be a Swiss-army knife for the SAM API." >> USAGE.md
	@echo "" >> USAGE.md
	@echo "usage:" >> USAGE.md
	@echo "------" >> USAGE.md
	@echo "" >> USAGE.md
	@echo '```' >> USAGE.md
	./bin/$(samcatd) -h  2>> USAGE.md; true
	@echo '```' >> USAGE.md
	@echo "" >> USAGE.md
	make key-management
	make example-config

key-management:
	@echo "managing $(samcatd) save-encryption keys" >> USAGE.md
	@echo "=====================================" >> USAGE.md
	@echo "" >> USAGE.md
	@echo "In order to keep from saving the .i2pkeys files in plaintext format, samcatd" >> USAGE.md
	@echo "can optionally generate a key and encrypt the .i2pkeys files securely. Of" >> USAGE.md
	@echo "course, to fully benefit from this arrangement, you need to move those keys" >> USAGE.md
	@echo "away from the machine where the tunnel keys(the .i2pkeys file) are located," >> USAGE.md
	@echo "or protect them in some other way(sandboxing, etc). If you want to use" >> USAGE.md
	@echo "encrypted .i2pkeys files, you can specify a key file to use with the -cr" >> USAGE.md
	@echo "option on the terminal or with keyfile option in the .ini file." >> USAGE.md
	@echo "" >> USAGE.md

example-config:
	@echo "example config - valid for both ephsite and samcat" >> USAGE.md
	@echo "==================================================" >> USAGE.md
	@echo "Options are still being added, pretty much as fast as I can put them" >> USAGE.md
	@echo "in. For up-to-the-minute options, see [the checklist](config/CHECKLIST.md)" >> USAGE.md
	@echo "" >> USAGE.md
	@echo "(**ephsite** will only use top-level options, but they can be labeled or" >> USAGE.md
	@echo "unlabeled)" >> USAGE.md
	@echo "" >> USAGE.md
	@echo "(**samcatd** treats the first set of options it sees as the default, and" >> USAGE.md
	@echo "does not start tunnels based on unlabeled options unless passed the" >> USAGE.md
	@echo "-s flag.)" >> USAGE.md
	@echo "" >> USAGE.md
	@echo '``` ini' >> USAGE.md
	cat etc/samcatd/tunnels.ini >> USAGE.md
	@echo '```' >> USAGE.md
	@echo "" >> USAGE.md
	mv USAGE.md docs/USAGE.md


docker-build:
	docker build --no-cache \
		--build-arg user=$(samcatd) \
		--build-arg path=example/www \
		-f Dockerfile \
		-t eyedeekay/$(samcatd) .

docker-volume:
	docker run -i -t -d \
		--name $(samcatd)-volume \
		--volume $(samcatd):/home/$(samcatd)/ \
		eyedeekay/$(samcatd); true
	docker stop $(samcatd)-volume; true

docker-run: docker-volume
	docker rm -f eephttpd; true
	docker run -i -t -d \
		--cap-add "net_bind_service" \
		--network $(network) \
		--env samhost=$(samhost) \
		--env samport=$(samport) \
		--env args=$(args) \
		--network-alias $(samcatd) \
		--hostname $(samcatd) \
		--name $(samcatd) \
		--restart always \
		--volumes-from $(samcatd)-volume \
		eyedeekay/$(samcatd)
	make follow

c:
	go build ./i2pkeys

follow:
	docker logs -f $(samcatd)

docker: docker-build docker-volume docker-run

index:
	pandoc USAGE.md -o example/www/index.html && cp example/www/index.html docs/index.html
	cp config/CHECKLIST.md docs/config

visit:
	http_proxy=http://127.0.0.1:4444 surf http://566niximlxdzpanmn4qouucvua3k7neniwss47li5r6ugoertzuq.b32.i2p

forward:
	./bin/ephsite -client -dest i2p-projekt.i2p

gojs:
	go get -u github.com/gopherjs/gopherjs

GOPHERJS=$(GOPATH)/bin/gopherjs

js:
	mkdir -p bin
	$(GOPHERJS) build -v --tags netgo \
		-o ./javascript/$(samcatd).js \
		./daemon/*.go

cleantar:
	rm -f ../$(packagename)_$(VERSION).orig.tar.xz

tar:
	tar --exclude .git \
		--exclude .go \
		--exclude bin \
		-cJvf ../$(packagename)_$(VERSION).orig.tar.xz .
