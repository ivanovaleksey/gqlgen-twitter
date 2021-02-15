package storage

import (
	"github.com/ivanovaleksey/twitter/graph/model"
	"math/rand"
	"time"
)

type userID = int64

type InMemory struct {
	storage map[userID][]model.Post
}

func NewInMemory() *InMemory {
	return &InMemory{
		storage: make(map[userID][]model.Post),
	}
}

func (impl *InMemory) CreatePost(newPost model.NewPost) (model.Post, error) {
	post := model.Post{
		ID:              rand.Int63(),
		PublicationDate: time.Now().Unix(),
		ContentText:     newPost.Text,
		UserID:          newPost.UserID,
	}
	impl.storage[post.UserID] = append(impl.storage[post.UserID], post)
	return post, nil
}
