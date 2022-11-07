package handler

import (
	"net/http"
	"tes-golang-ordent/auth"
	"tes-golang-ordent/helper"
	"tes-golang-ordent/user"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
	authService auth.Service
}

func NewUserHandler(userService user.Service, authService auth.Service) *userHandler {
	return &userHandler{userService, authService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {

	var input user.RegisterUserInput
	var email user.EmailInput
	err := c.ShouldBindJSON((&input))
	email.Email = input.Email

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.ApiResponse("register fail", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	available, err := h.userService.GetUserByEmail(email)

	if available == false {
		response := helper.ApiResponse("email already registered", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		response := helper.ApiResponse("Register Account failed", http.StatusBadRequest, "error", nil)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	token, err := h.authService.GenerateToken(newUser.ID)
	if err != nil {
		response := helper.ApiResponse("register fail", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := user.FormatUser(newUser, token)

	response := helper.ApiResponse("register success", http.StatusCreated, "success", formatter)

	c.JSON(http.StatusCreated, response)

}

func (h *userHandler) Login(c *gin.Context) {

	var input user.LoginInput

	err := c.ShouldBindJSON(&input)

	if err != nil {

		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.ApiResponse("Login Account failed", http.StatusUnprocessableEntity, "error", errorMessage)

		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedInUser, err := h.userService.Login(input)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.ApiResponse("Login Account failed", http.StatusUnprocessableEntity, "error", errorMessage)

		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	token, err := h.authService.GenerateToken(loggedInUser.ID)
	if err != nil {
		response := helper.ApiResponse("Login  failed", http.StatusBadRequest, "error", nil)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(loggedInUser, token)

	response := helper.ApiResponse("Login successfull", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)

}
