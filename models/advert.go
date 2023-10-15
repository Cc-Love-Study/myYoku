package models

type Advert struct {
	Id    int    `gorm:"primaryKey;column:id;" `
	Title string `gorm:"column:title;" `
	// SubTitle  string `gorm:"column:sub_title;" json:"sub_title"`
	// ChannelId int    `gorm:"column:channel_id;" json:"channel_id"`
	Img string `gorm:"column:img;" `
	// Sort      string `gorm:"column:sort;" json:"sort"`
	AddTime int    `gorm:"column:add_time;" `
	Url     string `gorm:"column:url;" `
	// Status    int    `gorm:"column:status;" json:"status"`
}

func (p Advert) TableName() string {
	return "advert"
}

// 工厂函数
func NewAdvert() *Advert {
	return &Advert{}
}
