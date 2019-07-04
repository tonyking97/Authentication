#!/usr/bin/env bash

## Authentication backend service proto file generate shell script  (go-lang)
#protoc ./authpb.proto -I. --go_out=plugins=grpc:.

## Authentication frontend service proto file generate shell script  (javascript)
#protoc ./authpb.proto --js_out=import_style=commonjs:.
#protoc ./authpb.proto --grpc-web_out=import_style=commonjs,mode=grpcwebtext:.

## CPP
#protoc -I . --cpp_out=. authpb.proto

## TODO() Dependency file installation
#go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
#go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
#
#go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
#go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger

## proto file and geteway file generation
protoc -I. -I. \
    -I${GOPATH}/src \
    -I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
    --go_out=plugins=grpc:. \
    authpb.proto

protoc -I. -I. \
    -I${GOPATH}/src \
    -I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
    --grpc-gateway_out=logtostderr=true:. \
    authpb.proto

protoc -I. -I. \
    -I${GOPATH}/src \
    -I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
    --swagger_out=logtostderr=true:. \
    authpb.proto



