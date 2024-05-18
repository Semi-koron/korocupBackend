package crud

import (
	"github.com/semikoron/korocupbackend/database"
)

func GetUsername(uid string) string {

	user := database.User{}
	database.DB.Where("firebase_uid = ?", uid).First(&user)

	return user.UserName

}
