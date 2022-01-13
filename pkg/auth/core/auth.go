package core

import "context"

type AuthenticationStorage interface {
	StoreUser(ctx context.Context, user User) error
	GetUser(ctx context.Context, id string) (*User, error)
}

type AuthorizationService interface {
	CreateUser(ctx context.Context, username string, password []byte) error
	AuthenticateUser(ctx context.Context, username string, password []byte) (accessToken, refreshToken string, err error)
	RefreshToken(token string) (accessToken, refreshToken string, err error)
}
