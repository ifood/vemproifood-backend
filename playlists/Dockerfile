FROM golang:alpine AS builder

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o main

FROM alpine:latest

COPY --from=builder /app/main /main

EXPOSE 80

ENTRYPOINT [ "/main" ]
