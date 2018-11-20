# STEP 1 build executable binary

FROM golang:1.10 as builder


ENV password=${arangoPassword}

#Download the service
RUN mkdir -p /go/src/github.com/iafoosball/matches-service
#WORKDIR /go/src/github.com/iafoosball
WORKDIR /go/src/github.com/iafoosball/matches-service
COPY . .
RUN touch $password
RUN ls
#RUN git fetch --tags
#RUN latestTag=$(git describe --tags `git rev-list --tags --max-count=1`) && git checkout $latestTag
#WORKDIR /go/src/github.com/iafoosball/matches-service

# Download and install the latest release of dep
ADD https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64 /usr/bin/dep
RUN chmod +x /usr/bin/dep

#Download and install swagger in go and run codegen
RUN go get -u github.com/go-swagger/go-swagger/cmd/swagger
RUN /go/bin/swagger generate server -f /go/src/github.com/iafoosball/matches-service/matches.yml -A matches
RUN go get -u golang.org/x/net/netutil
RUN dep ensure -vendor-only

#Install tests
WORKDIR /go/src/github.com/iafoosball/matches-service/matches
RUN CGO_ENABLED=0 GOOS=linux  go test -c -ldflags="-s -w" -v

#Install the service
WORKDIR /go/src/github.com/iafoosball/matches-service/cmd/matches-server/
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w" -a -installsuffix cgo -o matches-service .

# STEP 2 build a small image
# start from scratch
# FROM scratch
FROM alpine:latest
RUN apk --no-cache add ca-certificates

# Copy our static executable
COPY --from=builder /go/src/github.com/iafoosball/matches-service/cmd/matches-server/matches-service .
COPY --from=builder /go/src/github.com/iafoosball/matches-service/matches/matches.test .
CMD ["./matches-service","--port","8000","--host","0.0.0.0", "--dbhost=matches-arangodb-stag", "--dbuser=root", "--dbpassword=matchesPassword"]