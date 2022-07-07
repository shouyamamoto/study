package main

import (
	"github/shouyamamoto/study/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.HideBanner = true

	arrowOrigins := []string{"http://localhost:8080"}
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: arrowOrigins,
	}))

	e.GET("/users", controller.GetUsers)

	e.Logger.Fatal(e.Start(":8081"))
}
