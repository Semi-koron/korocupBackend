package crud

import (
	"github.com/semikoron/korocupbackend/database"
)

func CanDeleteUser(firebaseUid string, user database.User) error {

	if err := database.DB.Debug().Where("firebase_uid = ?", firebaseUid).First(&user).Error; err != nil {
		return err
	}

	return nil

}

func DeleteUserDb(firebaseUid string) database.User {

	user := database.User{}
	database.DB.Where("firebase_uid = ?", firebaseUid).Delete(&user)
	return user

}
