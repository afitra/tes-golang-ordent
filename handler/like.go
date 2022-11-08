package handler

import (
	"net/http"
	"tes-golang-ordent/helper"
	"tes-golang-ordent/like"
	"tes-golang-ordent/user"

	"github.com/gin-gonic/gin"
)

type likeHandler struct {
	service like.Service
}

func NewLikeHandler(service like.Service) *likeHandler {
	return &likeHandler{service}
}

func (h *likeHandler) Like(c *gin.Context) {
	var input like.LikeDataInput
	currentUser := c.MustGet("currentUser").(user.User)
	input.User_id = currentUser.ID

	c.ShouldBindUri(&input)

	available, err := h.service.FindLike(input)

	if available.ID != 0 {
		response := helper.ApiResponse("Failed this article has been be like ", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	newLike, err := h.service.CreateLike(input)

	if err != nil {
		response := helper.ApiResponse("Failed to like articleb ", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.ApiResponse("Succes  like  article", http.StatusCreated, "success", like.FormatLike(newLike))
	c.JSON(http.StatusCreated, response)
}
func (h *likeHandler) Dislike(c *gin.Context) {
	var input like.LikeDataInput
	currentUser := c.MustGet("currentUser").(user.User)
	input.User_id = currentUser.ID

	c.ShouldBindUri(&input)

	available, err := h.service.FindLike(input)

	if available.ID == 0 {
		response := helper.ApiResponse("Failed this article has been be like ", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	err = h.service.DeleteLike(input)

	if err != nil {
		response := helper.ApiResponse("Failed to dislike article ", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.ApiResponse("Succes  dislike  article", http.StatusCreated, "success", nil)
	c.JSON(http.StatusCreated, response)
}
