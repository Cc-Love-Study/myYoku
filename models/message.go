package models

type Message struct {
	Id      int
	Content string
	AddTime int64
}

func (c Message) TableName() string {
	return "message"
}

func NewMessage() *Message {
	return &Message{}
}

type MessageUser struct {
	Id        int
	MessageId int
	UserId    int
	AddTime   int64
	Status    int
}

func NewMessageUser() *MessageUser {
	return &MessageUser{}
}
func (c MessageUser) TableName() string {
	return "message_user"
}
