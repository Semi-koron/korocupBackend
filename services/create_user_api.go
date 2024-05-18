package services

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/semikoron/korocupbackend/crud"
	"github.com/semikoron/korocupbackend/database"
)

func NewUser(c echo.Context) error {
	type body struct {
		UserID      string `json:"user_id"`
		FirebaseUID string `json:"firebase_uid"`
		Icon        int    `json:"icon"`
		Profile     int    `json:"profile"`
	}
	obj := body{}
	user := database.User{}

	if err := c.Bind(&obj); err != nil {
		return err
	}
	user.UserID = obj.UserID
	user.FirebaseUID = obj.FirebaseUID
	user.Icon = obj.Icon
	user.Profile = obj.Profile
	user = crud.CreateUserDb(user)

	return c.JSON(http.StatusCreated, user)

}
