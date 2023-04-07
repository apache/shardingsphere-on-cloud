.PHONY:openssl-local test build lint

GOOS := $(shell go env GOOS)

openssl-local:
	mkdir -p certs && \
  	cd certs  && \
	openssl req -new -SHA256 -newkey rsa:2048 -nodes -keyout tls.key -out tls.csr -subj "/C=CN/ST=beijing/L=beijing/O=/OU=/" && \
	openssl x509 -req -sha256 -days 365 -in tls.csr -signkey tls.key -out tls.crt
test:
	go test ./... -coverprofile cover.out

build:
	GOOS=$(GOOS) go build -o pitr-agent
lint:
	golangci-lint run -v --timeout 5m
