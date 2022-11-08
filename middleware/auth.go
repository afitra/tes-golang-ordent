package Middleware

import (
	"fmt"
	"net/http"
	"tes-golang-ordent/article"
	"tes-golang-ordent/auth"
	"tes-golang-ordent/helper"
	"tes-golang-ordent/user"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func IsUser(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("token")
		token, err := authService.ValidateToken(authHeader)

		if err != nil {
			response := helper.ApiResponse("Unauthorized", http.StatusUnauthorized, "error", nil)

			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, errorValidationToken := token.Claims.(jwt.MapClaims)
		if !errorValidationToken || !token.Valid {
			response := helper.ApiResponse("Unauthorized", http.StatusUnauthorized, "error", nil)

			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		userID := int(claim["user_id"].(float64))

		user, err := userService.GetUserByID(userID)

		if err != nil {

			response := helper.ApiResponse("Unauthorized", http.StatusUnauthorized, "error", nil)

			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		if user.Role != "user" {
			response := helper.ApiResponse("user   have not access this api", http.StatusUnauthorized, "error", nil)

			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("currentUser", user)
	}
}
func IsAdmin(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("token")
		token, err := authService.ValidateToken(authHeader)

		if err != nil {
			response := helper.ApiResponse("Unauthorized", http.StatusUnauthorized, "error", nil)

			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, errorValidationToken := token.Claims.(jwt.MapClaims)
		if !errorValidationToken || !token.Valid {
			response := helper.ApiResponse("Unauthorized", http.StatusUnauthorized, "error", nil)

			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		userID := int(claim["user_id"].(float64))

		user, err := userService.GetUserByID(userID)

		if err != nil {

			response := helper.ApiResponse("Unauthorized", http.StatusUnauthorized, "error", nil)

			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		if user.Role != "admin" {
			response := helper.ApiResponse("user   have not access this api", http.StatusUnauthorized, "error", nil)

			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("currentUser", user)
	}
}
func Authorize(articleService article.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		currentUser := c.MustGet("currentUser").(user.User)
		var inputID article.ArticleParamInput
		err := c.ShouldBindUri(&inputID)
		if err != nil {
			response := helper.ApiResponse("Failed to update article", http.StatusBadRequest, "error", nil)

			c.JSON(http.StatusBadRequest, response)
			return
		}
		fmt.Println(currentUser.ID, "user >>>>>>> articleId", inputID.ID)

		article, err := articleService.CostomFind(currentUser.ID, inputID)
		fmt.Println("<<<<<<<<", article)
		if article.ID == 0 {
			fmt.Println("masokkkkkk")
			response := helper.ApiResponse("user have not access this article", http.StatusUnauthorized, "error", nil)

			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

	}
}
