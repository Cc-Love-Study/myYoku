package services

import (
	"myYoku/daos"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type MessageService struct {
	MessageDao *daos.MessageDao
	Utils      *Utils
}

// 工厂函数
func NewMessageService(messageDao *daos.MessageDao, utils *Utils) *MessageService {
	return &MessageService{MessageDao: messageDao, Utils: utils}
}

// 发送信息
func (m *MessageService) SendMessageDo(c *gin.Context) {
	var (
		uids    string
		content string
	)
	uids = c.PostForm("uids")
	content = c.PostForm("content")
	if uids == "" {
		c.JSON(http.StatusOK, m.Utils.ReturnError(4901, "uids不可为空"))
		return
	}
	if content == "" {
		c.JSON(http.StatusOK, m.Utils.ReturnError(4902, "content不可为空"))
		return
	}
	err, messageId := m.MessageDao.SaveMessage(content)
	if err != nil {
		c.JSON(http.StatusOK, m.Utils.ReturnError(4903, "message插入错误"))
		return
	} else {
		uidConfig := strings.Split(uids, ",")
		for _, v := range uidConfig {
			uidInt, err := strconv.Atoi(v)
			if err != nil {
				c.JSON(http.StatusOK, m.Utils.ReturnError(4903, "userId 转换失败"))
				return
			}
			err = m.MessageDao.SaveMessageUser(uidInt, messageId)
			if err != nil {
				c.JSON(http.StatusOK, m.Utils.ReturnError(4903, "messageUser插入错误"))
				return
			}
		}
		c.JSON(http.StatusOK, m.Utils.ReturnSucess(0, "发送成功", nil, 1))
		return
	}
}
