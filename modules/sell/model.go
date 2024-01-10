package sell

import "github.com/shopspring/decimal"

type SellTransaction struct {
	Id         int64           `json:"id"`
	CustomerId int64           `json:"customer_id"`
	ProductId  int64           `json:"product_id"`
	Quantity   decimal.Decimal `json:"quantity"`
	Price      decimal.Decimal `json:"price"`
}
