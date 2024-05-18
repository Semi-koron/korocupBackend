package crud

import (
	"github.com/semikoron/korocupbackend/database"
)

func GetUserProfileDb(
	userName string,
	userData *database.User,
	userPosts *[]database.Post) {

	database.DB.Where("user_name = ?", userName).Find(&userData)

	database.DB.Where("user_name = ?", userName).Order("created_at DESC").Find(&userPosts)

}
