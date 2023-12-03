package convert

import (
	"fmt"
	"github.com/go-faster/errors"
	"github.com/stefanicai/transact/internal/api"
	"github.com/stefanicai/transact/internal/model"
	"github.com/stefanicai/transact/internal/validate"
	"log/slog"
	"math/big"
)

func ToModelTransaction(req *api.CreateTransactionRequest) (*model.Transaction, error) {
	r := &big.Rat{}
	r, ok := r.SetString(req.Amount.Value)
	if !ok {
		// validation should cover this
		slog.Error(fmt.Sprintf("Failed converting string %s to big.Rat. Check validation logic!", req.Amount.Value))
		return nil, errors.New(fmt.Sprintf(validate.AmountFormat, req.Amount.Value))
	}

	return &model.Transaction{
		Description: req.Description.Value,
		AmountInUSD: *r,
	}, nil
}
