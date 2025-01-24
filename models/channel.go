/* 这个是分类 地区 类型的实体*/
package models

type Channel struct {
	Id   int    `gorm:"primaryKey;column:id;" `
	Name string `gorm:"column:name;" `
}

func (c Channel) TableName() string {
	return "channel"
}
