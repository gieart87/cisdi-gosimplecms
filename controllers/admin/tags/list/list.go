package list

import (
	"github.com/gin-gonic/gin"
	"gosimplecms/models"
	"gosimplecms/services"
	"gosimplecms/utils/response"
	"net/http"
)

type TagListController struct {
	TagService services.TagService
}

func NewTagListController(tagService services.TagService) *TagListController {
	return &TagListController{
		TagService: tagService,
	}
}

// GetTags godoc
// @Summary Get all tags
// @Tags Manage Tags
// @Security BearerAuth
// @Produce json
// @Success 200 {array} models.CreateUpdateCategoryResponse
// @Failure 401 {object} map[string]string
// @Router /api/v1/admin/tags [get]
func (ctl *TagListController) GetTags(c *gin.Context) {
	tags, err := ctl.TagService.GetTags()
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err)
	}

	if len(tags) == 0 {
		tags = make([]models.Tag, 0)
	}

	response.SuccessResponse(c, ctl.transformToResponse(tags), "successfully retrieve tags")
}
