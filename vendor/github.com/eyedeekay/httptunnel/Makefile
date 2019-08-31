
GO111MODULE=on

GO_COMPILER_OPTS = -a -tags netgo -ldflags '-w -extldflags "-static"'

httpall: fmt win lin linarm mac

include multiproxy/Makefile

opall: fmt opwin oplin oplinarm opmac

ball:
	cd multiproxy && make all

all: httpall opall ball

fmt:
	find . -name '*.go' -exec gofmt -w {} \;

dep:
	go get -u github.com/eyedeekay/httptunnel/httpproxy
	go get -u github.com/eyedeekay/httptunnel/multiproxy/browserproxy

win: win32 win64

win64:
	GOOS=windows GOARCH=amd64 go build \
		$(GO_COMPILER_OPTS) \
		-buildmode=exe \
		-o ./httpproxy.exe \
		./windows/main.go
	@echo "built"

win32:
	GOOS=windows GOARCH=386 go build \
		$(GO_COMPILER_OPTS) \
		-buildmode=exe \
		-o ./httpproxy.exe \
		./windows/main.go
	@echo "built"

lin: lin64 lin32

lin64:
	GOOS=linux GOARCH=amd64 go build \
		$(GO_COMPILER_OPTS) \
		-o ./httpproxy-64 \
		./httpproxy/main.go
	@echo "built"

lin32:
	GOOS=linux GOARCH=386 go build \
		$(GO_COMPILER_OPTS) \
		-buildmode=exe \
		-o ./httpproxy-32 \
		./httpproxy/main.go
	@echo "built"

linarm: linarm32 linarm64

linarm64:
	GOOS=linux GOARCH=arm64 go build \
		$(GO_COMPILER_OPTS) \
		-buildmode=exe \
		-o ./httpproxy-arm64 \
		./httpproxy/main.go
	@echo "built"

linarm32:
	GOOS=linux GOARCH=arm go build \
		$(GO_COMPILER_OPTS) \
		-buildmode=exe \
		-o ./httpproxy-arm32 \
		./httpproxy/main.go
	@echo "built"


mac: mac32 mac64

mac64:
	GOOS=darwin GOARCH=amd64 go build \
		$(GO_COMPILER_OPTS) \
		-o ./httpproxy-64.app \
		./httpproxy/main.go
	@echo "built"

mac32:
	GOOS=darwin GOARCH=amd64 go build \
		$(GO_COMPILER_OPTS) \
		-o ./httpproxy-32.app \
		./httpproxy/main.go
	@echo "built"

vet:
	go vet ./*.go
	go vet ./httpproxy/*.go
	go vet ./windows/*.go

clean: bclean
	rm -f httpproxy-* outproxy-* *.exe *.log *.js *.map

opwin: opwin32 opwin64

opwin64:
	GOOS=windows GOARCH=amd64 go build \
		$(GO_COMPILER_OPTS) \
		-buildmode=exe \
		-o ./outproxy.exe \
		./outproxy/windows/main.go
	@echo "built"

opwin32:
	GOOS=windows GOARCH=386 go build \
		$(GO_COMPILER_OPTS) \
		-buildmode=exe \
		-o ./outproxy.exe \
		./outproxy/windows/main.go
	@echo "built"

oplin: oplin64 oplin32

oplin64:
	GOOS=linux GOARCH=amd64 go build \
		$(GO_COMPILER_OPTS) \
		-o ./outproxy-64 \
		./outproxy/cmd/main.go
	@echo "built"

oplin32:
	GOOS=linux GOARCH=386 go build \
		$(GO_COMPILER_OPTS) \
		-buildmode=exe \
		-o ./outproxy-32 \
		./outproxy/cmd/main.go
	@echo "built"

oplinarm: oplinarm32 oplinarm64

oplinarm64:
	GOOS=linux GOARCH=arm64 go build \
		$(GO_COMPILER_OPTS) \
		-buildmode=exe \
		-o ./outproxy-arm64 \
		./outproxy/cmd/main.go
	@echo "built"

oplinarm32:
	GOOS=linux GOARCH=arm go build \
		$(GO_COMPILER_OPTS) \
		-buildmode=exe \
		-o ./outproxy-arm32 \
		./outproxy/cmd/main.go
	@echo "built"


opmac: opmac32 opmac64

opmac64:
	GOOS=darwin GOARCH=amd64 go build \
		$(GO_COMPILER_OPTS) \
		-o ./outproxy-64.app \
		./outproxy/cmd/main.go
	@echo "built"

opmac32:
	GOOS=darwin GOARCH=amd64 go build \
		$(GO_COMPILER_OPTS) \
		-o ./outproxy-32.app \
		./outproxy/cmd/main.go
	@echo "built"

ureq:
	http_proxy=http://127.0.0.1:7950 \
	wget -d --auth-no-challenge --proxy-user user --proxy-password password \
		http://i2p-projekt.i2p -O /dev/null 2>&1 | less -rN
