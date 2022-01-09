package user

import "context"

type service struct {
	store Storage
}

func New(s Storage) Service {
	return &service{store: s}
}

func (s *service) CreateUser(ctx context.Context, userID string, user User) (*User, error) {
	user.ID = userID

	return s.store.CreateUser(ctx, user)
}

func (s *service) CheckUserExists(ctx context.Context, userID string) bool {
	user, err := s.store.GetUser(ctx, userID)
	if err != nil || user == nil {
		return false
	}

	return true
}

func (s *service) UpdateUsername(ctx context.Context, userID, newUsername string) error {
	return nil
}

func (s *service) UpdateAlias(ctx context.Context, userID, newAlias string) error {
	return nil
}

func (s *service) DeleteUser(ctx context.Context, userID string) error {
	return nil
}
