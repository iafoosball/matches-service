# STEP 1 build executable binary

FROM golang:1.10 as builder

COPY cmd/matches-server/matches-server .
RUN ls
CMD ["./matches-server", "--port=9000", "--host=0.0.0.0"]
