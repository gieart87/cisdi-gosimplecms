package create

import (
	"github.com/gin-gonic/gin"
	"gosimplecms/models"
	"gosimplecms/services"
	"gosimplecms/utils/response"
	"net/http"
)

type CategoryCreateController struct {
	CategoryService services.CategoryService
}

func NewCategoryCreateController(categoryService services.CategoryService) *CategoryCreateController {
	return &CategoryCreateController{categoryService}
}

// Create Category godoc
// @Summary Create a new category
// @Description Create a new category
// @Tags Manage Categories
// @Accept json
// @Produce json
// @Param users body models.CreateCategoryRequest true "Category data"
// @Success 201 {object} models.CreateUpdateCategoryResponse
// @Failure 400 {object} map[string]string
// @Router /api/v1/admin/categories [post]
func (ctl *CategoryCreateController) Create(c *gin.Context) {
	var req models.CreateCategoryRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	if err := req.Validate(); err != nil {
		response.ErrorResponse(c, http.StatusUnprocessableEntity, err)
		return
	}

	createdCategory, err := ctl.CategoryService.Create(req)
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	response.CreatedResponse(c, ctl.transformToResponse(createdCategory), "successfully create category")
}
