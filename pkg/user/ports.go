package user

// TODO: storage very subject to change
type Storage interface {
	CreateUser(user User)
	ChangeAlias(userID string, newAlias string)
	ChangeUsername(userID string, newUsername string)
	AddSubscriber(userID string)
	RemoveSubscriber(userID string)
}

type Service interface {
	CreateUser(user User)
	ChangeAlias(userID string, newAlias string)
	ChangeUsername(userID string, newUsername string)
	AddSubscriber(userID string)
	RemoveSubscriber(userID string)
}
