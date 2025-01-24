package controllers

import (
	"myYoku/services"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserRouterGroup *gin.RouterGroup
	UserService     *services.UserService
}

// 工厂函数
func NewUserController(r *gin.Engine, name string, userService *services.UserService) *UserController {
	rGroup := r.Group("/" + name)
	return &UserController{UserRouterGroup: rGroup, UserService: userService}
}

// 路由注册
func (u *UserController) InitUserController() {
	u.UserRouterGroup.POST("/register/save", u.UserService.UserRegister)
	u.UserRouterGroup.POST("/login/do", u.UserService.UserLogin)
}
