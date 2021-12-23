package post

type service struct {
	s Storage
}

func NewService(s Storage) Service {
	return &service{s: s}
}

func (s *service) CreatePost(post Post) (*Post, error) {
	return nil, nil
}

func (s *service) GetPost(postID string) (*Post, error) {
	return nil, nil
}

func (s *service) GetPosts(userID string) (*[]Post, error) {
	return nil, nil
}

func (s *service) UpdatePost(post Post) (*Post, error) {
	return nil, nil
}

func (s *service) DeletePost(postID string) error {
	return nil
}
