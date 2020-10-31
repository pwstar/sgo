#!/bin/sh

mkdir -p bin
go build -o bin/server.darwin server/Server.go
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/server.linux server/Server.go
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o bin/server.exe server/Server.go
chmod +x bin/*
