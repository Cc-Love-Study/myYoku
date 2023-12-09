package controllers

import (
	"myYoku/services"

	"github.com/gin-gonic/gin"
)

type BarrageController struct {
	BarrageRouterGroup *gin.RouterGroup
	BarrageService     *services.BarrageService
}

// 工厂函数
func NewBarrageController(r *gin.Engine, name string, barrageService *services.BarrageService) *BarrageController {
	rGroup := r.Group("/" + name)
	return &BarrageController{BarrageRouterGroup: rGroup, BarrageService: barrageService}
}

// 路由注册
func (b *BarrageController) InitBarrageController() {
	// 获得弹幕
	b.BarrageRouterGroup.GET("/barrage/ws", b.BarrageService.BarrageWs)
	// 获得弹幕测试
	b.BarrageRouterGroup.GET("/barrage/wsss", b.BarrageService.BarrageGet)
	// 弹幕发送
	b.BarrageRouterGroup.POST("/barrage/save", b.BarrageService.BarrageSave)
}
