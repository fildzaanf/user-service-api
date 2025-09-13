package rest

import (
	gormRepository "user-service-api/internal/user/adapter/repository/gorm"
	"user-service-api/internal/user/application/service"
	"user-service-api/pkg/middleware"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RegisterUserRoutes(user *echo.Group, db *gorm.DB) {
	userQueryRepository := gormRepository.NewUserQueryRepository(db)
	userCommandRepository := gormRepository.NewUserCommandRepository(db)

	userQueryService := service.NewUserQueryService(userCommandRepository, userQueryRepository)
	userCommandService := service.NewUserCommandService(userCommandRepository, userQueryRepository)

	userQueryHandler := NewUserQueryHandler(userQueryService)
	userCommandHandler := NewUserCommandHandler(userCommandService)

	user.POST("/register", userCommandHandler.RegisterUser)
	user.POST("/login", userCommandHandler.LoginUser)
	user.GET("/:id", userQueryHandler.GetUserByID, middleware.JWTMiddleware())
}