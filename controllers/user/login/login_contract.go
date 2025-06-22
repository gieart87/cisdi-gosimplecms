package login

import (
	"gosimplecms/models"
)

func (uc *UserLoginController) transformToResponse(token string) *models.LoginResponse {
	return &models.LoginResponse{Token: token}
}
