.PHONY: build test

GOOS := $(shell go env GOOS)

build:
	GOOS=$(GOOS) go build -o gs_pitr main.go
test:
	go test ./... -coverprofile cover.out
