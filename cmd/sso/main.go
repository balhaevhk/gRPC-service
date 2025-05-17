package main

import (
	"grpc-service/internal/config"
	"grpc-service/internal/lib/logger/handlers/slogpretty"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"grpc-service/internal/app"
)

const (
	envLocal = "local"
	envDev = "dev"
	envProd = "prod"

)
func main() {
	// TODO: инициализировать объкт конфига
	cfg := config.MustLoad()

	// TODO: инициализировать логгер
	log := setupLogger(cfg.Env)

	log.Info("starting application", slog.Any("cfg", cfg))

	application := app.New(log, cfg.GRPC.Port, cfg.StoragePath, cfg.TokenTTL)

	go application.GRPCServer.MustRun()

	// TODO: инициализировать приложение (app)

	// TODO: запустить gRPC-сервер приложения

	// Gruceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop

	application.GRPCServer.Stop()
	log.Info("Gracefully stopped")

}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = setupPrettySlog()
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}

func setupPrettySlog() *slog.Logger {
opts := slogpretty.PrettyHandlerOptions{
	SlogOpts: &slog.HandlerOptions{
		Level: slog.LevelDebug,
	},
}

handler := opts.NewPrettyHandler(os.Stdout)
return slog. New(handler)
}


