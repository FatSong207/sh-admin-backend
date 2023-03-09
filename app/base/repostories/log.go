package repostories

import (
	"SH-admin/app/base/core"
	"SH-admin/app/base/interface_repostories"
	"SH-admin/app/models"
	"SH-admin/app/models/common"
	"gorm.io/gorm"
)

type LogRepostory struct {
	core.BaseRepostory[models.Log, models.LogOutDto]
}

// NewLogRepostory CTOR
func NewLogRepostory() interface_repostories.ILogRepostory {
	ins := &LogRepostory{}
	return ins
}

func (l LogRepostory) FindLoginlogWithPager(searchDto common.PageInfo, db *gorm.DB, order string, dest *[]*models.LoginlogOutDto, bind *[]*models.LoginlogOutDto) (int64, error) {
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
