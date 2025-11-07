package port // inbound

import (
	"context"
	entity "user-service-api/internal/user/domain"
)

type UserCommandServiceInterface interface {
	RegisterUser(ctx context.Context, user entity.User) (entity.User, error)
	LoginUser(ctx context.Context, email, password string) (entity.User, string, error)
}

type UserQueryServiceInterface interface {
	GetUserByID(ctx context.Context, id string) (entity.User, error)
}
