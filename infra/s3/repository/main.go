package repository

import (
	"base-gin-go/config"
	"base-gin-go/domain/repository"
	"base-gin-go/infra/s3"
)

func NewStorageRepositoryRepository(cfg *config.Environment, s3Client *s3.Client) repository.StorageRepository {
	return &storageRepository{
		cfg:      cfg,
		s3Client: s3Client,
	}
}

type storageRepository struct {
	cfg      *config.Environment
	s3Client *s3.Client
}
