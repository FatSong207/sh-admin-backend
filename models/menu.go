package models

type Menu struct {
	Id         int64  `gorm:"primaryKey"`
	MenuLevel  int64  `json:"-"`
	ParentId   string `json:"parentId"`
	Path       string `json:"path"`
	Name       string `json:"name"`
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

func (Menu) TableName() string {
	return "menu"
}
