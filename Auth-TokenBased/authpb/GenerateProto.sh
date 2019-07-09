#!/usr/bin/env bash

## Authentication backend service proto file generate shell script  (go-lang)
protoc authpb.proto --go_out=plugins=grpc:.
protoc -I=. auth.proto --js_out=import_style=commonjs:. --grpc-web_out=import_style=commonjs,mode=grpcwebtext:.