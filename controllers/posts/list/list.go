package list

import (
	"github.com/gin-gonic/gin"
	"gosimplecms/models"
	"gosimplecms/services"
	"gosimplecms/utils/response"
	"net/http"
	"strconv"
)

type PostListController struct {
	PostService services.PostService
}

func NewListPostController(postService services.PostService) *PostListController {
	return &PostListController{
		PostService: postService,
	}
}

// GetPosts godoc
// @Summary Get all posts
// @Tags Posts
// @Security BearerAuth
// @Produce json
// @Success 200 {array} Post
// @Failure 401 {object} map[string]string
// @Router /api/v1/posts [get]
func (ctl *PostListController) GetPosts(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("per_page", "10"))
	offset := (page - 1) * limit

	posts, total, err := ctl.PostService.GetActivePosts(limit, offset)
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err)
	}

	if len(posts) == 0 {
		posts = make([]models.Post, 0)
	}

	paginationMeta := models.PaginationMeta{
		Total: total,
		Page:  page,
		Limit: limit,
	}

	postResponse := struct {
		Posts      []*Post               `json:"posts"`
		Pagination models.PaginationMeta `json:"pagination"`
	}{
		Posts:      ctl.transformToResponse(posts),
		Pagination: paginationMeta,
	}

	response.SuccessResponse(c, postResponse, "successfully retrieve posts")
}
