package sell

import "github.com/shopspring/decimal"

type SellTransaction struct {
	Id         int64
	CustomerId int64
	ProductId  int64
	Quantity   decimal.Decimal
	Price      decimal.Decimal
}
