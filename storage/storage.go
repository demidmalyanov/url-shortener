package storage

import (
	"context"
	"errors"
)

type Storage interface {
	Save(ctx context.Context, t *Token) error
	Get(ctx context.Context, token string) (url string, err error)
}

type Token struct {
	Token string
	Url   string
}

var ErrNoSuchUrl = errors.New("no such url")
