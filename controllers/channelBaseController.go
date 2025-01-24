package controllers

import (
	"myYoku/services"

	"github.com/gin-gonic/gin"
)

type ChannelBaseController struct {
	ChannelBaseRouterGroup *gin.RouterGroup
	ChannelBaseService     *services.ChannelBaseService
}

// 工厂函数
func NewChannelBaseController(r *gin.Engine, name string, channelBaseService *services.ChannelBaseService) *ChannelBaseController {
	rGroup := r.Group("/" + name)
	return &ChannelBaseController{ChannelBaseRouterGroup: rGroup, ChannelBaseService: channelBaseService}
}

// 路由注册
func (c *ChannelBaseController) InitCannelBaseController() {
	c.ChannelBaseRouterGroup.GET("/channel/region", c.ChannelBaseService.GetChannelRegion)
	c.ChannelBaseRouterGroup.GET("/channel/type", c.ChannelBaseService.GetChannelType)
}
