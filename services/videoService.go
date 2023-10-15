package services

import (
	"myYoku/daos"
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
