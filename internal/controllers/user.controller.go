package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/tuanbui-n9/project-4/internal/services"
	"github.com/tuanbui-n9/project-4/pkg/response"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: services.NewUserService(),
	}
}

func (uc *UserController) GetUserById(c *gin.Context) {
	response.SuccessResponse(c, 20001, uc.userService.GetUserById(c.Param("id")))
}
