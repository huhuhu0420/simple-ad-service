FROM golang:1.21.7-alpine3.19
WORKDIR /app

ENV PG_HOST=pgbouncer
ENV GIN_MODE=release

COPY . .
RUN go mod download
RUN go build -o main .

EXPOSE 5000
CMD ["./main"]