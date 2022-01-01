package post

import (
	"strconv"

	"github.com/google/uuid"
)

type service struct {
	store Storage
}

func New(s Storage) Service {
	return &service{store: s}
}

func (s *service) CreatePost(creatorID string, post Post) (*Post, error) {
	post.CreatorID = creatorID
	post.ID = strconv.Itoa(int(uuid.New().ID()))

	return s.store.CreatePost(creatorID, post)
}

func (s *service) GetPosts(creatorID string) (*[]Post, error) {
	return s.store.GetPosts(creatorID)
}

func (s *service) GetPost(creatorID, postID string) (*Post, error) {
	return s.store.GetPost(creatorID, postID)
}
