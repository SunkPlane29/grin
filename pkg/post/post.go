package post

import "time"

type Post struct {
	ID           string    `json:"id"`
	Content      string    `json:"content"`
	CreatorID    string    `json:"creator_id"`
	CreationTime time.Time `json:"creation_time"`
}
