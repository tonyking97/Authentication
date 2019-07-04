#!/usr/bin/env bash

## Authentication backend service proto file generate shell script  (go-lang)
protoc authpb.proto --go_out=plugins=grpc:.