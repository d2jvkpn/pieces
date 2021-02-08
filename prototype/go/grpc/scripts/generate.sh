#! /bin/bash

#! /usr/bin/env bash
set -eu -o pipefail

go mod init x

go get -u google.golang.org/grpc
go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc

# protoc greet/greetpb/greet.proto --go_out==plugins=grpc:./
# protoc greet/greetpb/greet.proto --go-grpc_out=.

# protoc --go_out=plugins=grpc:. pkg/greetpb/greet.proto
protoc --go_out=./  --go-grpc_out=./  pkg/greetpb/greet.proto

protoc --go_out=./  --go-grpc_out=./  pkg/calculatorpb/calculator.proto
