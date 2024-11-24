# FROM golang:1.23.3-alpine3.20
FROM golang:1.23.3 AS builder
RUN useradd -ms /bin/sh -u 1001 app
USER APP

WORKDIR /app
COPY --chown=app:app . /app
RUN go mod download
RUN go build cmd/app.go 

EXPOSE 80
ENV GIN_MODE release
CMD ["./app"]

