package services

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/semikoron/korocupbackend/crud"
)

func GetPostDetail(c echo.Context) error {
	// ここにデータベースから投稿を取得する処理を書く
	postID := c.Param("postid")
	posts, err := crud.GetPostDetail(postID)
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}
	replys, _ := crud.GetReplyList(postID)
	posts = append(posts, replys...)
	return c.JSON(http.StatusOK, posts)
}
