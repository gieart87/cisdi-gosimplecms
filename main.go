package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"gosimplecms/configs"
	adminCategoryCreate "gosimplecms/controllers/admin/categories/create"
	adminCategoryList "gosimplecms/controllers/admin/categories/list"
	adminPostCreate "gosimplecms/controllers/admin/posts/create"
	adminPostList "gosimplecms/controllers/admin/posts/list"
	adminPostUpdate "gosimplecms/controllers/admin/posts/update"
	"gosimplecms/controllers/posts/detail"
	"gosimplecms/controllers/posts/list"
	tagScore "gosimplecms/controllers/posts/tagscore"
	"gosimplecms/controllers/users/login"
	"gosimplecms/controllers/users/profile"
	"gosimplecms/controllers/users/register"
	"gosimplecms/db/seeds"
	_ "gosimplecms/docs"
	"gosimplecms/models"
	"gosimplecms/repositories"
	"gosimplecms/routes"
	"gosimplecms/services"
	"log"
)

// @title Go Simple CMS API
// @version 1.0
// @description This is a simple API with JWT Auth, RBAC, and Clean Architecture.
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}

	mode := flag.String("mode", "serve", "Mode to run: serve | migrate | seed | score")
	port := flag.String("port", "8080", "Port to run server on")
	flag.Parse()

	switch *mode {
	case "serve":
		startServer(*port)
	case "migrate":
		runMigration()
	case "seed":
		runSeeder()
	case "score":
		runScoreSeeder()
	default:
		fmt.Println("Unknown mode:", *mode)
	}
}

func startServer(port string) {

	//gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	db := configs.ConnectDatabase()

	userRepo := repositories.NewUserRepository()
	postRepo := repositories.NewPostRepository(db)
	tagRepo := repositories.NewTagRepository()
	categoryRepo := repositories.NewCategoryRepository()

	userService := services.NewUserService(userRepo)
	postService := services.NewPostService(postRepo, tagRepo, categoryRepo)
	categoryService := services.NewCategoryService(categoryRepo)

	userRegisterController := register.NewUserRegisterController(userService)
	userLoginController := login.NewUserLoginController(userService)
	userProfileController := profile.NewUserProfileController(userService)
	listPostController := list.NewListPostController(postService)
	detailPostController := detail.NewDetailPostController(postService)
	tagScoreController := tagScore.NewPostTagScoreController(postService)

	adminCategoryCreateController := adminCategoryCreate.NewCategoryCreateController(categoryService)
	adminCategoryListController := adminCategoryList.NewCategoryListController(categoryService)
	adminPostCreateController := adminPostCreate.NewPostCreateController(postService)
	adminPostUpdateController := adminPostUpdate.NewPostUpdateController(postService)
	adminPostListController := adminPostList.NewPostListController(postService)

	routes.SetupRoutes(r,
		userRegisterController,
		userLoginController,
		userProfileController,
		listPostController,
		detailPostController,
		tagScoreController,
		adminCategoryCreateController,
		adminCategoryListController,
		adminPostCreateController,
		adminPostUpdateController,
		adminPostListController,
	)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":" + port)
}

func runMigration() {
	configs.ConnectDatabase()

	err := configs.DB.AutoMigrate(&models.User{}, &models.Post{}, &models.PostVersion{}, &models.Category{}, &models.Tag{}, &models.TagRelationship{})
	if err != nil {
		log.Fatal("❌ Migration failed:", err)
	}
	fmt.Println("✅ Migration completed")
}

func runSeeder() {
	configs.ConnectDatabase()
	err := seeds.SeedUsers(configs.DB)
	if err != nil {
		log.Fatal("Run SeedUsers failed:", err)
	}
	fmt.Println("✅ Seeder executed")
}

func runScoreSeeder() {
	configs.ConnectDatabase()
	err := seeds.SeedTagRelationships(configs.DB)
	if err != nil {
		log.Fatal("Run Score Seeder failed:", err)
	}
	fmt.Println("✅ Seeder executed")
}
