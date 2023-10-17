package models

type ChannelType struct {
	Id   int    `gorm:"primaryKey;column:id;" `
	Name string `gorm:"column:name;" `
}

func (c ChannelType) TableName() string {
	return "channel_type"
}
