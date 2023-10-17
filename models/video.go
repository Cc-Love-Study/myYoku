package models

type Video struct {
	Id                 int    `gorm:"primaryKey;column:id;" `
	Title              string `gorm:"column:title;" `
	SubTitle           string `gorm:"column:sub_title;"`
	Img                string `gorm:"column:img;" `
	Img1               string `gorm:"column:img1;" `
	AddTime            int    `gorm:"column:add_time;" `
	EpisodesCount      int    `gorm:"column:episodes_count;" `
	IsEnd              int    `gorm:"column:is_end;" `
	ChannelId          int    `gorm:"column:channel_id;" `
	Status             int    `gorm:"column:status;" `
	RegionId           int    `gorm:"column:region_id;" `
	TypeId             int    `gorm:"column:type_id;" `
	EpisodesUpdateTime int    `gorm:"column:episodes_update_time;" `
}

func (p Video) TableName() string {
	return "video"
}

// 工厂函数
func NewVideo() *Video {
	return &Video{}
}

type SelectVideoConditions struct {
	ChannelId int    `form:"channelId" binding:"required"`
	RegionId  int    `form:"regionId" binding:"-"`
	TypeId    int    `form:"typeId" binding:"-"`
	End       string `form:"end" binding:"-"`
	Sort      string `form:"sort" binding:"-"`
	Limit     int    `form:"limit" binding:"-"`
	Offset    int    `form:"offset" binding:"-"`
}
