package main

import (
	"fmt"
	"log"
	"os"

	"tes-golang-ordent/article"
	"tes-golang-ordent/auth"
	"tes-golang-ordent/handler"
	"tes-golang-ordent/like"
	Middleware "tes-golang-ordent/middleware"
	"tes-golang-ordent/user"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file, please check your .env file")
	}

	var DB_PORT = []byte(os.Getenv("DB_PORT"))
	var MYSQL_USERNAME = []byte(os.Getenv("MYSQL_USERNAME"))
	var MYSQL_PASSWORD = []byte(os.Getenv("MYSQL_PASSWORD"))
	var MYSQL_HOST = []byte(os.Getenv("MYSQL_HOST"))
	var DB_NAME = []byte(os.Getenv("DB_NAME"))

	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		MYSQL_USERNAME,
		MYSQL_PASSWORD,
		MYSQL_HOST,
		DB_PORT,
		DB_NAME,
	)
	dataBase, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	dataBase.Debug().AutoMigrate(
		&user.User{},
		&article.Article{},
		&like.Like{},
	)
	fmt.Println("\n koneksi dataBase berhasil *******\n")

	userRepository := user.NewRepository(dataBase)
	articleRepository := article.NewRepository(dataBase)
	likeRepository := like.NewRepository(dataBase)

	userService := user.NewService(userRepository)
	articleService := article.NewService(articleRepository)
	likeService := like.NewService(likeRepository)

	authService := auth.NewService()

	userHandler := handler.NewUserHandler(userService, authService)
	articleHandler := handler.NewArticleHandler(articleService, userService)
	likeHandler := handler.NewLikeHandler(likeService)

	router := gin.Default()
	api := router.Group("/api")

	api.POST("/register", userHandler.RegisterUser)
	api.POST("/login", userHandler.Login)
	// =================================================== Article API
	api.POST("/article", Middleware.IsUser(authService, userService), articleHandler.CreateArticle)
	api.GET("/article/:id", Middleware.IsUser(authService, userService), Middleware.Authorize(articleService), articleHandler.GetArticle)
	api.PUT("/article/:id", Middleware.IsUser(authService, userService), Middleware.Authorize(articleService), articleHandler.UpdateArticle)
	api.DELETE("/article/:id", Middleware.IsUser(authService, userService), Middleware.Authorize(articleService), articleHandler.DeleteArticle)
	// =================================================== Like API
	api.POST("/article/like/:id", Middleware.IsUser(authService, userService), likeHandler.Like)
	api.DELETE("/article/dislike/:id", Middleware.IsUser(authService, userService), likeHandler.Dislike)

	// =================================================== Article API
	api.GET("/article/all", Middleware.IsAdmin(authService, userService), articleHandler.GetAllArticle)

	router.Run()

}
