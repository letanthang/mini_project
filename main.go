package main

import (
	"fmt"
	"log"

	"app/config"
	userRepo "app/repo/user"

	mw "app/middleware"
	"app/route"

	mv "app/myvalidator"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	fmt.Printf("config app: %+v", config.Config)
	e := echo.New()
	e.Validator = mv.NewValidator()
	e.Use(middleware.Recover())
	e.Use(mw.SimpleLogger())
	route.All(e)

	users, err := userRepo.GetUsers()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v", users)
	log.Println(e.Start(":9090"))
}
