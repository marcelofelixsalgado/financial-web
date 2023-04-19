package find

import "time"

type InputFindSubCategoryDto struct {
	Id string
}

type OutputFindSubCategoryDto struct {
	Id        string    `json:"id"`
	Code      string    `json:"code"`
	Name      string    `json:"name"`
	Category  Category  `json:"category"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type Category struct {
	Id              string          `json:"id"`
	Code            string          `json:"code"`
	Name            string          `json:"name"`
	TransactionType TransactionType `json:"transaction_type"`
}

type TransactionType struct {
	Code string `json:"code"`
	Name string `json:"name"`
}
