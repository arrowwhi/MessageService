package main

import (
	"MessageService/internal/config"
	"MessageService/internal/ep"
	"MessageService/internal/logger"
	"context"
	"go.uber.org/zap"
	"log"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Загрузка конфигурации
	cfg, err := config.GetConfigFromEnv()
	if err != nil {
		log.Fatalf("Failed to load configuration: %s\n", err.Error())
	}

	// Инициализация логгера
	zapLogger := logger.NewClientZapLogger(cfg.LogLevel, cfg.ServiceName)

	// Запуск сервера
	if err = ep.Run(ctx, cfg, zapLogger); err != nil {
		zapLogger.Fatal("Run server failed: %s\n", zap.Error(err))
	}
}
