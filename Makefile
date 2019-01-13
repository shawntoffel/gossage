REPO=shawntoffel/gossage
GO=GO111MODULE=on go
BUILD=GOARCH=amd64 $(GO) build 

.PHONY: all deps test build

all: deps test build 
deps:
	$(GO) mod download

test:
	$(GO) vet ./...
	$(GO) test -v -race ./...

build:
	$(BUILD)
