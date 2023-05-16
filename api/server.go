package api

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"text/template"
	"time"

	"github.com/marcelofelixsalgado/financial-web/api/controllers/balance"
	"github.com/marcelofelixsalgado/financial-web/api/controllers/credentials"
	"github.com/marcelofelixsalgado/financial-web/api/controllers/health"
	"github.com/marcelofelixsalgado/financial-web/api/controllers/home"
	"github.com/marcelofelixsalgado/financial-web/api/controllers/login"
	"github.com/marcelofelixsalgado/financial-web/api/controllers/logout"
	"github.com/marcelofelixsalgado/financial-web/api/controllers/period"
	"github.com/marcelofelixsalgado/financial-web/api/controllers/transactiontype"
	"github.com/marcelofelixsalgado/financial-web/api/controllers/user"
	"github.com/marcelofelixsalgado/financial-web/api/cookies"
	"github.com/marcelofelixsalgado/financial-web/api/middlewares"
	"github.com/marcelofelixsalgado/financial-web/api/routes"
	"github.com/marcelofelixsalgado/financial-web/api/utils"
	"github.com/marcelofelixsalgado/financial-web/commons/logger"
	"github.com/marcelofelixsalgado/financial-web/settings"

	userCreate "github.com/marcelofelixsalgado/financial-web/pkg/usecase/user/create"
	userDelete "github.com/marcelofelixsalgado/financial-web/pkg/usecase/user/delete"
	userFind "github.com/marcelofelixsalgado/financial-web/pkg/usecase/user/find"
	userUpdate "github.com/marcelofelixsalgado/financial-web/pkg/usecase/user/update"

	userCredentialsCreate "github.com/marcelofelixsalgado/financial-web/pkg/usecase/credentials/create"
	userCredentialsLogin "github.com/marcelofelixsalgado/financial-web/pkg/usecase/credentials/login"
	userCredentialsUpdate "github.com/marcelofelixsalgado/financial-web/pkg/usecase/credentials/update"

	periodCreate "github.com/marcelofelixsalgado/financial-web/pkg/usecase/period/create"
	periodDelete "github.com/marcelofelixsalgado/financial-web/pkg/usecase/period/delete"
	periodFind "github.com/marcelofelixsalgado/financial-web/pkg/usecase/period/find"
	periodList "github.com/marcelofelixsalgado/financial-web/pkg/usecase/period/list"
	periodUpdate "github.com/marcelofelixsalgado/financial-web/pkg/usecase/period/update"

	transactionTypeList "github.com/marcelofelixsalgado/financial-web/pkg/usecase/transactiontype/list"

	balanceList "github.com/marcelofelixsalgado/financial-web/pkg/usecase/balances/list"

	logs "github.com/marcelofelixsalgado/financial-web/commons/logger"

	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
)

// Server this is responsible for running an http server
type Server struct {
	http   *echo.Echo
	routes *routes.Routes
	stop   chan struct{}
}

func NewServer() *Server {
	// Load environment variables
	settings.Load()

	server := &Server{
		stop: make(chan struct{}),
	}

	// Configure cookies
	cookies.Configure()

	return server
}

// Run is the procedure main for start the application
func (s *Server) Run() {
	s.startServer()
	<-s.stop
}

type TemplateRenderer struct {
	templates *template.Template
}

func (server *Server) startServer() {
	go server.watchStop()

	server.http = echo.New()
	logger := logs.GetLogger()
	logger.Infof("Server is starting now in %s.", settings.Config.Environment)

	// Load HTML templates
	server.http.Renderer = utils.LoadTemplates()

	// Setup static files (*.js *.css)
	server.http.Static("/web/assets/", "web/assets/")
	server.http.Static("/web/charts/", "web/charts/")

	// Prometheus
	p := prometheus.NewPrometheus("echo", nil)
	p.Use(server.http)

	// Middlewares
	server.http.Use(middlewares.Logger())

	loginRoutes := setupLoginRoutes()
	userCredentialsRoutes := setupUserCredentialsRoutes()
	userRoutes := setupUserRoutes()
	homeRoutes := setupHomeRoutes()
	periodRoutes := setupPeriodRoutes()
	transactionTypeRoutes := setupTransactionTypeRoutes()
	balanceRoutes := setupBalanceRoutes()
	logoutRoutes := setupLogoutRoutes()
	healthRoutes := setupHealthRoutes()

	// Setup all routes
	routes := routes.NewRoutes(loginRoutes, userCredentialsRoutes, userRoutes, homeRoutes, transactionTypeRoutes, periodRoutes, balanceRoutes, logoutRoutes, healthRoutes)

	routes.RouteMapping(server.http)

	server.routes = routes

	showRoutes(server.http)

	addr := fmt.Sprintf(":%v", settings.Config.WebHttpPort)
	go func() {
		if err := server.http.Start(addr); err != nil {
			logger.Errorf("Shutting down the server now: ", err)
		}
	}()
}

