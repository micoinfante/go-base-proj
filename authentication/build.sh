#!/bin/bash

go clean --cache && go test -v -cover authentication/...
go build -o authsvc main.go
go build -o api/authsvc api/main.go