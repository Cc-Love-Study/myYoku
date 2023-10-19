package services

import (
	"myYoku/daos"
	"myYoku/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CommentService struct {
	CommentDao *daos.CommentDao
	Utils      *Utils
}

// 工厂函数
func NewCommentService(commentDao *daos.CommentDao, utils *Utils) *CommentService {
	return &CommentService{CommentDao: commentDao, Utils: utils}
}

func (com *CommentService) GetCommentList(c *gin.Context) {
	episodesId := ""
	episodesId = c.Query("episodesId")
	offset := ""
	offset = c.Query("offset")
	limit := ""
	limit = c.Query("limit")
	episodesIdInt, err := strconv.Atoi(episodesId)
	offsetInt, err := strconv.Atoi(offset)
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		c.JSON(http.StatusOK, com.Utils.ReturnError(4081, "转换整型失败"))
		return
	}
	if limitInt == 0 {
		limitInt = 12
	}
	if episodesIdInt == 0 {
		c.JSON(http.StatusOK, com.Utils.ReturnError(4082, "指定未指定集数"))
		return
	}
	nums, err, comments := com.CommentDao.GetCommentList(episodesIdInt, offsetInt, limitInt)
	if err != nil {
		c.JSON(http.StatusOK, com.Utils.ReturnError(4083, "评论查询失败"))
		return
	} else {
		// 查询到了评论信息
		// 根据评论信息也需要用户信息 渲染用户头像等需要
		var showDatas []*models.CommentShowInfo
		data := models.NewCommentShowInfo()
		for _, v := range comments {
			data.Id = v.Id
			data.Content = v.Content
			data.AddTime = v.AddTime
			data.AddTimeTitle = com.Utils.FormatTime(int64(data.AddTime))
			data.UserId = v.UserId
			data.Stamp = v.Stamp
			data.PariseCount = v.PraiseCount
			// 获得用户信息
			err, data.UserInfo = com.CommentDao.FindCommentUser(v.UserId)
			if err != nil {
				c.JSON(http.StatusOK, com.Utils.ReturnError(4084, "用户查询失败"))
				return
			}
			showDatas = append(showDatas, data)
		}
		c.JSON(http.StatusOK, com.Utils.ReturnSucess(0, "success", showDatas, int64(nums)))
		return
	}
}

func (com *CommentService) CommentSave(c *gin.Context) {
	content := c.PostForm("content")
	uid := c.PostForm("uid")
	episodesId := c.PostForm("episodesId")
	videoId := c.PostForm("videoId")

	if content == "" {
		c.JSON(http.StatusOK, com.Utils.ReturnError(4091, "内容不能为空"))
		return
	}
	uidInt, err := strconv.Atoi(uid)
	episodesIdInt, err := strconv.Atoi(episodesId)
	videoIdInt, err := strconv.Atoi(videoId)
	if err != nil {
		c.JSON(http.StatusOK, com.Utils.ReturnError(4092, "Id转换错误"))
		return
	}
	if uidInt == 0 {
		c.JSON(http.StatusOK, com.Utils.ReturnError(4093, "未登录"))
		return
	}
	if episodesIdInt == 0 {
		c.JSON(http.StatusOK, com.Utils.ReturnError(4094, "必须指定集数"))
		return
	}
	if videoIdInt == 0 {
		c.JSON(http.StatusOK, com.Utils.ReturnError(4095, "必须指定番剧"))
		return
	}

	// 这里其实还需要事务来保持 原子性
	err = com.CommentDao.SaveComment(uidInt, episodesIdInt, videoIdInt, content)
	if err != nil {
		c.JSON(http.StatusOK, com.Utils.ReturnError(4096, "评论插入失败"))
		return
	}
	err = com.CommentDao.VideoAddCommentCount(videoIdInt)
	if err != nil {
		c.JSON(http.StatusOK, com.Utils.ReturnError(4096, "评论插入失败"))
		return
	}
	err = com.CommentDao.EpisodesAddCommentCount(episodesIdInt)
	if err != nil {
		c.JSON(http.StatusOK, com.Utils.ReturnError(4096, "评论插入失败"))
		return
	}
	c.JSON(http.StatusOK, com.Utils.ReturnSucess(0, "success", nil, 1))
	return
	// 这里其实还需要事务来保持 原子性
}
