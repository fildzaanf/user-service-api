package grpc // inbound

import (
	"context"
	"user-service-api/internal/user/application/port"
	"user-service-api/pkg/constant"
	"user-service-api/pkg/middleware"

	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

type userQueryHandler struct {
	UnimplementedUserQueryServiceServer
	userQueryService port.UserQueryServiceInterface
}

func NewUserQueryHandler(uqs port.UserQueryServiceInterface) *userQueryHandler {
	return &userQueryHandler{
		userQueryService: uqs,
	}
}

func (uh *userQueryHandler) GetUserByID(ctx context.Context, userRequest *GetUserByIDRequest) (*UserResponse, error) {
	if userRequest.GetId() == "" {
		return nil, status.Error(codes.InvalidArgument, "user id is required")
	}

	userID, role, errExtract := middleware.ExtractTokenFromContext(ctx)
	if errExtract != nil {
		return nil, status.Error(codes.Unauthenticated, "unauthorized access")
	}

	if role != constant.USER && role != constant.SELLER && role != constant.BUYER {
		return nil, status.Error(codes.PermissionDenied, constant.ERROR_ROLE_ACCESS)
	}

	if userRequest.GetId() != userID {
		return nil, status.Error(codes.PermissionDenied, "forbidden access")
	}

	user, err := uh.userQueryService.GetUserByID(userRequest.GetId())
	if err != nil {
		return nil, status.Error(codes.NotFound, "user not found")
	}

	UserResponse := UserEntityToResponse(user)

	return UserResponse, nil
}
