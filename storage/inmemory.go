package storage

import (
	"github.com/ivanovaleksey/twitter/graph/model"
	"math/rand"
	"strconv"
	"time"
)

type userID = string

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
		ID:              strconv.FormatInt(rand.Int63(), 10),
		PublicationDate: int(time.Now().Unix()),
		ContentText:     newPost.Text,
		UserID:          newPost.UserID,
	}
	impl.storage[post.UserID] = append(impl.storage[post.UserID], post)
	return post, nil
}
