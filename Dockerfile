FROM golang:1.25-alpine AS builder

FROM golang:1.25-alpine

WORKDIR /app

COPY . .

RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum ./
RUN go mod download

CMD ["air", "-c", ".air.toml"]