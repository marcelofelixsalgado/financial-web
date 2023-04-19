package create

import "time"

type InputCreateCategoryDto struct {
	TenantId        string
	Code            string               `json:"code"`
	Name            string               `json:"name"`
	TransactionType TransactionTypeInput `json:"transaction_type"`
}

type OutputCreateCategoryDto struct {
	Id              string                `json:"id"`
	Code            string                `json:"code"`
	Name            string                `json:"name"`
	TransactionType TransactionTypeOutput `json:"transaction_type"`
	CreatedAt       time.Time             `json:"created_at"`
}

type TransactionTypeInput struct {
	Code string `json:"code"`
}

type TransactionTypeOutput struct {
	Code string `json:"code"`
	Name string `json:"name"`
}
