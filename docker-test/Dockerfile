FROM golang:1.24-alpine AS builder

WORKDIR /app
ADD . /app
#WORKDIR /app

RUN go build -o /main 

####################################

FROM alpine:3.20

WORKDIR /root/
COPY --from=builder /main .

EXPOSE 8081

CMD ["./main"]