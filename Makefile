.PHONY: build install test

version=$(shell ./version.sh)

build:
	go build -ldflags "-X main.version=${version}" ./cmd/gofs

install:
	go install -ldflags "-X main.version=${version}" ./cmd/gofs

test:
	go test ./...
	truffle test
