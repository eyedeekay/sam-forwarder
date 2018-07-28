
GOPATH = $(PWD)/.go

echo:
	@echo "$(GOPATH)"
	find . -name "*.go" -exec gofmt -w {} \;

deps:
	go get -u github.com/eyedeekay/sam-forwarder
	go get -u github.com/kpetku/sam3

build: clean
	mkdir -p bin
	cd main && go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o ../bin/ephsite

clean:
	rm -f bin/ephsite

noopts: clean
	mkdir -p bin
	cd main && go build -o ../bin/ephsite

gendoc: deps build
	@echo "ephsite - Easy forwarding of local services to i2p" > USAGE.md
	@echo "==================================================" >> USAGE.md
	@echo "" >> USAGE.md
	@echo "ephsite is" >> USAGE.md
	@echo "" >> USAGE.md
	@echo "usage:" >> USAGE.md
	@echo "------" >> USAGE.md
	@echo "" >> USAGE.md
	./bin/ephsite -h | sed 's|  |       |g' 2>&1 | tee -a USAGE.md
	@echo "" >> USAGE.md
