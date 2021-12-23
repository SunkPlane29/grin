package post

//TODO: storage very subject to change
type Storage interface {
	CreatePost(post Post)
	GetPost(UserID string, postID string)
	GetPosts(userID string)
	UpdatePost(post Post)
	DeletePost(post Post)
}

type Service interface {
	CreatePost(post Post)
	GetPost(UserID string, postID string)
	GetPosts(userID string)
	UpdatePost(post Post)
	DeletePost(post Post)
}
