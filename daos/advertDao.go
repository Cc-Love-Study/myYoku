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
	err := a.DbOrm.Find(&adverts, "channel_id=?", channelId).Error
	return err, adverts
}
