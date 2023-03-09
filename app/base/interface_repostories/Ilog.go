package interface_repostories

import (
	"SH-admin/app/base/core"
	"SH-admin/app/models"
	"SH-admin/app/models/common"
	"gorm.io/gorm"
)

type ILogRepostory interface {
	core.IBaseRepostory[models.Log, models.LogOutDto]
	FindLoginlogWithPager(searchDto common.PageInfo, db *gorm.DB, order string, dest *[]*models.LoginlogOutDto, bind *[]*models.LoginlogOutDto) (int64, error)
}
