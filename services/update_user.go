package services

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/semikoron/korocupbackend/crud"
	"github.com/semikoron/korocupbackend/database"
)

func UpdateUser(c echo.Context) error {

	profileUpdateUid := c.Param("firebase_uid")
	obj := crud.Body{}
	if err := c.Bind(&obj); err != nil {
		return err
	}
	user := database.User{}
	if err := database.DB.Where("firebase_uid = ?", profileUpdateUid).First(&user).Error; err != nil {
		return err
	}

	user = crud.UpdateUserDb(user, obj)

	return c.JSON(http.StatusOK, user)

}
