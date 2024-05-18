package database

import (
	"context"
	"log"
)

func InitDB() {
	if _, err := DB.NewCreateTable().Model((*User)(nil)).Exec(context.TODO()); err != nil {
		log.Fatalln(err)
	}
	if _, err := DB.NewCreateTable().
		Model((*Post)(nil)).
		ForeignKey(`("user_id") REFERENCES "users" ("user_id") ON DELETE CASCADE`).
		Exec(context.TODO()); err != nil {
		log.Fatalln(err)
	}
	if _, err := DB.NewCreateTable().Model((*Like)(nil)).
		ForeignKey(`("user_id") REFERENCES "users" ("user_id") ON DELETE CASCADE`).
		ForeignKey(`("post_id") REFERENCES "posts" ("id") ON DELETE CASCADE`).
		Exec(context.TODO()); err != nil {
		log.Fatalln(err)
	}
}
