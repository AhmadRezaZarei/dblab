package employee

import "time"

type Employee struct {
	Id        int64     `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	BirthDate time.Time `json:"birth_date"`
	Rank      int64     `json:"rank"`
}
