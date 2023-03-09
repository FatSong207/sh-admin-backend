package models

import (
	"gorm.io/gorm"
	"time"
)

type Role struct {
	Id      int64     `gorm:"primaryKey"`
	Created time.Time `json:"created" form:"created"` // 創建時間
	Updated time.Time `json:"updated" form:"updated"` // 更新時間
	Name    string    `json:"name" form:"name"`
	EnName  string    `json:"enName" form:"enName"`
}

type RoleOutDto struct {
	Id      int64     `gorm:"primaryKey"`
	Created time.Time `json:"created" form:"created"` // 創建時間
	Updated time.Time `json:"updated" form:"updated"` // 更新時間
	Name    string    `json:"name" form:"name"`
	EnName  string    `json:"enName" form:"enName"`
}

type RoleInDto struct {
	Id     int64  `gorm:"primaryKey"`
	Name   string `json:"name" form:"name"`
	EnName string `json:"enName" form:"enName"`
}

func (r Role) TableName() string {
	return "role"
}

func (r *Role) BeforeCreate(db *gorm.DB) (err error) {
	r.Created = time.Now()
	r.Updated = time.Now()
	return
}

func (r *Role) BeforeUpdate(db *gorm.DB) (err error) {
	db.Statement.SetColumn("updated", time.Now())
	return
}
