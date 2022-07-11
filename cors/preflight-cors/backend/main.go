package main

import (
	"github/shouyamamoto/study/shared/infra/database"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.HideBanner = true

	arrowOrigins := []string{"http://localhost:1234"}
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: arrowOrigins,
		AllowMethods: []string{"GET", "DELETE", "OPTIONS"},
	}))

	e.GET("/", database.GetAlbums)
	e.DELETE("/albums/delete/:id", database.Delete)

	e.Logger.Fatal(e.Start(":8888"))
}
