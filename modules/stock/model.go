package stock

type Stock struct {
	Id    int64  `json:"id"`
	Title string `json:"title"`
}

type StockItems struct {
	Id                 int64
	StockId            int64
	PrimarySubstanceId int64
}
