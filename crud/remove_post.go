package crud

import (
	"github.com/semikoron/korocupbackend/database"
)

func CanDeletePost(firebaseUid string, post database.Post) error {

	userName := GetUsername(firebaseUid)
	if err := database.DB.Debug().Where("user_name = ?", userName).First(&post).Error; err != nil {
		return err
	}

	return nil

}

func DeletePostDb(firebaseUid string, imageId int) database.Post {

	post := database.Post{}
	userName := GetUsername(firebaseUid)
	database.DB.Where("user_name = ? and ID = ?", userName, imageId).Delete(&post)
	return post

}
