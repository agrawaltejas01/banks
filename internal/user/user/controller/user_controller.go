package user_controller

import (
	"net/http"

	user_interfaces "github.com/agrawaltejas01/banks/internal/user/user/interfaces"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	service user_interfaces.UserService
}

func NewUserController(service user_interfaces.UserService) *UserController {
	return &UserController{service: service}
}

func (c *UserController) CreateUser(ctx *gin.Context) {
	var req struct {
		Name  string `json:"name" binding:"required"`
		Email string `json:"email" binding:"required"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := c.service.CreateUser(req.Name, req.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (c *UserController) GetUserById(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := c.service.GetUserById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	ctx.JSON(http.StatusOK, user)
}
