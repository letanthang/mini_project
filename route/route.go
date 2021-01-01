package route

import (
	"app/handler"

	"github.com/labstack/echo/v4"
)

func All(e *echo.Echo) {
	Private(e)
	Staff(e)
	Public(e)
}

func Private(e *echo.Echo) {
}

func Staff(e *echo.Echo) {
}

func Public(e *echo.Echo) {
	g := e.Group("/api")
	g.GET("/health", handler.HealthCheck)
	g.Any("/user", handler.GetUsers)
	g.Any("/user/create", handler.CreateUser)

	g.Any("/team", handler.GetTeams)
	g.Any("/team/create", handler.CreateTeam)

	g.Any("/hub", handler.GetHubs)
	g.Any("/hub/create", handler.CreateHub)

}
