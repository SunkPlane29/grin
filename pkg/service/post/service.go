package post

import "github.com/segmentio/ksuid"

type service struct {
	store Storage
}

func New(s Storage) Service {
	return &service{store: s}
}

func (s *service) CreatePost(creatorID string, post Post) (*Post, error) {
	post.CreatorID = creatorID
	post.ID = ksuid.New().String()

	return s.store.CreatePost(creatorID, post)
}

func (s *service) GetPosts(creatorID string) (*[]Post, error) {
	return s.store.GetPosts(creatorID)
}

func (s *service) GetPost(creatorID, postID string) (*Post, error) {
	return s.store.GetPost(creatorID, postID)
}
