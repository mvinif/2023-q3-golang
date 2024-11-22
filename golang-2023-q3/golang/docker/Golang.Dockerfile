# FROM golang:1.23.3-alpine3.20
FROM golang:1.23.3

WORKDIR \app

COPY go.* ./

RUN go mod download


