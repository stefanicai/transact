package persistence

import (
	"github.com/stefanicai/transact/internal/model"
)

type TransactionDao interface {
	Store(tr *model.Transaction) error
	Get(ID string) (*model.Transaction, error)
}
