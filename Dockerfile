# Build stage
FROM golang:1.19-alpine3.16 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go
RUN apk add curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.1/migrate.linux-amd64.tar.gz | tar xvz

# Run Stage
FROM alpine:3.16
WORKDIR /app
COPY  --from=builder /app/main .
COPY  --from=builder /app/migrate ./migrate
COPY app.env .
COPY start.sh .
COPY wait-for-it.sh .
RUN ["chmod", "+x", "/app/wait-for-it.sh", "/app/start.sh"]
COPY db/migration ./migration

EXPOSE 8081
EXPOSE 8082
EXPOSE 8083
