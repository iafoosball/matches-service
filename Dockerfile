# STEP 1 build executable binary

FROM golang:1.10 as builder

COPY cmd/matches-server/main .
RUN ls
CMD ["./main", "--port=9000", "--host=0.0.0.0"]
