# Build the manager binary
FROM golang:1.17 as builder

WORKDIR /workspace

# Copy the Go Modules manifests
COPY shardingsphere-operator/go.mod go.mod
COPY shardingsphere-operator/go.sum go.sum
# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
#ENV GOPROXY=https://proxy.golang.com.cn
RUN go mod download

# Copy the go source
COPY shardingsphere-operator/main.go main.go
COPY shardingsphere-operator/api api/
COPY shardingsphere-operator/pkg pkg/
RUN mkdir -p certs && cd certs && openssl req -new -SHA256 -newkey rsa:2048 -nodes -keyout tls.key -out tls.csr -subj "/C=CN/ST=beijing/L=beijing/O=/OU=/" && openssl x509 -req -sha256 -days 365 -in tls.csr -signkey tls.key -out tls.crt
# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -ldflags "-w -s" -gcflags "-N -l" -a -o manager main.go

# Use distroless as minimal base image to package the manager binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM alpine:3.15.4
WORKDIR /
COPY --from=builder /workspace/manager .
USER root
COPY --from=builder /workspace/certs /etc/operator/certs

ENTRYPOINT ["/manager"]
