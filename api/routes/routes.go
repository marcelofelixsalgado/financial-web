package routes

import (
	"github.com/marcelofelixsalgado/financial-web/api/controllers"
	"github.com/marcelofelixsalgado/financial-web/api/controllers/balance"
	"github.com/marcelofelixsalgado/financial-web/api/controllers/credentials"
	"github.com/marcelofelixsalgado/financial-web/api/controllers/health"
	"github.com/marcelofelixsalgado/financial-web/api/controllers/home"
	"github.com/marcelofelixsalgado/financial-web/api/controllers/login"
	"github.com/marcelofelixsalgado/financial-web/api/controllers/logout"
	"github.com/marcelofelixsalgado/financial-web/api/controllers/period"
	"github.com/marcelofelixsalgado/financial-web/api/controllers/transactiontype"
	"github.com/marcelofelixsalgado/financial-web/api/controllers/user"
	"github.com/marcelofelixsalgado/financial-web/api/middlewares"

	"github.com/labstack/echo/v4"
)

type Routes struct {
	loginRoutes           login.LoginRoutes
	userCredentialRoutes  credentials.UserCredentialsRoutes
	userRoutes            user.UserRoutes
	homeRoutes            home.HomeRoutes
	transactionTypeRoutes transactiontype.TransactionTypeRoutes
	periodRoutes          period.PeriodRoutes
	balanceRoutes         balance.BalanceRoutes
	logoutRoutes          logout.LogoutRoutes
	healthRoutes          health.HealthRoutes
}

func NewRoutes(loginRoutes login.LoginRoutes, userCredentialRoutes credentials.UserCredentialsRoutes, userRoutes user.UserRoutes, homeRoutes home.HomeRoutes,
	transactionTypeRoutes transactiontype.TransactionTypeRoutes,
	periodRoutes period.PeriodRoutes, balanceRoutes balance.BalanceRoutes, logoutRoutes logout.LogoutRoutes, healthRoutes health.HealthRoutes) *Routes {
	return &Routes{
		loginRoutes:           loginRoutes,
		userCredentialRoutes:  userCredentialRoutes,
		userRoutes:            userRoutes,
		homeRoutes:            homeRoutes,
		transactionTypeRoutes: transactionTypeRoutes,
		periodRoutes:          periodRoutes,
		balanceRoutes:         balanceRoutes,
		logoutRoutes:          logoutRoutes,
		healthRoutes:          healthRoutes,
	}
}

func (routes *Routes) RouteMapping(http *echo.Echo) {

	// login routes
	basePath, loginRoutes := routes.loginRoutes.LoginRouteMapping()
	setupRoute(http, basePath, loginRoutes)

	// user credentials routes
	basePath, userCredentialsRoutes := routes.userCredentialRoutes.UserCredentialsRouteMapping()
	setupRoute(http, basePath, userCredentialsRoutes)

	// user routes
	basePath, userRoutes := routes.userRoutes.UserRouteMapping()
	setupRoute(http, basePath, userRoutes)

	// home routes
	basePath, homeRoutes := routes.homeRoutes.HomeRouteMapping()
	setupRoute(http, basePath, homeRoutes)

	// transaction type routes
	basePath, transactionTypeRoutes := routes.transactionTypeRoutes.TransactionTypeRouteMapping()
	setupRoute(http, basePath, transactionTypeRoutes)

	// period routes
	basePath, periodRoutes := routes.periodRoutes.PeriodRouteMapping()
	setupRoute(http, basePath, periodRoutes)

	// balance routes
	basePath, balanceRoutes := routes.balanceRoutes.BalanceRouteMapping()
	setupRoute(http, basePath, balanceRoutes)

	// logout routes
	basePath, logoutRoutes := routes.logoutRoutes.LogoutRouteMapping()
	setupRoute(http, basePath, logoutRoutes)

	// health routes
	basePath, healthRoutes := routes.healthRoutes.HealthRouteMapping()
	setupRoute(http, basePath, healthRoutes)

	// Setting the path for static files in assets folder
	// fileServer := http.FileServer(http.Dir("./web/assets/"))
	// router.PathPrefix("/web/assets/").Handler(http.StripPrefix("/web/assets/", fileServer))

	// return router
}

func setupRoute(http *echo.Echo, basePath string, routes []controllers.Route) {
	group := http.Group(basePath)

	for _, route := range routes {
		if route.RequiresAuthentication {
			group.Add(route.Method, route.URI, route.Function, middlewares.Authenticate)
		} else {
			group.Add(route.Method, route.URI, route.Function)
		}
	}
}
