#!/bin/bash

# install generator
go install github.com/golang/protobuf/protoc-gen-go

protoc voucher/voucherpb/voucher.proto --go_out=plugins=grpc:.