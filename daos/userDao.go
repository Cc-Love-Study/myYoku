package daos

import (
	"myYoku/models"
	"time"

	"github.com/jinzhu/gorm"
)

type UserDao struct {
	DbOrm *gorm.DB
}

// 工厂函数
func NewUserDao(db *gorm.DB) *UserDao {
	return &UserDao{DbOrm: db}
}

//判断手机号是否存在 如果不存在 返回True
func (u *UserDao) IsUserMobile(mobile string) bool {
	is := u.DbOrm.Find(models.NewUser(), "mobile=?", mobile).RecordNotFound()
	return is
}

//插入用户
func (u *UserDao) SaveUser(mobile string, password string) error {
	user := models.NewUser()
	user.Mobile = mobile
	user.Password = password
	user.AddTime = time.Now().Unix()
	err := u.DbOrm.Create(user).Error
	return err
}

//查询用户
func (u *UserDao) FindUser(mobile string) (error, *models.User) {
	user := models.NewUser()
	err := u.DbOrm.Find(user, "mobile=?", mobile).Error
	return err, user
}
