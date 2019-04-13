NAME=vpngate-cli
VERSION=$(shell git describe --tags --abbrev=0)
LDFLAGS="'-w -s -X main.version=$(VERSION)'"
PREFIX=/usr/local/bin/
BIN_DIR=bin

.PHONY: build $(BIN_DIR) clean install uninstall deps test

build:$(BIN_DIR)
	@GO111MODULE=on go build -ldflags=$(LDFLAGS) -o $(BIN_DIR)/$(NAME) ./cmd/$(NAME)/main.go

$(BIN_DIR):
	@mkdir -p $(BIN_DIR)

clean:
	-rm -rf $(BIN_DIR)

install:
	@cp -i $(BIN_DIR)/$(NAME) $(PREFIX)

uninstall:
	@rm -i $(PREFIX)/$(NAME)

deps:
	@GO111MODULE=on go mod download

test: deps
	@go test ./...

