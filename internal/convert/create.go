package convert

import (
	"fmt"
	"github.com/go-faster/errors"
	"github.com/stefanicai/transact/internal/api"
	"github.com/stefanicai/transact/internal/model"
	"math/big"
	"regexp"
	"strings"
	"time"
)

const (
	maxDescriptionLength = 50
	AmountFormat         = "amount must be a positive float value with max 2 decimals e.g. 10.50, found %s"
)

var AcceptedDateFormats = [2]string{
	"2006-01-02",
	"2006-01-02T15:04:05",
}
var amountRegEx *regexp.Regexp

func init() {
	// if this fails it's a programming error, we want it to panic early
	amountRegEx = regexp.MustCompile(`^[0-9]+\.[0-9]{1,2}$`)
}

func ToModelTransaction(req *api.CreateTransactionRequest) (*model.Transaction, error) {
	validationErrs := validateCreateTransactionRequest(req)
	if len(validationErrs) > 0 {
		return nil, errors.New(strings.Join(validationErrs, "\n"))
	}
	r := &big.Rat{}
	r, _ = r.SetString(req.Amount.Value)

	// parse
	date, _ := parseDate(req.Date.Value)

	return &model.Transaction{
		Description: req.Description.Value,
		AmountInUSD: r,
		Date:        *date,
	}, nil
}

// ParseDate checks if the date is in an accepted format and return the corresponding Time object
func parseDate(dateString string) (*time.Time, error) {
	for _, format := range AcceptedDateFormats {
		date, err := time.Parse(format, dateString)
		if err == nil {
			return &date, err
		}
	}
	return nil, errors.New(fmt.Sprintf("incorrect date passed: %s, accepted formats: %s", dateString, strings.Join(AcceptedDateFormats[:], ", ")))
}

func validateCreateTransactionRequest(req *api.CreateTransactionRequest) []string {
	var fails []string

	// description
	if !req.Description.IsSet() {
		fails = append(fails, "description must be provided")
	} else if len(req.Description.Value) > maxDescriptionLength {
		fails = append(fails, fmt.Sprintf("description cannot be longer than %d chars", maxDescriptionLength))
	}

	// amount
	if !req.Amount.IsSet() {
		fails = append(fails, "amount must be provided")
	} else {
		if ok := amountRegEx.MatchString(req.Amount.Value); !ok {
			fails = append(fails, fmt.Sprintf(AmountFormat, req.Amount.Value))
		}
	}

	// date
	_, err := parseDate(req.Date.Value)
	if err != nil {
		fails = append(fails, err.Error())
	}

	return fails
}
