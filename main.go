package main

import (
	"github.com/stefanicai/transact/internal/api"
	"github.com/stefanicai/transact/internal/handler"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	service := handler.NewTransactionService()

	transactionServer, err := api.NewServer(service)
	if err != nil {
		slog.Error("server couldn't be created", err)
		os.Exit(1)
	}

	httpServer := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: transactionServer,
	}

	slog.Info("starting server")
	slog.Error("failed to start server", httpServer.ListenAndServe())
}
