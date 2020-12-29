package main

import (
	"fmt"
	"log"

	"app/config"
	"app/repo"

	mw "app/middleware"
	"app/route"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	fmt.Printf("config app: %+v", config.Config)
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(mw.SimpleLogger())
	route.All(e)

	levels, err := repo.GetUsers()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v", levels)
	log.Println(e.Start(":9090"))
}
