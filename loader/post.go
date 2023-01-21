package loader

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/durmusrasit/kampus-gql/models"
)

func ReadPosts() ([]*models.Post, error) {
	var posts []*models.Post

	content, err := os.ReadFile("./posts.json")
	if err != nil {
		return nil, errors.New("An error occurred while reading file. Error:" + err.Error())
	}

	err = json.Unmarshal(content, &posts)

	if err != nil {
		return nil, errors.New("An error occurred while unmarshal posts content. Error:" + err.Error())
	}

	return posts, nil
}
