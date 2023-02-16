package Repostories

import (
	"SH-admin/common/Core"
	"SH-admin/common/IRepostories"
	"SH-admin/models"
	response "SH-admin/models/common"
	"gorm.io/gorm"
)

type LogRepostory struct {
	Core.BaseRepostory[models.Log, models.LogOutDto]
}

// NewLogRepostory CTOR
func NewLogRepostory() IRepostories.ILogRepostory {
	ins := &LogRepostory{}
	return ins
}

func (l LogRepostory) FindLoginlogWithPager(searchDto response.PageInfo, db *gorm.DB, order string, dest *[]*models.LoginlogOutDto, bind *[]*models.LoginlogOutDto) (int64, error) {
	limit := searchDto.PageSize
	offset := searchDto.PageSize * (searchDto.PageNum - 1)
	var t models.Log
	name := t.TableName()
	//global.Db.Offset(offset).Limit(limit).Table(name).Where(query).Order(order).Find(dest)
	//res := global.Db.Table(name).Where(query).Find(bind)
	res := db.Table(name).Find(bind)
	total := res.RowsAffected
	db = db.Offset(offset).Limit(limit).Table(name).Order(order).Find(dest)
	return total, res.Error
}
