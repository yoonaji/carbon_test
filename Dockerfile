# 1. Go 빌드용 베이스 이미지
FROM golang:1.20 AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./
RUN go build -o main .

# 2. 런타임 이미지 (작고 가볍게)
FROM debian:bullseye-slim

WORKDIR /app
COPY --from=builder /app/main .

CMD ["./main"]
