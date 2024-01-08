package stock

type Stock struct {
	Id int64
	Title string
}

type StockItems struct {
	Id int64
	StockId int64
	PrimarySubstanceId int64
}

