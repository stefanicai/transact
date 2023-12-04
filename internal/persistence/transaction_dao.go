package persistence

import (
	"context"
	"github.com/stefanicai/transact/internal/model"
)

type TransactionDao interface {
	Store(ctx context.Context, tr *model.Transaction) error
	Get(ctx context.Context, ID string) (*model.Transaction, error)
}
