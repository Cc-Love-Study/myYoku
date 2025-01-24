package models

type Comment struct {
	Id          int    `gorm:"primaryKey;column:id;" `
	Content     string `gorm:"column:content;" `
	AddTime     int64  `gorm:"column:add_time;" `
	UserId      int    `gorm:"column:user_id;"`
	Status      int    `gorm:"column:status;" `
	Stamp       int    `gorm:"column:stamp;" `
	PraiseCount int    `gorm:"column:praise_count;" `
	EpisodesId  int    `gorm:"column:episodes_id;" `
	VideoId     int    `gorm:"column:video_id;"`
}

func (p Comment) TableName() string {
	return "comment"
}

// 工厂函数
func NewComment() *Comment {
	return &Comment{}
}

type CommentUserInfo struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	AddTime int    `json:"addTime"`
	Avatar  string `json:"avatar"`
}

func (p CommentUserInfo) TableName() string {
	return "user"
}

// 工厂函数
func NewCommentUserInfo() *CommentUserInfo {
	return &CommentUserInfo{}
}

type CommentShowInfo struct {
	Id           int              `json:"id"`
	Content      string           `json:"content"`
	AddTime      int64            `json:"addTime"`
	AddTimeTitle string           `json:"addTimeTitle"`
	UserId       int              `json:"userId"`
	Stamp        int              `json:"stamp"`
	PariseCount  int              `json:"praiseCount"`
	UserInfo     *CommentUserInfo `json:"userinfo"`
}

// 工厂函数
func NewCommentShowInfo() *CommentShowInfo {
	return &CommentShowInfo{}
}
