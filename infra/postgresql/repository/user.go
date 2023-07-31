package repository

import (
	"base-gin-go/domain/entity"
	"base-gin-go/domain/repository"
	"base-gin-go/infra/postgresql"
	"base-gin-go/infra/postgresql/model"
	dataPkg "base-gin-go/pkg/data"
	"context"
)

type userRepository struct {
	db          *postgresql.Database
	dataService dataPkg.Service
}

func NewUserRepository(db *postgresql.Database, dataService dataPkg.Service) repository.UserRepository {
	return &userRepository{
		db:          db,
		dataService: dataService,
	}
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	user := &model.User{}
	err := r.db.WithContext(ctx).Model(&model.User{}).
		Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	result := &entity.User{}
	err = r.dataService.Copy(result, user)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *userRepository) GetByID(ctx context.Context, id int) (*entity.User, error) {
	user := &model.User{}
	err := r.db.WithContext(ctx).Model(&model.User{}).
		Where("id = ?", id).
		Error
	if err != nil {
		return nil, err
	}
	result := &entity.User{}
	err = r.dataService.Copy(result, user)
	if err != nil {
		return nil, err
	}
	return result, nil
}
