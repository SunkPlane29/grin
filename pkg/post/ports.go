package post

//TODO: storage very subject to change
type Storage interface {
	CreatePost(post Post) (*Post, error)
	GetPost(userID string, postID string) (*Post, error)
	GetPosts(userID string) (*[]Post, error)
	UpdatePost(post Post) (*Post, error)
	DeletePost(postID string) error
}

type Service interface {
	CreatePost(post Post) (*Post, error)
	GetPost(postID string) (*Post, error)
	GetPosts(userID string) (*[]Post, error)
	UpdatePost(post Post) (*Post, error)
	DeletePost(postID string) error
}
