package handler

import (
	"context"
	"github.com/stefanicai/transact/internal/api"
	"github.com/stefanicai/transact/internal/config"
	"github.com/stefanicai/transact/internal/forex"
	"github.com/stefanicai/transact/internal/persistence"
	"github.com/stefanicai/transact/internal/persistence/inmem"
	"github.com/stefanicai/transact/internal/persistence/mongo"
	"github.com/stefanicai/transact/internal/transaction"
	"log/slog"
)

// transactionHandler is the implementation of swagger/ogen generated http handler. It handles logging and error wrapping, the functionality is delegated to a service
// In our case it's only one service, but for larger applications, we would likely be splitting the functionality between several service packages.
type transactionHandler struct {
	transactions transaction.Service
}

func (t *transactionHandler) Create(ctx context.Context, req *api.CreateTransactionRequest) (*api.CreateTransactionResponse, error) {
	slog.Debug("create transaction request received", "request", req)
	resp, err := t.transactions.Create(ctx, req)
	if err != nil {
		// fixme: add error wrapping. E.g. we don't want to pass on the internal error except for when it's a validation error. skipping for now.
		slog.Debug("error processing create transaction", "request", req, "error", err)
	}
	slog.Debug("transaction created successfully", "request", req, "resp", resp)
	return resp, err
}

func (t *transactionHandler) Get(ctx context.Context, req *api.GetTransactionRequest) (*api.GetTransactionResponse, error) {
	slog.Debug("get transaction request received", "request", req)
	resp, err := t.transactions.Get(ctx, req)
	if err != nil {
		// fixme: add error wrapping
		slog.Debug("error processing get transaction", "request", req, "error", err)
	}
	slog.Debug("transaction retrieved successfully", "request", req, "resp", resp)
	return resp, err
}

// NewError creates *ErrorStatusCode from error returned by handler.
//
// Used for common default response.
func (t *transactionHandler) NewError(ctx context.Context, err error) *api.ErrorStatusCode {
	// this isn't quite useful. We'd need to introduce our own errors to make it useful.
	return &api.ErrorStatusCode{
		Response: api.Error{Message: err.Error()},
	}
}

func NewTransactionService(cfg config.Config) (api.Handler, error) {
	var dao *persistence.TransactionDao
	if cfg.Mongo.UseMock {
		inmemDao := inmem.MakeTransactionDao()
		dao = &inmemDao
	} else {
		mongoDao, err := mongo.MakeTransactionDao(cfg.Mongo)
		if err != nil {
			return nil, err
		}
		dao = &mongoDao
	}
	tService := transaction.NewService(*dao, forex.MakeForexService(cfg.Forex))
	th := transactionHandler{
		transactions: tService,
	}
	return &th, nil
}
