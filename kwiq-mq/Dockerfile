FROM golang:alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

RUN git clone --depth 1 --branch develop https://github.com/thebetar/kwik-mq.git .

RUN go build -o kwik-mq cmd/kwik-mq/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/kwik-mq /app/

EXPOSE 10526

CMD ["./kwik-mq"]
