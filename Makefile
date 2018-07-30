
GOPATH = $(PWD)/.go

echo:
	@echo "$(GOPATH)"
	find . -name "*.go" -exec gofmt -w {} \;

test:
	go test
	cd udp && go test


deps:
	go get -u github.com/zieckey/goini
	go get -u github.com/eyedeekay/sam-forwarder
	go get -u github.com/eyedeekay/sam-forwarder/udp
	go get -u github.com/eyedeekay/sam-forwarder/config
	go get -u github.com/kpetku/sam3

build: clean
	mkdir -p bin
	cd main && go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o ../bin/ephsite

clean:
	rm -f bin/ephsite

noopts: clean
	mkdir -p bin
	cd main && go build -o ../bin/ephsite

gendoc: build
	@echo "ephsite - Easy forwarding of local services to i2p" > USAGE.md
	@echo "==================================================" >> USAGE.md
	@echo "" >> USAGE.md
	@echo "ephsite is a forwarding proxy designed to configure a tunnel for use" >> USAGE.md
	@echo "with i2p. It can be used to easily forward a local service to the" >> USAGE.md
	@echo "i2p network using i2p's SAM API instead of the tunnel interface." >> USAGE.md
	@echo "" >> USAGE.md
	@echo "usage:" >> USAGE.md
	@echo "------" >> USAGE.md
	@echo "" >> USAGE.md
	@echo '```' >> USAGE.md
	./bin/ephsite -h  2>> USAGE.md; true
	@echo '```' >> USAGE.md
