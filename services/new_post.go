package services

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/semikoron/korocupbackend/crud"
	"github.com/semikoron/korocupbackend/database"
)

func NewPost(c echo.Context) error {
	post := database.Post{}
	if err := c.Bind(&post); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	ctx := c.Request().Context()
	if err := crud.CreatePost(post, ctx); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusCreated, post)
}
