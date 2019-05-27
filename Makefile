BIN := 'go-excel'

VERSION := '0.9.4'
REVISION := '$(shell git rev-parse --short HEAD)'

BUILD_TAGS_PRODUCTION := 'production'
BUILD_TAGS_DEVELOPMENT := 'development unittest'

MAIN := apple-x-co/go-excel/cmd/go-excel

all: clean dev-mac linux

.PHONY: version
version:
	echo $(VERSION).$(REVISION)

.PHONY: base
base:
	go build -o $(BIN_NAME) -tags '$(BUILD_TAGS) netgo' -installsuffix netgo -ldflags '-s -w -X main.version=$(VERSION) -X main.revision=$(REVISION) -extldflags "-static"' $(MAIN)

.PHONY: dev-mac
dev-mac:
	go mod tidy
	go fmt $(MAIN)
	if [ ! -d bin/darwin ]; then mkdir -p bin/darwin; fi
	$(MAKE) base BUILD_TAGS=$(BUILD_TAGS_DEVELOPMENT) BIN_NAME=bin/darwin/$(BIN)

.PHONY: linux
linux:
	if [ ! -d bin/linux ]; then mkdir -p bin/linux; fi
	$(MAKE) base BUILD_TAGS=$(BUILD_TAGS_PRODUCTION) CGO_ENABLED=0 GOOS=linux GOARCH=amd64 BIN_NAME=bin/linux/$(BIN)

.PHONY: clean
clean:
	rm -rf bin/*/$(BIN)
	go clean $(MAIN)

.PHONY: ci-test
ci-test:
	if [ ! -d work ]; then mkdir work; fi
	./bin/darwin/$(BIN) --in ./book.json --out ./work/book.xlsx