# FROM golang:1.23.3-alpine3.20
FROM golang:1.23.3 AS builder

WORKDIR /app
COPY . .
RUN go mod download
RUN go build app.go 

EXPOSE 8080
CMD ["./app"]

