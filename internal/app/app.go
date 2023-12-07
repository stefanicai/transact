package app

import (
	"context"
	"github.com/stefanicai/transact/internal/api"
	"github.com/stefanicai/transact/internal/client"
	"github.com/stefanicai/transact/internal/config"
	"github.com/stefanicai/transact/internal/handler"
	"log/slog"
	"net/http"
	"os"
)

// Start starts the application server
func Start(cfg config.Config) {
	ctx := context.Background()

	//create clients and destroy them at server exit
	clients, err := client.MakeClients(ctx, cfg)
	if err != nil {
		slog.Error("clients couldn't be created, application cannot run", err)
		os.Exit(1)
	}
	defer func(clients client.Clients, ctx context.Context) {
		err := clients.Disconnect(ctx)
		slog.Error("error disconnecting clients", err)
	}(clients, ctx)

	service, err := handler.NewTransactionService(ctx, cfg, clients)
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
