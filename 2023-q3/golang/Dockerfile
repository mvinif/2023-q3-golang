# FROM golang:1.23.3-alpine3.20
FROM golang:1.23.3 AS builder

WORKDIR /app
COPY . /app
RUN go mod download
RUN go build cmd/app.go 

ENV GIN_MODE=release
CMD ["./app"]

