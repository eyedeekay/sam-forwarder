
GOPATH = $(PWD)/.go

appname = ephsite

echo:
	@echo "$(GOPATH)"
	find . -name "*.go" -exec gofmt -w {} \;
	find . -name "*.i2pkeys" -exec rm {} \;

test:
	go test
	cd udp && go test


deps:
	go get -u github.com/zieckey/goini
	go get -u github.com/eyedeekay/sam-forwarder
	go get -u github.com/eyedeekay/sam-forwarder/udp
	go get -u github.com/eyedeekay/sam-forwarder/config
	go get -u github.com/kpetku/sam3

build: clean bin/$(appname)

bin/$(appname):
	mkdir -p bin
	cd main && go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o ../bin/$(appname)

server: clean-server bin/eephttpd

bin/eephttpd:
	mkdir -p bin
	go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o ../bin/eephttpd ./example/serve.go

all: build server

clean:
	rm -f bin/$(appname)

clean-server:
	rm -f bin/eephttpd

noopts: clean
	mkdir -p bin
	cd main && go build -o ../bin/$(appname)

gendoc: build
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
	./bin/$(appname) -h  2>> USAGE.md; true
	@echo '```' >> USAGE.md
