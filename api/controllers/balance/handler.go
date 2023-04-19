package balance

import (
	"net/http"

	"github.com/marcelofelixsalgado/financial-web/api/responses"
	"github.com/marcelofelixsalgado/financial-web/api/responses/faults"
	"github.com/marcelofelixsalgado/financial-web/commons/logger"

	listBalance "github.com/marcelofelixsalgado/financial-web/pkg/usecase/balances/list"
	listPeriod "github.com/marcelofelixsalgado/financial-web/pkg/usecase/period/list"

	"github.com/labstack/echo/v4"
)

type IBalanceHandler interface {
	ListPeriod(ctx echo.Context) error
	ListBalance(ctx echo.Context) error
}

type BalanceHandler struct {
	listPeriodUseCase  listPeriod.IListPeriodUseCase
	listBalanceUseCase listBalance.IListBalanceUseCase
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

	// Response ok
	return ctx.Render(http.StatusOK, "balance.html", struct {
		PeriodId   string
		PeriodName string
		PeriodYear string
		Balances   []listBalance.OutputBalance
		Total      listBalance.OutputBalanceTotal
	}{
		PeriodId:   periodId,
		PeriodName: periodName,
		PeriodYear: periodYear,
		Balances:   output.Balances,
		Total:      output.BalanceTotal,
	})
}
