FROM golang:1.21.7-alpine3.19
WORKDIR /app/server

COPY go.sum go.mod ./
RUN go mod download
