package transaction

import (
	"context"
	"github.com/google/uuid"
	"github.com/stefanicai/transact/internal/api"
	"github.com/stefanicai/transact/internal/convert"
	"github.com/stefanicai/transact/internal/forex"
	"github.com/stefanicai/transact/internal/persistence"
	"github.com/stefanicai/transact/internal/validate"
	"golang.org/x/text/currency"
	"time"
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

	// add missing fields
	tr.ID = uuid.NewString()
	tr.Date = time.Now() //use time of the container

	err = t.dao.Store(tr)
	if err != nil {
		return nil, err
	}

	resp := &api.CreateTransactionResponse{}
	resp.ID.SetTo(tr.ID)
	resp.Date.SetTo(tr.Date)

	return resp, nil
}

func (t *service) Get(ctx context.Context, req *api.GetTransactionRequest) (*api.GetTransactionResponse, error) {
	tr, err := t.dao.Get(req.ID.Value)
	if err != nil {
		return nil, err
	}

	unit, err := currency.ParseISO(req.GetCurrency().Value)
	if err != nil {
		return nil, err
	}

	amountInCurrency, err := t.forexService.Convert(unit, tr.AmountInUSD)
	if err != nil {
		return nil, err
	}

	// build response message
	var resp api.GetTransactionResponse
	resp.ID.SetTo(tr.ID)
	resp.Amount.SetTo(amountInCurrency.FloatString(validate.AmountPrecision))
	resp.Description.SetTo(tr.Description)
	resp.AmountUSD.SetTo(tr.AmountInUSD.FloatString(validate.AmountPrecision))

	return &resp, nil
}

func NewService(dao persistence.TransactionDao, forexService forex.Service) Service {
	return &service{
		dao:          dao,
		forexService: forexService,
	}
}
