package find

import "time"

type InputFindUserDto struct {
	Id string
}

type OutputFindUserDto struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Phone     string    `json:"phone"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
