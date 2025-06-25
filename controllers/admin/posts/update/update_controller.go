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

type PostUpdateController struct {
	PostService services.PostService
}

func NewPostUpdateController(postService services.PostService) *PostUpdateController {
	return &PostUpdateController{postService}
}

// Update Post godoc
// @Summary Update existing post
// @Description Update existing post
// @Tags Manage Posts
// @Accept json
// @Produce json
// @Param id path string true "Post ID"
// @Param users body models.CreatePostRequest true "Post data"
// @Success 201 {object} models.CreateUpdatePostResponse
// @Failure 400 {object} map[string]string
// @Router /api/v1/admin/posts/{id} [put]
func (ctl *PostUpdateController) Update(c *gin.Context) {
	var updatePostRequest models.UpdatePostRequest

	if err := c.ShouldBindJSON(&updatePostRequest); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	if err := updatePostRequest.Validate(); err != nil {
		response.ErrorResponse(c, http.StatusUnprocessableEntity, err)
		return
	}

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

	// 1. Get Categories & Tags
	// Get Categories
	var categories []models.Category
	if len(updatePostRequest.CategoryIDs) > 0 {
		categories, err = ctl.PostService.FindCategoriesByIDs(updatePostRequest.CategoryIDs)
		if err != nil {
			response.ErrorResponse(c, http.StatusBadRequest, errors.New("error: Some categories invalid"))
			return
		}
		if len(categories) != len(updatePostRequest.CategoryIDs) {
			response.ErrorResponse(c, http.StatusBadRequest, errors.New("error: Some categories invalid"))
			return
		}
	}

	// Get Tags
	var tags []models.Tag
	if len(updatePostRequest.TagIDs) > 0 {
		tags, err = ctl.PostService.FindTagsByIDs(updatePostRequest.TagIDs)
		if err != nil {
			response.ErrorResponse(c, http.StatusBadRequest, errors.New("error: Some tags invalid"))
			return
		}
		if len(tags) != len(updatePostRequest.TagIDs) {
			response.ErrorResponse(c, http.StatusBadRequest, errors.New("error: Some tags invalid"))
			return
		}
	}

	post, err := ctl.PostService.Update(uint(id), updatePostRequest)
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	response.CreatedResponse(c, ctl.transformToResponse(post), "successfully create post")
}
