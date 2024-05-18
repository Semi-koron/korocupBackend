package services

import (
	"errors"
	"net/http"

	"github.com/semikoron/korocupbackend/crud"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Login(c echo.Context) error {
	uid := c.Get("uid").(string)
	user, err := crud.GetUser(uid)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
	}
	return c.JSON(http.StatusOK, user)
}
