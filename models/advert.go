package models

type Advert struct {
	Id        int    `gorm:"primaryKey;column:id;"`
	Title     string `gorm:"column:title;" `
	SubTitle  string `gorm:"column:sub_title;" `
	ChannelId int    `gorm:"column:channel_id;" `
	Img       string `gorm:"column:img;" `
	Sort      string `gorm:"column:sort;" `
	AddTime   int    `gorm:"column:add_time;" `
	Url       string `gorm:"column:url;" `
	Status    int    `gorm:"column:status;" `
}

func (p Advert) TableName() string {
	return "advert"
}

// 工厂函数
func NewAdvert() *Advert {
	return &Advert{}
}
