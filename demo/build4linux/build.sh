#! /bin/bash

go build -o ../server ../geoindex
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ../server.exe ../geoindex