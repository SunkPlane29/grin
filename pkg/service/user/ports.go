package user

type Storage interface {
	CreateUser(user User) (*User, error)
	GetUser(userID string) (*User, error)
	UpdateUsername(id string, newUsername string) error
	UpdateAlias(id string, newAlias string) error
	DeleteUser(id string) error
}

type Service interface {
	CreateUser(userID string, user User) (*User, error)
	CheckUserExists(userID string) bool
	UpdateUsername(userID, newUsername string) error
	UpdateAlias(userID, newAlias string) error
	DeleteUser(userID string) error
}
