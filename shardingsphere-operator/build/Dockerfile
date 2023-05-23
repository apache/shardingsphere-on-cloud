#
# Licensed to the Apache Software Foundation (ASF) under one or more
# contributor license agreements.  See the NOTICE file distributed with
# this work for additional information regarding copyright ownership.
# The ASF licenses this file to You under the Apache License, Version 2.0
# (the "License"); you may not use this file except in compliance with
# the License.  You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#


# Stage1: Build Go binary for AMD64 
FROM --platform=linux/amd64 golang:1.19-alpine AS builder-amd64
WORKDIR /workspace
COPY go.mod go.mod
COPY go.sum go.sum
#ENV GOPROXY=https://goproxy.cn
RUN go mod download

# Copy Go source files
COPY cmd/shardingsphere-operator/manager cmd/shardingsphere-operator/manager/
COPY cmd/shardingsphere-operator/main.go cmd/shardingsphere-operator/main.go
COPY api api/
COPY pkg pkg/

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -ldflags "-w -s" -gcflags "-N -l" -a -o operator-bin-amd64 cmd/shardingsphere-operator/main.go

# Step2: Build Go binary for ARM64
FROM --platform=linux/arm64,linux/arm64/v8 golang:1.19-alpine AS builder-arm64
WORKDIR /workspace
COPY go.mod go.mod
COPY go.sum go.sum
#ENV GOPROXY=https://goproxy.cn
RUN go mod download

# Copy the go source
COPY cmd/shardingsphere-operator/manager cmd/shardingsphere-operator/manager/
COPY cmd/shardingsphere-operator/main.go cmd/shardingsphere-operator/main.go
COPY api api/
COPY pkg pkg/

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build  -ldflags "-w -s" -gcflags "-N -l" -a -o operator-bin-arm64 cmd/shardingsphere-operator/main.go

# Step3: Build a multi-platform image
FROM alpine:3.15.4
COPY --from=builder-amd64 /workspace/operator-bin-amd64 . 
COPY --from=builder-arm64 /workspace/operator-bin-arm64 . 

# Add a scripts for start 
COPY start.sh /start.sh
RUN chmod +x /start.sh
ENTRYPOINT ["/start.sh"]
