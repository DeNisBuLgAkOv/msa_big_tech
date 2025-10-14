package main

import (
	"log/slog"
	"msa_big_tech/users/internal/di"
	"os"
)

func main() {
	// Создание DI контейнера
	container := di.NewContainer()
	defer container.Close()

	// Запуск приложения
	if err := container.Run(); err != nil {
		slog.Error("failed to serve: %v", err)
		os.Exit(1)
	}
}
