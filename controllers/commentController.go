package controllers

import (
	"myYoku/services"

	"github.com/gin-gonic/gin"
)

type CommentController struct {
	CommentRouterGroup *gin.RouterGroup
	CommentService     *services.CommentService
}

// 工厂函数
func NewCommentController(r *gin.Engine, name string, commentService *services.CommentService) *CommentController {
	rGroup := r.Group("/" + name)
	return &CommentController{CommentRouterGroup: rGroup, CommentService: commentService}
}

// 路由注册
func (c *CommentController) InitCommentController() {
	// 获得广告
	c.CommentRouterGroup.GET("/comment/list", c.CommentService.GetCommentList)
	// 保存评论
	c.CommentRouterGroup.POST("/comment/save", c.CommentService.CommentSave)
}
