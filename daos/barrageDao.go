package daos

import (
	"fmt"
	"myYoku/models"
	"time"

	"github.com/jinzhu/gorm"
)

type BarrageDao struct {
	DbOrm *gorm.DB
}

// 工厂函数
func NewBarrageDao(db *gorm.DB) *BarrageDao {
	return &BarrageDao{DbOrm: db}
}

func (b *BarrageDao) BarrageList(episodesId int, startTime int, endTime int) (err error, barrageList []models.BarrageData) {
	err = b.DbOrm.Where("episodes_id=? and `current_time`>? and `current_time`<? and status=1", episodesId, startTime, endTime).Order("`current_time` ASC").Find(&barrageList).Error
	return
}

func (b *BarrageDao) BarrageSave(episodesId int, videoId int, currentTime int, userId int, content string) (err error) {
	var barrage models.Barrage
	barrage.AddTime = time.Now().Unix()
	barrage.Content = content
	barrage.CurrentTime = currentTime
	barrage.EpisodesId = episodesId
	barrage.UserId = userId
	barrage.Status = 1
	barrage.VideoId = videoId
	fmt.Println(barrage)
	err = b.DbOrm.Create(&barrage).Error
	return
}
