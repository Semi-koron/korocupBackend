package crud

import (
	"context"
	"log"

	"github.com/semikoron/korocupbackend/database"
)

func Test() {
	user := &database.User{
		Name:     "鈴木太郎",
		Email:    "example.com",
		Password: "password",
		Icon:     1,
		Profile:  1,
	}
	if _, err := database.DB.NewInsert().Model(user).Exec(context.TODO()); err != nil {
		log.Fatalln(err)
	}
}
