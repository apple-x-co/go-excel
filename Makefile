BIN := 'go-excel'

VERSION := '0.9.1'
REVISION := '$(shell git rev-parse --short HEAD)'

BUILD_TAGS_PRODUCTION := 'production'
BUILD_TAGS_DEVELOPMENT := 'development unittest'

all: clean dev-mac linux

.PHONY: version
version:
	echo $(VERSION).$(REVISION)

.PHONY: base
base:
	go build -o $(BIN_NAME) -tags '$(BUILD_TAGS) netgo' -installsuffix netgo -ldflags '-s -w -X main.version=$(VERSION) -X main.revision=$(REVISION) -extldflags "-static"' main.go

.PHONY: dev-mac
dev-mac:
	go mod tidy
	go fmt
	$(MAKE) base BUILD_TAGS=$(BUILD_TAGS_DEVELOPMENT) BIN_NAME=bin/$(BIN)-dev-mac

.PHONY: linux
linux:
	$(MAKE) base BUILD_TAGS=$(BUILD_TAGS_PRODUCTION) CGO_ENABLED=0 GOOS=linux GOARCH=amd64 BIN_NAME=bin/$(BIN)-linux64

.PHONY: clean
clean:
	rm -rf bin/*
	go clean

.PHONY: ci-test
ci-test:
	if [ ! -d work ]; then mkdir work; fi
	./bin/$(BIN)-dev-mac --in ./book.json --out ./work/book.xlsx