// watchStop wait for the interrupt signal.
func (server *Server) watchStop() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	logger.GetLogger().Info(<-stop)
	server.stopServer()
}

// stopServer stops the server http
func (s *Server) stopServer() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(settings.Config.ServerCloseWait))
	defer cancel()

	logger := logs.GetLogger()
	logger.Info("Server is stoping...")
	s.http.Shutdown(ctx)
	close(s.stop)
}

func setupLoginRoutes() login.LoginRoutes {

	// setup Use Cases (services)
	loginUseCase := userCredentialsLogin.NewLoginUseCase()

	// setup router handlers
	loginHandler := login.NewLoginHandler(loginUseCase)

	// setup routes
	loginRoutes := login.NewLoginRoutes(loginHandler)

	return loginRoutes
}

func setupUserRoutes() user.UserRoutes {

	// setup Use Cases (services)
	createUseCase := userCreate.NewCreateUseCase()
	updateUseCase := userUpdate.NewUpdateUseCase()
	findUseCase := userFind.NewFindUseCase()
	deleteUseCase := userDelete.NewDeleteUseCase()

	// setup router handlers
	userHandler := user.NewUserHandler(createUseCase, updateUseCase, findUseCase, deleteUseCase)

	// setup routes
	userRoutes := user.NewUserRoutes(userHandler)

	return userRoutes
}

func setupUserCredentialsRoutes() credentials.UserCredentialsRoutes {

	// setup Use Cases (services)
	createUseCase := userCredentialsCreate.NewCreateUseCase()
	updateUseCase := userCredentialsUpdate.NewUpdateUseCase()
	loginUseCase := userCredentialsLogin.NewLoginUseCase()

	// setup router handlers
	userCredentialsHandler := credentials.NewUserCredentialsHandler(createUseCase, updateUseCase, loginUseCase)

	// setup routes
	userCredentialsRoutes := credentials.NewUserCredentialsRoutes(userCredentialsHandler)

	return userCredentialsRoutes
}

func setupPeriodRoutes() period.PeriodRoutes {

	// setup Use Cases (services)
	createUseCase := periodCreate.NewCreatePeriodUseCase()
	listUseCase := periodList.NewListPeriodUseCase()
	findUseCase := periodFind.NewFindPeriodUseCase()
	updateUseCase := periodUpdate.NewUpdatePeriodUseCase()
	deleteUseCase := periodDelete.NewDeletePeriodUseCase()

	// setup router handlers
	periodHandler := period.NewPeriodHandler(createUseCase, listUseCase, findUseCase, updateUseCase, deleteUseCase)

	// setup routes
	periodRoutes := period.NewPeriodRoutes(periodHandler)

	return periodRoutes
}

func setupTransactionTypeRoutes() transactiontype.TransactionTypeRoutes {

	// setup Use Cases (services)
	listUseCase := transactionTypeList.NewListTransactionTypeUseCase()

	// setup router handlers
	handler := transactiontype.NewTransactionTypeHandler(listUseCase)

	// setup routes
	routes := transactiontype.NewTransactionTypeRoutes(handler)

	return routes
}

func setupBalanceRoutes() balance.BalanceRoutes {
	// setup Use Cases (services)
	balanceListUseCase := balanceList.NewListBalanceUseCase()
	periodListUseCase := periodList.NewListPeriodUseCase()

	// setup router handlers
	balanceHandler := balance.NewBalanceHandler(balanceListUseCase, periodListUseCase)

	// setup routes
	balanceRoutes := balance.NewBalanceRoutes(balanceHandler)

	return balanceRoutes
}

func setupHomeRoutes() home.HomeRoutes {

	// setup router handlers
	homeHandler := home.NewHomeHandler()

	// setup routes
	homeRoutes := home.NewHomeRoutes(homeHandler)

	return homeRoutes
}

func setupLogoutRoutes() logout.LogoutRoutes {
	// setup router handlers
	logoutHandler := logout.NewLogoutHandler()

	// setup routes
	logoutRoutes := logout.NewLogoutRoutes(logoutHandler)

	return logoutRoutes
}

func setupHealthRoutes() health.HealthRoutes {
	// setup router handlers
	healthHandler := health.NewHealthHandler()

	// setup routes
	healthRoutes := health.NewHealthRoutes(healthHandler)

	return healthRoutes
}

func showRoutes(e *echo.Echo) {
	var routes = e.Routes()
	logger := logger.GetLogger()

	if len(routes) > 0 {
		for _, route := range routes {
			// if strings.Contains(route.Name, "forklift-api") {
			logger.Infof("%6s: %s \n", route.Method, route.Path)
			// }
		}
	}
}
