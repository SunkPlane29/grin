package memory

import (
	"errors"

	"github.com/SunkPlane29/grin/pkg/service/user"
)

//TODO: implement user storage for CreateUser
type UserStorage struct {
	users map[string]*user.User
}

func NewUserStorage() user.Storage {
	return &UserStorage{users: make(map[string]*user.User)}
}

func (us *UserStorage) CreateUser(user user.User) (*user.User, error) {
	us.users[user.ID] = &user
	return &user, nil
}

func (us *UserStorage) GetUser(userID string) (*user.User, error) {
	user, ok := us.users[userID]
	if !ok {
		return nil, errors.New("no user found matching id")
	}
	return user, nil
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
