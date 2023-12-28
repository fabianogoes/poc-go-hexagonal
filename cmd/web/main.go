package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/demo/go-hexagonal/internal/adapter/handler"
	"github.com/demo/go-hexagonal/internal/adapter/repository"
	"github.com/demo/go-hexagonal/internal/core/service"
	"github.com/joho/godotenv"
)

func init() {
	// Init logger
	var logHandler *slog.JSONHandler

	env := os.Getenv("APP_ENV")
	if env == "production" {
		logHandler = slog.NewJSONHandler(os.Stdout, nil)
	} else {
		logHandler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		})

		// Load .env file
		err := godotenv.Load()
		if err != nil {
			slog.Error("Error loading .env file", "error", err)
			os.Exit(1)
		}
	}

	logger := slog.New(logHandler)
	slog.SetDefault(logger)
}

func main() {
	fmt.Println("Starting...")

	httpUrl := os.Getenv("HTTP_URL")
	httpPort := os.Getenv("HTTP_PORT")
	listenAddr := httpUrl + ":" + httpPort

	productRepository := repository.NewProductRepository()
	productService := service.NewProductService(productRepository)
	productHandler := handler.NewProductHandler(productService)

	router, err := handler.NewRouter(
		*productHandler,
	)
	if err != nil {
		slog.Error("Error initializing router", "error", err)
		os.Exit(1)
	}

	// Start server
	slog.Info("Starting the HTTP server", "listen_address", listenAddr)
	err = router.Serve(listenAddr)
	if err != nil {
		slog.Error("Error starting the HTTP server", "error", err)
		os.Exit(1)
	}
}
