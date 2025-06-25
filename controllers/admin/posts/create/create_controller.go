package create

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
	"gorm.io/gorm"
	"gosimplecms/models"
	"gosimplecms/services"
	"gosimplecms/utils/response"
	"net/http"
)

type PostCreateController struct {
	PostService services.PostService
}

func NewPostCreateController(postService services.PostService) *PostCreateController {
	return &PostCreateController{postService}
}

// Create Post godoc
// @Summary Create a new post
// @Description Create a new post
// @Tags Manage Posts
// @Accept json
// @Produce json
// @Param users body models.CreatePostRequest true "Post data"
// @Success 201 {object} models.CreateUpdatePostResponse
// @Failure 400 {object} map[string]string
// @Router /api/v1/admin/posts [post]
func (ctl *PostCreateController) Create(c *gin.Context) {
	var createPostRequest models.CreatePostRequest

	if err := c.ShouldBindJSON(&createPostRequest); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	if err := createPostRequest.Validate(); err != nil {
		response.ErrorResponse(c, http.StatusUnprocessableEntity, err)
		return
	}

	userID, exist := c.Get("UserID")
	if !exist {
		response.ErrorResponse(c, http.StatusForbidden, errors.New("invalid user"))
		return
	}

	existPost, err := ctl.PostService.FindBySlug(slug.Make(createPostRequest.Title))
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		response.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	if existPost != nil {
		response.ErrorResponse(c, http.StatusConflict, errors.New("post title already exists"))
		return
	}

	createPostRequest.UserID = userID.(uint)

	// 1. Get Categories & Tags
	// Get Categories
	var categories []models.Category
	if len(createPostRequest.CategoryIDs) > 0 {
		categories, err = ctl.PostService.FindCategoriesByIDs(createPostRequest.CategoryIDs)
		if err != nil {
			response.ErrorResponse(c, http.StatusBadRequest, errors.New("error: Some categories invalid"))
			return
		}
		if len(categories) != len(createPostRequest.CategoryIDs) {
			response.ErrorResponse(c, http.StatusBadRequest, errors.New("error: Some categories invalid"))
			return
		}
	}

	// Get Tags
	var tags []models.Tag
	if len(createPostRequest.TagIDs) > 0 {
		tags, err = ctl.PostService.FindTagsByIDs(createPostRequest.TagIDs)
		if err != nil {
			response.ErrorResponse(c, http.StatusBadRequest, errors.New("error: Some tags invalid"))
			return
		}
		if len(tags) != len(createPostRequest.TagIDs) {
			response.ErrorResponse(c, http.StatusBadRequest, errors.New("error: Some tags invalid"))
			return
		}
	}

	post, err := ctl.PostService.Create(createPostRequest)
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	response.CreatedResponse(c, ctl.transformToResponse(post), "successfully create post")
}
