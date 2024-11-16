package ep

import (
	"MessageService/grpcserver"
	"MessageService/internal/config"
	cvtimpl "MessageService/internal/converter/converter"
	"MessageService/internal/ep/postgres"
	"MessageService/internal/handlers/users_handler"
	"MessageService/internal/infra/postgres/users_repo_impl"
	"MessageService/internal/interceptors"
	serviceimpl "MessageService/internal/service"
	"context"
	"fmt"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
)

// Run запускает основной процесс сервера, включая инициализацию всех зависимостей.
func Run(ctx context.Context, cfg *config.Config, zapLogger *zap.Logger) error {
	// Инициализация соединения с PostgreSQL
	masterPool, err := postgres.New(ctx, &cfg.PG, zapLogger)
	if err != nil {
		return fmt.Errorf("create PostgreSQL master pool: %w", err)
	}
	defer masterPool.Close()

	userRepo := users_repo_impl.New(zapLogger, masterPool)

	serviceCvt := cvtimpl.ServiceConverterImpl{}
	repoCvt := cvtimpl.RepoConverterImpl{}

	messageService := serviceimpl.New(zapLogger, userRepo, nil, nil, repoCvt)

	srv, err := grpcserver.NewServer(
		cfg.GRPCPort,
		cfg.GatewayPort,
		cfg.ServiceName,
		cfg.PrometheusPort,
		zapLogger,
		grpcserver.WithImplementationAdapters(
			users_handler.New(zapLogger, messageService, &serviceCvt)),
		grpcserver.WithGrpcUnaryServerInterceptors(
			interceptors.APIErrorInterceptor(zapLogger),
		),
	)
	if err != nil {
		return fmt.Errorf("create gRPC server: %w", err)
	}
	zapLogger.Info("Starting gRPC server", zap.String("gRPC port", cfg.GRPCPort))

	// Обработка системных сигналов для корректного завершения работы
	quit := setupSignalChannel()
	serverErrors := make(chan error, 1)

	// Запуск gRPC сервера в горутине
	go func() {
		serverErrors <- srv.Start(ctx)
	}()

	// Ожидание ошибки сервера или сигнала завершения
	select {
	case err = <-serverErrors:
		return fmt.Errorf("gRPC server error: %w", err)
	case sig := <-quit:
		zapLogger.Info(fmt.Sprintf("Received termination signal: %s", sig))
	}

	// Корректное завершение работы gRPC сервера
	zapLogger.Info("Shutting down gRPC server gracefully...")
	srv.Stop()
	zapLogger.Info("gRPC server stopped")
	return nil
}

// setupSignalChannel настраивает канал для прослушивания системных сигналов завершения.
func setupSignalChannel() chan os.Signal {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	return quit
}
