package app

import (
	"github.com/stefanicai/transact/internal/api"
	"github.com/stefanicai/transact/internal/config"
	"github.com/stefanicai/transact/internal/handler"
	"log/slog"
	"net/http"
	"os"
)

// Start starts the application server
func Start(cfg config.Config) {
	service, err := handler.NewTransactionService(cfg)
	if err != nil {
		slog.Error("failed to create the api handler", err)
	}

	transactionServer, err := api.NewServer(service)
	if err != nil {
		slog.Error("server couldn't be created", err)
		os.Exit(1)
	}

	httpServer := http.Server{
		Addr:    ":8080",
		Handler: transactionServer,
	}

	slog.Info("starting server", "addr", httpServer.Addr)
	slog.Error("failed to start server", httpServer.ListenAndServe())
}
