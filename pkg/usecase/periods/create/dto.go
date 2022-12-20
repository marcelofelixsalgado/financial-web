package create

import "time"

type InputCreatePeriodDto struct {
	Code      string    `json:"code"`
	Name      string    `json:"name"`
	Year      int       `json:"year"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}

type OutputCreatePeriodDto struct {
	Id        string    `json:"id"`
	Code      string    `json:"code"`
	Name      string    `json:"name"`
	Year      int       `json:"year"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	CreatedAt time.Time `json:"created_at"`
}
