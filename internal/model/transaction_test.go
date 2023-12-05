package model

import (
	"github.com/ogen-go/ogen/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"math/big"
	"testing"
	"time"
)

// adding these tests to confirm the json marshalling/unmarshalling for big.Rat works fine when using a pointer.

func TestMarshall(t *testing.T) {
	tr := Transaction{
		ID:          "123-123",
		Description: "some descr",
		Date:        time.Date(2023, 12, 05, 0, 0, 0, 0, time.UTC),
		AmountInUSD: big.NewRat(10, 3),
	}

	result, err := json.Marshal(tr)
	require.Nil(t, err)
	assert.Equal(t, string(result), `{"ID":"123-123","Description":"some descr","Date":"2023-12-05T00:00:00Z","AmountInUSD":"10/3"}`, "it should marshall correctly all fields")
}

func TestUnmarshal(t *testing.T) {
	tr := Transaction{}

	err := json.Unmarshal([]byte(`{"ID":"123-123","Description":"some descr","Date":"2023-12-05T00:00:00Z","AmountInUSD":"10/3"}`), &tr)
	require.Nil(t, err)
	assert.Equal(t, tr.ID, "123-123")
	assert.Equal(t, tr.Description, "some descr")
	assert.Equal(t, tr.Date, time.Date(2023, 12, 05, 0, 0, 0, 0, time.UTC))
	assert.Equal(t, tr.AmountInUSD, big.NewRat(10, 3))
}
