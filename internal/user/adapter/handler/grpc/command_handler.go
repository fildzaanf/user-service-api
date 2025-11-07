package grpc // inbound

import (
	"context"
	"user-service-api/internal/user/application/port"
)

type userCommandHandler struct {
	UnimplementedUserCommandServiceServer
	userCommandService port.UserCommandServiceInterface
}

func NewUserCommandHandler(ucs port.UserCommandServiceInterface) UserCommandServiceServer {
	return &userCommandHandler{
		userCommandService: ucs,
	}
}

func (h *userCommandHandler) RegisterUser(ctx context.Context, userRequest *UserRegisterRequest) (*UserRegisterResponse, error) {
	userEntity := UserRegisterRequestToEntity(userRequest)

	registeredUser, err := h.userCommandService.RegisterUser(ctx, userEntity)
	if err != nil {
		return nil, err
	}

	userResponse := UserRegisterEntityToResponse(registeredUser)

	return userResponse, nil
}

func (h *userCommandHandler) LoginUser(ctx context.Context, userRequest *UserLoginRequest) (*UserLoginResponse, error) {
	userEntity := UserLoginRequestToEntity(userRequest)

	loginUser, token, err := h.userCommandService.LoginUser(ctx, userEntity.Email, userEntity.Password)
	if err != nil {
		return nil, err
	}

	userResponse := UserEntityToLoginResponse(loginUser, token)

	return userResponse, nil
}
