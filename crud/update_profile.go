package crud

import (
	"github.com/semikoron/korocupbackend/database"
)

type Body struct {
	UserName string `json:"user_id"`
	Icon     int    `json:"icon"`
	Profile  int    `json:"profile"`
}

func UpdateUserDb(user database.User, tmp Body) database.User {

	database.DB.Model(&user).Updates(
		database.User{
			UserName: tmp.UserName,
			Icon:     tmp.Icon,
			Profile:  tmp.Profile})

	return user

}
