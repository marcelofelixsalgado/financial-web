package routes

import (
	"marcelofelixsalgado/financial-web/api/controllers"
	"marcelofelixsalgado/financial-web/api/controllers/credentials"
	"marcelofelixsalgado/financial-web/api/controllers/health"
	"marcelofelixsalgado/financial-web/api/controllers/home"
	"marcelofelixsalgado/financial-web/api/controllers/login"
	"marcelofelixsalgado/financial-web/api/controllers/logout"
	"marcelofelixsalgado/financial-web/api/controllers/period"
	"marcelofelixsalgado/financial-web/api/controllers/user"
	"marcelofelixsalgado/financial-web/api/middlewares"

	"github.com/labstack/echo/v4"
)

type Routes struct {
	loginRoutes          login.LoginRoutes
	userCredentialRoutes credentials.UserCredentialsRoutes
	userRoutes           user.UserRoutes
	homeRoutes           home.HomeRoutes
	periodRoutes         period.PeriodRoutes
	logoutRoutes         logout.LogoutRoutes
	healthRoutes         health.HealthRoutes
}

func NewRoutes(loginRoutes login.LoginRoutes, userCredentialRoutes credentials.UserCredentialsRoutes, userRoutes user.UserRoutes, homeRoutes home.HomeRoutes,
	periodRoutes period.PeriodRoutes, logoutRoutes logout.LogoutRoutes, healthRoutes health.HealthRoutes) *Routes {
	return &Routes{
		loginRoutes:          loginRoutes,
		userCredentialRoutes: userCredentialRoutes,
		userRoutes:           userRoutes,
		homeRoutes:           homeRoutes,
		periodRoutes:         periodRoutes,
		logoutRoutes:         logoutRoutes,
		healthRoutes:         healthRoutes,
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

	// period routes
	basePath, periodRoutes := routes.periodRoutes.PeriodRouteMapping()
	setupRoute(http, basePath, periodRoutes)

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
