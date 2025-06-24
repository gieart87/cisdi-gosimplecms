package detail

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gosimplecms/models"
	"gosimplecms/services"
	"gosimplecms/utils/response"
	"net/http"
	"strconv"
)

type PostDetailController struct {
	PostService services.PostService
}

func NewDetailPostController(postService services.PostService) *PostDetailController {
	return &PostDetailController{
		PostService: postService,
	}
}

// Show post godoc
// @Summary Get detail post
// @Tags Posts
// @Security BearerAuth
// @Produce json
// @Success 200 {object} models.CreateUpdatePostResponse
// @Failure 401 {object} map[string]string
// @Router /api/v1/posts/{id} [get]
func (ctl *PostDetailController) Show(c *gin.Context) {
	_, exist := c.Get("UserID")
	if !exist {
		response.ErrorResponse(c, http.StatusForbidden, errors.New("invalid user"))
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.ErrorResponse(c, http.StatusNotFound, errors.New("invalid id"))
		return
	}

	post, err := ctl.PostService.FindByID(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.ErrorResponse(c, http.StatusNotFound, errors.New("post not found"))
			return
		}

		response.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	if post.Status != models.PostStatusPublished {
		response.ErrorResponse(c, http.StatusNotFound, errors.New("post not found"))
		return
	}

	response.SuccessResponse(c, ctl.transformToResponse(post), "successfully retrieve posts")
}
