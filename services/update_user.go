package services

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/semikoron/korocupbackend/crud"
	"github.com/semikoron/korocupbackend/database"
)

func UpdateUser(c echo.Context) error {

	type body struct {
		UserName string `json:"user_id"`
		Icon     int    `json:"icon"`
		Profile  int    `json:"profile"`
	}

	profileUpdateUid := c.Get("uid").(string)
	obj := body{}
	if err := c.Bind(&obj); err != nil {
		return err
	}
	user := database.User{}

	if err := crud.CanUpdateUser(profileUpdateUid, user); err != nil {
		return err
	}

	user = crud.UpdateUserDb(
		user,
		obj.UserName,
		obj.Icon,
		obj.Profile,
		profileUpdateUid)

	return c.JSON(http.StatusOK, user)

}
