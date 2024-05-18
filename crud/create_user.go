package crud

import (
	"github.com/semikoron/korocupbackend/database"
)

func CreateUserDb(user database.User) database.User {

	database.DB.Create(&user)
	return user

}
