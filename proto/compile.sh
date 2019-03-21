#!/usr/bin/env sh

protodir=../proto

for target in ../services/*;
do
    mkdir -p ../services/$target/pb
    protoc --go_out=plugins=grpc:../services/$target/pb -I $protodir $protodir/portanizer.proto
done