package forex

import (
	"golang.org/x/text/currency"
	"math/big"
)

type Service interface {
	Convert(unit currency.Unit, amount big.Rat) (big.Rat, error)
}

type forexService struct {
}

func (s *forexService) Convert(unit currency.Unit, amount big.Rat) (big.Rat, error) {
	return amount, nil
}

func MakeForexService() Service {
	return &forexService{}
}
