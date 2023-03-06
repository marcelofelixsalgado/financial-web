package balance

import (
	"marcelofelixsalgado/financial-web/api/responses"
	"marcelofelixsalgado/financial-web/api/responses/faults"
	"marcelofelixsalgado/financial-web/commons/logger"
	"net/http"
	"sort"

	listBalance "marcelofelixsalgado/financial-web/pkg/usecase/balances/list"
	listPeriod "marcelofelixsalgado/financial-web/pkg/usecase/periods/list"

	"github.com/labstack/echo/v4"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

type IBalanceHandler interface {
	ListPeriod(ctx echo.Context) error
	ListBalance(ctx echo.Context) error
}

type BalanceHandler struct {
	listPeriodUseCase  listPeriod.IListPeriodUseCase
	listBalanceUseCase listBalance.IListBalanceUseCase
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

type OutputBalanceTotal struct {
	ActualAmount       string
	LimitAmount        string
	DifferenceAmount   string
	DifferenceNegative bool
}

func NewBalanceHandler(listBalanceUseCase listBalance.IListBalanceUseCase, listPeriodUseCase listPeriod.IListPeriodUseCase) IBalanceHandler {
	return &BalanceHandler{
		listBalanceUseCase: listBalanceUseCase,
		listPeriodUseCase:  listPeriodUseCase,
	}
}

func (balanceHandler *BalanceHandler) ListPeriod(ctx echo.Context) error {
	log := logger.GetLogger()

	input := listPeriod.InputListPeriodDto{}

	// Calling use case
	output, faultMessage, httpStatusCode, err := balanceHandler.listPeriodUseCase.Execute(input, ctx)
	if err != nil {
		log.Errorf("Error trying to convert the output to response body: %v", err)
		responseMessage := responses.NewResponseMessage().AddMessageByErrorCode(faults.InternalServerError)
		ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}

	// Return error response
	if httpStatusCode != http.StatusOK && httpStatusCode != http.StatusNotFound {
		log.Errorf("Internal error: %d %v", httpStatusCode, faultMessage)
		return ctx.JSON(httpStatusCode, faultMessage)
	}

	// Response ok
	return ctx.Render(http.StatusOK, "balance-period.html", struct {
		Periods []listPeriod.Period
	}{
		Periods: output.Periods,
	})
}

func (balanceHandler *BalanceHandler) ListBalance(ctx echo.Context) error {
	log := logger.GetLogger()

	periodId := ctx.QueryParam("period_id")
	periodName := ctx.QueryParam("period_name")
	periodYear := ctx.QueryParam("period_year")

	input := listBalance.InputListBalanceDto{
		PeriodId: periodId,
	}

	// Calling use case
	output, faultMessage, httpStatusCode, err := balanceHandler.listBalanceUseCase.Execute(input, ctx)
	if err != nil {
		log.Errorf("Error trying to convert the output to response body: %v", err)
		responseMesage := responses.NewResponseMessage().AddMessageByErrorCode(faults.InternalServerError)
		ctx.JSON(responseMesage.HttpStatusCode, responseMesage)
	}

	// Return error response
	if httpStatusCode != http.StatusOK && httpStatusCode != http.StatusNotFound {
		log.Errorf("Internal error: %d %v", httpStatusCode, faultMessage)
		return ctx.JSON(httpStatusCode, faultMessage)
	}

	p := message.NewPrinter(language.BrazilianPortuguese)
	var outputBalance []OutputBalance
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
		outputBalance = append(outputBalance, balanceDto)
	}

	outputBalanceTotal := OutputBalanceTotal{
		ActualAmount:       p.Sprintf("%.2f", output.BalanceTotal.ActualAmount),
		LimitAmount:        p.Sprintf("%.2f", output.BalanceTotal.LimitAmount),
		DifferenceAmount:   p.Sprintf("%.2f", output.BalanceTotal.LimitAmount-output.BalanceTotal.ActualAmount),
		DifferenceNegative: ((output.BalanceTotal.LimitAmount - output.BalanceTotal.ActualAmount) < 0),
	}

	sort.SliceStable(outputBalance, func(i, j int) bool {
		return outputBalance[i].CategoryId < outputBalance[j].CategoryId
	})

	// Response ok
	return ctx.Render(http.StatusOK, "balance.html", struct {
		PeriodId   string
		PeriodName string
		PeriodYear string
		Balances   []OutputBalance
		Total      OutputBalanceTotal
	}{
		PeriodId:   periodId,
		PeriodName: periodName,
		PeriodYear: periodYear,
		Balances:   outputBalance,
		Total:      outputBalanceTotal,
	})
}
