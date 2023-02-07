package database

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type Store struct {
	db *sqlx.DB
}

func New() *Store {
	return &Store{}
}

func (s *Store) GetAllChallenges(ctx context.Context) ([]Challenge, error) {
	return []Challenge{}, nil
}

func (s *Store) GetChallenge(ctx context.Context, slug string) ([]Challenge, error) {
	return []Challenge{}, nil
}
