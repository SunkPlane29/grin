package sqlite

import (
	"context"
	"database/sql"
	"errors"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Storage struct {
	db *sql.DB
}

func (s *Storage) Close() {
	s.db.Close()
}

func Open(filepath string) (*Storage, error) {
	db, err := sql.Open("sqlite3", filepath)
	if err != nil {
		return nil, err
	}

	return &Storage{
		db: db,
	}, nil
}

func New(filepath string) (*Storage, error) {
	return Open(filepath)
}

func NewScratch(ctx context.Context, filepath string) (*Storage, error) {
	if err := os.Remove(filepath); err != nil && !errors.Is(err, os.ErrNotExist) {
		return nil, err
	}

	storage, err := Open(filepath)
	if err != nil {
		return nil, err
	}

	if err := storage.PrepareDB(ctx); err != nil {
		return nil, err
	}

	return storage, nil
}

func (s *Storage) PrepareDB(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second*30)
	defer cancel()

	const createUsersTable = `
	CREATE TABLE users (
		ID TEXT PRIMARY KEY,
		username TEXT,
		password BLOB
	)
	`

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	if _, err := tx.ExecContext(ctx, createUsersTable); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
