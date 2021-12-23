package user

type User struct {
	ID          string   `json:"id"`
	Username    string   `json:"username"`
	Alias       string   `json:"alias"`
	Subscribers []string `json:"subscribers"`
}
