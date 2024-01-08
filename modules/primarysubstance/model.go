package primarysubstance

import (
	"time"

	"github.com/shopspring/decimal"
)

type PrimarySubstance struct {
	Id    int64
	Title string
}

type PrimarySubstanceTransaction struct {
	Id                 int64
	PrimarySubstanceId int64
	Quantity           decimal.Decimal
	Price              decimal.Decimal
	StockId            int64
	SupplierId         int64
	Date               time.Time
	Type               int64
}
