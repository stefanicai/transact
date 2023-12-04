package inmem

import (
	"context"
	"errors"
	"github.com/stefanicai/transact/internal/model"
	"github.com/stefanicai/transact/internal/persistence"
)

type transactionDao struct {
	store map[string]model.Transaction
}

// Store persists a create.go
func (t *transactionDao) Store(ctx context.Context, tr *model.Transaction) error {
	t.store[tr.ID] = *tr
	return nil
}

func (t *transactionDao) Get(ctx context.Context, ID string) (*model.Transaction, error) {
	tr, ok := t.store[ID]
	if !ok {
		return nil, errors.New("can't find id")
	}

	return &tr, nil
}

func MakeTransactionDao() persistence.TransactionDao {
	return &transactionDao{
		store: make(map[string]model.Transaction),
	}
}
