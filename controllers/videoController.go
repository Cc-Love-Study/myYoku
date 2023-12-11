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
	// 获得正在热播
	v.VideoRouterGroup.GET("/channel/hot", v.VideoService.ChannelHotVideo)
	// 根据地区推荐地区视频
	v.VideoRouterGroup.GET("/channel/recommend/region", v.VideoService.ChannelRecommendRegionVideo)
	// 根据类型推荐地区视频
	v.VideoRouterGroup.GET("/channel/recommend/type", v.VideoService.ChannelRecommendTypeVideo)
	// 根据筛选 获取结果
	v.VideoRouterGroup.GET("/channel/video", v.VideoService.ChannelSelectVideo)
	// 获得视频详细消息
	v.VideoRouterGroup.GET("/video/info", v.VideoService.GetVideoInfo)
	// 得到视频的所以集
	v.VideoRouterGroup.GET("/video/episodes/list", v.VideoService.GetVideoEpisodesInfo)
	// 获得channel排行榜信息
	v.VideoRouterGroup.GET("/channel/top", v.VideoService.GetChannelTop)
	// 获得type排行榜信息
	v.VideoRouterGroup.GET("/type/top", v.VideoService.GetTypeTop)
	// 获得用户视频信息
	v.VideoRouterGroup.GET("/user/video", v.VideoService.UserVideo)
	// 保存用户上传视频信息
	v.VideoRouterGroup.POST("/video/save", v.VideoService.VideoSave)
	// // 保存用户上传到阿里云视频信息
	// v.VideoRouterGroup.POST("/video/save", v.VideoService.VideoSave)
}
