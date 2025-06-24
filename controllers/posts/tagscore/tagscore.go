package tagscore

import (
	"github.com/gin-gonic/gin"
	"gosimplecms/services"
	"gosimplecms/utils/response"
	"net/http"
)

type PostTagScoreController struct {
	PostService services.PostService
}

func NewPostTagScoreController(postService services.PostService) *PostTagScoreController {
	return &PostTagScoreController{
		PostService: postService,
	}
}

// GetScores godoc
// @Summary Get tag scores
// @Tags Posts
// @Security BearerAuth
// @Produce json
// @Success 200 {array} models.TagRelationship
// @Failure 401 {object} map[string]string
// @Router /api/v1/posts/tag-scores [get]
func (ctl *PostTagScoreController) GetScores(c *gin.Context) {
	scores, err := ctl.PostService.GetTagRelationshipScores()
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	response.SuccessResponse(c, scores, "Retrieve scores successfully")

}
