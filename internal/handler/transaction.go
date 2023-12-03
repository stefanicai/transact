package handler

import (
	"context"
	"fmt"
	"github.com/stefanicai/transact/internal/api"
	"github.com/stefanicai/transact/internal/forex"
	"github.com/stefanicai/transact/internal/persistence/inmem"
	"github.com/stefanicai/transact/internal/transaction"
	"log/slog"
)

// transactionHandler is
type transactionHandler struct {
	transactions transaction.Service
}

func (t *transactionHandler) Create(ctx context.Context, req *api.CreateTransactionRequest) (*api.CreateTransactionResponse, error) {
	slog.Info(fmt.Sprintf("%+v\n", req))
	return t.transactions.Create(ctx, req)
}

func (t *transactionHandler) Get(ctx context.Context, req *api.GetTransactionRequest) (*api.GetTransactionResponse, error) {
	return t.transactions.Get(ctx, req)
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

func NewTransactionService() api.Handler {
	tService := transaction.NewService(inmem.MakeTransactionDao(), forex.MakeForexService())
	th := transactionHandler{
		transactions: tService,
	}
	return &th
}
