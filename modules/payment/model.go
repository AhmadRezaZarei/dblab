package payment

import (
	"time"

	"github.com/shopspring/decimal"
)

type Payment struct {
	Id         int64
	EmployeeId int64
	Date       time.Time
	Salary     decimal.Decimal
}
