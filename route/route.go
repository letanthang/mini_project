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
	// g := e.Group("/api/user/v1/private")
	// JWTConfig := middleware.JWTConfig{
	// 	SigningKey: []byte(config.Config.JWTSecret.JWTKey),
	// 	Claims:     &model.UserClaims{},
	// }
	// g.Use(middleware.JWTWithConfig(JWTConfig))
	// g.PUT("/user", handler.UpdateUser)
}

func Staff(e *echo.Echo) {
	//g := e.Group("/api/user/v1/staff")
	//g.POST("/student", handler.AddStudent)
	//g.PUT("/student", handler.UpdateStudent)
}

func Public(e *echo.Echo) {
	g := e.Group("/api/user/v1/public")
	g.GET("/health", handler.HealthCheck)
	// g.POST("/user/register", handler.RegisterUser)
	// g.PATCH("/user/login", handler.LoginUser)

}
