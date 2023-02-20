package models

import (
	"gorm.io/gorm"
	"time"
)

type Api struct {
	Id          int64     `gorm:"primaryKey"`
	Created     time.Time `json:"created" form:"created"`
	Updated     time.Time `json:"updated" form:"updated"`
	Path        string    `json:"path" form:"path"`
	Description string    `json:"description" form:"description"`
	ApiGroup    string    `json:"apiGroup" form:"apiGroup"`
	Method      string    `json:"method" form:"method"`
}

type ApiOutDto struct {
	Id          int64     `gorm:"primaryKey"`
	Created     time.Time `json:"created" form:"created"`
	Updated     time.Time `json:"updated" form:"updated"`
	Path        string    `json:"path" form:"path"`
	Description string    `json:"description" form:"description"`
	ApiGroup    string    `json:"apiGroup" form:"apiGroup"`
	Method      string    `json:"method" form:"method"`
}

type ApiInDto struct {
	Id          int64  `gorm:"primaryKey"`
	Path        string `json:"path" form:"path"`
	Description string `json:"description" form:"description"`
	ApiGroup    string `json:"apiGroup" form:"apiGroup"`
	Method      string `json:"method" form:"method"`
}

func (a Api) TableName() string {
	return "api"
}

func (a *Api) BeforeCreate(db *gorm.DB) (err error) {
	a.Created = time.Now()
	a.Updated = time.Now()
	return
}

func (a *Api) BeforeUpdate(db *gorm.DB) (err error) {
	db.Statement.SetColumn("updated", time.Now())
	return
}
