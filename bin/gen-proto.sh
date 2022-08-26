#!/bin/bash

set -e

PROTOBUF_PATH=./proto
PROTO_THIRD_PARTY=${PROTOBUF_PATH}/third_party
SERVICES=`ls -F proto|grep /|grep -v "third_party"|sed 's/\/$//'`
for service in ${SERVICES}; do
  rm -rf proto/${service}/protobuf
  mkdir -p proto/${service}/protobuf
  exists=`ls ${PROTOBUF_PATH}/${service}/*.proto >/dev/null 2>&1; echo $?`
  if [ ! $exists -ne 0 ]; then
    protoc \
    -I ${PROTO_THIRD_PARTY} \
    --proto_path=${PROTOBUF_PATH}/${service} ${PROTOBUF_PATH}/${service}/*.proto \
    --go-grpc_out=require_unimplemented_servers=false:./proto/${service}/protobuf \
    --go_out=./proto/${service}/protobuf;
  fi
done
