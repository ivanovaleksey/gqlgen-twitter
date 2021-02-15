package pubsub

import (
	"context"
	"github.com/ivanovaleksey/twitter/graph/model"
	"math/rand"
	"sync"
)

type (
	subID  = int64
	userID = int64
)

type Subscription struct {
	ID    subID
	Posts chan *model.Post
}

type PubSub struct {
	subscriptions map[userID]map[subID]Subscription
	mu            sync.RWMutex
}

func NewPubSub() *PubSub {
	return &PubSub{
		subscriptions: make(map[userID]map[subID]Subscription),
	}
}

func (ps *PubSub) Pub(ctx context.Context, post model.Post) error {
	ps.mu.RLock()
	defer ps.mu.RUnlock()

	subs, ok := ps.subscriptions[post.UserID]
	if !ok {
		return nil
	}
	for _, sub := range subs {
		sub.Posts <- &post
	}
	return nil
}

func (ps *PubSub) Sub(ctx context.Context, userID userID) (Subscription, error) {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	subs, ok := ps.subscriptions[userID]
	if !ok {
		subs = make(map[subID]Subscription)
	}
	sub := Subscription{
		ID:    rand.Int63(),
		Posts: make(chan *model.Post),
	}
	subs[sub.ID] = sub
	ps.subscriptions[userID] = subs

	return sub, nil
}

func (ps *PubSub) Unsub(ctx context.Context, id, userID int64) error {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	subs, ok := ps.subscriptions[userID]
	if !ok {
		return nil
	}

	delete(subs, id)
	if len(subs) == 0{
		delete(ps.subscriptions, userID)
	}

	return nil
}
