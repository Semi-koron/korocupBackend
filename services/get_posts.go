package services

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/semikoron/korocupbackend/crud"
)

func GetPosts(c echo.Context) error {
	// ここにデータベースから投稿を取得する処理を書く
	posts := crud.GetAllPost()
	return c.JSON(http.StatusOK, posts)
}
