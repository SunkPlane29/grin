package user

type Storage interface {
	CreateUser(userID string, user User) (*User, error)
	UpdateUsername(id string, newUsername string) error
	UpdateAlias(id string, newAlias string) error
	DeleteUser(id string) error
}

type Service interface {
	CreateUser(userID string, user User) (*User, error)
	UpdateUsername(userID, newUsername string) error
	UpdateAlias(userID, newAlias string) error
	DeleteUser(userID string) error
}
