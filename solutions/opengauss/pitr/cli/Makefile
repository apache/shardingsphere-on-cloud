.PHONY: build test cover lint

GOOS := $(shell go env GOOS)

build:
	GOOS=$(GOOS) go build -o gs_pitr main.go
test:
	go test -gcflags=-l -v ./... -cover -coverprofile cover.out
cover:
	go tool cover -html cover.out
lint:
	golangci-lint run -v --timeout 5m
