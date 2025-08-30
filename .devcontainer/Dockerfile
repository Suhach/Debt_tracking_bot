FROM golang:1.24.6 AS builder

WORKDIR /bot

COPY . .

RUN go build -o loan-bot

ENTRYPOINT [ "./loan-bot" ]

FROM alpine

WORKDIR /ln-bot

COPY --from=builder ./bot/loan-bot ./binary

ENTRYPOINT [ "./binary" ]