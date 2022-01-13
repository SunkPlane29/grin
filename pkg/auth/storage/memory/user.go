package memory

import (
	"context"
	"errors"

	"github.com/SunkPlane29/grin/pkg/auth/core"
)

type AuthenticationStorage struct {
	users map[string]core.User
}

func NewAuthorizationStorage() core.AuthenticationStorage {
	return &AuthenticationStorage{
		users: make(map[string]core.User),
	}
}

func (as *AuthenticationStorage) StoreUser(ctx context.Context, user core.User) error {
	as.users[user.Username] = user

	return nil
}

func (as *AuthenticationStorage) GetUser(ctx context.Context, username string) (*core.User, error) {
	user, ok := as.users[username]
	if !ok {
		return nil, errors.New("user not found")
	}

	return &user, nil
}
