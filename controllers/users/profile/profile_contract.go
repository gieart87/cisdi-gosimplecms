package profile

import (
	"gosimplecms/models"
	"time"
)

func (uc *UserProfileController) transformToResponse(user *models.User) *models.RegisterResponse {
	return &models.RegisterResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
	}
}
