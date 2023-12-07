package forex

import (
	"context"
	"github.com/go-faster/errors"
	"github.com/stefanicai/transact/internal/treasury"
	"log/slog"
	"math/big"
	"time"
)

type Service interface {
	Convert(ctx context.Context, countryName string, amount big.Rat) (*big.Rat, error)
}

type forexService struct {
	treasuryClient treasury.Client
}

func (s *forexService) Convert(ctx context.Context, countryName string, amount big.Rat) (*big.Rat, error) {
	record, err := s.treasuryClient.GetRate(ctx, countryName, time.Now().UTC())
	if err != nil {
		return nil, err
	}
	var exchangeRate, convertedAmount big.Rat
	_, ok := exchangeRate.SetString(record.ExchangeRate)
	if !ok {
		slog.Error("exchange rate could not be parsed", "rate", record.ExchangeRate)
		return nil, errors.New("invalid exchange rate")
	}
	slog.Info("record found", "record", record)
	return convertedAmount.Mul(&amount, &exchangeRate), nil
}

func MakeForexService(cfg Config) Service {
	return &forexService{
		treasuryClient: treasury.MakeClient(cfg.TreasuryURL),
	}
}
