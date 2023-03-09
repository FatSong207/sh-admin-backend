package models

import (
	"gorm.io/gorm"
	"time"
)

type TaskJob struct {
	Id          int64     `gorm:"primaryKey"`
	TaskName    string    `json:"taskName" form:"taskName" gorm:"taskName"`
	Description string    `json:"description" from:"description"`
	Cron        string    `json:"cron" form:"cron" gorm:"cron"`
	Status      int       `json:"status" gorm:"status" form:"status"`
	Created     time.Time `json:"created" gorm:"created" form:"created"`
	Updated     time.Time `json:"updated" gorm:"updated" form:"updated"`
}

func (t TaskJob) TableName() string {
	return "task_job"
}

func (t *TaskJob) BeforeCreate(db *gorm.DB) (err error) {
	t.Created = time.Now()
	t.Updated = time.Now()
	return
}

func (t *TaskJob) BeforeUpdate(db *gorm.DB) (err error) {
	db.Statement.SetColumn("updated", time.Now())
	return
}
