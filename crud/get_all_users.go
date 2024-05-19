package crud

import (
	"github.com/semikoron/korocupbackend/database"
)

func GetUsersDb(users []database.User) []database.User {

	database.DB.Find(&users)
	return users

}
