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

func (ctl *PostListController) GetPosts(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit

	posts, total, err := ctl.PostService.GetPosts(limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	if len(posts) == 0 {
		posts = make([]models.Post, 0)
	}

	_ = models.PaginationMeta{
		Total: total,
		Page:  0,
		Limit: 0,
	}

	response.SuccessResponse(c, posts, "successfully retrieve posts")
}
