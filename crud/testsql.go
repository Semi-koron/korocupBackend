package crud

import (
	"context"
	"log"

	"github.com/semikoron/korocupbackend/database"
	"github.com/semikoron/korocupbackend/schema"
)

func Test() {
	if _, err := database.DB.NewCreateTable().Model((*schema.User)(nil)).Exec(context.TODO()); err != nil {
		log.Fatalln(err)
	}
	user := &schema.User{Name: "鈴木太郎"}
	if _, err := database.DB.NewInsert().Model(user).Exec(context.TODO()); err != nil {
		log.Fatalln(err)
	}
}
