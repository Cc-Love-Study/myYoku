package services

import (
	"myYoku/daos"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ChannelBaseService struct {
	ChannelBaseDao *daos.ChannelBaseDao
	Utils          *Utils
}

// 工厂函数
func NewChannelBaseService(channelBaseDao *daos.ChannelBaseDao, utils *Utils) *ChannelBaseService {
	return &ChannelBaseService{ChannelBaseDao: channelBaseDao, Utils: utils}
}

func (ch *ChannelBaseService) GetChannelRegion(c *gin.Context) {
	channelId := ""
	channelId = c.Query("channelId")
	if channelId == "" {
		c.JSON(http.StatusOK, ch.Utils.ReturnError(4041, "channelId为空"))
		return
	}
	channelIdInt, err := strconv.Atoi(channelId)
	if err != nil {
		c.JSON(http.StatusOK, ch.Utils.ReturnError(4042, "channelId不为数字"))
		return
	} else {
		err, regions := ch.ChannelBaseDao.FindChannelRegion(channelIdInt)
		if err != nil {
			c.JSON(http.StatusOK, ch.Utils.ReturnError(4043, "ChannelRegion查询错误"))
			return
		} else {
			c.JSON(http.StatusOK, ch.Utils.ReturnSucess(0, "success", regions, int64(len(regions))))
			return
		}
	}
}

func (ch *ChannelBaseService) GetChannelType(c *gin.Context) {
	channelId := ""
	channelId = c.Query("channelId")
	if channelId == "" {
		c.JSON(http.StatusOK, ch.Utils.ReturnError(4051, "channelId为空"))
		return
	}
	channelIdInt, err := strconv.Atoi(channelId)
	if err != nil {
		c.JSON(http.StatusOK, ch.Utils.ReturnError(4052, "channelId不为数字"))
		return
	} else {
		err, types := ch.ChannelBaseDao.FindChannelType(channelIdInt)
		if err != nil {
			c.JSON(http.StatusOK, ch.Utils.ReturnError(4053, "ChannelType查询错误"))
			return
		} else {
			c.JSON(http.StatusOK, ch.Utils.ReturnSucess(0, "success", types, int64(len(types))))
			return
		}
	}
}
