#!/bin/bash

parent_path=$( cd "$(dirname "${BASH_SOURCE[0]}")" ; pwd -P )

protoc \
    -I=$PROJECT_ROOT/proto/
    --go_out=$parent_path/$1 \
    --go_opt=paths=source_relative \
    --go-grpc_out=$1 \
    --go-grpc_opt=paths=source_relative \
    $2
    
