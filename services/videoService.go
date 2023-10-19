package services

import (
	"fmt"
	"myYoku/daos"
	"myYoku/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type VideoService struct {
	VideoDao *daos.VideoDao
	Utils    *Utils
}

// 工厂函数
func NewVideoService(videoDao *daos.VideoDao, utils *Utils) *VideoService {
	return &VideoService{VideoDao: videoDao, Utils: utils}
}

func (v *VideoService) ChannelHotVideo(c *gin.Context) {
	channelId := ""
	channelId = c.Query("channelId")
	if channelId == "" {
		c.JSON(http.StatusOK, v.Utils.ReturnError(4011, "channelId为空"))
		return
	}
	channelIdInt, err := strconv.Atoi(channelId)
	if err != nil {
		c.JSON(http.StatusOK, v.Utils.ReturnError(4012, "channelId不为数字"))
		return
	} else {
		err, videos := v.VideoDao.FindHotVideo(channelIdInt)
		if err != nil {
			c.JSON(http.StatusOK, v.Utils.ReturnError(4013, "广告查询错误"))
			return
		} else {
			c.JSON(http.StatusOK, v.Utils.ReturnSucess(0, "success", videos, int64(len(videos))))
			return
		}
	}
}

// 推荐视频 根据地区
func (v *VideoService) ChannelRecommendRegionVideo(c *gin.Context) {
	channelId := ""
	channelId = c.Query("channelId")
	regionId := ""
	regionId = c.Query("regionId")

	if channelId == "" {
		c.JSON(http.StatusOK, v.Utils.ReturnError(4021, "channelId为空"))
		return
	}
	if regionId == "" {
		c.JSON(http.StatusOK, v.Utils.ReturnError(4022, "regionId为空"))
		return
	}
	channelIdInt, err1 := strconv.Atoi(channelId)
	regionIdInt, err2 := strconv.Atoi(regionId)
	if err1 != nil || err2 != nil {
		c.JSON(http.StatusOK, v.Utils.ReturnError(4023, "channelId or regionId不为数字"))
		return
	} else {
		err, videos := v.VideoDao.FindRecommendVideoByRegion(channelIdInt, regionIdInt)
		if err != nil {
			c.JSON(http.StatusOK, v.Utils.ReturnError(4024, "地区推荐查询错误"))
			return
		} else {
			c.JSON(http.StatusOK, v.Utils.ReturnSucess(0, "success", videos, int64(len(videos))))
			return
		}
	}
}

// 推荐视频 根据类型
func (v *VideoService) ChannelRecommendTypeVideo(c *gin.Context) {
	channelId := ""
	channelId = c.Query("channelId")
	typeId := ""
	typeId = c.Query("typeId")

	if channelId == "" {
		c.JSON(http.StatusOK, v.Utils.ReturnError(4031, "channelId为空"))
		return
	}
	if typeId == "" {
		c.JSON(http.StatusOK, v.Utils.ReturnError(4032, "type"))
		return
	}
	channelIdInt, err1 := strconv.Atoi(channelId)
	typeIdInt, err2 := strconv.Atoi(typeId)
	if err1 != nil || err2 != nil {
		c.JSON(http.StatusOK, v.Utils.ReturnError(4033, "channelId or regionId不为数字"))
		return
	} else {
		err, videos := v.VideoDao.FindRecommendVideoByType(channelIdInt, typeIdInt)
		if err != nil {
			c.JSON(http.StatusOK, v.Utils.ReturnError(4034, "类型推荐查询错误"))
			return
		} else {
			c.JSON(http.StatusOK, v.Utils.ReturnSucess(0, "success", videos, int64(len(videos))))
			return
		}
	}
}

// 根据条件筛选视频
func (v *VideoService) ChannelSelectVideo(c *gin.Context) {
	var selectConditions models.SelectVideoConditions
	err := c.ShouldBindQuery(&selectConditions)
	fmt.Println(selectConditions)
	if err != nil {
		c.JSON(http.StatusOK, v.Utils.ReturnError(4061, err.Error()))
		return
	}
	if selectConditions.Limit == 0 {
		selectConditions.Limit = 12
	}
	err, videos := v.VideoDao.FindVideo(selectConditions)
	if err != nil {
		c.JSON(http.StatusOK, v.Utils.ReturnError(4062, "视频查询错误"))
		return
	} else {
		c.JSON(http.StatusOK, v.Utils.ReturnSucess(0, "success", videos, int64(len(videos))))
		return
	}
}

// 获得视频详情
func (v *VideoService) GetVideoInfo(c *gin.Context) {
	videoId := ""
	videoId = c.Query("videoId")
	if videoId == "" {
		c.JSON(http.StatusOK, v.Utils.ReturnError(4071, "videoId 为空"))
		return
	}
	videoIdInt, err := strconv.Atoi(videoId)
	if err != nil {
		c.JSON(http.StatusOK, v.Utils.ReturnError(4072, "videoId 不为数字"))
		return
	}
	err, videoInfo := v.VideoDao.GetVideoInfo(videoIdInt)
	if err != nil {
		c.JSON(http.StatusOK, v.Utils.ReturnError(4073, "videoInfo 查询失败"))
		return
	} else {
		c.JSON(http.StatusOK, v.Utils.ReturnSucess(0, "success", videoInfo, 1))
		return
	}
}

// 获得视频详情
func (v *VideoService) GetVideoEpisodesInfo(c *gin.Context) {
	videoId := ""
	videoId = c.Query("videoId")
	if videoId == "" {
		c.JSON(http.StatusOK, v.Utils.ReturnError(4081, "videoId 为空"))
		return
	}
	videoIdInt, err := strconv.Atoi(videoId)
	if err != nil {
		c.JSON(http.StatusOK, v.Utils.ReturnError(4082, "videoId 不为数字"))
		return
	}
	err, videoInfos := v.VideoDao.GetVideoEpisodes(videoIdInt)
	if err != nil {
		c.JSON(http.StatusOK, v.Utils.ReturnError(4083, "videoInfo 查询失败"))
		return
	} else {
		c.JSON(http.StatusOK, v.Utils.ReturnSucess(0, "success", videoInfos, int64(len(videoInfos))))
		return
	}
}
