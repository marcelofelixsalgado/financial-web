package list

import "time"

type InputListPeriodDto struct {
}

type Period struct {
	Id        string    `json:"id"`
	Code      string    `json:"code"`
	Name      string    `json:"name"`
	Year      int       `json:"year"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}

type OutputListPeriodDto struct {
	Periods []Period `json:"-"`
}
