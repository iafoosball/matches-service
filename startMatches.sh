#!/bin/bash
cd cmd/matches-server
go build main.go
cd ../..

docker-compose up --build --force-recreate