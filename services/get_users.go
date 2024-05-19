package services

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/semikoron/korocupbackend/crud"
	"github.com/semikoron/korocupbackend/database"
)

func GetUsers(c echo.Context) error {

	users := []database.User{}
	users = crud.GetUsersDb(users)
	return c.JSON(http.StatusOK, users)

}
