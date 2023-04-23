#!/usr/bin/env bash

protoc --go_out=./pkg/models --go_opt=paths=source_relative \
       --go-grpc_out=./pkg/service --go-grpc_opt=paths=source_relative \
       --grpc-gateway_out=./pkg/service --grpc-gateway_opt=logtostderr=true --grpc-gateway_opt=paths=source_relative \
       --openapiv2_out=./api/swagger --openapiv2_opt=logtostderr=true \
       proto/*.proto
