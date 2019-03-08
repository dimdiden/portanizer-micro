#!/usr/bin/env sh

protoc users.proto --go_out=plugins=grpc:.