package controller

import (
	"github/shouyamamoto/study/db"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetUsers(c echo.Context) error {
	users := []*db.User{
		{
			Name: "Yamamoto",
			Age:  26,
		},
		{
			Name: "Suzuki",
			Age:  30,
		},
	}
	if err := c.Bind(users); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, users)
}
