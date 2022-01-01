package post

type Post struct {
	ID        string `json:"id"`
	CreatorID string `json:"creator_id"`
	Content   string `json:"content"`
}
