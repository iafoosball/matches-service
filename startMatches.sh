#!/bin/bash
#go run ../../go-swagger/go-swagger/cmd/swagger/swagger.go generate server -f matches.yml -A matches


cd cmd/matches-server
go build .
cd ../..

docker-compose up --build
#docker-compose up --build --force-recreate