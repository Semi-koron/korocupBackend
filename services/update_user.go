package services

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/semikoron/korocupbackend/crud"
)

func UpdateUser(c echo.Context) error {

	type body struct {
		UserName string `json:"user_id"`
		Icon     string `json:"icon"`
		Profile  int    `json:"profile"`
	}

	profileUpdateUid := c.Get("uid").(string)
	obj := body{}
	if err := c.Bind(&obj); err != nil {
		return err
	}

	user, err := crud.GetUser(profileUpdateUid)
	if err != nil {
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
