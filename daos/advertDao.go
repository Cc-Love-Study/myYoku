package daos

import (
	"myYoku/models"

	"github.com/jinzhu/gorm"
)

type AdvertDao struct {
	DbOrm *gorm.DB
}

// 工厂函数
func NewAdvertDao(db *gorm.DB) *AdvertDao {
	return &AdvertDao{DbOrm: db}
}

func (a *AdvertDao) FindAdvert(channelId int) (error, []models.Advert) {
	var adverts []models.Advert
	err := a.DbOrm.Where("channel_id=? AND status=1", channelId).Order("sort DESC").
		Select([]string{"id", "title", "sub_title", "img", "add_time", "url"}).
		First(&adverts).Error

	return err, adverts
}
