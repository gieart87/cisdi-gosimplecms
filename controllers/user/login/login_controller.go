package login

import (
	"errors"
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"gosimplecms/services"
	"gosimplecms/utils/jwt"
	"gosimplecms/utils/password"
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
// @Summary Login user
// @Description Create a new user account
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body login.Request true "User login"
// @Success 200 {object} login.Response
// @Failure 400 {object} map[string]string
// @Router /login [post]
func (uc *UserLoginController) Login(c *gin.Context) {
	var req Request

	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	if err := req.Validate(); err != nil {
		response.ErrorResponse(c, http.StatusUnprocessableEntity, err)
		return
	}

	existUser, err := uc.UserService.FindByEmail(req.Email)
	if err != nil || !password.CheckPassword(existUser.Password, req.Password) {
		response.ErrorResponse(c, http.StatusForbidden, errors.New("invalid email or password"))
		return
	}

	token, err := jwt.GenerateToken(existUser.ID, existUser.Role)
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, errors.New("failed to generate token"))
		return
	}

	response.SuccessResponse(c, token, "successfully generate token")
}

func (r Request) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Email, validation.Required, validation.Length(5, 50), is.Email),
		validation.Field(&r.Password, validation.Required, validation.Length(6, 20)),
	)
}
