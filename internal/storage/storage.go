package storage

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
)

type Storage struct {
	db *sql.DB
}

func New(db *sql.DB) *Storage {
	return &Storage{
		db: db,
	}
}

func (s *Storage) AddPlayer(ctx context.Context, name string) error {
	_, err := s.db.ExecContext(ctx, "insert into players values (null, ?)", name)
	if err != nil {
		return errors.Wrap(err, "can't insert new player")
	}
	return nil
}

func (s *Storage) UpdatePlayer(ctx context.Context, id string, name string) error {
	_, err := s.db.ExecContext(ctx, "update players set name = ? where id = ?", name, id)
	if err != nil {
		return errors.Wrapf(err, "can't update player %d:%s", id, name)

	}
	return nil
}
