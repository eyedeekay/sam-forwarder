
GOPATH = $(PWD)/.go

appname = ephsite
eephttpd = eephttpd
network = si
samhost = sam-host
samport = 7656
args = -r

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

server: clean-server bin/$(eephttpd)

bin/$(eephttpd):
	mkdir -p bin
	go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o ./bin/$(eephttpd) ./example/serve.go

all: build server

clean:
	rm -f bin/$(appname)

clean-server:
	rm -f bin/$(eephttpd)

noopts: clean
	mkdir -p bin
	cd main && go build -o ../bin/$(appname)

gendoc: all
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
	@echo "" >> USAGE.md
	@echo "$(eephttpd) - Static file server automatically forwarded to i2p" >> USAGE.md
	@echo "============================================================" >> USAGE.md
	@echo "" >> USAGE.md
	@echo "usage:" >> USAGE.md
	@echo "------" >> USAGE.md
	@echo "" >> USAGE.md
	@echo "$(eephttpd) is a static http server which automatically runs on i2p with" >> USAGE.md
	@echo "the help of the SAM bridge. By default it will only be available from" >> USAGE.md
	@echo "the localhost and it's i2p tunnel. It can be masked from the localhost" >> USAGE.md
	@echo "using a container." >> USAGE.md
	@echo "" >> USAGE.md
	@echo '```' >> USAGE.md
	./bin/$(eephttpd) -h  2>> USAGE.md; true
	@echo '```' >> USAGE.md
	@echo "" >> USAGE.md
	make docker-cmd
	@echo "" >> USAGE.md
	@cat USAGE.md

docker-build:
	docker build --build-arg user=$(eephttpd) -f Dockerfile -t eyedeekay/$(eephttpd) .

docker-run:
	docker rm -f eephttpd
	docker run -i -t -d \
		--network $(network) \
		--env samhost=$(samhost) \
		--env samport=$(samport) \
		--env args=$(args) \
		--network-alias $(eephttpd) \
		--hostname $(eephttpd) \
		--name $(eephttpd) \
		--restart always \
		--volume $(eephttpd):/home/$(eephttpd)/www \
		eyedeekay/$(eephttpd)

docker-cmd:
	@echo "### build in docker" >> USAGE.md
	@echo "" >> USAGE.md
	@echo '```' >> USAGE.md
	@echo "docker build --build-arg user=$(eephttpd) -f Dockerfile -t eyedeekay/$(eephttpd) ." >> USAGE.md
	@echo '```' >> USAGE.md
	@echo "" >> USAGE.md
	@echo "### Run in docker" >> USAGE.md
	@echo "" >> USAGE.md
	@echo '```' >> USAGE.md
	@echo "docker run -i -t -d --network $(network) \\" >> USAGE.md
	@echo "    --env samhost=$(samhost) \\" >> USAGE.md
	@echo "    --env samport=$(samport) \\" >> USAGE.md
	@echo "    --env args=$(args) # Additional arguments to pass to eephttpd\\" >> USAGE.md
	@echo "    --network-alias $(eephttpd) \\" >> USAGE.md
	@echo "    --hostname $(eephttpd) \\" >> USAGE.md
	@echo "    --name $(eephttpd) \\" >> USAGE.md
	@echo "    --restart always \\" >> USAGE.md
	@echo "    --volume $(eephttpd):/home/$(eephttpd)/ \\" >> USAGE.md
	@echo "    eyedeekay/$(eephttpd)" >> USAGE.md
	@echo '```' >> USAGE.md

index:
	markdown USAGE.md | tee example/www/index.html
