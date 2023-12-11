package controllers

import (
	"myYoku/services"

	"github.com/gin-gonic/gin"
)

type AliyunController struct {
	AliyunRouterGroup *gin.RouterGroup
	AliyunService     *services.AliyunService
}

// 工厂函数
func NewAliyunController(r *gin.Engine, name string, aliyunService *services.AliyunService) *AliyunController {
	rGroup := r.Group("/" + name)
	return &AliyunController{AliyunRouterGroup: rGroup, AliyunService: aliyunService}
}

// 路由注册
func (a *AliyunController) InitAliyunController() {
	// 阿里云上传地址和凭证
	a.AliyunRouterGroup.POST("/aliyun/create/upload/video", a.AliyunService.GetAliyunUploadAddrAndAuth)
	// 阿里云刷新凭证
	a.AliyunRouterGroup.POST("/aliyun/refresh/upload/video", a.AliyunService.ReflashAliyunUploadAuth)
	// 阿里云播放凭证
	a.AliyunRouterGroup.POST("/aliyun/video/play/auth", a.AliyunService.PlayAuth)
}
