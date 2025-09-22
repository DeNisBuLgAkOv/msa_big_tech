# Стадия сборки
FROM golang:1.25-alpine AS builder

RUN apk add --no-cache git protobuf make

WORKDIR /app

# Копируем ВСЕ файлы перед выполнением make
COPY . .

RUN make vendor
RUN make generate

# Компилируем бинарник
RUN go build -o bin/users cmd/main.go

# Стадия выполнения
FROM alpine:latest

WORKDIR /app

# Копируем бинарник
COPY --from=builder /app/bin/users .

# Указываем порт
EXPOSE 50055

# Запускаем сервис
CMD ["./users"]