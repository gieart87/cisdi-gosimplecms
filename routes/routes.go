package routes

import (
	"github.com/gin-gonic/gin"
	adminCategoryCreate "gosimplecms/controllers/admin/categories/create"
	adminCategoryList "gosimplecms/controllers/admin/categories/list"
	"gosimplecms/controllers/posts/list"
	"gosimplecms/controllers/users/login"
	"gosimplecms/controllers/users/register"
)

func SetupRoutes(
	r *gin.Engine,

	userRegisterController *register.UserRegisterController,
	userLoginController *login.UserLoginController,
	listPostsController *list.PostListController,
	adminCategoryCreateController *adminCategoryCreate.CategoryCreateController,
	adminCategoryListController *adminCategoryList.CategoryListController,
) {

	r.POST("/register", userRegisterController.Register)
	r.POST("/login", userLoginController.Login)

	// Protected routes with JWT
	apiV1 := r.Group("/api/v1")

	// Admin only routes
	adminV1 := apiV1.Group("/admin")

	adminV1Categories := adminV1.Group("/categories")
	{
		adminV1Categories.POST("", adminCategoryCreateController.Create)
		adminV1Categories.GET("", adminCategoryListController.GetCategories)
	}
	//postRoutes := r.Group("/posts")
	//{
	//	postRoutes.GET("", listPostsController.GetPosts)
	//}
}
