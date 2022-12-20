package update

import "time"

type InputUpdatePeriodDto struct {
	Id        string    `json:"_"`
	Code      string    `json:"code"`
	Name      string    `json:"name"`
	Year      int       `json:"year"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}

type OutputUpdatePeriodDto struct {
	Id        string    `json:"id"`
	Code      string    `json:"code"`
	Name      string    `json:"name"`
	Year      int       `json:"year"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
