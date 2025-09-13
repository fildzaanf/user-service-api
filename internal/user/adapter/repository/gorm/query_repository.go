package gorm

import (
	"errors"
	"user-service-api/internal/user/adapter/model"
	"user-service-api/internal/user/application/port"
	entity "user-service-api/internal/user/domain"

	"gorm.io/gorm"
)

type userQueryRepository struct {
	db *gorm.DB
}

func NewUserQueryRepository(db *gorm.DB) port.UserQueryRepositoryInterface {
	return &userQueryRepository{
		db: db,
	}
}

func (uqr *userQueryRepository) GetUserByID(id string) (entity.User, error) {
	var userModel model.User
	result := uqr.db.Where("id = ?", id).First(&userModel)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return entity.User{}, errors.New("user not found")
		}
		return entity.User{}, result.Error
	}

	userDomain := entity.UserModelToEntity(userModel)

	return userDomain, nil
}

func (uqr *userQueryRepository) GetUserByEmail(email string) (entity.User, error) {
	var user entity.User
	result := uqr.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return entity.User{}, errors.New("user not found")
		}
		return entity.User{}, result.Error
	}

	return user, nil
}
