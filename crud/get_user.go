package crud

import "github.com/semikoron/korocupbackend/database"

func GetUser(uid string) (database.User, error) {
	var user database.User
	if err := database.DB.Where("firebase_uid = ?", uid).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}
