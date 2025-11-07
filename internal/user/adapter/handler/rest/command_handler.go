package rest // inbound

import (
	"net/http"
	"user-service-api/internal/user/application/port"
	"user-service-api/pkg/constant"
	"user-service-api/pkg/response"

	"github.com/labstack/echo/v4"
)

type userCommandHandler struct {
	userCommandService port.UserCommandServiceInterface
}

func NewUserCommandHandler(ucs port.UserCommandServiceInterface) *userCommandHandler {
	return &userCommandHandler{
		userCommandService: ucs,
	}
}

// command
func (uh *userCommandHandler) RegisterUser(c echo.Context) error {
	ctx := c.Request().Context()

	var userRequest UserRegisterRequest

	if err := c.Bind(&userRequest); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(err.Error()))
	}

	userEntity := UserRegisterRequestToEntity(userRequest)

	registeredUser, err := uh.userCommandService.RegisterUser(ctx, userEntity)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(err.Error()))
	}

	userResponse := UserRegisterEntityToResponse(registeredUser)

	return c.JSON(http.StatusCreated, response.SuccessResponse(constant.SUCCESS_REGISTER, userResponse))
}

func (uh *userCommandHandler) LoginUser(c echo.Context) error {
	ctx := c.Request().Context()

	var userRequest UserLoginRequest

	if err := c.Bind(&userRequest); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(err.Error()))
	}

	loginUser, token, err := uh.userCommandService.LoginUser(ctx, userRequest.Email, userRequest.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, response.ErrorResponse(err.Error()))
	}

	userResponse := UserEntityToLoginResponse(loginUser, token)

	return c.JSON(http.StatusOK, response.SuccessResponse(constant.SUCCESS_LOGIN, userResponse))
}
