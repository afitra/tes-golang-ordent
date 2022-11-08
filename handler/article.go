package handler

import (
	"net/http"
	"tes-golang-ordent/article"
	"tes-golang-ordent/helper"
	"tes-golang-ordent/user"

	"github.com/gin-gonic/gin"
)

type articleHandler struct {
	service     article.Service
	userService user.Service
}

func NewArticleHandler(service article.Service, userService user.Service) *articleHandler {
	return &articleHandler{service, userService}
}

func (h *articleHandler) CreateArticle(c *gin.Context) {
	var input article.ArticleDataInput
	currentUser := c.MustGet("currentUser").(user.User)

	input.User_id = currentUser.ID
	err := c.ShouldBindJSON(&input)

	if err != nil {

		errors := helper.FormatValidationError(err)

		errorMessage := gin.H{"errors": errors}

		response := helper.ApiResponse("Failed to create article", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return

	}

	newArticle, err := h.service.CreateArticle(input)

	if err != nil {
		response := helper.ApiResponse("Failed to create article", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.ApiResponse("Succes  create  article", http.StatusCreated, "success", article.FormatArticle(newArticle))
	c.JSON(http.StatusCreated, response)
}
func (h *articleHandler) GetArticle(c *gin.Context) {
	var inputID article.ArticleParamInput
	err := c.ShouldBindUri(&inputID)

	if err != nil {
		response := helper.ApiResponse("Failed to update article", http.StatusBadRequest, "error", nil)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	articleData, err := h.service.GetArticleByID(inputID)

	if articleData.ID == 0 {
		response := helper.ApiResponse("article data not found", http.StatusBadRequest, "error", nil)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ApiResponse("Succes  get article data", http.StatusOK, "success", article.FormatArticle(articleData))
	c.JSON(http.StatusOK, response)
}
func (h *articleHandler) UpdateArticle(c *gin.Context) {
	var inputID article.ArticleParamInput

	err := c.ShouldBindUri(&inputID)

	if err != nil {
		response := helper.ApiResponse("Failed to update article", http.StatusBadRequest, "error", nil)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	articleData, err := h.service.GetArticleByID(inputID)

	if articleData.ID == 0 {
		response := helper.ApiResponse("article data not found", http.StatusBadRequest, "error", nil)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	var input article.ArticleDataInput
	currentUser := c.MustGet("currentUser").(user.User)

	input.User_id = currentUser.ID
	err = c.ShouldBindJSON(&input)

	if err != nil {

		errors := helper.FormatValidationError(err)

		errorMessage := gin.H{"errors": errors}

		response := helper.ApiResponse("Failed to update article", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return

	}

	newArticle, err := h.service.UpdateArticleData(inputID, input)

	if err != nil {
		response := helper.ApiResponse("Failed to create article", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.ApiResponse("Succes  create  article", http.StatusOK, "success", article.FormatArticle(newArticle))
	c.JSON(http.StatusOK, response)
}
func (h *articleHandler) DeleteArticle(c *gin.Context) {
	var inputID article.ArticleParamInput
	err := c.ShouldBindUri(&inputID)

	if err != nil {
		response := helper.ApiResponse("Failed to update article", http.StatusBadRequest, "error", nil)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	articleData, err := h.service.GetArticleByID(inputID)

	if articleData.ID == 0 {
		response := helper.ApiResponse("article data not found", http.StatusBadRequest, "error", nil)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	err = h.service.DestroyArticle(inputID)

	if err != nil {
		response := helper.ApiResponse("Failed to delete article", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.ApiResponse("Succes  delete  article", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}

func (h *articleHandler) GetAllArticle(c *gin.Context) {
	allData, err := h.service.GetAllArticleData()
	if err != nil {
		response := helper.ApiResponse("Failed to get data article", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.ApiResponse("Succes to get  article", http.StatusOK, "success", article.FormatAllArticle(allData))
	c.JSON(http.StatusOK, response)
}
