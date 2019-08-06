#!/bin/bash

curDir=$(pwd)
dir=$1
pbName=$2
if [ "${dir}" == "" ] || [ "${pbName}" == "" ]; then
    echo "Usage: ./gensrpc.sh config config"
    exit 1
fi

protoc -I=. -I=../ -I=$GOPATH/src --gogo_out=. ./${dir}/${pbName}.proto
