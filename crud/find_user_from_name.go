package crud

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/semikoron/korocupbackend/database"
)

func FindUsersFromName(name string) ([]database.User, error) {
	var users []database.User
	if err := database.DB.Where("user_name LIKE ?", "%"+name+"%").Find(&users).Error; err != nil {
		return users, echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	return users, nil
}
