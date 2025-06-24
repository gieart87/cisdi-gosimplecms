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

//func (ctl *PostCreateController) Create(c *gin.Context) {
//	var input struct {
//		Slug        string      `json:"slug" binding:"required"`
//		Title       string      `json:"title" binding:"required"`
//		Body        string      `json:"body"`
//		TagIDs      []uuid.UUID `json:"tag_ids"`
//		CategoryIDs []uuid.UUID `json:"category_ids"`
//		UserID      string      `json:"user_id"`
//	}
//	userID, _ := c.Get("UserID")
//	input.UserID = userID.(string)
//
//	if err := c.ShouldBindJSON(&input); err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	// 1. Create Post if not exists
//	var post models.Post
//	if err := configs.DB.Where("slug = ?", input.Slug).First(&post).Error; err != nil {
//		post = models.Post{Slug: input.Slug}
//		if err := configs.DB.Create(&post).Error; err != nil {
//			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create post"})
//			return
//		}
//	}
//
//	// 2. Get Categories & Tags
//	// Get Categories
//	var categories []models.Category
//	if len(input.CategoryIDs) > 0 {
//		if err := configs.DB.Where("id IN ?", input.CategoryIDs).Find(&categories).Error; err != nil {
//			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch categories"})
//			return
//		}
//		if len(categories) != len(input.CategoryIDs) {
//			c.JSON(http.StatusBadRequest, gin.H{"error": "Some categories not found"})
//			return
//		}
//	}
//
//	// Get Tags
//	var tags []models.Tag
//	if len(input.TagIDs) > 0 {
//		if err := configs.DB.Where("id IN ?", input.TagIDs).Find(&tags).Error; err != nil {
//			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tags"})
//			return
//		}
//		if len(tags) != len(input.TagIDs) {
//			c.JSON(http.StatusBadRequest, gin.H{"error": "Some tags not found"})
//			return
//		}
//	}
//	// 3. Create PostVersion
//	postVersion := models.PostVersion{
//		Title:  input.Title,
//		Body:   input.Body,
//		PostID: post.ID,
//	}
//
//	if err := configs.DB.Create(&postVersion).Error; err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create post version"})
//		return
//	}
//
//	if err := configs.DB.Model(&postVersion).Association("Categories").Append(categories); err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to associate categories"})
//		return
//	}
//
//	if err := configs.DB.Model(&postVersion).Association("Tags").Append(tags); err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to associate tags"})
//		return
//	}
//
//	c.JSON(http.StatusCreated, postVersion)
//
//	//response.CreatedResponse(c, ctl.transformToResponse(createdPost), "successfully create post")
//}
