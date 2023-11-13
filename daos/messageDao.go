package daos

import (
	"myYoku/models"
	"time"

	"github.com/jinzhu/gorm"
)

type MessageDao struct {
	DbOrm *gorm.DB
}

// 工厂函数
func NewMessageDao(db *gorm.DB) *MessageDao {
	return &MessageDao{DbOrm: db}
}

// 保存发送信息
func (m *MessageDao) SaveMessage(content string) (error, int) {
	message := models.NewMessage()
	message.Content = content
	message.AddTime = time.Now().Unix()
	err := m.DbOrm.Create(message).Error
	return err, message.Id
}

// 保存消息接收人
func (m *MessageDao) SaveMessageUser(userId int, messageId int) error {
	messageUser := models.NewMessageUser()
	messageUser.AddTime = time.Now().Unix()
	messageUser.UserId = userId
	messageUser.MessageId = messageId
	messageUser.Status = 1
	err := m.DbOrm.Create(messageUser).Error
	return err
}
