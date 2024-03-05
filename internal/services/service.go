package services

import (
	"github.com/Vafelence/fifa_project/internal/storage"
)

type Service struct {
	storage *storage.Storage
}

func New(storage *storage.Storage) *Service {
	return &Service{
		storage: storage,
	}
}
