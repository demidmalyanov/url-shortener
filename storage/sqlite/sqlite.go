package sqlite

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/demidmalyanov/url-shortener/storage"
)

type Storage struct {
	db *sql.DB
}

// New creates new sqlite storage.
func New(path string) (*Storage, error) {
	db, err := sql.Open("sqlite3", path)

	if err != nil {
		return nil, fmt.Errorf("can`t open database %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("can`t connect to database %w", err)
	}

	return &Storage{db: db}, nil

}

// Save saves token to storage.
func (s *Storage) Save(ctx context.Context, t *storage.Token) error {
	q := `INSERT INTO tokens (url,token) VALUES (?, ?)`

	if _, err := s.db.ExecContext(ctx, q, t.Url, t.Token); err != nil {
		return fmt.Errorf("can't save token: %w", err)
	}

	return nil
}

// Get retrieves a url string by token from storage.
func (s *Storage) Get(ctx context.Context, token string) (*storage.Token, error) {
	q := `SELECT url FROM tokens WHERE token = ?`

	var url string

	err := s.db.QueryRowContext(ctx, q, token).Scan(&url)
	if err == sql.ErrNoRows {
		return nil, storage.ErrNoSuchUrl
	}
	if err != nil {
		return nil, fmt.Errorf("can't get token: %w", err)
	}

	return &storage.Token{
		Url:   url,
		Token: token,
	}, nil
}

// Init initializes new table.
func (s *Storage) Init(ctx context.Context) error {
	q := `CREATE TABLE IF NOT EXISTS tokens (url TEXT, token TEXT)`

	_, err := s.db.ExecContext(ctx, q)
	if err != nil {
		return fmt.Errorf("can't create table: %w", err)
	}

	return nil
}
