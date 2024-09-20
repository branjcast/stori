package types

import "time"

type TMovements struct {
	Id            int
	Date          time.Time
	Transaction   string
	CreditOrDebit string
}
