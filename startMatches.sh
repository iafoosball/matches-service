#!/bin/bash
#go run ../../go-swagger/go-swagger/cmd/swagger/swagger.go generate server -f matches.yml -A matches


cd cmd/matches-server
go build .
cd ../..

docker-compose -f docker-compose.dev.yml up --build --force-recreate
#docker-compose up --build --force-recreate