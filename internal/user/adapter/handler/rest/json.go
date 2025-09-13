package rest // inbound

import entity "user-service-api/internal/user/domain"


// request
type UserRegisterRequest struct {
	Name            string `json:"name" form:"name"`
	Email           string `json:"email" form:"email"`
	Role            string `json:"role" form:"role"`
	Password        string `json:"password" form:"password"`
	ConfirmPassword string `json:"confirm_password" form:"confirm_password"`
}

type UserLoginRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func UserRegisterRequestToEntity(request UserRegisterRequest) entity.User {
	return entity.User{
		Name:            request.Name,
		Email:           request.Email,
		Role:            request.Role,
		Password:        request.Password,
		ConfirmPassword: request.ConfirmPassword,
	}
}

func UserLoginRequestToEntity(request UserLoginRequest) entity.User {
	return entity.User{
		Email:    request.Email,
		Password: request.Password,
	}
}

// response
type UserRegisterResponse struct {
	ID    string `json:"id"`
	Role  string `json:"role"`
	Email string `json:"email"`
}

type UserLoginResponse struct {
	ID    string `json:"id"`
	Role  string `json:"role"`
	Token string `json:"token"`
}

type UserResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

func UserRegisterEntityToResponse(user entity.User) UserRegisterResponse {
	return UserRegisterResponse{
		ID:    user.ID,
		Email: user.Email,
		Role:  user.Role,
	}
}

func UserEntityToLoginResponse(user entity.User, token string) UserLoginResponse {
	return UserLoginResponse{
		ID:    user.ID,
		Role:  user.Role,
		Token: token,
	}
}

func UserEntityToResponse(user entity.User) UserResponse {
	return UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
	}
}
