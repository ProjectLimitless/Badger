.PHONY: build fmt run run_race test clean vendor_get

GOPATH_ORIG := $(GOPATH)
GOPATH := $(PWD)/vendor
export GOPATH

default: build

build:
	go build -o ./bin/badger ./src/main.go

fmt:
	go fmt ./src/...

run: build
	./bin/badger

run_race:
	go run -race ./src/main.go

test:
	go test ./src/... -v

test_cover:
	go test ./src/badger/parsers -v -cover

test_cover_submit:
	go test ./src/badger/parsers -v -cover -covermode=count -coverprofile=coverage.out

clean:
	rm ./bin/*
	rm ./logs/*

vendor_get:
	GOPATH=${PWD}/vendor go get -d -u -v \
	github.com/gorilla/mux \
	github.com/op/go-logging \
	github.com/tylerb/graceful \
	github.com/donovansolms/lumberjack
