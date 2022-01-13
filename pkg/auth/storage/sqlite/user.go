package sqlite

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/SunkPlane29/grin/pkg/auth/core"
)

func (s *Storage) StoreUser(ctx context.Context, user core.User) error {
	const createUser = `
	INSERT INTO users (ID, username, password)
	VALUES (?, ?, ?)
	`

	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	if _, err := tx.ExecContext(ctx, createUser, user.ID, user.Username, user.PasswordHash); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

//FIXME: GetUser returning ErrUserNotFound after creating user
func (s *Storage) GetUser(ctx context.Context, username string) (*core.User, error) {
	const getUser = `
	SELECT *
	FROM users
	WHERE username=?
	`

	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	result := tx.QueryRowContext(ctx, getUser, username)
	if err != nil {
		return nil, err
	}

	var dbUser core.User
	if err := result.Scan(&dbUser.ID, &dbUser.Username, &dbUser.PasswordHash); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, core.ErrUserNotFound
		}
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &dbUser, nil
}
