FROM golang:1.25-alpine AS builder
RUN apk add --no-cache git
WORKDIR /app
ENV GOPROXY=https://proxy.golang.org,direct
# Копируем go.mod и go.sum для gateway
COPY gateway/go.mod gateway/go.sum ./
# Копируем локальные модули целиком
COPY auth ../auth
COPY chat ../chat
COPY social ../social
COPY users ../users
# Загружаем зависимости
RUN go mod download
# Копируем код gateway
COPY gateway .
# Компилируем приложение
RUN go build -o ./bin/gateway ./cmd/main.go
FROM alpine:latest
RUN apk add --no-cache ca-certificates
WORKDIR /app
COPY --from=builder /app/bin/gateway ./
EXPOSE 8080
CMD ["./gateway"]