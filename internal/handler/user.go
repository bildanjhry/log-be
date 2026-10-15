package handler

import (
	"fmt"
	"net/http"

	"github.com/bildanjhry/log-be/internal/service"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) GetUsers(c *gin.Context) {
	users, err := h.service.GetUsers(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Something went wrong",
		})
		fmt.Print(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Success get data",
		"data":    users,
	})

}
