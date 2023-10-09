package controllers

import (
	"myYoku/services"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserRouterGroup *gin.RouterGroup
}

// 声明userService的结构体
var userFunc *services.UserService = services.NewUserService()

// 工厂函数
func NewUserController(r *gin.Engine, name string) *UserController {
	rGroup := r.Group("/" + name)
	return &UserController{UserRouterGroup: rGroup}
}

// 路由注册
func (u *UserController) InitUserController() {
	u.UserRouterGroup.GET("/register/save", userFunc.UserRegister)
}
