# Docker for building and running a Go application
# Multi-stage build to keep the final image small

FROM golang:1.22-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN GOWORK=off go mod tidy

COPY . .

RUN go build -o /app/main ./cmd/api

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 8080

CMD ["/app/main"]