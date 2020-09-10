.PHONY: all clean test coverage
all: build test

# go list is the canonical utility to find go files
GOFILES := $(shell go list -f '{{ join .GoFiles "\n" }}' ./...)

build: .bin-stamp
	go build -o bin/federator

ci-build: .bin-stamp
	go build -ldflags="-s -w" -o bin/federator-$(platform)

test:
	go test -covermode=atomic -coverpkg=all ./...

coverage: .artifacts-stamp
	go-acc -o artifacts/c.out ./...
	go tool cover -html=artifacts/c.out -o artifacts/coverage.html

# directories do werid things in make, so we can use a stamp
.bin-stamp:
	mkdir -p bin
	touch .bin-stamp

.artifacts-stamp:
	mkdir artifacts
	touch .artifacts-stamp

clean:
	rm -rf bin
	rm .bin-stamp
	rm -rf artifacts
	rm .artifacts-stamp
