package list

import (
	"encoding/json"
	"fmt"
	"io"
	"marcelofelixsalgado/financial-web/pkg/infrastructure/requests"
	"marcelofelixsalgado/financial-web/pkg/usecase/responses/faults"
	"marcelofelixsalgado/financial-web/settings"
	"net/http"
	"sort"

	"github.com/labstack/echo/v4"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

type Balance struct {
	Id           string  `json:"id"`
	PeriodId     string  `json:"period_id"`
	CategoryId   string  `json:"category_id"`
	ActualAmount float32 `json:"actual_amount"`
	LimitAmount  float32 `json:"limit_amout"`
}

type IListBalanceUseCase interface {
	Execute(InputListBalanceDto, echo.Context) (OutputListBalanceDto, faults.IFaultMessage, int, error)
}

type ListBalanceUseCase struct {
}

func NewListBalanceUseCase() IListBalanceUseCase {
	return &ListBalanceUseCase{}
}

func (listBalanceUseCase *ListBalanceUseCase) Execute(input InputListBalanceDto, ctx echo.Context) (OutputListBalanceDto, faults.IFaultMessage, int, error) {

	var balances []Balance

	balance, err := json.Marshal(input)
	if err != nil {
		return OutputListBalanceDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
	}

	url := fmt.Sprintf("%s/v1/balances?period_id=%s", settings.Config.BalanceApiURL, input.PeriodId)

	response, err := requests.MakeUpstreamRequest(ctx, http.MethodGet, url, balance, true)
	if err != nil {
		return OutputListBalanceDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
	}
	defer response.Body.Close()

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return OutputListBalanceDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
	}

	if response.StatusCode >= 400 {
		var faultMessage faults.FaultMessage
		err := json.Unmarshal(bodyBytes, &faultMessage)
		if err != nil {
			return OutputListBalanceDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
		}
		return OutputListBalanceDto{}, faultMessage, response.StatusCode, nil
	}

	if err := json.Unmarshal(bodyBytes, &balances); err != nil {
		return OutputListBalanceDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
	}

	outputListBalanceDto := OutputListBalanceDto{
		Balances:     formatBalances(balances),
		BalanceTotal: formatBalancesTotal(balances),
	}

	return outputListBalanceDto, faults.FaultMessage{}, response.StatusCode, nil
}

func formatBalances(balances []Balance) []OutputBalance {

	p := message.NewPrinter(language.BrazilianPortuguese)
	var outputBalance []OutputBalance
	for _, balance := range balances {
		balanceDto := OutputBalance{
			Id:                 balance.Id,
			PeriodId:           balance.PeriodId,
			CategoryId:         balance.CategoryId,
			ActualAmount:       p.Sprintf("%.2f", balance.ActualAmount),
			LimitAmount:        p.Sprintf("%.2f", balance.LimitAmount),
			DifferenceAmount:   p.Sprintf("%.2f", balance.LimitAmount-balance.ActualAmount),
			DifferenceNegative: ((balance.LimitAmount - balance.ActualAmount) < 0),
		}
		outputBalance = append(outputBalance, balanceDto)
	}

	sort.SliceStable(outputBalance, func(i, j int) bool {
		return outputBalance[i].CategoryId < outputBalance[j].CategoryId
	})

	return outputBalance
}

func sumBalancesTotal(balances []Balance) (float32, float32) {
	var actualAmountTotal float32
	var limitAmountTotal float32

	for _, balance := range balances {
		actualAmountTotal += balance.ActualAmount
		limitAmountTotal += balance.LimitAmount
	}
	return actualAmountTotal, limitAmountTotal
}

func formatBalancesTotal(balances []Balance) OutputBalanceTotal {

	actualAmountTotal, limitAmountTotal := sumBalancesTotal(balances)

	p := message.NewPrinter(language.BrazilianPortuguese)

	return OutputBalanceTotal{
		ActualAmount:       p.Sprintf("%.2f", actualAmountTotal),
		LimitAmount:        p.Sprintf("%.2f", limitAmountTotal),
		DifferenceAmount:   p.Sprintf("%.2f", limitAmountTotal-actualAmountTotal),
		DifferenceNegative: ((limitAmountTotal - actualAmountTotal) < 0),
	}
}
