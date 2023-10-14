package models

type Video struct {
	Id       int    `gorm:"primaryKey;column:id;"`
	Name     string `gorm:"column:name;" `
	Password string `gorm:"column:password;" `
	Status   int    `gorm:"column:status;" `
	AddTime  int64  `gorm:"column:add_time;" `
	Mobile   string `gorm:"column:mobile;" `
	Avatar   string `gorm:"column:avatar;" `
}

func (p Video) TableName() string {
	return "video"
}

// 工厂函数
func NewVideo() *Video {
	return &Video{}
}
