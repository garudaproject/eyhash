VERSION = $(shell describe --tags --abbrev=0)
LDFLAGS = -ldflags="-s -w -X main.version=$(VERSION)"

build:
	go build $(LDFLAGS) -o ./bin/eyhash ./cmd/eyhash

install:
	CGO_ENABLED=0 go install $(LDFLAGS) ./cmd/eyhash

run: build
	./bin/eyhash

run-file: build
	./bin/eyhash test.txt -f

run-folder: build
	./bin/eyhash tests -d

run-unknown: build
	./bin/eyhash test.txt -u

test:
	go test -v ./...

test-race:
	go test -v ./... --race

clean:
	rm -rf ./bin
