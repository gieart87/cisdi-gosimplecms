package profile

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gosimplecms/services"
	"gosimplecms/utils/response"
	"net/http"
)

type UserProfileController struct {
	UserService services.UserService
}

func NewUserProfileController(userService services.UserService) *UserProfileController {
	return &UserProfileController{userService}
}

// Profile godoc
// @Summary Profile user
// @Description Get profile user
// @Tags Auth
// @Accept json
// @Produce json
// @Success 200 {object} models.RegisterResponse
// @Failure 401 {object} map[string]string
// @Router /api/v1/users/profile [get]
func (uc *UserProfileController) Profile(c *gin.Context) {
	userID, exist := c.Get("UserID")
	if !exist {
		response.ErrorResponse(c, http.StatusForbidden, errors.New("invalid user"))
		return
	}

	user, err := uc.UserService.FindByID(userID.(uint))
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		response.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	response.SuccessResponse(c, uc.transformToResponse(user), "successfully get profile user")
}
