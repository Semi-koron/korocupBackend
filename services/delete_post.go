package services

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/semikoron/korocupbackend/crud"
)

func DeletePost(c echo.Context) error {

	type body struct {
		Id int `json:"id"`
	}
	imageId := body{}

	if err := c.Bind(&imageId); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	firebaseUid := c.Get("uid").(string)

	post := crud.DeletePostDb(firebaseUid, imageId.Id)
	return c.JSON(http.StatusOK, post)

}
