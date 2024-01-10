package drivertrip

import (
	"time"

	"github.com/shopspring/decimal"
)

type DriverTrip struct {
	Id          int64     `json:"id"`
	SupplierId  int64     `json:"supplier_id"`
	Source      string    `json:"source"`
	Destination string    `json:"destination"`
	DriverId    int64     `json:"driver_id"`
	Date        time.Time `json:"date"`
}

type DriverTripItem struct {
	Id                 int64           `json:"id"`
	TripId             int64           `json:"trip_id"`
	PrimarySubstanceId int64           `json:"primary_substance_id"`
	Quantity           decimal.Decimal `json:"quantity"`
	Price              decimal.Decimal `json:"price"`
}

type DriverTripRequest struct {
	Model DriverTrip       `json:"model"`
	Items []DriverTripItem `json:"items"`
}
