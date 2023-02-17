package balance

import (
	"log"
	"marcelofelixsalgado/financial-web/api/responses"
	"marcelofelixsalgado/financial-web/api/responses/faults"
	"marcelofelixsalgado/financial-web/pkg/usecase/balances/list"
	"net/http"
	"sort"

	"github.com/labstack/echo/v4"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

type IBalanceHandler interface {
	ListBalance(ctx echo.Context) error
}

type BalanceHandler struct {
	listBalanceUseCase list.IListBalanceUseCase
}

type OutputBalance struct {
	Id                 string
	PeriodId           string
	CategoryId         string
	ActualAmount       string
	LimitAmount        string
	DifferenceAmount   string
	DifferenceNegative bool
}

func NewBalanceHandler(listBalanceUseCase list.IListBalanceUseCase) IBalanceHandler {
	return &BalanceHandler{
		listBalanceUseCase: listBalanceUseCase,
	}
}

func (balanceHandler *BalanceHandler) ListBalance(ctx echo.Context) error {
	input := list.InputListBalanceDto{}

	// Calling use case
	output, faultMessage, httpStatusCode, err := balanceHandler.listBalanceUseCase.Execute(input, ctx)
	if err != nil {
		log.Printf("Error trying to convert the output to response body: %v", err)
		responseMesage := responses.NewResponseMessage().AddMessageByErrorCode(faults.InternalServerError)
		ctx.JSON(responseMesage.HttpStatusCode, responseMesage)
	}

	// Return error response
	if httpStatusCode != http.StatusOK && httpStatusCode != http.StatusNotFound {
		log.Printf("Internal error: %d %v", httpStatusCode, faultMessage)
		return ctx.JSON(httpStatusCode, faultMessage)
	}

	p := message.NewPrinter(language.BrazilianPortuguese)
	var result []OutputBalance
	for _, balance := range output.Balances {

		balanceDto := OutputBalance{
			Id:                 balance.Id,
			PeriodId:           balance.PeriodId,
			CategoryId:         balance.CategoryId,
			ActualAmount:       p.Sprintf("%.2f", balance.ActualAmount),
			LimitAmount:        p.Sprintf("%.2f", balance.LimitAmount),
			DifferenceAmount:   p.Sprintf("%.2f", balance.LimitAmount-balance.ActualAmount),
			DifferenceNegative: ((balance.LimitAmount - balance.ActualAmount) < 0),
		}
		result = append(result, balanceDto)
	}

	sort.SliceStable(result, func(i, j int) bool {
		return result[i].CategoryId < result[j].CategoryId
	})

	// Response ok
	return ctx.Render(http.StatusOK, "balance.html", struct {
		Balances []OutputBalance
	}{
		Balances: result,
	})
}
