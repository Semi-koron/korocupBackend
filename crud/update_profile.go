package crud

import (
	"github.com/semikoron/korocupbackend/database"
)

func UpdateUserDb(
	user database.User,
	newName string,
	newIcon string,
	newProfile int,
	profileUpdateUid string) database.User {

	database.DB.Where("firebase_uid = ?", profileUpdateUid).First(&user)

	database.DB.Model(&user).
		Updates(
			database.User{
				UserName: newName,
				Icon:     newIcon,
				Profile:  newProfile})
	return user

}
