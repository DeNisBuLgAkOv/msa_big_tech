# Стадия сборки
FROM golang:1.25-alpine AS builder

RUN apk add --no-cache git protobuf make

WORKDIR /app

COPY Makefile vendor.proto.mk go.mod go.sum ./
RUN go mod download

COPY . .

RUN make vendor
RUN make generate

# Компилируем бинарник
RUN go build -o bin/chat cmd/main.go

# Стадия выполнения
FROM alpine:latest

WORKDIR /app

# Копируем бинарник
COPY --from=builder /app/bin/chat .

# Указываем порт
EXPOSE 50057

# Запускаем сервис
CMD ["./chat"]