package crud

import (
	"github.com/semikoron/korocupbackend/database"
)

func Can_UpdateUser(profileUpdateUid string, user database.User) error {

	if err := database.DB.Debug().Where("firebase_uid = ?", profileUpdateUid).First(&user).Error; err != nil {
		return err
	}

	return nil

}

func UpdateUserDb(
	user database.User,
	newName string,
	newIcon int,
	newProfile int,
	profileUpdateUid string) database.User {

	database.DB.Where("firebase_uid = ?", profileUpdateUid).First(&user)

	database.DB.Debug().Model(&user).
		Updates(
			database.User{
				UserName: newName,
				Icon:     newIcon,
				Profile:  newProfile})
	return user

}
