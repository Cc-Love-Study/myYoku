package daos

import (
	"myYoku/models"
	"time"

	"github.com/jinzhu/gorm"
)

type CommentDao struct {
	DbOrm *gorm.DB
}

// 工厂函数
func NewCommentDao(db *gorm.DB) *CommentDao {
	return &CommentDao{DbOrm: db}
}

func (c *CommentDao) GetCommentList(episodesId int, offset int, limit int) (count int, err error, commentList []models.Comment) {
	c.DbOrm.Where("episodes_id=? AND status=1", episodesId).Find(&commentList).Count(&count)
	err = c.DbOrm.Where("episodes_id=? AND status=1", episodesId).
		Order("add_time DESC").Limit(limit).Offset(offset).Find(&commentList).Error
	return
}

//查询用户
func (c *CommentDao) FindCommentUser(userId int) (error, *models.CommentUserInfo) {
	user := models.NewCommentUserInfo()
	err := c.DbOrm.Select([]string{"id", "name", "add_time", "avatar"}).Find(user, "id=?", userId).Error
	return err, user
}

// 保存评论
func (c *CommentDao) SaveComment(uid int, episodesId int, videoId int, content string) error {
	comment := models.NewComment()
	comment.UserId = uid
	comment.EpisodesId = episodesId
	comment.VideoId = videoId
	comment.Content = content
	comment.Stamp = 0
	comment.Status = 1
	comment.AddTime = time.Now().Unix()
	comment.PraiseCount = 0
	err := c.DbOrm.Create(comment).Error
	return err
}

// video表增加一条评论数
func (c *CommentDao) VideoAddCommentCount(videoId int) error {
	// fmt.Println("的video评论数进行+1", videoId)
	err := c.DbOrm.Model(models.NewVideo()).Where("id=?", videoId).Update("comment", gorm.Expr("comment + ?", 1)).Error
	// fmt.Println("错误:", err)
	return err
}

// video_epis表增加一条评论数
func (c *CommentDao) EpisodesAddCommentCount(episodesId int) error {
	err := c.DbOrm.Model(models.NewVideoDetail()).Where("id=?", episodesId).Update("comment", gorm.Expr("comment + ?", 1)).Error
	return err
}
