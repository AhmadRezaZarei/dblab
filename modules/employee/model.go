package employee

import "time"

type Employee struct {
	Id int64
	FirstName string
	LastName string
	BirthDate time.Time
	Rank int64
}