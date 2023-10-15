package daos

import (
	"myYoku/models"

	"github.com/jinzhu/gorm"
)

type VideoDao struct {
	DbOrm *gorm.DB
}

// 工厂函数
func NewVideoDao(db *gorm.DB) *VideoDao {
	return &VideoDao{DbOrm: db}
}

func (v *VideoDao) FindHotVideo(channelId int) (error, []models.Video) {
	var videos []models.Video
	err := v.DbOrm.
		Where("channel_id=? AND is_hot=1 AND status=1", channelId).
		Order("episodes_update_time DESC").
		Select([]string{"id", "title", "sub_title", "img", "img1", "add_time", "episodes_count", "is_end"}).
		Limit(9).Find(&videos).Error

	return err, videos
}
