package post

type Storage interface {
	CreatePost(creatorID string, post Post) (*Post, error)
}

type Service interface {
	CreatePost(creatorID string, post Post) (*Post, error)
}
