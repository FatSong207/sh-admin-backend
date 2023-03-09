package models

import (
	"gorm.io/gorm"
	"time"
)

type Menu struct {
	Id         int64  `gorm:"primaryKey"`
	MenuLevel  int64  `json:"-"`
	MenuType   int64  `json:"menuType"`
	ParentId   string `json:"parentId"`
	Path       string `json:"path"`
	Name       string `json:"name"`
	ChName     string `json:"chname"`
	Hidden     bool   `json:"hidden"`
	Component  string `json:"component"`
	Sort       int64  `json:"sort"`
	ActiveName string `json:"activeName"`
	Meta       `json:"meta"`
	Children   []Menu `json:"children" gorm:"-"`
	Created    int64  `gorm:"created" form:"created"`
	Updated    int64  `gorm:"updated" form:"updated"`
}

type Meta struct {
	KeepAlive   bool   `json:"keepAlive" gorm:"comment:是否缓存"`           // 是否缓存
	DefaultMenu bool   `json:"defaultMenu" gorm:"comment:是否是基础路由（开发中）"` // 是否是基础路由（开发中）
	Title       string `json:"title" gorm:"comment:菜单名"`                // 菜单名
	Icon        string `json:"icon" gorm:"comment:菜单图标"`                // 菜单图标
	CloseTab    bool   `json:"closeTab" gorm:"comment:自动关闭tab"`         // 自动关闭tab
}

type MenuOutDto struct {
	Menu
	MenuId   int64        `json:"menuId" gorm:"comment:菜单ID"`
	RoleId   int64        `json:"-" gorm:"comment:角色ID"`
	Children []MenuOutDto `json:"children" gorm:"-"`
}

type MenuOutDto2 struct {
	Id        int64  `gorm:"primaryKey"`
	ParentId  string `json:"parentId"`
	ParentIds string `json:"parentIds"`
	MenuType  int64  `json:"menuType"`
	Path      string `json:"path"`
	Name      string `json:"name"`
	ChName    string `json:"chname"`
	Hidden    bool   `json:"hidden"`
	Component string `json:"component"`
	Sort      int64  `json:"sort"`
	Title     string `json:"title" gorm:"comment:菜单名"` // 菜单名
	Icon      string `json:"icon" gorm:"comment:菜单图标"` // 菜单图标
	Created   int64  `gorm:"created" form:"created"`
	Updated   int64  `gorm:"updated" form:"updated"`
}

type MenuInDto struct {
	Id        int64  `gorm:"primaryKey"`
	MenuType  int64  `json:"menuType"`
	ParentId  string `json:"parentId"`
	Path      string `json:"path"`
	Name      string `json:"name"`
	ChName    string `json:"chname"`
	Hidden    bool   `json:"hidden"`
	Component string `json:"component"`
	Sort      int64  `json:"sort"`
	KeepAlive bool   `json:"keepAlive" gorm:"comment:是否缓存"` // 是否缓存
	Title     string `json:"title" gorm:"comment:菜单名"`      // 菜单名
	Icon      string `json:"icon" gorm:"comment:菜单图标"`      // 菜单图标
}

func (Menu) TableName() string {
	return "menu"
}

// BeforeUpdate 更新前鉤子
func (m *Menu) BeforeUpdate(db *gorm.DB) (err error) {
	db.Statement.SetColumn("updated", time.Now().Unix())
	return
}

func (m *Menu) BeforeCreate(db *gorm.DB) (err error) {
	m.Created = time.Now().Unix()
	return
}
