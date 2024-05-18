package crud

import (
	"strconv"

	"github.com/semikoron/korocupbackend/database"
)

func GetPostDetail(postID string) ([]database.Post, error) {
	// ここにデータベースから投稿を取得する処理を書く
	posts := []database.Post{}
	post := database.Post{}
	if err := database.DB.Table("posts").Where("id = ?", postID).Take(&post).Error; err != nil {
		return []database.Post{}, err
	}
	// 返信の場合、返信先の投稿を取得
	if post.Reply != 0 {
		rootPost := database.Post{}
		if err := database.DB.Table("posts").Where("id = ?", strconv.Itoa(post.Reply)).Take(&rootPost).Error; err != nil {
			return []database.Post{}, err
		}
		posts = append(posts, rootPost)
	}
	posts = append(posts, post)
	return posts, nil
}
