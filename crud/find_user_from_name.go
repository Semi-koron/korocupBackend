package crud

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/semikoron/korocupbackend/database"
)

func FindUsersFromName(name string) ([]database.User, error) {
	var users []database.User
	database.DB.Where("user_name LIKE ?", "%"+name+"%").Find(&users)
	fmt.Println(len(users))
	if len(users) == 0 {
		return users, echo.NewHTTPError(http.StatusNotFound, "No user found")
	}
	return users, nil
}
