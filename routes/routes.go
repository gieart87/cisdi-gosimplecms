package routes

import (
	"github.com/gin-gonic/gin"
	adminCategoryCreate "gosimplecms/controllers/admin/categories/create"
	adminCategoryList "gosimplecms/controllers/admin/categories/list"
	adminPostCreate "gosimplecms/controllers/admin/posts/create"
	adminPostList "gosimplecms/controllers/admin/posts/list"
	"gosimplecms/controllers/posts/detail"
	"gosimplecms/controllers/posts/list"
	tagScore "gosimplecms/controllers/posts/tagscore"
	"gosimplecms/controllers/users/login"
	"gosimplecms/controllers/users/profile"
	"gosimplecms/controllers/users/register"
	"gosimplecms/middlewares"
	"gosimplecms/models"
)

func SetupRoutes(
	r *gin.Engine,

	userRegisterController *register.UserRegisterController,
	userLoginController *login.UserLoginController,
	userProfileController *profile.UserProfileController,
	listPostsController *list.PostListController,
	detailPostController *detail.PostDetailController,
	tagScorePostController *tagScore.PostTagScoreController,
	adminCategoryCreateController *adminCategoryCreate.CategoryCreateController,
	adminCategoryListController *adminCategoryList.CategoryListController,
	adminPostCreateController *adminPostCreate.PostCreateController,
	adminPostListController *adminPostList.PostListController,

) {

	r.POST("/register", userRegisterController.Register)
	r.POST("/login", userLoginController.Login)

	// Protected routes with JWT
	api := r.Group("/api")
	api.Use(middlewares.JWTAuthMiddleware())

	apiV1 := api.Group("/v1")

	apiV1Users := apiV1.Group("/users")
	{
		apiV1Users.GET("/profile", userProfileController.Profile)
	}

	apiV1Posts := apiV1.Group("/posts")
	{
		apiV1Posts.GET("", listPostsController.GetPosts)
		apiV1Posts.GET("/:id", detailPostController.Show)
		apiV1Posts.GET("/tag-scores", tagScorePostController.GetScores)
	}

	// Admin only routes
	adminV1 := apiV1.Group("/admin")

	adminV1Categories := adminV1.Group("/categories")
	{
		adminV1Categories.POST("",
			middlewares.AllowRoleMiddleware(models.RoleAdmin),
			adminCategoryCreateController.Create,
		)
		adminV1Categories.GET("",
			middlewares.AllowRoleMiddleware(models.RoleAdmin, models.RoleAuthor, models.RoleEditor),
			adminCategoryListController.GetCategories,
		)
	}

	adminV1Posts := adminV1.Group("/posts")
	{
		adminV1Posts.POST("",
			middlewares.AllowRoleMiddleware(models.RoleAdmin, models.RoleAuthor, models.RoleEditor),
			adminPostCreateController.Create,
		)
		adminV1Posts.GET("",
			middlewares.AllowRoleMiddleware(models.RoleAdmin, models.RoleAuthor, models.RoleEditor),
			adminPostListController.GetPosts,
		)
	}
}
