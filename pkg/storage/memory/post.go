package memory

import (
	"errors"

	"github.com/SunkPlane29/solitude/pkg/post"
)

type PostStorage struct {
	users map[string]_user
}

func NewPostStorage(defaultUserID string) *PostStorage {
	users := make(map[string]_user)
	users[defaultUserID] = _user{d: make(map[string]*post.Post)}

	return &PostStorage{users: users}
}

type _user struct {
	d map[string]*post.Post
}

func (ps *PostStorage) CreatePost(p post.Post) (*post.Post, error) {
	ps.users[p.CreatorID].d[p.ID] = &p
	return &p, nil
}

func (ps *PostStorage) GetPost(userID string, postID string) (*post.Post, error) {
	p := ps.users[userID].d[postID]
	if p == nil {
		return nil, errors.New("post not found") //TODO: create custom error defined int ports
	}

	return p, nil
}

func (ps *PostStorage) GetPosts(userID string) (*[]post.Post, error) {
	posts := []post.Post{}
	postsMap := ps.users[userID].d
	for _, p := range postsMap {
		posts = append(posts, *p)
	}

	return &posts, nil
}

func (ps *PostStorage) UpdatePost(p post.Post) (*post.Post, error) {
	ps.users[p.CreatorID].d[p.ID] = &p
	return &p, nil
}

func (ps *PostStorage) DeletePost(creatorID string, postID string) error {
	delete(ps.users[creatorID].d, postID)
	return nil
}
