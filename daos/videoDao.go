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

func (v *VideoDao) FindRecommendVideoByRegion(channelId int, reginId int) (error, []models.Video) {
	var videos []models.Video
	err := v.DbOrm.
		Where("channel_id=? AND is_recommend=1 AND status=1 AND region_id=?", channelId, reginId).
		Order("episodes_update_time DESC").
		Select([]string{"id", "title", "sub_title", "img", "img1", "add_time", "episodes_count", "is_end"}).
		Limit(9).Find(&videos).Error

	return err, videos
}

func (v *VideoDao) FindRecommendVideoByType(channelId int, typeId int) (error, []models.Video) {
	var videos []models.Video
	err := v.DbOrm.
		Where("channel_id=? AND is_recommend=1 AND status=1 AND type_id=?", channelId, typeId).
		Order("episodes_update_time DESC").
		Select([]string{"id", "title", "sub_title", "img", "img1", "add_time", "episodes_count", "is_end"}).
		Limit(9).Find(&videos).Error

	return err, videos
}

func (v *VideoDao) FindVideo(selectConditions models.SelectVideoConditions) (error, []models.Video) {
	var videos []models.Video
	se := v.DbOrm.Where("channel_id=? AND status=1", selectConditions.ChannelId)
	if selectConditions.RegionId > 0 {
		se = se.Where("region_id=?", selectConditions.RegionId)
	}
	if selectConditions.TypeId > 0 {
		se = se.Where("type_id=?", selectConditions.TypeId)
	}
	if selectConditions.End == "n" {
		se = se.Where("is_end=?", 0)
	} else if selectConditions.End == "y" {
		se = se.Where("is_end=?", 1)
	}
	se = se.Limit(selectConditions.Limit).Offset(selectConditions.Offset)
	if selectConditions.Sort == "episodesUpdateTime" {
		se = se.Order("episodes_update_time DESC")
	} else if selectConditions.Sort == "comment" {
		se = se.Order("comment DESC")
	} else if selectConditions.Sort == "addTime" {
		se = se.Order("add_time DESC")
	} else {
		se = se.Order("add_time DESC")
	}
	err := se.Limit(selectConditions.Limit).Offset(selectConditions.Offset).Find(&videos).Error
	return err, videos
}
