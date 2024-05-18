package services

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/semikoron/korocupbackend/crud"
)

func NewPost(c echo.Context) error {
	type body struct {
		UserName string
		Image    string `json:"image"`
		Reply    int    `json:"reply"`
		Likes    int    `json:"likes"`
	}
	obj := body{}

	// UIDからusernameの取得
	uid := c.Get("uid").(string)

	if err := c.Bind(&obj); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	obj.UserName = crud.GetUsername(uid)

	CreatedPost, _ := crud.CreatePost(obj.UserName, obj.Image, obj.Reply, obj.Likes, c)
	return c.JSON(http.StatusCreated, CreatedPost)
}
