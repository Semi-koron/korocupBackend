package services

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/semikoron/korocupbackend/crud"
)

func SearchUser(c echo.Context) error {
	searchUserName := c.Param("user_name")
	users, _ := crud.FindUsersFromName(searchUserName)
	return c.JSON(http.StatusOK, users)
}
