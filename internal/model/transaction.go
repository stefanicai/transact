package model

import (
	"math/big"
	"time"
)

type Transaction struct {
	ID          string
	Description string
	Date        time.Time
	AmountInUSD *big.Rat
}
