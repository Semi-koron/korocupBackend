package crud

import (
	"github.com/semikoron/korocupbackend/database"
)

func GetAllPost() []database.Post {
	// ここにデータベースから投稿を取得する処理を書く
	posts := []database.Post{}
	database.DB.Find(&posts)
	return posts
}
