package memory

import "github.com/SunkPlane29/grin/pkg/post"

type PostStorage struct {
	posts map[string]map[string]*post.Post
}

func NewPostStorage() post.Storage {
	return &PostStorage{posts: make(map[string]map[string]*post.Post)}
}

func (ps *PostStorage) CreatePost(creatorID string, p post.Post) (*post.Post, error) {
	if ps.posts[creatorID] == nil {
		ps.posts[creatorID] = make(map[string]*post.Post)
	}

	ps.posts[creatorID][p.ID] = &p
	return &p, nil
}
