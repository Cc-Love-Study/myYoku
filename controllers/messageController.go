package controllers

import (
	"myYoku/services"

	"github.com/gin-gonic/gin"
)

type MessageController struct {
	MessageRouterGroup *gin.RouterGroup
	MessageService     *services.MessageService
}

// 工厂函数
func NewMessageController(r *gin.Engine, name string, messageService *services.MessageService) *MessageController {
	rGroup := r.Group("/" + name)
	return &MessageController{MessageRouterGroup: rGroup, MessageService: messageService}
}

// 路由注册
func (m *MessageController) InitMessageController() {
	m.MessageRouterGroup.POST("/send/message", m.MessageService.SendMessageDo)
}
