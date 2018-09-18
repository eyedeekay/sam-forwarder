
GOPATH = $(PWD)/.go

appname = ephsite
eephttpd = eephttpd
samcatd = samcatd
network = si
samhost = sam-host
samport = 7656
args = -r

#WEB_INTERFACE = -tags webface

echo:
	@echo "$(GOPATH)"
	find . -name "*.go" -exec gofmt -w {} \;
	find . -name "*.i2pkeys" -exec rm {} \;

mng:
	cd manager && go test

test:
	go test
	cd udp && go test
	cd config && go test
	cd manager && go test

deps:
	go get -u github.com/gtank/cryptopasta
	go get -u golang.org/x/crypto/openpgp
	go get -u github.com/zieckey/goini
	go get -u github.com/eyedeekay/sam-forwarder
	go get -u github.com/eyedeekay/sam-forwarder/udp
	go get -u github.com/eyedeekay/sam-forwarder/config
	go get -u github.com/eyedeekay/sam-forwarder/manager
	go get -u github.com/kpetku/sam3
	go get -u github.com/eyedeekay/sam3
	go get -u github.com/eyedeekay/samcatd-web

build: clean bin/$(appname)

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
	go build -a -tags netgo $(WEB_INTERFACE) \
		-ldflags '-w -extldflags "-static"' \
		-o ./bin/$(samcatd) \
		./daemon/*.go

all: daemon build server

clean-all: clean clean-server

clean:
	rm -f bin/$(appname)

clean-server:
	rm -f bin/$(eephttpd)

clean-daemon:
	rm -f bin/$(samcatd)

noopts: clean
	mkdir -p bin
	cd main && go build -o ../bin/$(appname)

install:
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
	make example-config

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
	@echo '```' >> USAGE.md
	cat etc/sam-forwarder/tunnels.ini >> USAGE.md
	@echo '```' >> USAGE.md
	@echo "" >> USAGE.md


docker-build:
	docker build --no-cache \
		--build-arg user=$(eephttpd) \
		--build-arg path=example/www \
		-f Dockerfile \
		-t eyedeekay/$(eephttpd) .

docker-volume:
	docker run -i -t -d \
		--name $(eephttpd)-volume \
		--volume $(eephttpd):/home/$(eephttpd)/ \
		eyedeekay/$(eephttpd); true
	docker stop $(eephttpd)-volume; true

docker-run: docker-volume
	docker rm -f eephttpd; true
	docker run -i -t -d \
		--network $(network) \
		--env samhost=$(samhost) \
		--env samport=$(samport) \
		--env args=$(args) \
		--network-alias $(eephttpd) \
		--hostname $(eephttpd) \
		--name $(eephttpd) \
		--restart always \
		--volumes-from $(eephttpd)-volume \
		eyedeekay/$(eephttpd)
	make follow

follow:
	docker logs -f $(eephttpd)

docker: docker-build docker-volume docker-run

index:
	pandoc USAGE.md -o example/www/index.html

visit:
	http_proxy=http://127.0.0.1:4444 surf http://566niximlxdzpanmn4qouucvua3k7neniwss47li5r6ugoertzuq.b32.i2p

forward:
	./bin/ephsite -client -dest i2p-projekt.i2p
