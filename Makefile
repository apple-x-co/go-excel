.PHONY: base develop production clean ci-test

BUILD_TAGS_PRODUCTION = 'production'
BUILD_TAGS_DEVELOPMENT = 'development unittest'

all: clean develop production

base:
	go build -o $(BIN_NAME) -tags '$(BUILD_TAGS) netgo' -installsuffix netgo -ldflags '-s -w -extldflags "-static"' main.go

develop:
	go mod tidy
	go fmt
	$(MAKE) base BUILD_TAGS=$(BUILD_TAGS_DEVELOPMENT) BIN_NAME=bin/go-excel-dev-mac

production:
	$(MAKE) base BUILD_TAGS=$(BUILD_TAGS_PRODUCTION) CGO_ENABLED=0 GOOS=linux GOARCH=amd64 BIN_NAME=bin/go-excel-linux64

clean:
	rm -rf bin/*

ci-test:
	if [ ! -d work ]; then mkdir work; fi
	./bin/go-excel-dev-mac --file ./book.json --out ./work/book.xlsx