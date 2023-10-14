package services

import (
	"myYoku/daos"
)

type VideoService struct {
	VideoDao *daos.VideoDao
	Utils    *Utils
}

// 工厂函数
func NewVideoService(videoDao *daos.VideoDao, utils *Utils) *VideoService {
	return &VideoService{VideoDao: videoDao, Utils: utils}
}
