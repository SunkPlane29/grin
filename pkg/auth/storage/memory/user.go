package memory

import (
	"errors"

	"github.com/SunkPlane29/grin/pkg/auth/core"
)

type AuthenticationStorage struct {
	users map[string]*core.User
}

func NewAuthorizationStorage() core.AuthenticationStorage {
	return &AuthenticationStorage{
		users: make(map[string]*core.User),
	}
}

func (as *AuthenticationStorage) StoreUser(user *core.User) error {
	as.users[user.Username] = user

	return nil
}

func (as *AuthenticationStorage) GetUser(username string) (*core.User, error) {
	user, ok := as.users[username]
	if !ok {
		return nil, errors.New("user not found")
	}

	return user, nil
}
