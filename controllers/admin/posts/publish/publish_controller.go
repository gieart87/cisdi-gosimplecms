package create

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

type PostPublishController struct {
	PostService services.PostService
}

func NewPostPublishController(postService services.PostService) *PostPublishController {
	return &PostPublishController{postService}
}

// Publish Post godoc
// @Summary Publish existing post
// @Description Publish existing post
// @Tags Manage Posts
// @Accept json
// @Produce json
// @Param id path string true "Post ID"
// @Success 200 {object} models.CreateUpdatePostResponse
// @Failure 400 {object} map[string]string
// @Router /api/v1/admin/posts/{id}/publish [put]
func (ctl *PostPublishController) Publish(c *gin.Context) {
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

	_, err = ctl.PostService.FindByID(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.ErrorResponse(c, http.StatusNotFound, err)
			return
		}
		response.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	var updatePostRequest models.UpdatePostRequest
	updatePostRequest.Status = models.PostStatusPublished
	post, err := ctl.PostService.Update(uint(id), updatePostRequest)
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	response.SuccessResponse(c, ctl.transformToResponse(post), "successfully create post")
}
