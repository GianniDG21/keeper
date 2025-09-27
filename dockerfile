FROM golang:1.23-alpine as builder

WORKDIR /app

COPY . .

RUN ls -la
RUN cat go.mod

RUN go mod tidy -v

RUN ls -la ./cmd/
RUN ls -la ./cmd/api/

RUN go build -x -v -o /app/main ./cmd/api

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["/app/main"]