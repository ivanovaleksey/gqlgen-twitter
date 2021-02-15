package graph

import (
	"github.com/ivanovaleksey/twitter/graph/model"
	"github.com/ivanovaleksey/twitter/storage"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	storage Storage
}

func NewResolver() *Resolver {
	return &Resolver{storage: storage.NewInMemory()}
}

type Storage interface {
	CreatePost(model.NewPost) (model.Post, error)
}
