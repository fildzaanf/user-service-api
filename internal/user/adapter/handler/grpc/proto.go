package grpc

import entity "user-service-api/internal/user/domain"

// request
func UserRegisterRequestToEntity(request *UserRegisterRequest) entity.User {
	return entity.User{
		Name:            request.GetName(),
		Email:           request.GetEmail(),
		Role:            request.GetRole(),
		Password:        request.GetPassword(),
		ConfirmPassword: request.GetConfirmPassword(),
	}
}

func UserLoginRequestToEntity(request *UserLoginRequest) entity.User {
	return entity.User{
		Email:    request.GetEmail(),
		Password: request.GetPassword(),
	}
}

// response
func UserRegisterEntityToResponse(entity entity.User) *UserRegisterResponse {
	return &UserRegisterResponse{
		Id:    entity.ID,
		Email: entity.Email,
		Role:  entity.Role,
	}
}		

func UserEntityToLoginResponse(entity entity.User, token string) *UserLoginResponse {
	return &UserLoginResponse{
		Id:    entity.ID,
		Role:  entity.Role,
		Token: token,
	}
}	

func UserEntityToResponse(entity entity.User) *UserResponse {
	return &UserResponse{
		Id:    entity.ID,	
		Name:  entity.Name,
		Email: entity.Email,
		Role:  entity.Role,
	}
}
