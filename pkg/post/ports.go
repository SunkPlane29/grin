package post

//TODO: storage very subject to change
type Storage interface {
	CreatePost(post Post)
	UpdatePost(post Post)
	DeletePost(post Post)
	GetPost(UserID string, postID string)
	GetPosts(userID string)
}

type Service interface {
	CreatePost(post Post)
	UpdatePost(post Post)
	DeletePost(post Post)
	GetPost(UserID string, postID string)
	GetPosts(userID string)
}
