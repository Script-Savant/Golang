# 1. Build the go binary
FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY *.go .
RUN go build -o /my-go-app

# 2. minimal runtime
FROM alpine:3.20

WORKDIR /root/
COPY --from=builder /my-go-app .

EXPOSE  8081

CMD ["./my-go-app"]
