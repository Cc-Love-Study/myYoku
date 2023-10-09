package services

import (
	"fmt"
	"myYoku/daos"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
)

type UserService struct {
	UserDao *daos.UserDao
	Utils   *Utils
}

// 工厂函数
func NewUserService(userDao *daos.UserDao, utils *Utils) *UserService {
	return &UserService{UserDao: userDao, Utils: utils}
}

/*用户注册*/
func (u *UserService) UserRegister(c *gin.Context) {
	var (
		mobile   string
		password string
		err      error
	)
	mobile = c.PostForm("mobile")
	password = c.PostForm("password")

	fmt.Println("手机号:", mobile)
	fmt.Println("密码:", password)

	if mobile == "" {
		c.JSON(http.StatusOK, u.Utils.ReturnError(4001, "手机号不可为空"))
		return
	}
	ok, _ := regexp.MatchString(`^1(3|4|5|7|8)[0-9]\d{8}$`, mobile)
	if !ok {
		c.JSON(http.StatusOK, u.Utils.ReturnError(4002, "手机号格式不对"))
		return
	}

	if password == "" {
		c.JSON(http.StatusOK, u.Utils.ReturnError(4003, "密码为空"))
		return
	}

	ok = u.UserDao.IsUserMobile(mobile)
	if !ok {
		c.JSON(http.StatusOK, u.Utils.ReturnError(4005, "手机号已经存在"))
		return
	} else {
		err = u.UserDao.SaveUser(mobile, u.Utils.MD5V(password))
		if err != nil {
			c.JSON(http.StatusOK, u.Utils.ReturnError(4006, "创建失败"))
			return
		} else {
			c.JSON(http.StatusOK, u.Utils.ReturnError(2000, "创建账户成功"))
			return
		}
	}

}
