// Package validate package handles validation of request objects
// NOTE: If we generated the swagger from the model (instead of model from swagger), we could have used go-playground/validate package to do the validation via annotations
package validate

import (
	"fmt"
	"github.com/stefanicai/transact/internal/api"
	"regexp"
)

const (
	maxDescriptionLength = 50
	AmountFormat         = "amount must be a positive float value with max 2 decimals e.g. 10.50, found %s"
	AmountPrecision      = 2
)

var amountRegEx *regexp.Regexp

func init() {
	// if this fails it's a programming error, we want it to panic early
	amountRegEx = regexp.MustCompile(`^[0-9]+.[0-9]{1,2}$`)
}

func ValidateCreateTransactionRequest(req api.CreateTransactionRequest) []string {
	var fails []string

	// description
	if !req.Description.IsSet() {
		fails = append(fails, "description must be provided")
	} else if len(req.Description.Value) > maxDescriptionLength {
		fails = append(fails, fmt.Sprintf("description cannot be longer than %d chars", maxDescriptionLength))
	}

	if !req.Amount.IsSet() {
		fails = append(fails, "amount must be provided")
	} else {
		if ok := amountRegEx.MatchString(req.Amount.Value); ok {
			fails = append(fails, fmt.Sprintf(AmountFormat, req.Amount.Value))
		}
	}

	return fails
}
