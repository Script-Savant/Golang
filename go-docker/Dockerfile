FROM golang:1.25.3-alpine3.22 AS builder

ENV CGO_ENABLED=1
ENV GOOS=linux

WORKDIR /app

RUN apk add --no-cache build-base postgresql-dev

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

FROM alpine:latest

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /root/

COPY --from=builder /app/main .

EXPOSE 8080

CMD [ "./main" ]