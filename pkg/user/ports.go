package user

// TODO: storage very subject to change
type Storage interface {
	CreateUser(user User) (*User, error)
	GetUser(userID string) (*User, error)
	ChangeAlias(userID string, newAlias string) (*User, error)
	ChangeUsername(userID string, newUsername string) (*User, error)
	DeleteUser(userID string) error
	AddSubscriber(userID string, subscriberID string)
	GetSubscribed(userID string) (*[]User, error)
	RemoveSubscriber(userID string, subscriberID string)
}

type Service interface {
	CreateUser(user User) (*User, error)
	GetUser(userID string) (*User, error)
	ChangeAlias(userID string, newAlias string) (*User, error)
	ChangeUsername(userID string, newUsername string) (*User, error)
	DeleteUser(userID string) error
	AddSubscriber(userID string, subscriberID string) error
	GetSubscribed(userID string) (*[]User, error)
	RemoveSubscriber(userID string, subscriberID string) error
}
