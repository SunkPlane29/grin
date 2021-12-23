package post

type service struct {
	s Storage
}

func NewService(s Storage) Service {
	return &service{s: s}
}

func (s *service) CreatePost(post Post) {

}

func (s *service) GetPost(userID string, postID string) {
	//TODO: check private stuff
}

func (s *service) GetPosts(userID string) {
	//TODO: check private stuff
}

func (s *service) UpdatePost(post Post) {

}

func (s *service) DeletePost(post Post) {

}
