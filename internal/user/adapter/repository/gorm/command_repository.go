package gorm

import (
	"errors"
	entity "user-service-api/internal/user/domain"
	"user-service-api/internal/user/adapter/model"
	"user-service-api/internal/user/application/port"
	"user-service-api/pkg/constant"
	"user-service-api/pkg/crypto"

	"gorm.io/gorm"
)

type userCommandRepository struct {
	db *gorm.DB
}

func NewUserCommandRepository(db *gorm.DB) port.UserCommandRepositoryInterface {
	return &userCommandRepository{
		db: db,
	}
}

func (ucr *userCommandRepository) RegisterUser(user entity.User) (entity.User, error) {
	tx := ucr.db.Begin()
	if tx.Error != nil {
		return entity.User{}, tx.Error
	}

	userModel := entity.UserEntityToModel(user)

	if err := tx.Create(&userModel).Error; err != nil {
		tx.Rollback()
		return entity.User{}, err
	}

	userEntity := entity.UserModelToEntity(userModel)

	if err := tx.Commit().Error; err != nil {
		return entity.User{}, err
	}

	return userEntity, nil
}

func (ucr *userCommandRepository) LoginUser(email, password string) (entity.User, error) {
	tx := ucr.db.Begin()

	if tx.Error != nil {
		return entity.User{}, tx.Error
	}

	userModel := model.User{}

	result := tx.Where("email = ?", email).First(&userModel)
	if result.Error != nil {
		tx.Rollback()
		return entity.User{}, result.Error
	}

	if errComparePass := crypto.ComparePassword(userModel.Password, password); errComparePass != nil {
		tx.Rollback()
		return entity.User{}, errors.New(constant.ERROR_PASSWORD_INVALID)
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return entity.User{}, err
	}

	userEntity := entity.UserModelToEntity(userModel)

	return userEntity, nil
}
