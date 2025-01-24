package services

import (
	"crypto/md5"
	"encoding/hex"
	"time"
)

// 这个文件存放了各种返回值的格式和类型
type Utils struct {
	Md5Key string
}

type JsonStruct struct {
	Code  int         `json:"code"`
	Msg   interface{} `json:"msg"`
	Items interface{} `json:"items"`
	Count int64       `json:"count"`
}

func NewUtils(md5key string) *Utils {
	return &Utils{Md5Key: md5key}
}

// 成功返回的应包含 代码 信息 数据 计数
func (u *Utils) ReturnSucess(code int, msg interface{}, items interface{}, count int64) *JsonStruct {
	json := &JsonStruct{Code: code, Msg: msg, Items: items, Count: count}
	return json
}

// 失败返回 代码 信息
func (u *Utils) ReturnError(code int, msg interface{}) *JsonStruct {
	json := &JsonStruct{Code: code, Msg: msg}
	return json
}

// 加密函数
func (u *Utils) MD5V(password string) string {
	h := md5.New()
	h.Write([]byte(password + u.Md5Key))
	return hex.EncodeToString(h.Sum(nil))
}

func (u *Utils) FormatTime(otime int64) string {
	video_time := time.Unix(otime, 0)
	return video_time.Format("2006-01-02")
}
