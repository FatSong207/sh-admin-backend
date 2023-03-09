package models

import "gorm.io/gorm"

type Customer struct {
	Id       int64  `gorm:"primaryKey" json:"id,omitempty"`
	Name     string `gorm:"name" json:"name,omitempty"`
	Source   string `gorm:"source" json:"source,omitempty"`
	Phone    string `gorm:"phone" json:"phone,omitempty"`
	Email    string `gorm:"email" json:"email,omitempty"`
	Industry string `gorm:"industry" json:"industry,omitempty"`
	Level    string `gorm:"level" json:"level,omitempty"`
	Remarks  string `gorm:"remarks" json:"remarks,omitempty"`
	Region   string `gorm:"region" json:"region,omitempty"`
	Address  string `gorm:"address" json:"address,omitempty"`
	Status   int    `gorm:"status" json:"status,omitempty"`
	Creator  int64  `gorm:"creator" json:"creator,omitempty"`
	Created  int64  `gorm:"created" json:"created,omitempty"`
	Updated  int64  `gorm:"updated" json:"updated,omitempty"`
}

type CustomerOutDto struct {
	Id       int64  `gorm:"primaryKey"`
	Name     string `gorm:"name" json:"name"`
	Source   string `gorm:"source" json:"source"`
	Phone    string `gorm:"phone" json:"phone"`
	Email    string `gorm:"email" json:"email"`
	Industry string `gorm:"industry" json:"industry"`
	Level    string `gorm:"level" json:"level"`
	Remarks  string `gorm:"remarks" json:"remarks"`
	Region   string `gorm:"region" json:"region"`
	Address  string `gorm:"address" json:"address"`
	Status   int    `gorm:"status" json:"status"`
	//Creator  int64  `gorm:"creator"`
	Created int64 `gorm:"created" json:"created"`
	//Updated  int64  `gorm:"updated"`
}

func (c Customer) TableName() string {
	return "customer"
}

func (c Customer) BeforeCreate(db *gorm.DB) (err error) {
	return
}

func (c Customer) BeforeUpdate(db *gorm.DB) (err error) {
	return
}

