package storage

import (
	"context"
	"github.com/ivanovaleksey/twitter/graph/model"
	"math/rand"
	"sync"
	"time"
)

type userID = int64

type InMemory struct {
	storage map[userID][]model.Post
	mu      sync.RWMutex
}

func NewInMemory() *InMemory {
	return &InMemory{
		storage: make(map[userID][]model.Post),
	}
}

func (impl *InMemory) CreatePost(ctx context.Context, newPost model.NewPost) (model.Post, error) {
	post := model.Post{
		ID:              rand.Int63(),
		PublicationDate: time.Now().Unix(),
		ContentText:     newPost.Text,
		UserID:          newPost.UserID,
	}
	impl.mu.Lock()
	defer impl.mu.Unlock()
	posts := impl.storage[post.UserID]
	posts = append(posts, post)
	impl.storage[post.UserID] = posts
	return post, nil
}

type GetLatestPostsByUserParams struct {
	UserID userID
	Limit  int
	Offset int
}

func (impl *InMemory) GetLatestPostsByUser(ctx context.Context, params GetLatestPostsByUserParams) ([]model.Post, int, error) {
	const (
		defaultLimit  = 2
	)

	impl.mu.RLock()
	defer impl.mu.RUnlock()

	posts, ok := impl.storage[params.UserID]
	if !ok {
		return nil, 0, nil
	}

	if params.Limit <= 0 {
		params.Limit = defaultLimit
	}

	total := len(posts)

	if params.Offset > len(posts) {
		return nil, total, nil
	}
	posts = posts[params.Offset:]

	if params.Limit > len(posts) {
		params.Limit = len(posts)
	}
	posts = posts[:params.Limit]

	return posts, total, nil
}
