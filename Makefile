.PHONY: base develop production test-run

BUILD_TAGS_PRODUCTION = 'production'
BUILD_TAGS_DEVELOPMENT = 'development unittest'

base:
	go build -o $(BIN_NAME) -tags '$(BUILD_TAGS) netgo' -installsuffix netgo -ldflags '-s -w' main.go

develop:
	go mod tidy
	$(MAKE) base BUILD_TAGS=$(BUILD_TAGS_DEVELOPMENT) BIN_NAME=bin/go-excel-dev-mac

production:
	$(MAKE) base BUILD_TAGS=$(BUILD_TAGS_PRODUCTION) CGO_ENABLED=0 GOOS=linux GOARCH=amd64 BIN_NAME=bin/go-excel

test-run:
	./bin/go-excel-dev-mac -file book.json -out work/book.xlsx