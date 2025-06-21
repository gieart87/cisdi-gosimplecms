package routes

import (
	"github.com/gin-gonic/gin"
	"gosimplecms/controllers/posts/list"
	"gosimplecms/controllers/user/login"
	"gosimplecms/controllers/user/register"
)

func SetupRoutes(
	r *gin.Engine,

	userRegisterController *register.UserRegisterController,
	userLoginController *login.UserLoginController,
	listPostsController *list.PostListController,
) {

	r.POST("/register", userRegisterController.Register)
	r.POST("/login", userLoginController.Login)

	postRoutes := r.Group("/posts")
	{
		postRoutes.GET("", listPostsController.GetPosts)
	}
}
