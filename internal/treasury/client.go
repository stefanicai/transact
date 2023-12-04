package treasury

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"time"
)

const (
	queryStringFormat = "%s?format=json" +
		"&page[number]=1" +
		"&page[size]=1" +
		"&fields=effective_date,exchange_rate,currency" +
		"&sort=-effective_date" +
		"&filter=effective_date:gte:%s,effective_date:lte:%s,country:eq:%s"
	dateFormat             = "2006-01-02"
	rateNotOlderThanMonths = 6
)

type Client interface {
	GetRate(country string, effectiveDate time.Time) (*Rate, error)
}

// Rate is an exchange rate record at a specific time
type Rate struct {
	EffectiveDate string `json:"effective_date"`
	Currency      string `json:"currency"`
	ExchangeRate  string `json:"exchange_rate"`
}

type ApiResponse struct {
	Data []Rate
}

type client struct {
	baseURL string
}

// GetRate loads the latest rate in the last 6 months. Error if we can't find one.
func (c *client) GetRate(country string, rateDate time.Time) (*Rate, error) {
	// find the oldest date accepted
	startDate := rateDate.AddDate(0, -rateNotOlderThanMonths, 0)
	formattedStartDate := startDate.Format(dateFormat)

	formattedDate := rateDate.Format(dateFormat)

	url := fmt.Sprintf(queryStringFormat, c.baseURL, formattedStartDate, formattedDate, country)
	response, err := http.Get(url)
	if err != nil {
		slog.Error("failed to get rates", "date", formattedDate, "country", country, "error", err)
		return nil, err
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		slog.Error("failed to read response body from treasury service", "date", formattedDate, "country", country, "error", err)
		return nil, err
	}

	var ar ApiResponse
	err = json.Unmarshal(responseData, &ar)
	if err != nil {
		slog.Error("failed to parse response from treasury service", "date", formattedDate, "country", country, "error", err)
		return nil, err
	}

	if len(ar.Data) == 0 {
		return nil, errors.New("no records found")
	}

	return &ar.Data[0], nil
}

func MakeClient(baseURL string) Client {
	return &client{
		baseURL: baseURL,
	}
}
