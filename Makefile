
GOPATH = $(PWD)/.go

echo:
	@echo "$(GOPATH)"

deps:
	go get -u github.com/eyedeekay/sam-forwarder
	go get -u github.com/kpetku/sam3

build: clean
	mkdir -p bin
	cd main && go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o ../bin/ephsite

clean:
	rm -f bin/ephsite

run:
	./bin/ephsite -addr="127.0.0.1:8081"

noopts: clean
	mkdir -p bin
	cd main && go build -o ../bin/ephsite
