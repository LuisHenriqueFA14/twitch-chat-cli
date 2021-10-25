#!/bin/bash

if test -d "./build/"; then
	rm -r ./build
fi

mkdir build
go build -o ./build/twitch ./cmd/main.go 
