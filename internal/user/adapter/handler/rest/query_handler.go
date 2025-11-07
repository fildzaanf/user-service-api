package rest // inbound

import (
	"net/http"
	"user-service-api/internal/user/application/port"
	"user-service-api/pkg/constant"
	"user-service-api/pkg/middleware"
	"user-service-api/pkg/response"

	"github.com/labstack/echo/v4"
)

type userQueryHandler struct {
	userQueryService port.UserQueryServiceInterface
}

func NewUserQueryHandler(uqs port.UserQueryServiceInterface) *userQueryHandler {
	return &userQueryHandler{
		userQueryService: uqs,
	}
}

// query
func (uh *userQueryHandler) GetUserByID(c echo.Context) error {
	ctx := c.Request().Context()

	userIDParam := c.Param("id")
	if userIDParam == "" {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse("user id is required"))
	}

	userID, role, errExtract := middleware.ExtractToken(c)
	if errExtract != nil {
		return c.JSON(http.StatusUnauthorized, response.ErrorResponse("unauthorized access"))
	}

	if role != constant.USER && role != constant.SELLER && role != constant.BUYER {
		return c.JSON(http.StatusUnauthorized, response.ErrorResponse(constant.ERROR_ROLE_ACCESS))
	}

	if userIDParam != userID {
		return c.JSON(http.StatusForbidden, response.ErrorResponse("forbidden access"))
	}

	user, err := uh.userQueryService.GetUserByID(ctx, userIDParam)
	if err != nil {
		return c.JSON(http.StatusNotFound, response.ErrorResponse("user not found"))
	}

	userResponse := UserEntityToResponse(user)

	return c.JSON(http.StatusOK, response.SuccessResponse(constant.SUCCESS_RETRIEVED, userResponse))
}
