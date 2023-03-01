package models

import "gorm.io/gorm"

type RoleAuthorize struct {
	RoleId      int64 `json:"roleId"`
	AuthorizeId int64 `json:"-"`
}

type RoleAuthorizeUpdateDto struct {
	RoleId       int64   `json:"roleId" form:"roleId"`
	AuthorizeIds []int64 `json:"authorizeIds" form:"authorizeIds"`
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
