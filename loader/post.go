package loader

import (
	"github.com/durmusrasit/kampus-gql/models"
)

func (db *DB) ReadPosts() ([]*models.Post, error) {
	var posts []*models.Post

	result := db.DB.Find(&posts)
	if result == nil {
		return nil, result.Error
	}

	return posts, nil
}
