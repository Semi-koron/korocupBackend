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
	newProfile int) database.User {

	database.DB.Model(&user).Updates(
		database.User{
			UserName: newName,
			Icon:     newIcon,
			Profile:  newProfile})

	return user

}
