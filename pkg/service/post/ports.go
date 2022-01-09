package post

import "context"

type Storage interface {
	CreatePost(ctx context.Context, post Post) (*Post, error)
	GetPosts(ctx context.Context, creatorID string) (*[]Post, error)
	GetPost(ctx context.Context, creatorID, postID string) (*Post, error)
}

type Service interface {
	CreatePost(ctx context.Context, creatorID string, post Post) (*Post, error)
	GetPosts(ctx context.Context, creatorID string) (*[]Post, error)
	GetPost(ctx context.Context, creatorID, postID string) (*Post, error)
}
