package crud

import (
	"github.com/semikoron/korocupbackend/database"
)

// GetReplyDetail is a function to get reply detail
func GetReplyList(rootPostID string) ([]database.Post, error) {
	// ここにデータベースから返信を取得する処理を書く
	replys := []database.Post{}
	database.DB.Table("posts").Where("reply = ?", rootPostID).Order("created_at ASC").Find(&replys)
	return replys, nil
}
