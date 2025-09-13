package grpc

import (
	"user-service-api/internal/user/application/service"
	gormRepository "user-service-api/internal/user/adapter/repository/gorm"

	grpc "google.golang.org/grpc"
	"gorm.io/gorm"
)

func RegisterUserServices(server *grpc.Server, db *gorm.DB) {
	userQueryRepository := gormRepository.NewUserQueryRepository(db)
	userCommandRepository := gormRepository.NewUserCommandRepository(db)

	userQueryService := service.NewUserQueryService(userCommandRepository, userQueryRepository)
	userCommandService := service.NewUserCommandService(userCommandRepository, userQueryRepository)

	userQueryHandler := NewUserQueryHandler(userQueryService)
	userCommandHandler := NewUserCommandHandler(userCommandService)

	RegisterUserQueryServiceServer(server, userQueryHandler)
	RegisterUserCommandServiceServer(server, userCommandHandler)
}
