#!/usr/bin/env bash

## Authentication backend service proto file generate shell script  (go-lang)
protoc authdashboardpb.proto --go_out=plugins=grpc:.
protoc -I=. authdashboardpb.proto --js_out=import_style=commonjs:. --grpc-web_out=import_style=commonjs,mode=grpcwebtext:.