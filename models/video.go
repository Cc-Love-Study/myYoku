package models

type Video struct {
	Id            int    `gorm:"primaryKey;column:id;" `
	Title         string `gorm:"column:title;" `
	SubTitle      string `gorm:"column:sub_title;"`
	Img           string `gorm:"column:img;" `
	Img1          string `gorm:"column:img1;" `
	AddTime       int    `gorm:"column:add_time;" `
	EpisodesCount int    `gorm:"column:episodes_count;" `
	IsEnd         int    `gorm:"column:is_end;" `
}

func (p Video) TableName() string {
	return "video"
}

// 工厂函数
func NewVideo() *Video {
	return &Video{}
}
