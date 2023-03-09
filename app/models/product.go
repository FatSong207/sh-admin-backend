package models

import (
	"gorm.io/gorm"
	"time"
)

type Product struct {
	Id int64 `gorm:"primaryKey"`
	//Name        string  `gorm:"name" form:"name" binding:"required,oneof=运动鞋 洗衣液"`
	Name        string  `gorm:"name" form:"name"`
	Type        int     `gorm:"type" form:"type"`
	Unit        string  `gorm:"unit" form:"unit"`
	Code        string  `gorm:"code" form:"code"`
	Price       float64 `gorm:"price" form:"price"`
	Description string  `gorm:"description" form:"description"`
	Status      int     `gorm:"status" form:"status"`
	Creator     int64   `gorm:"creator" form:"creator"`
	Created     int64   `gorm:"created" form:"created"`
	Updated     int64   `gorm:"updated" form:"updated"`
}

type ProductOutDto struct {
	Id          int64   `gorm:"primaryKey"`
	Name        string  `gorm:"name"`
	Type        int     `gorm:"type"`
	Unit        string  `gorm:"unit"`
	Code        string  `gorm:"code"`
	Price       float64 `gorm:"price"`
	Description string  `gorm:"description"`
	Status      int     `gorm:"status"`
	//Creator     int64   `gorm:"creator"`
	Created int64 `gorm:"created"`
	//Updated     int64   `gorm:"updated"`
}

// TableName 繼承自Entity
func (p Product) TableName() string {
	return "product"
}

// BeforeCreate 創建時鉤子
func (p *Product) BeforeCreate(db *gorm.DB) (err error) {
	p.Creator = 207
	p.Created = time.Now().Unix()
	return
}
