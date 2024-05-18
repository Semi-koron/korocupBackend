package crud

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/semikoron/korocupbackend/database"
)

func CreatePost(userName, image string, reply, likes int, c echo.Context) (database.Post, error) {
	if err := database.DB.Where("user_name = ?", userName).First(&database.User{}).Error; err != nil {
		// エラーを直接クライアントに送信
		c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Database Error: " + err.Error(),
		})
		return database.Post{}, err // 空のPostとエラーを返す
	}
	post := database.Post{
		UserName: userName,
		Image:    image,
		Reply:    reply,
		Likes:    likes,
	}
	if err := database.DB.Create(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Create Post Error: " + err.Error(),
		})
		return database.Post{}, err // 空のPostとエラーを返す
	}
	return post, nil // 成功した場合、作成されたPostとnilを返す
}
