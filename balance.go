package main

import "fmt"

type balanceIn struct {
	Name   string
	Limit  float64
	Actual float64
}

type balanceOut struct {
	Name       string
	Limit      string
	Actual     string
	Difference string
	IsNegative bool
}

var balancesIn = []balanceIn{
	{Name: "Açougue", Limit: 400, Actual: 593.97},
	{Name: "Alimentação", Limit: 700, Actual: 853.14},
	{Name: "Casa", Limit: 300, Actual: 475.12},
	{Name: "Cabeleleira/Manicure", Limit: 400, Actual: 561},
	{Name: "Invest. carreira", Limit: 200, Actual: 162.30},
	{Name: "Consultório", Limit: 200.00, Actual: 0},
	{Name: "Desconhecido", Limit: 200.00, Actual: 3},
	{Name: "Diversos", Limit: 600.00, Actual: 284.90},
	{Name: "Farmácia", Limit: 1000, Actual: 1302.52},
	{Name: "Mercado", Limit: 1600, Actual: 2657.69},
	{Name: "Padaria", Limit: 600.00, Actual: 661.87},
	{Name: "Papelaria", Limit: 200.00, Actual: 70.18},
	{Name: "Perfumaria", Limit: 200.00, Actual: 279.57},
	{Name: "Presente", Limit: 300, Actual: 616.81},
	{Name: "Roupas", Limit: 500, Actual: 488.89},
	{Name: "Transporte", Limit: 600.00, Actual: 699.47},
}

var totalIn = balanceIn{
	Name: "Total", Limit: 8000, Actual: 9710.43,
}

func getBalances() []balanceOut {
	var balancesOut []balanceOut

	for _, v := range balancesIn {

		difference := v.Limit - v.Actual

		balanceItem := balanceOut{
			Name:       v.Name,
			Limit:      fmt.Sprintf("%.2f", v.Limit),
			Actual:     fmt.Sprintf("%.2f", v.Actual),
			Difference: fmt.Sprintf("%.2f", difference),
			IsNegative: (difference < 0),
		}

		balancesOut = append(balancesOut, balanceItem)
	}
	return balancesOut
}

func getTotal() balanceOut {

	difference := totalIn.Limit - totalIn.Actual

	return balanceOut{
		Name:       totalIn.Name,
		Limit:      fmt.Sprintf("%.2f", totalIn.Limit),
		Actual:     fmt.Sprintf("%.2f", totalIn.Actual),
		Difference: fmt.Sprintf("%.2f", difference),
		IsNegative: (difference < 0),
	}
}
