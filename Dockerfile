# 1. Go 빌드용 베이스 이미지
FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./
RUN go build -o main .

# 2. 런타임 이미지 (작고 가볍게)
FROM ubuntu:22.04

WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/app.env .

EXPOSE 8080

CMD ["./main"]
