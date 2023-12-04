package convert

import (
	"github.com/stefanicai/transact/internal/api"
	"github.com/stefanicai/transact/internal/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"math/big"
	"testing"
	"time"
)

func TestToModelTransaction(t *testing.T) {
	t.Parallel()
	t.Run("should produce the equivalent model when no validaton error", func(t *testing.T) {
		expectedDate := time.Date(2023, 12, 2, 0, 0, 0, 0, time.UTC)
		expectedAmount := big.NewRat(21, 2)
		r := api.CreateTransactionRequest{
			Description: mocks.OptString("some description"),
			Amount:      mocks.OptString("10.5"),
			Date:        mocks.OptString("2023-12-02"),
		}
		tr, err := ToModelTransaction(&r)
		assert.Nil(t, err)
		require.NotNil(t, tr)
		assert.Equal(t, expectedDate, tr.Date)
		assert.Equal(t, "some description", tr.Description)
		assert.Equal(t, *expectedAmount, tr.AmountInUSD)
	})

}

func TestValidateCreateTransactionRequest(t *testing.T) {
	t.Parallel()
	t.Run("should report all missing fields", func(t *testing.T) {
		t.Parallel()
		r := api.CreateTransactionRequest{}
		errs := validateCreateTransactionRequest(&r)
		assert.Equal(t, []string{
			"description must be provided",
			"amount must be provided",
			"incorrect date passed: , accepted formats: 2006-01-02, 2006-01-02T15:04:05",
		}, errs)
	})

	t.Run("should return empty errors array if all values are valid", func(t *testing.T) {
		t.Parallel()
		r := api.CreateTransactionRequest{
			Description: mocks.OptString("some description"),
			Amount:      mocks.OptString("10.32"),
			Date:        mocks.OptString("2023-12-02"),
		}
		errs := validateCreateTransactionRequest(&r)
		assert.Equal(t, 0, len(errs))
	})

	t.Run("should fail if date and amount are provided but invalid", func(t *testing.T) {
		t.Parallel()
		r := api.CreateTransactionRequest{
			Description: mocks.OptString("some description"),
			Amount:      mocks.OptString("10.322"),
			Date:        mocks.OptString("2023-30-02"),
		}
		errs := validateCreateTransactionRequest(&r)
		assert.Equal(t, []string{
			"amount must be a positive float value with max 2 decimals e.g. 10.50, found 10.322",
			"incorrect date passed: 2023-30-02, accepted formats: 2006-01-02, 2006-01-02T15:04:05",
		}, errs)
	})

	t.Run("should accept long date format", func(t *testing.T) {
		t.Parallel()
		r := api.CreateTransactionRequest{
			Description: mocks.OptString("some description"),
			Amount:      mocks.OptString("10.32"),
			Date:        mocks.OptString("2023-11-02T10:20:21"),
		}
		errs := validateCreateTransactionRequest(&r)
		assert.Equal(t, 0, len(errs))
	})
}
