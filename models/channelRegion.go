package models

type ChannelRegion struct {
	Id   int    `gorm:"primaryKey;column:id;" `
	Name string `gorm:"column:name;" `
}

func (c ChannelRegion) TableName() string {
	return "channel_region"
}
