package list

type InputListBalanceDto struct {
	PeriodId string
}

type OutputListBalanceDto struct {
	Balances     []OutputBalance
	BalanceTotal OutputBalanceTotal
}

// type BalanceTotal struct {
// 	ActualAmount float32
// 	LimitAmount  float32
// }

type OutputBalance struct {
	Id                 string
	PeriodId           string
	CategoryId         string
	ActualAmount       string
	LimitAmount        string
	DifferenceAmount   string
	DifferenceNegative bool
}

type OutputBalanceTotal struct {
	ActualAmount       string
	LimitAmount        string
	DifferenceAmount   string
	DifferenceNegative bool
}
