PACKAGE_NAME=friendlyhash

all: test analyze

ci: install-dependencies install-tools test-ci analyze

doc:
	@echo "http://localhost:6060/pkg/github.com/boreq/${PACKAGE_NAME}/"
	@echo "In order to display unexported declarations append ?m=all to an url after"
	@echo "opening docs for a specific package."
	godoc -http=:6060 -play

install-dependencies:
	go get -t ./...

install-tools:
	go get -v -u honnef.co/go/tools/cmd/staticcheck

analyze: analyze-vet analyze-staticcheck analyze-gofmt

analyze-vet:
	go vet github.com/boreq/${PACKAGE_NAME}/...

analyze-staticcheck:
	# https://github.com/dominikh/go-tools/tree/master/cmd/staticcheck
	staticcheck github.com/boreq/${PACKAGE_NAME}/...

analyze-gofmt:
	./confirm_gofmt.sh

test:
	go test ./...

test-ci:
	go test -coverprofile=coverage.txt -covermode=atomic ./...

test-verbose:
	go test -v ./...

.PHONY: all doc install-dependencies install-tools analyze analyze-vet analyze-staticcheck analyze-gofmt test test-ci test-verbose
