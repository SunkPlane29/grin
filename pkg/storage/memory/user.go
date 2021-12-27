package memory

import "github.com/SunkPlane29/grin/pkg/user"

//TODO: implement user storage for CreateUser
type UserStorage struct {
	users map[string]*user.User
}

func NewUserStorage() user.Storage {
	return &UserStorage{users: make(map[string]*user.User)}
}

func (us *UserStorage) CreateUser(userID string, user user.User) (*user.User, error) {
	us.users[userID] = &user
	return &user, nil
}

func (us *UserStorage) UpdateUsername(userID, newUsername string) error {
	us.users[userID].Username = newUsername
	return nil
}

func (us *UserStorage) UpdateAlias(userID, newAlias string) error {
	us.users[userID].Alias = newAlias
	return nil
}

func (us *UserStorage) DeleteUser(userID string) error {
	delete(us.users, userID)
	return nil
}
