package memory

import (
	"errors"

	"github.com/SunkPlane29/grin/pkg/post"
)

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

func (ps *PostStorage) GetPosts(creatorID string) (*[]post.Post, error) {
	posts := []post.Post{}

	for _, v := range ps.posts[creatorID] {
		posts = append(posts, *v)
	}

	return &posts, nil
}

func (ps *PostStorage) GetPost(creatorID, postID string) (*post.Post, error) {
	p, ok := ps.posts[creatorID][postID]
	if !ok {
		return nil, errors.New("post not found")
	}

	return p, nil
}
