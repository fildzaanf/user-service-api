package port // outbond

import (
	entity "user-service-api/internal/user/domain"
)

type UserCommandRepositoryInterface interface {
	RegisterUser(user entity.User) (entity.User, error)
	LoginUser(email, password string) (entity.User, error)
}

type UserQueryRepositoryInterface interface {
	GetUserByID(id string) (entity.User, error)
	GetUserByEmail(email string) (entity.User, error)
}