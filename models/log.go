package models

import (
	"gorm.io/gorm"
	"time"
)

type Log struct {
	Id           int64         `gorm:"primaryKey"`
	Created      time.Time     `json:"created" form:"created"`             // 創建時間
	Updated      time.Time     `json:"updated" form:"updated"`             // 更新時間
	Ip           string        `json:"ip" form:"ip" `                      // 來源ip
	Method       string        `json:"method" form:"method" `              // 方法
	Path         string        `json:"path" form:"path"`                   // api路徑
	Type         string        `json:"type" form:"type"`                   // 類型（登入或一般操作）
	Status       int           `json:"status" form:"status" `              // 狀態
	Latency      time.Duration `json:"latency" form:"latency" `            // 延遲
	Agent        string        `json:"agent" form:"agent"`                 // 代理
	ErrorMessage string        `json:"error_message" form:"error_message"` // 錯誤訊息
	Body         string        `json:"body" form:"body"`                   // RequestBody
	Response     string        `json:"response" form:"resp"`               // ResponseBody
	Code         int           `json:"code" form:"code"`                   // 響應的code
	UserID       int           `json:"user_id" form:"user_id"`             // 用户id
	//User         SysUser       `json:"user"`
}

type LogOutDto struct {
	Id           int64     `gorm:"primaryKey"`
	Created      time.Time `json:"created" form:"created"`             // 創建時間
	Updated      time.Time `json:"updated" form:"updated"`             // 更新時間
	Ip           string    `json:"ip" form:"ip" `                      // 來源ip
	Method       string    `json:"method" form:"method" `              // 方法
	Path         string    `json:"path" form:"path"`                   // api路徑
	Type         string    `json:"type" form:"type"`                   // 類型（登入或一般操作）
	Status       int       `json:"status" form:"status" `              // 狀態
	ErrorMessage string    `json:"error_message" form:"error_message"` // 錯誤訊息
	Body         string    `json:"body" form:"body"`                   // RequestBody
	Response     string    `json:"response" form:"resp"`               // ResponseBody
	Code         int       `json:"code" form:"code"`                   // 響應的code
	UserID       int       `json:"user_id" form:"user_id"`             // 用户id
}

func (u Log) TableName() string {
	return "log"
}

func (u *Log) BeforeCreate(db *gorm.DB) (err error) {
	u.Created = time.Now()
	u.Updated = time.Now()
	return
}

func (u *Log) BeforeUpdate(db *gorm.DB) (err error) {
	return
}
