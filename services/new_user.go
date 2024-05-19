package services

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/semikoron/korocupbackend/crud"
	"github.com/semikoron/korocupbackend/database"
)

func NewUser(c echo.Context) error {
	type body struct {
		UserName    string `json:"user_name"`
		FirebaseUID string
		Icon        int `json:"icon"`
		Profile     int `json:"profile"`
	}
	obj := body{}
	user := database.User{}

	if err := c.Bind(&obj); err != nil {
		return err
	}

	user.UserName = obj.UserName
	user.FirebaseUID = c.Get("uid").(string)
	user.Icon = obj.Icon
	user.Profile = obj.Profile
	if user.UserName == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "User name is required.")
	}
	user = crud.CreateUserDb(user)

	if user.ID == 0 {
		return echo.NewHTTPError(http.StatusConflict, "This username is already exist.")
	}

	return c.JSON(http.StatusCreated, user)

}
