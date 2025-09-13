package port // inbound

import entity "user-service-api/internal/user/domain"

type UserCommandServiceInterface interface {
	RegisterUser(user entity.User) (entity.User, error)
	LoginUser(email, password string) (entity.User, string, error)
}

type UserQueryServiceInterface interface {
	GetUserByID(id string) (entity.User, error)
}