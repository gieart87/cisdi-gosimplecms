package register

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gosimplecms/models"
	"gosimplecms/services"
	"gosimplecms/utils/response"
	"net/http"
)

type UserRegisterController struct {
	UserService services.UserService
}

func NewUserRegisterController(userService services.UserService) *UserRegisterController {
	return &UserRegisterController{userService}
}

// Register godoc
// @Summary Register a new user
// @Description Create a new user account
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body register.Request true "User registration data"
// @Success 201 {object} register.Response
// @Failure 400 {object} map[string]string
// @Router /register [post]
func (uc *UserRegisterController) Register(c *gin.Context) {
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
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		response.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	if existUser != nil {
		response.ErrorResponse(c, http.StatusConflict, errors.New("email already registered"))
		return
	}

	user := models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	createdUser, err := uc.UserService.CreateUser(user)
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	response.CreatedResponse(c, uc.transformToResponse(createdUser), "successfully create user")
}
