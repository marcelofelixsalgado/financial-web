package find

import "time"

type InputFindCategoryDto struct {
	Id string
}

type OutputFindCategoryDto struct {
	Id              string                `json:"id"`
	Code            string                `json:"code"`
	Name            string                `json:"name"`
	TransactionType TransactionTypeOutput `json:"transaction_type"`
	CreatedAt       time.Time             `json:"created_at"`
	UpdatedAt       time.Time             `json:"updated_at,omitempty"`
}

type TransactionTypeOutput struct {
	Code string `json:"code"`
	Name string `json:"name"`
}
