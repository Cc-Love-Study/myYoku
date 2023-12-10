package models

type VideoEpisodes struct {
	Id            int    `gorm:"primaryKey;column:id;" `
	Title         string `gorm:"column:title;" `
	Num           int    `gorm:"column:num;" `
	PlayUrl       string `gorm:"column:play_url;" `
	Comment       int    `gorm:"column:comment;" `
	AddTime       int64  `gorm:"column:add_time;" `
	VideoId       int    `gorm:"column:video_id;" `
	Status        int    `gorm:"column:status;" `
	AliYunVideoId string `gorm:"column:aliyun_video_id;" `
}

func (p VideoEpisodes) TableName() string {
	return "video_episodes"
}

// 工厂函数
func NewVideoDetail() *VideoEpisodes {
	return &VideoEpisodes{}
}
