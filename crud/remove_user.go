package crud

import (
	"github.com/semikoron/korocupbackend/database"
)

func DeleteUserDb(firebaseUid string) database.User {

	user := database.User{}
	database.DB.Where("firebase_uid = ?", firebaseUid).Delete(&user)
	return user

}
