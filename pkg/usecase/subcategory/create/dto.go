package create

import "time"

type InputCreateSubCategoryDto struct {
	TenantId string
	Code     string        `json:"code"`
	Name     string        `json:"name"`
	Category CategoryInput `json:"category"`
}

type OutputCreateSubCategoryDto struct {
	Id        string         `json:"id"`
	Code      string         `json:"code"`
	Name      string         `json:"name"`
	Category  CategoryOutput `json:"category"`
	CreatedAt time.Time      `json:"created_at"`
}

type CategoryInput struct {
	Id string `json:"id"`
}

type CategoryOutput struct {
	Id              string          `json:"id"`
	Code            string          `json:"code"`
	Name            string          `json:"name"`
	TransactionType TransactionType `json:"transaction_type"`
}

type TransactionType struct {
	Code string `json:"code"`
	Name string `json:"name"`
}
