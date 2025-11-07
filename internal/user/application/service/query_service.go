package service

import (
	"context"
	"errors"
	"user-service-api/internal/user/application/port"
	entity "user-service-api/internal/user/domain"
	"user-service-api/pkg/constant"
)

type userQueryService struct {
	userCommandRepository port.UserCommandRepositoryInterface
	userQueryRepository   port.UserQueryRepositoryInterface
}

func NewUserQueryService(ucr port.UserCommandRepositoryInterface, uqr port.UserQueryRepositoryInterface) port.UserQueryServiceInterface {
	return &userQueryService{
		userCommandRepository: ucr,
		userQueryRepository:   uqr,
	}
}

func (uqs *userQueryService) GetUserByID(ctx context.Context, id string) (entity.User, error) {
	if id == "" {
		return entity.User{}, errors.New(constant.ERROR_ID_INVALID)
	}

	userEntity, errGetID := uqs.userQueryRepository.GetUserByID(ctx, id)
	if errGetID != nil {
		return entity.User{}, errors.New(constant.ERROR_DATA_EMPTY)
	}

	return userEntity, nil
}
