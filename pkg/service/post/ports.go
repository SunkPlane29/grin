package post

type Storage interface {
	CreatePost(creatorID string, post Post) (*Post, error)
	GetPosts(creatorID string) (*[]Post, error)
	GetPost(creatorID, postID string) (*Post, error)
}

type Service interface {
	CreatePost(creatorID string, post Post) (*Post, error)
	GetPosts(creatorID string) (*[]Post, error)
	GetPost(creatorID, postID string) (*Post, error)
}
