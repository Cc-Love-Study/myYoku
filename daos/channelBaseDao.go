package daos

import (
	"myYoku/models"

	"github.com/jinzhu/gorm"
)

type ChannelBaseDao struct {
	DbOrm *gorm.DB
}

// 工厂函数
func NewChannelBaseDao(db *gorm.DB) *ChannelBaseDao {
	return &ChannelBaseDao{DbOrm: db}
}

func (ch *ChannelBaseDao) FindChannelRegion(channelId int) (error, []models.ChannelRegion) {
	var regions []models.ChannelRegion
	err := ch.DbOrm.Where("channel_id=? AND status=1", channelId).Order("sort DESC").
		Select([]string{"id", "name"}).
		Find(&regions).Error

	return err, regions
}

func (ch *ChannelBaseDao) FindChannelType(channelId int) (error, []models.ChannelType) {
	var types []models.ChannelType
	err := ch.DbOrm.Where("channel_id=? AND status=1", channelId).Order("sort DESC").
		Select([]string{"id", "name"}).
		Find(&types).Error

	return err, types
}
