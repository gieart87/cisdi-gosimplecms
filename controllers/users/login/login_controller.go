package login

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gosimplecms/models"
	"gosimplecms/services"
	"gosimplecms/utils/errs"
	"gosimplecms/utils/response"
	"net/http"
)

type UserLoginController struct {
	UserService services.UserService
}

func NewUserLoginController(userService services.UserService) *UserLoginController {
	return &UserLoginController{userService}
}

// Login godoc
// @Summary Login users
// @Description Create a new users account
// @Tags Auth
// @Accept json
// @Produce json
// @Param users body models.LoginRequest true "User login"
// @Success 200 {object} models.LoginResponse
// @Failure 400 {object} map[string]string
// @Router /login [post]
func (uc *UserLoginController) Login(c *gin.Context) {
	var req models.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	if err := req.Validate(); err != nil {
		response.ErrorResponse(c, http.StatusUnprocessableEntity, err)
		return
	}

	var appErr *errs.AppError
	token, err := uc.UserService.Login(req)
	if err != nil {
		if errors.As(err, &appErr) && appErr.Code == errs.ErrCodeLoginFailed {
			response.ErrorResponse(c, http.StatusForbidden, errors.New(appErr.Message))
			return
		}
		response.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	response.SuccessResponse(c, uc.transformToResponse(token), "successfully generate token")
}
