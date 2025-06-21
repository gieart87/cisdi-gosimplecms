package list

import (
	"github.com/gin-gonic/gin"
	"gosimplecms/models"
	"gosimplecms/services"
	"gosimplecms/utils/response"
	"net/http"
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
	posts, err := ctl.PostService.GetPosts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	if len(posts) == 0 {
		posts = make([]models.Post, 0)
	}

	response.SuccessResponse(c, posts, "successfully retrieve posts")
}
