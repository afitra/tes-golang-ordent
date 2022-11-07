package main

import (
	"fmt"
	"log"
	"os"

	"tes-golang-ordent/auth"
	"tes-golang-ordent/handler"
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
		// &article.Article{}
	)
	fmt.Println("\n koneksi dataBase berhasil *******\n")

	userRepository := user.NewRepository(dataBase)
	// articleRepository := article.NewRepository(dataBase)

	userService := user.NewService(userRepository)
	// articleService := asset.NewService(articleRepository)

	authService := auth.NewService()
	userHandler := handler.NewUserHandler(userService, authService)
	// articleHandler := handler.NewWalletHandler(articleService, authService)

	router := gin.Default()
	api := router.Group("/api")

	api.POST("/register", userHandler.RegisterUser)
	api.POST("/login", userHandler.Login)
	// =================================================== Wallet API
	// api.POST("/article", articleHandler.Login)

	router.Run()

}
