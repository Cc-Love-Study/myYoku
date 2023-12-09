/* 这个是分类 地区 类型的实体*/
package models

type Barrage struct {
	Id          int    `gorm:"primaryKey;column:id;" `
	Content     string `gorm:"column:content;" `
	AddTime     int64  `gorm:"column:add_time;" `
	UserId      int    `gorm:"column:user_id;" `
	Status      int    `gorm:"column:status;" `
	EpisodesId  int    `gorm:"column:episodes_id;" `
	VideoId     int    `gorm:"column:video_id;" `
	CurrentTime int    `gorm:"column:current_time;" `
}

type BarrageData struct {
	Id          int    `json:"id" gorm:"primaryKey;column:id;"`
	Content     string `json:"content" gorm:"column:content;"`
	CurrentTime int    `json:"currentTime" gorm:"column:current_time;"`
}

type BarrageGetInfo struct {
	CurrentTime int `json:"currentTime"`
	EpisodesId  int `json:"episodesId"`
}

func (c Barrage) TableName() string {
	return "barrage"
}

func (c BarrageData) TableName() string {
	return "barrage"
}
