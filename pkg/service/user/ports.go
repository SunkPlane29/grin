package user

import "context"

type Storage interface {
	CreateUser(ctx context.Context, user User) (*User, error)
	GetUser(ctx context.Context, userID string) (*User, error)
	UpdateUsername(ctx context.Context, id, newUsername string) error
	UpdateAlias(ctx context.Context, id, newAlias string) error
	DeleteUser(ctx context.Context, id string) error
}

type Service interface {
	//TODO: get user function???
	CreateUser(ctx context.Context, userID string, user User) (*User, error)
	CheckUserExists(ctx context.Context, userID string) bool
	UpdateUsername(ctx context.Context, userID, newUsername string) error
	UpdateAlias(ctx context.Context, userID, newAlias string) error
	DeleteUser(ctx context.Context, userID string) error
}
