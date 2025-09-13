package service


import (
	"errors"
	entity "user-service-api/internal/user/domain"
	"user-service-api/internal/user/application/port"
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

func (uqs *userQueryService) GetUserByID(id string) (entity.User, error) {
	if id == "" {
		return entity.User{}, errors.New(constant.ERROR_ID_INVALID)
	}

	userEntity, errGetID := uqs.userQueryRepository.GetUserByID(id)
	if errGetID != nil {
		return entity.User{}, errors.New(constant.ERROR_DATA_EMPTY)
	}

	return userEntity, nil
}
