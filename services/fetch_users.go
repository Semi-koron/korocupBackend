package services

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/semikoron/korocupbackend/crud"
	"github.com/semikoron/korocupbackend/database"
)

func FetchUsers(c echo.Context) error {

	users := []database.User{}
	users = crud.FetchUsersDb(users)
	return c.JSON(http.StatusOK, users)

}
