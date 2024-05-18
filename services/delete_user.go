package services

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/semikoron/korocupbackend/crud"
	"github.com/semikoron/korocupbackend/database"
)

func DeleteUser(c echo.Context) error {

	user := database.User{}
	firebaseUid := c.Get("uid").(string)
	if err := crud.CanDeleteUser(firebaseUid, user); err != nil {
		return err
	}

	user = crud.DeleteUserDb(firebaseUid)
	return c.JSON(http.StatusOK, user)

}
