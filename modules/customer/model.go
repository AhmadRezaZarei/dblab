package customer

type Customer struct {
	Id          int64
	Title       string `json:"title"`
	PhoneNumber string `json:"phone_number"`
}
