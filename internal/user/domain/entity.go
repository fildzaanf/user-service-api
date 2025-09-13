package domain

import (
	"user-service-api/internal/user/adapter/model"
	"time"
)

type User struct {
	ID              string
	Name            string
	Email           string
	Password        string
	ConfirmPassword string
	Role            string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func UserEntityToModel(userEntity User) model.User {
	return model.User{
		ID:        userEntity.ID,
		Name:      userEntity.Name,
		Email:     userEntity.Email,
		Password:  userEntity.Password,
		Role:      userEntity.Role,
		CreatedAt: userEntity.CreatedAt,
		UpdatedAt: userEntity.UpdatedAt,
	}
}

func UserModelToEntity(userModel model.User) User {
	return User{
		ID:        userModel.ID,
		Name:      userModel.Name,
		Email:     userModel.Email,
		Password:  userModel.Password,
		Role:      userModel.Role,
		CreatedAt: userModel.CreatedAt,
		UpdatedAt: userModel.UpdatedAt,
	}
}
