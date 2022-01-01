package user

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Alias    string `json:"alias"` //could be cool if the user had more than 1 alias
}
