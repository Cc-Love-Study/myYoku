package services

import (
	"myYoku/daos"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AdvertService struct {
	AdvertDao *daos.AdvertDao
	Utils     *Utils
}

// 工厂函数
func NewAdvertService(advertDao *daos.AdvertDao, utils *Utils) *AdvertService {
	return &AdvertService{AdvertDao: advertDao, Utils: utils}
}

func (a *AdvertService) ChannelAdvert(c *gin.Context) {
	channelId := ""
	channelId = c.Query("channelId")
	if channelId == "" {
		c.JSON(http.StatusOK, a.Utils.ReturnError(4011, "channelId为空"))
		return
	}
	channelIdInt, err := strconv.Atoi(channelId)
	if err != nil {
		c.JSON(http.StatusOK, a.Utils.ReturnError(4012, "channelId不为数字"))
		return
	} else {
		err, adverts := a.AdvertDao.FindAdvert(channelIdInt)
		if err != nil {
			c.JSON(http.StatusOK, a.Utils.ReturnError(4013, "广告查询错误"))
			return
		} else {
			c.JSON(http.StatusOK, a.Utils.ReturnSucess(0, "success", adverts, int64(len(adverts))))
			return
		}
	}
}
