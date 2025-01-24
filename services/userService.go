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
			c.JSON(http.StatusOK, u.Utils.ReturnSucess(2000, "创建账户成功", nil, 0))
			return
		}
	}
}

// 用户登录
func (u *UserService) UserLogin(c *gin.Context) {
	var (
		mobile   string
		password string
	)
	mobile = c.PostForm("mobile")
	password = c.PostForm("password")

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
		c.JSON(http.StatusOK, u.Utils.ReturnError(4003, "密码不可为空"))
		return
	}
	ok = u.UserDao.IsUserMobile(mobile)
	if ok {
		c.JSON(http.StatusOK, u.Utils.ReturnError(4005, "账号不存在"))
		return
	} else {
		err, user := u.UserDao.FindUser(mobile)
		if err != nil {
			c.JSON(http.StatusOK, u.Utils.ReturnError(4006, "未知错误！"))
			return
		} else {
			if user.Password == u.Utils.MD5V(password) {
				c.JSON(http.StatusOK, u.Utils.ReturnSucess(0, "登录成功", gin.H{"uid": user.Id, "username": user.Name}, 1))
				return
			} else {
				c.JSON(http.StatusOK, u.Utils.ReturnError(4007, "密码错误"))
				return
			}
		}
	}

}
