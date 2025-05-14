package main

import (
	"fmt"
	"gRPC_service/internal/config"
)

func main() {
	// TODO: инициализировать объкт конфига
	cfg := config.MustLoad()

	fmt.Println(cfg)
	// TODO: инициализировать логгер

	// TODO: инициализировать приложение (app)

	// TODO: запустить gRPC-сервер приложения

}
