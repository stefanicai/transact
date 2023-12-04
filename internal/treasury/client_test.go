package treasury

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

var testDate = time.Date(2023, time.Month(12), 2, 1, 10, 30, 0, time.UTC)

func TestClient_GetRate_FailedAPICall(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	c := MakeClient(server.URL)
	rate, err := c.GetRate("Australia", time.Now())
	assert.NotNil(t, err)
	assert.Nil(t, rate)
}

func TestClient_GetRate_SuccessfulAPICall(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// test query is correct
		expectedQuery := "format=json&page[number]=1&page[size]=1&fields=effective_date,exchange_rate,currency&sort=-effective_date&filter=effective_date:gte:2023-06-02,effective_date:lte:2023-12-02,country:eq:Australia"
		assert.Equal(t, r.URL.RawQuery, expectedQuery)
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `
			{
			  "data": [
				{
				  "effective_date": "2023-09-30",
				  "exchange_rate": "77.86",
				  "currency": "Afghani"
				}
			  ]
			}
		`)
	}))
	defer server.Close()

	c := MakeClient(server.URL)
	rate, err := c.GetRate("Australia", testDate)
	assert.Nil(t, err)
	require.NotNil(t, rate, "returned rate must not be null")
	assert.Equal(t, "2023-09-30", rate.EffectiveDate)
	assert.Equal(t, "77.86", rate.ExchangeRate)
	assert.Equal(t, "Afghani", rate.Currency)
}
