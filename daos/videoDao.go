package daos

import (
	"myYoku/models"
	"time"

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
	if selectConditions.Sort == "episodesUpdateTime" {
		se = se.Order("episodes_update_time DESC")
	} else if selectConditions.Sort == "comment" {
		se = se.Order("comment DESC")
	} else if selectConditions.Sort == "addTime" {
		se = se.Order("add_time DESC")
	} else {
		se = se.Order("add_time DESC")
	}
	se = se.Limit(selectConditions.Limit).Offset(selectConditions.Offset)
	err := se.Find(&videos).Error
	return err, videos
}

func (v *VideoDao) GetVideoInfo(videoId int) (error, *models.Video) {
	videoDetail := models.NewVideo()
	err := v.DbOrm.Where("id=? AND status=1", videoId).Limit(1).Find(videoDetail).Error
	return err, videoDetail
}

func (v *VideoDao) GetVideoEpisodes(videoId int) (error, []models.VideoEpisodes) {
	var videoEpisodes []models.VideoEpisodes
	err := v.DbOrm.Where("video_id=? AND status=1", videoId).Order("num").Find(&videoEpisodes).Error
	return err, videoEpisodes
}

func (v *VideoDao) GetChannelTop(channelId int) (error, []models.Video) {
	var videos []models.Video
	err := v.DbOrm.
		Where("channel_id=? AND status=1", channelId).
		Order("comment DESC").
		Select([]string{"id", "title", "sub_title", "img", "img1", "add_time", "episodes_count", "is_end"}).
		Limit(10).Find(&videos).Error

	return err, videos
}

func (v *VideoDao) GetTypeTop(typeId int) (error, []models.Video) {
	var videos []models.Video
	err := v.DbOrm.
		Where("type_id=? AND status=1", typeId).
		Order("comment DESC").
		Select([]string{"id", "title", "sub_title", "img", "img1", "add_time", "episodes_count", "is_end"}).
		Limit(10).Find(&videos).Error

	return err, videos
}

func (v *VideoDao) GetUserVideo(uid int) (err error, videos []models.Video) {
	err = v.DbOrm.
		Where("user_id = ?", uid).
		Order("add_time DESC").
		Select([]string{"id", "title", "sub_title", "img", "img1", "add_time", "episodes_count", "is_end"}).
		Find(&videos).Error
	return
}

func (v *VideoDao) SaveVideo(uid int, title string, subTitle string, channleId int, regionId int, typeId int, playUrl string, aliyunVideoId string) (err error) {
	video := models.NewVideo()
	ttime := time.Now().Unix()

	video.SubTitle = subTitle
	video.Title = title
	video.AddTime = ttime
	video.Img = ""
	video.Img1 = ""
	video.EpisodesCount = 1
	video.IsEnd = 1
	video.Status = 1
	video.ChannelId = channleId
	video.RegionId = regionId
	video.TypeId = typeId
	video.EpisodesUpdateTime = ttime
	video.Comment = 0
	video.UserId = uid

	err = v.DbOrm.Create(video).Error

	videoEp := models.NewVideoDetail()
	if err == nil {
		videoEp.VideoId = video.Id
		videoEp.Comment = 0
		videoEp.Num = 1
		videoEp.Status = 1
		videoEp.PlayUrl = playUrl
		videoEp.Title = subTitle
		videoEp.AddTime = ttime
		videoEp.AliYunVideoId = aliyunVideoId
		if aliyunVideoId != "" {
			videoEp.PlayUrl = ""
		}
		err = v.DbOrm.Create(videoEp).Error
	}
	return
}
