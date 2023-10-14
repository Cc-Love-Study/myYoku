package controllers

import (
	"myYoku/services"

	"github.com/gin-gonic/gin"
)

type VideoController struct {
	VideoRouterGroup *gin.RouterGroup
	VideoService     *services.VideoService
}

// 工厂函数
func NewVideoController(r *gin.Engine, name string, videoService *services.VideoService) *VideoController {
	rGroup := r.Group("/" + name)
	return &VideoController{VideoRouterGroup: rGroup, VideoService: videoService}
}

// 路由注册
func (v *VideoController) InitVideoController() {
	// 获得广告
	v.VideoRouterGroup.GET("/channel/advert")

}
