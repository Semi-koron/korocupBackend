package crud

import (
	"github.com/semikoron/korocupbackend/database"
)

func DeletePostDb(firebaseUid string, imageId int) database.Post {

	post := database.Post{}
	userName := GetUsername(firebaseUid)
	database.DB.Where("user_name = ? and ID = ?", userName, imageId).Delete(&post)
	return post

}
