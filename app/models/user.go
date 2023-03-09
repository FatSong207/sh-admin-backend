package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id       int64  `gorm:"primaryKey" form:"id"`
	Email    string `gorm:"email" form:"email"`
	Password string `gorm:"password" form:"password" json:"-"`
	Name     string `gorm:"name" form:"name"`
	Version  int    `gorm:"version" form:"version"`
	Expired  int64  `gorm:"expired" form:"expired"`
	Status   int    `gorm:"status" form:"status"`
	RoleId   int64  `gorm:"role_id" form:"roleId"`
	Created  int64  `gorm:"created" form:"created"`
	Updated  int64  `gorm:"updated" form:"updated"`
}

type UserOutDto struct {
	Id       int64  `gorm:"primaryKey"  form:"id" json:"id"`
	Email    string `gorm:"email" form:"email" json:"email"`
	Name     string `gorm:"name" form:"name" json:"name"`
	Version  int    `gorm:"version" form:"version" json:"version"`
	Expired  int64  `gorm:"expired" form:"expired" json:"expired"`
	Status   int    `gorm:"status" form:"status" json:"status"`
	RoleId   int64  `gorm:"role_id" form:"roleId" json:"roleId"`
	RoleName string `gorm:"roleName" form:"roleName" json:"roleName"`
	Created  int64  `gorm:"created" form:"created" json:"created"`
	Updated  int64  `gorm:"updated" form:"updated" json:"updated"`
}

// UserLoginReq 登入參數
type UserLoginReq struct {
	//Email    string `json:"email" binding:"required,email"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// UserLoginRes 登入返回結果
type UserLoginRes struct {
	Token string `json:"token"`
	User  User   `json:"user"`
	//ExpiresAt *jwt.NumericDate `json:"expiresAt"`
}

// UserRegisterReq 註冊參數
type UserRegisterReq struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Code     string `json:"code" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (u User) TableName() string {
	return "user"
}

func (u *User) BeforeCreate(db *gorm.DB) (err error) {
	u.Created = time.Now().Unix()
	u.Updated = time.Now().Unix()
	return
}

func (u *User) BeforeUpdate(db *gorm.DB) (err error) {
	db.Statement.SetColumn("updated", time.Now().Unix())
	return
}
