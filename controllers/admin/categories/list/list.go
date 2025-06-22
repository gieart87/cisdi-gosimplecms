package list

import (
	"github.com/gin-gonic/gin"
	"gosimplecms/models"
	"gosimplecms/services"
	"gosimplecms/utils/response"
	"net/http"
)

type CategoryListController struct {
	CategoryService services.CategoryService
}

func NewCategoryListController(categoryService services.CategoryService) *CategoryListController {
	return &CategoryListController{
		CategoryService: categoryService,
	}
}

// GetCategories godoc
// @Summary Get all categories
// @Tags Manage Categories
// @Security BearerAuth
// @Produce json
// @Success 200 {array} models.CreateUpdateCategoryResponse
// @Failure 401 {object} map[string]string
// @Router /api/v1/admin/categories [get]
func (ctl *CategoryListController) GetCategories(c *gin.Context) {
	categories, err := ctl.CategoryService.GetCategories()
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err)
	}

	if len(categories) == 0 {
		categories = make([]models.Category, 0)
	}

	response.SuccessResponse(c, ctl.transformToResponse(categories), "successfully retrieve categories")
}
