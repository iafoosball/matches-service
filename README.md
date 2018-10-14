# Matches Service

This service handles all data relevant for matches, this includes the goals,
the player IDs, how many players per side, the goal speed and more. 
All endpoints are specific to a match and for the moment all endpoints can be used.
In the future we plan to introduce a live match service, based on Websockets and
upload information as batch opererations to this service.
<br />
<br />
On a side note, the service does not handle any personal data and when 
excluding the player IDs the data can be used for statistics without further
considerations.   


## How to build and run this service
As a prerequiste you need to have docker and docker-compose installed. 
Installation guides can be found online.  <br />
The service is based on Golang and thus requires a working installation.
Download the repo to your local machine. <br />
```go get github.com/iafoosball/matches-service``` <br />
We use dep as our dependency management tool. To install it use <br /> 
```go get -u github.com/golang/dep/cmd/dep``` <br />
which should put a binary in your go `bin` folder.  <br />
Next get the go-swagger library, which is used to produce a sevrer 
from a openAPI spec file.  <br />
```go get github.com/go-swagger/go-swagger/cmd/swagger``` <br />
From inside the matches-service folder, where the `matches.yml` file and
 all other configuration files are, ensure all dependencies and produce the server  <br />
 ```../../../../bin/dep ensure -vendor-only```  <br />
 and for the server <br />
 ```go run ../../go-swagger/go-swagger/cmd/swagger/swagger.go generate server -f matches.yml -A matches``` <br />
Finally, make the `startMatches.sh` file executable and execute it by using `chmod +x startMatches.sh && ./startMatches.sh`  <br />
This will start an the database (ArangoDB) and 




## TODOS
Should we integrate godocs? Can be easily produced via 
`godoc -http=:3333 -goroot /home/joe/go/src/github.com/iafoosball/matches-service/matches
` 
They don't provide any real value. Just leaving it here as food for thought.
<br />
<br />
Add central swagger ui serer and documentation



