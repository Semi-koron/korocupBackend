package crud

import (
	"context"
	"log"

	"github.com/semikoron/korocupbackend/database"
)

func Test() {
	if _, err := database.DB.NewCreateTable().Model((*database.User)(nil)).Exec(context.TODO()); err != nil {
		log.Fatalln(err)
	}
	user := &database.User{Name: "鈴木太郎"}
	if _, err := database.DB.NewInsert().Model(user).Exec(context.TODO()); err != nil {
		log.Fatalln(err)
	}
}
