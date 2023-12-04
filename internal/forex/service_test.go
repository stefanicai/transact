package forex

import (
	"fmt"
	"github.com/go-faster/errors"
	"github.com/stefanicai/transact/internal/treasury"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"math/big"
	"testing"
	"time"
)

type clientMock struct {
	getRate func(country string, effectiveDate time.Time) (*treasury.Rate, error)
}

func (c clientMock) GetRate(country string, effectiveDate time.Time) (*treasury.Rate, error) {
	return c.getRate(country, effectiveDate)
}

func TestConvert_Successful(t *testing.T) {
	amount := big.NewRat(12, 5)
	exchangeRate := "1.5"
	expectedAmount := big.NewRat(18, 5)
	c := clientMock{
		getRate: func(country string, effectiveDate time.Time) (*treasury.Rate, error) {
			rate := treasury.Rate{
				EffectiveDate: "not relevant",
				Currency:      "not relevant",
				ExchangeRate:  exchangeRate,
			}
			return &rate, nil
		},
	}

	s := forexService{
		treasuryClient: c,
	}

	actualAmount, err := s.Convert("not relevant", *amount)
	require.Nil(t, err)
	require.NotNil(t, actualAmount)
	assert.Equal(t, expectedAmount, actualAmount, fmt.Sprintf("mismatching values actual %s, expected %s", actualAmount.FloatString(2), expectedAmount.FloatString(2)))
}

func TestConvert_FailsWhenClientErrors(t *testing.T) {
	amount := big.NewRat(12, 5)
	c := clientMock{
		getRate: func(country string, effectiveDate time.Time) (*treasury.Rate, error) {
			return nil, errors.New("failure")
		},
	}

	s := forexService{
		treasuryClient: c,
	}

	actualAmount, err := s.Convert("not relevant", *amount)
	require.NotNil(t, err)
	require.Nil(t, actualAmount)
}
