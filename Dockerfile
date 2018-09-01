# STEP 1 build executable binary

FROM golang:1.10 as builder

# Download and install the latest release of dep
ADD https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64 /usr/bin/dep
RUN chmod +x /usr/bin/dep

#Download the service
RUN mkdir -p /go/src/github.com/iafoosball
WORKDIR /go/src/github.com/iafoosball
RUN git clone https://github.com/iafoosball/matches-service.git
WORKDIR /go/src/github.com/iafoosball/matches-service

#Download and install swagger in go and run codegen
RUN go get -u github.com/go-swagger/go-swagger/cmd/swagger
RUN /go/bin/swagger generate server -f /go/src/github.com/iafoosball/matches-service/swagger.yml -A matches

RUN dep ensure -vendor-only

#Install the service
WORKDIR /go/src/github.com/iafoosball/matches-service/cmd/matches-server/
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w" -a -installsuffix cgo -o matches-service .



# Copy the code from the host and compile it
#RUN mkdir -p /go/src/github.com/iafoosball/matches-service
#WORKDIR /go/src/github.com/iafoosball/matches-service
#COPY Gopkg.toml Gopkg.lock ./
#RUN dep ensure -vendor-only
#COPY . .





#WORKDIR /go/src/github.com/iafoosball/matches-service
#RUN /usr/local/bin/swagger generate server -f matches.yml -A matches
#RUN dep ensure -vendor-only
#
#WORKDIR /go/src/github.com/iafoosball/matches-service/cmd/matches-server/




#iafoosball/users-service/cmd/iafoosball-server

# STEP 2 build a small image
# start from scratch
#FROM scratch

FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy our static executable
COPY --from=builder /go/src/github.com/iafoosball/matches-service/cmd/matches-server/matches-service .
CMD ["./matches-service","--port","4444","--host","0.0.0.0"]