package port // outbond

import (
	"context"
	entity "user-service-api/internal/user/domain"
)

type UserCommandRepositoryInterface interface {
	RegisterUser(ctx context.Context, user entity.User) (entity.User, error)
	LoginUser(ctx context.Context, email, password string) (entity.User, error)
}

type UserQueryRepositoryInterface interface {
	GetUserByID(ctx context.Context, id string) (entity.User, error)
	GetUserByEmail(ctx context.Context, email string) (entity.User, error)
}
