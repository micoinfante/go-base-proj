#!/bin/bash

go clean --cache && go test -v -cover authentication/...
go build -o authsvc main.go
sudo chmod +x authsvc