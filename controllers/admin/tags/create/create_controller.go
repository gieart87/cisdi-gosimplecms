package create

import (
	"github.com/gin-gonic/gin"
	"gosimplecms/models"
	"gosimplecms/services"
	"gosimplecms/utils/response"
	"net/http"
)

type TagCreateController struct {
	TagService services.TagService
}

func NewTagCreateController(tagService services.TagService) *TagCreateController {
	return &TagCreateController{tagService}
}

// Create Tag godoc
// @Summary Create a new tag
// @Description Create a new tag
// @Tags Manage Tags
// @Accept json
// @Produce json
// @Param users body models.CreateTagRequest true "Tag data"
// @Success 201 {object} models.CreateUpdateTagResponse
// @Failure 400 {object} map[string]string
// @Router /api/v1/admin/tags [post]
func (ctl *TagCreateController) Create(c *gin.Context) {
	var req models.CreateTagRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	if err := req.Validate(); err != nil {
		response.ErrorResponse(c, http.StatusUnprocessableEntity, err)
		return
	}

	tag, err := ctl.TagService.Create(req)
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	response.CreatedResponse(c, ctl.transformToResponse(tag), "successfully create tag")
}
