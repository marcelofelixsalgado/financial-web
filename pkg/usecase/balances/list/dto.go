package list

type InputListBalanceDto struct {
	PeriodId string
}

type OutputListBalanceDto struct {
	Balances     []Balance `json:"-"`
	BalanceTotal BalanceTotal
}

type Balance struct {
	Id           string  `json:"id"`
	PeriodId     string  `json:"period_id"`
	CategoryId   string  `json:"category_id"`
	ActualAmount float32 `json:"actual_amount"`
	LimitAmount  float32 `json:"limit_amout"`
}

type BalanceTotal struct {
	ActualAmount float32
	LimitAmount  float32
}
