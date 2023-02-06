package models

import "gorm.io/gorm"

type RoleAuthorize struct {
	RoleId      int64 `json:"roleId"`
	AuthorizeId int64 `json:"-"`
}

func (RoleAuthorize) TableName() string {
	return "role_authorize"
}

func (a RoleAuthorize) BeforeCreate(db *gorm.DB) (err error) {
	return
}

func (a RoleAuthorize) BeforeUpdate(db *gorm.DB) (err error) {
	return
}
