# STEP 1 build executable binary

FROM golang:1.10 as builder

# Download and install the latest release of dep
ADD https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64 /usr/bin/dep
RUN chmod +x /usr/bin/dep

# Copy the code from the host and compile it
RUN mkdir -p /go/src/github.com/iafoosball/users-service
WORKDIR /go/src/github.com/iafoosball/users-service
COPY Gopkg.toml Gopkg.lock ./
RUN dep ensure -vendor-only
COPY . .
WORKDIR /go/src/github.com/iafoosball/users-service/cmd/iafoosball-server/
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w" -a -installsuffix cgo -o users-service .
#iafoosball/users-service/cmd/iafoosball-server

# STEP 2 build a small image
# start from scratch
#FROM scratch

FROM scratch
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy our static executable
COPY --from=builder /go/src/github.com/iafoosball/users-service/cmd/iafoosball-server/users-service .
CMD ["./users-service","--port","4444","--host","0.0.0.0"]