package transaction

import (
	"context"
	"github.com/google/uuid"
	"github.com/stefanicai/transact/internal/api"
	"github.com/stefanicai/transact/internal/convert"
	"github.com/stefanicai/transact/internal/forex"
	"github.com/stefanicai/transact/internal/persistence"
)

const (
	amountPrecision = 2
)

type Service interface {
	Create(ctx context.Context, req *api.CreateTransactionRequest) (*api.CreateTransactionResponse, error)
	Get(ctx context.Context, req *api.GetTransactionRequest) (*api.GetTransactionResponse, error)
}

type service struct {
	dao          persistence.TransactionDao
	forexService forex.Service
}

func (t *service) Create(ctx context.Context, req *api.CreateTransactionRequest) (*api.CreateTransactionResponse, error) {
	tr, err := convert.ToModelTransaction(req)
	if err != nil {
		return nil, err
	}

	// generate ID
	tr.ID = uuid.NewString()

	err = t.dao.Store(ctx, tr)
	if err != nil {
		return nil, err
	}

	var resp api.CreateTransactionResponse
	resp.ID.SetTo(tr.ID)

	return &resp, nil
}

func (t *service) Get(ctx context.Context, req *api.GetTransactionRequest) (*api.GetTransactionResponse, error) {
	tr, err := t.dao.Get(ctx, req.ID.Value)
	if err != nil {
		return nil, err
	}

	amountInCurrency, err := t.forexService.Convert(req.Country.Value, tr.AmountInUSD)
	if err != nil {
		return nil, err
	}

	// build response message
	var resp api.GetTransactionResponse
	resp.ID.SetTo(tr.ID)
	resp.Amount.SetTo(amountInCurrency.FloatString(amountPrecision))
	resp.Description.SetTo(tr.Description)
	resp.AmountUSD.SetTo(tr.AmountInUSD.FloatString(amountPrecision))

	return &resp, nil
}

func NewService(dao persistence.TransactionDao, forexService forex.Service) Service {
	return &service{
		dao:          dao,
		forexService: forexService,
	}
}
