package services

import (
	"encoding/json"
	"fmt"
	"myYoku/daos"
	"myYoku/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type BarrageService struct {
	BarrageDao *daos.BarrageDao
	Utils      *Utils
}

// 工厂函数
func NewBarrageService(barrageDao *daos.BarrageDao, utils *Utils) *BarrageService {
	return &BarrageService{BarrageDao: barrageDao, Utils: utils}
}

// 配置ws
var upGrande = websocket.Upgrader{
	//设置允许跨域
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (b *BarrageService) BarrageWs(c *gin.Context) {

	ws, err := upGrande.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		ws.Close()
		return
	}
	// 这里ws循环进行信息处理
	for {
		_, data, err := ws.ReadMessage()
		if err != nil {
			ws.Close()
			return
		}

		var getInfo models.BarrageGetInfo
		json.Unmarshal(data, &getInfo)
		endTime := getInfo.CurrentTime + 60
		err, barrages := b.BarrageDao.BarrageList(getInfo.EpisodesId, getInfo.CurrentTime, endTime)
		if err != nil {
			ws.Close()
			return
		}
		err = ws.WriteJSON(barrages)
		if err != nil {
			ws.Close()
			return
		}
	}
}

/*测试查询*/
func (b *BarrageService) BarrageGet(c *gin.Context) {
	startTime := ""
	startTime = c.Query("startTime")
	startTimeInt, _ := strconv.Atoi(startTime)

	EpisodesId := ""
	EpisodesId = c.Query("EpisodesId")
	EpisodesIdInt, _ := strconv.Atoi(EpisodesId)

	endTime := startTimeInt + 60
	fmt.Println(EpisodesIdInt)
	fmt.Println(startTimeInt)
	fmt.Println(endTime)
	_, barrages := b.BarrageDao.BarrageList(EpisodesIdInt, startTimeInt, endTime)
	c.JSON(http.StatusOK, b.Utils.ReturnSucess(0, "success", barrages, int64(len(barrages))))
	return
}

func (b *BarrageService) BarrageSave(c *gin.Context) {

	content := c.PostForm("content")
	uid := c.PostForm("uid")
	episodesId := c.PostForm("episodesId")
	videoId := c.PostForm("videoId")
	currentTime := c.PostForm("currentTime")

	uidInt, err := strconv.Atoi(uid)
	episodesIdInt, err := strconv.Atoi(episodesId)
	videoIdInt, err := strconv.Atoi(videoId)
	currentTimeInt, err := strconv.Atoi(currentTime)

	if err != nil {
		c.JSON(http.StatusOK, b.Utils.ReturnError(4444, "未知错误"))
		return
	}
	if content == "" {
		c.JSON(http.StatusOK, b.Utils.ReturnError(4444, "弹幕不可为空"))
		return
	}
	if uidInt == 0 {
		c.JSON(http.StatusOK, b.Utils.ReturnError(4002, "未登录"))
		return
	}
	if episodesIdInt == 0 {
		c.JSON(http.StatusOK, b.Utils.ReturnError(4003, "未指定集数"))
		return
	}
	if videoIdInt == 0 {
		c.JSON(http.StatusOK, b.Utils.ReturnError(4004, "未指定视频ID"))
		return
	}
	if currentTimeInt == 0 {
		c.JSON(http.StatusOK, b.Utils.ReturnError(4004, "未指定视频时间"))
		return
	}
	err = b.BarrageDao.BarrageSave(episodesIdInt, videoIdInt, currentTimeInt, uidInt, content)
	if err != nil {
		c.JSON(http.StatusOK, b.Utils.ReturnError(4005, "插入失败"))
		return
	}
	c.JSON(http.StatusOK, b.Utils.ReturnSucess(0, "success", "", 1))
	return
}
