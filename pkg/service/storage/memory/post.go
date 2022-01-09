package memory

import (
	"context"
	"errors"

	"github.com/SunkPlane29/grin/pkg/service/post"
)

type PostStorage struct {
	posts map[string]map[string]*post.Post
}

func NewPostStorage() post.Storage {
	return &PostStorage{posts: make(map[string]map[string]*post.Post)}
}

func (ps *PostStorage) CreatePost(ctx context.Context, p post.Post) (*post.Post, error) {
	if ps.posts[p.CreatorID] == nil {
		ps.posts[p.CreatorID] = make(map[string]*post.Post)
	}

	ps.posts[p.CreatorID][p.ID] = &p
	return &p, nil

}

func (ps *PostStorage) GetPosts(ctx context.Context, creatorID string) (*[]post.Post, error) {
	posts := []post.Post{}

	for _, v := range ps.posts[creatorID] {
		posts = append(posts, *v)
	}

	return &posts, nil
}

func (ps *PostStorage) GetPost(ctx context.Context, creatorID, postID string) (*post.Post, error) {
	p, ok := ps.posts[creatorID][postID]
	if !ok {
		return nil, errors.New("post not found")
	}

	return p, nil
}
