package post

import (
	"context"

	"github.com/segmentio/ksuid"
)

type service struct {
	store Storage
}

func New(s Storage) Service {
	return &service{store: s}
}

func (s *service) CreatePost(ctx context.Context, creatorID string, p Post) (*Post, error) {
	p.CreatorID = creatorID
	p.ID = ksuid.New().String()

	return s.store.CreatePost(ctx, p)
}

func (s *service) GetPosts(ctx context.Context, creatorID string) (*[]Post, error) {
	return s.store.GetPosts(ctx, creatorID)
}

func (s *service) GetPost(ctx context.Context, creatorID, postID string) (*Post, error) {
	return s.store.GetPost(ctx, creatorID, postID)
}
