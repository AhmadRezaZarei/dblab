package drivertrip

import (
	"time"

	"github.com/shopspring/decimal"
)

type DriverTrip struct {
	Id          int64
	SupplierId  int64
	Source      string
	Destination string
	DriverId    int64
	Date        time.Time
}

type DriverTripItem struct {
	Id                 int64
	TripId             int64
	PrimarySubstanceId int64
	Quantity           decimal.Decimal
	Price              decimal.Decimal
}
