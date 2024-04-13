package handlers

import (
	"github.com/awwabi/monorepo-academics/cuti/internal/services"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) *UserHandler {
	return &UserHandler{userService}
}

func (h *UserHandler) SayHello(c *gin.Context) {
	username := c.Query("username")

	c.String(200, "Hello, %s", username)
}
