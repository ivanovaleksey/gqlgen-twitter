package graph

import (
	"context"
	"github.com/ivanovaleksey/gqlgen-twitter/graph/model"
	"github.com/ivanovaleksey/gqlgen-twitter/pubsub"
	"github.com/ivanovaleksey/gqlgen-twitter/storage"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	storage Storage
	pubSub  PubSub
}

func NewResolver() *Resolver {
	return &Resolver{
		storage: storage.NewInMemory(),
		pubSub:  pubsub.NewPubSub(),
	}
}

type Storage interface {
	CreatePost(ctx context.Context, newPost model.NewPost) (model.Post, error)
	GetLatestPostsByUser(ctx context.Context, params storage.GetLatestPostsByUserParams) ([]model.Post, int, error)
}

type PubSub interface {
	Pub(ctx context.Context, post model.Post) error
	Sub(ctx context.Context, userID int64) (pubsub.Subscription, error)
	Unsub(ctx context.Context, id, userID int64) error
}
