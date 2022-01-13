package sqlite

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/SunkPlane29/grin/pkg/service/user"
)

//TODO: implement at least create and get user
func (s *Storage) CreateUser(ctx context.Context, user user.User) (*user.User, error) {
	const insertUser = `
	INSERT INTO users (ID, username, alias)
	VALUES (?, ?, ?)
	`

	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	if _, err := tx.ExecContext(ctx, insertUser, user.ID, user.Username, user.Alias); err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *Storage) GetUser(ctx context.Context, userID string) (*user.User, error) {
	const getUser = `
	SELECT *
	FROM users
	WHERE id=?  
	`

	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	result := tx.QueryRowContext(ctx, getUser, userID)

	var dbUser user.User
	if err := result.Scan(&dbUser.ID, &dbUser.Username, &dbUser.Alias); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, user.ErrUserNotFound
		}

		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &dbUser, nil
}

func (s *Storage) UpdateUsername(ctx context.Context, id string, newUsername string) error {
	return nil
}

func (s *Storage) UpdateAlias(ctx context.Context, id string, newAlias string) error {
	return nil
}

func (s *Storage) DeleteUser(ctx context.Context, id string) error {
	return nil
}
