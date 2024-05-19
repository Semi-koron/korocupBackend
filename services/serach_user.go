package services

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/semikoron/korocupbackend/crud"
)

func SearchUser(c echo.Context) error {
	searchUserName := c.Param("user_name")
	users, err := crud.FindUsersFromName(searchUserName)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "No user found")
	}
	return c.JSON(http.StatusOK, users)
}
