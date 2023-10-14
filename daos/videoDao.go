package daos

import (
	"github.com/jinzhu/gorm"
)

type VideoDao struct {
	DbOrm *gorm.DB
}

// 工厂函数
func NewVideoDao(db *gorm.DB) *VideoDao {
	return &VideoDao{DbOrm: db}
}
