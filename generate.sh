#!/bin/bash
if [ -z "$1" ]
  then
    echo "Usage: ./generate.sh BINARYFILE"
    exit
fi

BINARY=$1 go generate
go build -o $1"_mem" -ldflags "-X main.BinaryName=$1"  memrun.go bindata.go 
rm bindata.go