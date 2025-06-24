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
	"gosimplecms/controllers/posts/list"
	"gosimplecms/controllers/users/login"
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

	mode := flag.String("mode", "serve", "Mode to run: serve | migrate | seed")
	port := flag.String("port", "8080", "Port to run server on")
	flag.Parse()

	switch *mode {
	case "serve":
		startServer(*port)
	case "migrate":
		runMigration()
	case "seed":
		runSeeder()
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
	listPostController := list.NewListPostController(postService)

	adminCategoryCreateController := adminCategoryCreate.NewCategoryCreateController(categoryService)
	adminCategoryListController := adminCategoryList.NewCategoryListController(categoryService)
	adminPostCreateController := adminPostCreate.NewPostCreateController(postService)
	adminPostListController := adminPostList.NewPostListController(postService)

	routes.SetupRoutes(r,
		userRegisterController,
		userLoginController,
		listPostController,
		adminCategoryCreateController,
		adminCategoryListController,
		adminPostCreateController,
		adminPostListController,
	)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":" + port)
}

func runMigration() {
	configs.ConnectDatabase()

	err := configs.DB.AutoMigrate(&models.User{}, &models.Post{}, &models.PostVersion{}, &models.Category{}, &models.Tag{})
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
