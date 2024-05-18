package services

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/semikoron/korocupbackend/crud"
	"github.com/semikoron/korocupbackend/database"
)

func GetUserProfile(c echo.Context) error {

	type profile struct {
		UserData  database.User   `json:"user_data"`
		UserPosts []database.Post `json:"user_posts"`
	}

	var userData database.User
	var userPosts []database.Post

	userName := c.Param("user_id")
	crud.GetUserProfileDb(userName, &userData, &userPosts)

	userProfile := profile{userData, userPosts}
	return c.JSON(http.StatusOK, userProfile)

}
