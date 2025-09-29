FROM golang:1.23-alpine as builder

WORKDIR /app

RUN go install github.com/swaggo/swag/cmd/swag@latest

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

RUN swag init -g cmd/api/main.go

RUN go build -o /app/main ./cmd/api

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .
COPY --from=builder /app/docs ./docs

EXPOSE 8080

CMD ["/app/main"]
