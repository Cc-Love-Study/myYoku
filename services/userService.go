package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}
func (u *UserService) UserRegister(c *gin.Context) {
	c.String(http.StatusOK, "注册成功！！")
}
