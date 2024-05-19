package services

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/semikoron/korocupbackend/crud"
)

func DeleteUser(c echo.Context) error {

	firebaseUid := c.Get("uid").(string)
	_, err := crud.GetUser(firebaseUid)
	if err != nil {
		return err
	}

	user := crud.DeleteUserDb(firebaseUid)
	return c.JSON(http.StatusOK, user)

}
