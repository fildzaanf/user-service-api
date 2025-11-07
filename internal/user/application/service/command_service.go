package service

import (
	"context"
	"errors"
	"user-service-api/internal/user/application/port"
	entity "user-service-api/internal/user/domain"
	"user-service-api/pkg/constant"
	"user-service-api/pkg/crypto"
	"user-service-api/pkg/middleware"
	"user-service-api/pkg/validator"
)

type userCommandService struct {
	userCommandRepository port.UserCommandRepositoryInterface
	userQueryRepository   port.UserQueryRepositoryInterface
}

func NewUserCommandService(ucr port.UserCommandRepositoryInterface, uqr port.UserQueryRepositoryInterface) port.UserCommandServiceInterface {
	return &userCommandService{
		userCommandRepository: ucr,
		userQueryRepository:   uqr,
	}
}

func (ucs *userCommandService) RegisterUser(ctx context.Context, user entity.User) (entity.User, error) {

	errEmpty := validator.IsDataEmpty(
		[]string{"name", "email", "role", "password", "confirm_password"},
		user.Name, user.Email, user.Role, user.Password, user.ConfirmPassword,
	)

	if errEmpty != nil {
		return entity.User{}, errEmpty
	}

	errRole := validator.IsRoleValid(user.Role)
	if errRole != nil {
		return entity.User{}, errRole
	}

	errEmailValid := validator.IsEmailValid(user.Email)
	if errEmailValid != nil {
		return entity.User{}, errEmailValid
	}

	errLength := validator.IsMinLengthValid(8, map[string]string{"password": user.Password})
	if errLength != nil {
		return entity.User{}, errLength
	}

	_, errGetEmail := ucs.userQueryRepository.GetUserByEmail(ctx, user.Email)
	if errGetEmail == nil {
		return entity.User{}, errors.New(constant.ERROR_EMAIL_EXIST)
	}

	if user.Password != user.ConfirmPassword {
		return entity.User{}, errors.New(constant.ERROR_PASSWORD_CONFIRM)
	}

	hashedPassword, errHash := crypto.HashPassword(user.Password)
	if errHash != nil {
		return entity.User{}, errors.New(constant.ERROR_PASSWORD_HASH)
	}

	user.Password = hashedPassword

	userEntity, errRegister := ucs.userCommandRepository.RegisterUser(ctx, user)
	if errRegister != nil {
		return entity.User{}, errRegister
	}

	return userEntity, nil
}

func (ucs *userCommandService) LoginUser(ctx context.Context, email, password string) (entity.User, string, error) {
	errEmpty := validator.IsDataEmpty([]string{"email", "password"}, email, password)
	if errEmpty != nil {
		return entity.User{}, "", errEmpty
	}

	errEmailValid := validator.IsEmailValid(email)
	if errEmailValid != nil {
		return entity.User{}, "", errEmailValid
	}

	userEntity, errGetEmail := ucs.userQueryRepository.GetUserByEmail(ctx, email)
	if errGetEmail != nil {
		return entity.User{}, "", errors.New(constant.ERROR_EMAIL_UNREGISTERED)
	}

	comparePassword := crypto.ComparePassword(userEntity.Password, password)
	if comparePassword != nil {
		return entity.User{}, "", errors.New(constant.ERROR_LOGIN)
	}

	token, errCreate := middleware.GenerateToken(userEntity.ID, userEntity.Role)
	if errCreate != nil {
		return entity.User{}, "", errors.New(constant.ERROR_TOKEN_GENERATE)
	}

	return userEntity, token, nil
}
