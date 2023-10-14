package controllers

import (
	"myYoku/services"

	"github.com/gin-gonic/gin"
)

type AdvertController struct {
	AdvertRouterGroup *gin.RouterGroup
	AdvertService     *services.AdvertService
}

// 工厂函数
func NewAdvertController(r *gin.Engine, name string, advertService *services.AdvertService) *AdvertController {
	rGroup := r.Group("/" + name)
	return &AdvertController{AdvertRouterGroup: rGroup, AdvertService: advertService}
}

// 路由注册
func (a *AdvertController) InitAdvertController() {
	// 获得广告
	a.AdvertRouterGroup.GET("/channel/advert", a.AdvertService.ChannelAdvert)
}
