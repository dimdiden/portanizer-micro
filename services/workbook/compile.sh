#!/usr/bin/env sh

protodir=../../proto
rm -f ./pb/*
mkdir -p ./pb
protoc --go_out=plugins=grpc:./pb -I $protodir $protodir/workbook.proto