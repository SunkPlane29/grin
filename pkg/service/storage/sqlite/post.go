package sqlite

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/SunkPlane29/grin/pkg/service/post"
)

//TODO: test em all

func (s *Storage) CreatePost(ctx context.Context, p post.Post) (*post.Post, error) {
	const createPost = `
	INSERT INTO posts
	VALUES (?, ?, ?)
	`

	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	if _, err := tx.ExecContext(ctx, createPost, p.ID, p.CreatorID, p.Content); err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &p, nil
}

func (s *Storage) GetPosts(ctx context.Context, creatorID string) (*[]post.Post, error) {
	const getPosts = `
	SELECT *
	FROM posts
	WHERE creatorID=?
	`

	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	result, err := tx.QueryContext(ctx, getPosts, creatorID)
	if err != nil {
		return nil, err
	}
	defer result.Close()

	var posts = []post.Post{}
	for result.Next() {
		var p post.Post

		if err := result.Scan(&p.ID, &p.CreatorID, &p.Content); err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, post.ErrNoPosts
			}

			return nil, err
		}

		posts = append(posts, p)
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &posts, nil
}

func (s *Storage) GetPost(ctx context.Context, creatorID string, postID string) (*post.Post, error) {
	const getPost = `
	SELECT *
	FROM posts
	WHERE ID=? AND creatorID=?
	`

	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	result := tx.QueryRowContext(ctx, getPost, postID, creatorID)

	var dbPost post.Post
	if err := result.Scan(&dbPost.ID, &dbPost.CreatorID, &dbPost.Content); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, post.ErrNoPosts
		}

		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &dbPost, nil
}
