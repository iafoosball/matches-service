# STEP 1 build executable binary

FROM golang:1.10 as builder

COPY cmd/matches-server/main .
RUN ls
CMD ["./main", "--port=4444", "--host=0.0.0.0"]
